// Phase 3 integration test: combat resolution, creatures, incapacitation, clone.
//
// Exit criteria under test: "A player can fight and defeat a spawned creature using
// a skill-gated ability, take damage against the correct HAM pool, and be
// incapacitated/revived-or-cloned correctly."
//
// Test 5 waits on genuine tick-driven retaliation (no mocks); allow ~2 minutes.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"swg-server/internal/protocol"
)

const serverURL = "http://localhost:8080"
const wsURL = "ws://localhost:8080/ws"

type wsClient struct {
	conn     *websocket.Conn
	mu       sync.Mutex
	messages []protocol.WSMessage
}

func (wc *wsClient) readLoop() {
	for {
		_, data, err := wc.conn.ReadMessage()
		if err != nil {
			return
		}
		var msg protocol.WSMessage
		json.Unmarshal(data, &msg)
		wc.mu.Lock()
		wc.messages = append(wc.messages, msg)
		wc.mu.Unlock()
	}
}

func (wc *wsClient) send(msgType string, data interface{}) {
	msg := protocol.WSMessage{Type: msgType, Data: data}
	payload, _ := json.Marshal(msg)
	wc.conn.WriteMessage(websocket.TextMessage, payload)
}

func (wc *wsClient) waitFor(msgType string, timeout time.Duration) (protocol.WSMessage, bool) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		wc.mu.Lock()
		for _, m := range wc.messages {
			if m.Type == msgType {
				wc.mu.Unlock()
				return m, true
			}
		}
		wc.mu.Unlock()
		time.Sleep(50 * time.Millisecond)
	}
	return protocol.WSMessage{}, false
}

func (wc *wsClient) clearMessages() {
	wc.mu.Lock()
	wc.messages = wc.messages[:0]
	wc.mu.Unlock()
}

func (wc *wsClient) lastMessage(msgType string) (protocol.WSMessage, bool) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	for i := len(wc.messages) - 1; i >= 0; i-- {
		if wc.messages[i].Type == msgType {
			return wc.messages[i], true
		}
	}
	return protocol.WSMessage{}, false
}

func main() {
	fmt.Println("=== Phase 3 Integration Test (testbed fork, generic content) ===")
	fmt.Println("Combat resolution, creatures, incapacitation, clone")
	fmt.Println()

	if !waitForServer() {
		fmt.Println("FAIL: Server not reachable")
		os.Exit(1)
	}
	fmt.Println("Server is reachable.")

	accountID, token := registerAccount()
	_ = accountID
	charID := createCharacter(token, "human", "CombatTestChar")
	fmt.Printf("Created character %s\n", charID)

	wc := connectWS(token)
	wc.send(protocol.MsgEnterWorld, protocol.EnterWorldMsg{CharacterID: charID})
	if _, ok := wc.waitFor(protocol.MsgWorldEnter, 5*time.Second); !ok {
		fmt.Println("  [FAIL] did not receive world_enter")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Entered world")

	// Collect creature spawn IDs seen at entry (lairs are within interest radius).
	time.Sleep(500 * time.Millisecond)
	creatureIDs := []string{}
	wc.mu.Lock()
	for _, m := range wc.messages {
		if m.Type == protocol.MsgEntitySpawn {
			data, _ := json.Marshal(m.Data)
			var sp protocol.EntitySpawnMsg
			json.Unmarshal(data, &sp)
			if sp.EntityType == "creature" {
				creatureIDs = append(creatureIDs, sp.EntityID)
			}
		}
	}
	wc.mu.Unlock()
	if len(creatureIDs) == 0 {
		fmt.Println("  [FAIL] no creature spawns visible at entry")
		os.Exit(1)
	}
	fmt.Printf("  Visible creature instances: %d\n", len(creatureIDs))

	// --- Test 1: skill-gated ability ---
	fmt.Println("\n--- Test 1: Skill-gated ability ---")
	wc.clearMessages()
	wc.send(protocol.MsgCombatAction, protocol.CombatActionMsg{TargetID: creatureIDs[0]})
	errMsg, ok := wc.waitFor(protocol.MsgError, 3*time.Second)
	if !ok {
		fmt.Println("  [FAIL] expected skill-gate error, got none")
		os.Exit(1)
	}
	fmt.Println("  [PASS] combat_action rejected with no combat skill trained")
	trainSkill(token, charID, "trainer_brawler_z1", "brawler_novice")
	fmt.Println("  [PASS] Novice Brawler trained")

	// --- Test 2: unknown target rejected ---
	fmt.Println("\n--- Test 2: Unknown target rejected ---")
	wc.clearMessages()
	wc.send(protocol.MsgCombatAction, protocol.CombatActionMsg{TargetID: "does-not-exist"})
	if _, ok := wc.waitFor(protocol.MsgError, 3*time.Second); !ok {
		fmt.Println("  [FAIL] expected unknown-target error, got none")
		os.Exit(1)
	}
	_ = errMsg
	fmt.Println("  [PASS] Unknown target rejected")

	// --- Test 3: posture change ---
	// Runs BEFORE any lair exposure: retaliation starts the moment we close to
	// attack range, and a mid-test incapacitation would reject set_posture and
	// flake the test. At spawn (20,5,0) we are outside creature attack range.
	fmt.Println("\n--- Test 3: Posture change ---")
	wc.clearMessages()
	wc.send(protocol.MsgSetPosture, protocol.PostureMsg{Posture: "kneeling"})
	hamMsg, ok := wc.waitFor(protocol.MsgHAMUpdate, 3*time.Second)
	if !ok {
		fmt.Println("  [FAIL] no ham_update after set_posture")
		os.Exit(1)
	}
	data, _ := json.Marshal(hamMsg.Data)
	var ham protocol.HAMUpdateMsg
	json.Unmarshal(data, &ham)
	if ham.Posture != "kneeling" {
		fmt.Printf("  [FAIL] posture not reflected: got %q\n", ham.Posture)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Posture change reflected in ham_update")
	wc.send(protocol.MsgSetPosture, protocol.PostureMsg{Posture: "standing"})
	time.Sleep(300 * time.Millisecond)

	// --- Test 4: fight and defeat a spawned creature ---
	fmt.Println("\n--- Test 4: Fight and defeat a spawned creature ---")
	// Move onto the first lair cluster (10,5,0) — 10m from spawn, one legal move.
	time.Sleep(1200 * time.Millisecond)
	wc.send(protocol.MsgMove, protocol.MoveMsg{X: 10, Y: 5, Z: 0, Heading: 0})
	time.Sleep(500 * time.Millisecond)
	target := creatureIDs[0]
	killed := false
	// Fast poll loop (300 ms, non-blocking scan): the old 2 s blocking wait per
	// attack stretched the kill over ~25 s, losing a damage race against six
	// retaliators and flaking the test. Kills land in ~5 s now.
	wc.clearMessages()
	for i := 0; i < 60 && !killed; i++ {
		wc.send(protocol.MsgCombatAction, protocol.CombatActionMsg{TargetID: target})
		time.Sleep(300 * time.Millisecond)
		wc.mu.Lock()
		for _, m := range wc.messages {
			if m.Type == protocol.MsgCreatureDeath {
				data, _ := json.Marshal(m.Data)
				var dm protocol.CreatureDeathMsg
				json.Unmarshal(data, &dm)
				if dm.InstanceID == target {
					killed = true
				}
			}
		}
		wc.mu.Unlock()
	}
	if !killed {
		fmt.Println("  [FAIL] creature not defeated within 60 attacks")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Creature defeated (creature_death received)")
	xp := getCombatXP(token, charID)
	if xp <= 0 {
		fmt.Println("  [FAIL] no combat XP awarded")
		os.Exit(1)
	}
	fmt.Printf("  [PASS] Combat XP awarded (combat XP pool: %d)\n", xp)

	// --- Test 5: sustained retaliation → incapacitation (real tick combat) ---
	fmt.Println("\n--- Test 5: Take damage from creature retaliation, incapacitate ---")
	fmt.Println("  Waiting for retaliation damage to incapacitate the player (real tick-driven combat, not simulated)...")
	time.Sleep(1200 * time.Millisecond)
	wc.send(protocol.MsgMove, protocol.MoveMsg{X: 11, Y: 5, Z: 0, Heading: 0})
	incapMsg, ok := wc.waitFor(protocol.MsgIncapacitated, 300*time.Second)
	if !ok {
		fmt.Println("  [FAIL] player never incapacitated within 300s")
		os.Exit(1)
	}
	_ = incapMsg
	fmt.Println("  [PASS] Player incapacitated")
	lastHam, ok := wc.lastMessage(protocol.MsgHAMUpdate)
	if !ok {
		fmt.Println("  [FAIL] no ham_update around incapacitation")
		os.Exit(1)
	}
	data, _ = json.Marshal(lastHam.Data)
	json.Unmarshal(data, &ham)
	if ham.HealthCurrent != 0 {
		fmt.Printf("  [FAIL] expected health_current 0, got %d\n", ham.HealthCurrent)
		os.Exit(1)
	}
	fmt.Println("  [PASS] health_current is 0 (correct HAM pool took the damage)")

	// --- Test 6: clone and verify state ---
	fmt.Println("\n--- Test 6: Clone and verify state ---")
	wc.clearMessages()
	wc.send(protocol.MsgClone, map[string]interface{}{})
	clonedMsg, ok := wc.waitFor(protocol.MsgCloned, 5*time.Second)
	if !ok {
		fmt.Println("  [FAIL] no cloned message")
		os.Exit(1)
	}
	data, _ = json.Marshal(clonedMsg.Data)
	var cloned protocol.ClonedMsg
	json.Unmarshal(data, &cloned)
	if cloned.Zone != "zone-0001" {
		fmt.Printf("  [FAIL] cloned to wrong zone: %q\n", cloned.Zone)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Cloned message received with correct zone")
	hamMsg2, ok := wc.waitFor(protocol.MsgHAMUpdate, 5*time.Second)
	if !ok {
		fmt.Println("  [FAIL] no ham_update after clone")
		os.Exit(1)
	}
	data, _ = json.Marshal(hamMsg2.Data)
	json.Unmarshal(data, &ham)
	if ham.HealthCurrent != 1000 || ham.WoundsHealth != 125 || ham.BattleFatiguePct != 2.0 {
		fmt.Printf("  [FAIL] post-clone state wrong: health=%d wounds=%d bf=%.1f\n",
			ham.HealthCurrent, ham.WoundsHealth, ham.BattleFatiguePct)
		os.Exit(1)
	}
	fmt.Printf("  [PASS] Post-clone state: health=%d, wounds_health=%d, battle_fatigue=%.1f%%\n",
		ham.HealthCurrent, ham.WoundsHealth, ham.BattleFatiguePct)

	fmt.Println("\n=== ALL PHASE 3 TESTS PASSED ===")
	wc.conn.Close()
}

func waitForServer() bool {
	for i := 0; i < 30; i++ {
		resp, err := http.Get(serverURL + "/health")
		if err == nil && resp.StatusCode == 200 {
			resp.Body.Close()
			return true
		}
		if resp != nil {
			resp.Body.Close()
		}
		time.Sleep(500 * time.Millisecond)
	}
	return false
}

func registerAccount() (string, string) {
	username := fmt.Sprintf("tu%d", time.Now().UnixNano()%1000000)
	registerBody, _ := json.Marshal(map[string]string{"username": username, "password": "testpass123"})
	resp, err := http.Post(serverURL+"/api/register", "application/json", bytes.NewReader(registerBody))
	if err != nil {
		log.Fatal("register failed:", err)
	}
	resp.Body.Close()

	loginBody, _ := json.Marshal(map[string]string{"username": username, "password": "testpass123"})
	resp2, err := http.Post(serverURL+"/api/login", "application/json", bytes.NewReader(loginBody))
	if err != nil {
		log.Fatal("login failed:", err)
	}
	var loginResult protocol.LoginResponse
	json.NewDecoder(resp2.Body).Decode(&loginResult)
	resp2.Body.Close()
	return loginResult.AccountID, loginResult.Token
}

func createCharacter(token, speciesID, name string) string {
	body, _ := json.Marshal(map[string]interface{}{
		"species": speciesID,
		"name":    name,
		"appearance": map[string]interface{}{
			"body_type": "average", "skin_color": "default", "height": 1.0,
		},
	})
	req, _ := http.NewRequest("POST", serverURL+"/api/characters", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("create character failed:", err)
	}
	respBody, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode != 201 {
		log.Fatalf("create character %s failed (HTTP %d): %s", name, resp.StatusCode, string(respBody))
	}
	var result protocol.CreateCharacterResponse
	json.Unmarshal(respBody, &result)
	return result.CharacterID
}

func trainSkill(token, charID, trainerID, skillBoxID string) {
	body, _ := json.Marshal(map[string]string{"trainer_id": trainerID, "skill_box_id": skillBoxID})
	req, _ := http.NewRequest("POST", serverURL+"/api/characters/"+charID+"/train", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("train failed:", err)
	}
	respBody, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		log.Fatalf("train %s failed (HTTP %d): %s", skillBoxID, resp.StatusCode, string(respBody))
	}
}

func getCombatXP(token, charID string) int {
	req, _ := http.NewRequest("GET", serverURL+"/api/characters/"+charID+"/skills", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("get skills failed:", err)
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	pools, ok := result["xp_pools"].(map[string]interface{})
	if !ok {
		return 0
	}
	v, ok := pools["combat"].(float64)
	if !ok {
		return 0
	}
	return int(v)
}

func connectWS(token string) *wsClient {
	d := websocket.Dialer{HandshakeTimeout: 5 * time.Second}
	url := fmt.Sprintf("%s?token=%s", wsURL, token)
	conn, _, err := d.Dial(url, nil)
	if err != nil {
		log.Fatal("WebSocket dial failed:", err)
	}
	wc := &wsClient{conn: conn}
	go wc.readLoop()
	return wc
}
