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

// Phase 3 integration test: HAM-driven combat resolution, creatures,
// incapacitation, and clone.
//
// GDD Phase 3 exit criteria: "A player can fight and defeat a spawned
// creature using a skill-gated ability, take damage against the correct
// HAM pool, and be incapacitated/revived-or-cloned correctly (Section 9.6)."

const serverURL = "http://localhost:8080"
const wsURL = "ws://localhost:8080/ws"

// Matches internal/handlers/combat.go's SeedCombatWorld() coordinates.
const wompRatLairX, wompRatLairZ = 3560.0, -4790.0
const worrtLairX, worrtLairZ = 3560.0, -4798.0

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

// waitForMatch polls until a message of msgType arrives *after messageOffset*
// whose decoded data satisfies pred. Searching from an offset (not always
// from index 0) matters here: without it, a call made after an earlier
// matching message was already seen (e.g. Test 4's ham_update with
// incapacitated=false) would immediately re-match that stale message
// instead of waiting for a fresh one.
func (wc *wsClient) waitForMatch(msgType string, messageOffset int, timeout time.Duration, pred func(map[string]interface{}) bool) (map[string]interface{}, bool) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		wc.mu.Lock()
		for _, m := range wc.messages[messageOffset:] {
			if m.Type != msgType {
				continue
			}
			b, _ := json.Marshal(m.Data)
			var d map[string]interface{}
			json.Unmarshal(b, &d)
			if pred(d) {
				wc.mu.Unlock()
				return d, true
			}
		}
		wc.mu.Unlock()
		time.Sleep(50 * time.Millisecond)
	}
	return nil, false
}

func (wc *wsClient) len() int {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	return len(wc.messages)
}

func main() {
	fmt.Println("=== SWG Pre-CU Phase 3 Integration Test ===")
	fmt.Println("Combat resolution, creatures, incapacitation, clone")
	fmt.Println()

	if !waitForServer() {
		fmt.Println("FAIL: Server not reachable")
		os.Exit(1)
	}
	fmt.Println("Server is reachable.")

	// --- Setup ---
	token, _ := registerAccount()
	charID := createCharacter(token, "human", "CombatTester")
	fmt.Printf("Created character %s\n", charID)

	conn, _, err := websocket.DefaultDialer.Dial(wsURL+"?token="+token, nil)
	if err != nil {
		log.Fatal("ws dial failed:", err)
	}
	wc := &wsClient{conn: conn}
	go wc.readLoop()

	wc.send(protocol.MsgEnterWorld, protocol.EnterWorldMsg{CharacterID: charID})
	if _, ok := wc.waitFor(protocol.MsgWorldEnter, 3*time.Second); !ok {
		fmt.Println("FAIL: world_enter not received")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Entered world")

	// --- Test 1: Skill gate — combat_action rejected before training ---
	fmt.Println("\n--- Test 1: Skill-gated ability (GDD Section 30 Phase 3 exit criteria) ---")
	wc.send(protocol.MsgCombatAction, protocol.CombatActionMsg{TargetID: "lair_womp_rat_c0"})
	if _, ok := wc.waitFor(protocol.MsgError, 2*time.Second); !ok {
		fmt.Println("FAIL: expected an error before any combat skill is trained")
		os.Exit(1)
	}
	fmt.Println("  [PASS] combat_action rejected with no combat skill trained")

	// Train Novice Brawler (free — GDD 6.3.1)
	trainers := getTrainers()
	var brawlerTrainerID string
	for _, t := range trainers {
		if t["profession_id"] == "brawler" {
			brawlerTrainerID = t["trainer_id"].(string)
		}
	}
	if brawlerTrainerID == "" {
		fmt.Println("FAIL: no Brawler trainer found")
		os.Exit(1)
	}
	trainSkill(token, charID, brawlerTrainerID, "brawler_novice")
	fmt.Println("  [PASS] Novice Brawler trained")

	// --- Test 2: Unknown target rejected ---
	fmt.Println("\n--- Test 2: Unknown target rejected ---")
	wc.send(protocol.MsgCombatAction, protocol.CombatActionMsg{TargetID: "not_a_real_creature"})
	if _, ok := wc.waitFor(protocol.MsgError, 2*time.Second); !ok {
		fmt.Println("FAIL: expected error for unknown target")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Unknown target rejected")

	// --- Test 3: Walk to the womp rat lair, then fight and defeat a creature ---
	fmt.Println("\n--- Test 3: Fight and defeat a spawned creature ---")
	// Wait before moving (Phase 1 lesson: dt since world-enter must be
	// large enough that the implied speed of the move is plausible).
	time.Sleep(4 * time.Second)
	wc.send(protocol.MsgMove, protocol.MoveMsg{X: wompRatLairX, Y: 5, Z: wompRatLairZ, Heading: 0})
	time.Sleep(500 * time.Millisecond)

	targetID := "lair_womp_rat_c0"
	killed := false
	// 60 attempts at ~35% hit chance needing ~6 hits to kill a 60HP womp
	// rat gives ample margin over the ~17-attempt expectation -- 20 was
	// too tight and produced a false failure on an unlucky miss streak.
	for attempt := 0; attempt < 60; attempt++ {
		startLen := len(wc.messages)
		wc.send(protocol.MsgCombatAction, protocol.CombatActionMsg{TargetID: targetID})
		time.Sleep(300 * time.Millisecond)

		wc.mu.Lock()
		for _, m := range wc.messages[startLen:] {
			if m.Type == protocol.MsgCreatureDeath {
				killed = true
			}
		}
		wc.mu.Unlock()
		if killed {
			break
		}
	}
	if !killed {
		fmt.Println("FAIL: creature was not defeated within 20 attack attempts")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Creature defeated (creature_death received)")

	skills := getCharacterSkills(token, charID)
	xpPools, _ := skills["xp_pools"].(map[string]interface{})
	combatXP, _ := xpPools["combat"].(float64)
	if combatXP <= 0 {
		fmt.Println("FAIL: expected combat XP reward after kill")
		os.Exit(1)
	}
	fmt.Printf("  [PASS] Combat XP awarded (combat XP pool: %.0f)\n", combatXP)

	// --- Test 4: Posture change ---
	fmt.Println("\n--- Test 4: Posture change ---")
	offset4 := wc.len()
	wc.send(protocol.MsgSetPosture, protocol.SetPostureMsg{Posture: "kneeling"})
	ham, ok := wc.waitForMatch(protocol.MsgHAMUpdate, offset4, 2*time.Second, func(d map[string]interface{}) bool {
		return d["posture"] == "kneeling"
	})
	if !ok {
		fmt.Println("FAIL: posture change not reflected in ham_update")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Posture change reflected in ham_update")
	_ = ham

	// --- Test 5: Take damage to the correct HAM pool, incapacitate ---
	fmt.Println("\n--- Test 5: Take damage from creature retaliation, incapacitate ---")
	wc.send(protocol.MsgSetPosture, protocol.SetPostureMsg{Posture: "standing"})
	// Both lairs are clustered within ~8m of each other and of where Test 3
	// already left the player, so no relocation is needed. Attack each
	// remaining creature several times (not once) to reliably establish
	// aggro despite a ~35-45% hit chance per attempt -- a single roll per
	// creature would leave a meaningful chance some never engage at all.
	for _, id := range []string{"lair_womp_rat_c1", "lair_womp_rat_c2", "lair_worrt_c0", "lair_worrt_c1"} {
		for i := 0; i < 5; i++ {
			wc.send(protocol.MsgCombatAction, protocol.CombatActionMsg{TargetID: id})
			time.Sleep(200 * time.Millisecond)
		}
	}

	fmt.Println("  Waiting for retaliation damage to incapacitate the player (real tick-driven combat, not simulated)...")
	offset5 := wc.len()
	_, incapacitated := wc.waitFor(protocol.MsgIncapacitated, 150*time.Second)
	if !incapacitated {
		fmt.Println("FAIL: player was not incapacitated within 150s of sustained combat")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Player incapacitated")

	lastHAM, _ := wc.waitForMatch(protocol.MsgHAMUpdate, offset5, 2*time.Second, func(d map[string]interface{}) bool {
		inc, _ := d["incapacitated"].(bool)
		return inc
	})
	if hc, ok := lastHAM["health_current"].(float64); !ok || hc != 0 {
		fmt.Println("FAIL: expected health_current == 0 while incapacitated")
		os.Exit(1)
	}
	fmt.Println("  [PASS] health_current is 0 (correct HAM pool took the damage)")

	// --- Test 6: Clone and verify Wounds/BF applied ---
	fmt.Println("\n--- Test 6: Clone and verify state ---")
	offset6 := wc.len()
	wc.send(protocol.MsgClone, nil)
	clonedMsg, ok := wc.waitFor(protocol.MsgCloned, 5*time.Second)
	if !ok {
		fmt.Println("FAIL: cloned message not received")
		os.Exit(1)
	}
	clonedData, _ := json.Marshal(clonedMsg.Data)
	var cloned protocol.ClonedMsg
	json.Unmarshal(clonedData, &cloned)
	if cloned.Planet != "tatooine" {
		fmt.Println("FAIL: expected clone at tatooine spawn")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Cloned message received with correct planet")

	postCloneHAM, ok := wc.waitForMatch(protocol.MsgHAMUpdate, offset6, 2*time.Second, func(d map[string]interface{}) bool {
		inc, _ := d["incapacitated"].(bool)
		return !inc
	})
	if !ok {
		fmt.Println("FAIL: post-clone ham_update not received")
		os.Exit(1)
	}
	hc, _ := postCloneHAM["health_current"].(float64)
	wounds, _ := postCloneHAM["wounds_health"].(float64)
	bf, _ := postCloneHAM["battle_fatigue_pct"].(float64)
	if hc <= 0 {
		fmt.Println("FAIL: expected positive health after clone")
		os.Exit(1)
	}
	if wounds <= 0 {
		fmt.Println("FAIL: expected Wounds gained on clone (GDD 9.6.2)")
		os.Exit(1)
	}
	if bf <= 0 {
		fmt.Println("FAIL: expected Battle Fatigue gained on clone (GDD 9.6.2)")
		os.Exit(1)
	}
	fmt.Printf("  [PASS] Post-clone state: health=%.0f, wounds_health=%.0f, battle_fatigue=%.1f%%\n", hc, wounds, bf)

	fmt.Println("\n=== ALL PHASE 3 TESTS PASSED ===")
}

// --- HTTP helpers (mirrors cmd/phase2test/main.go's conventions) ---

func waitForServer() bool {
	for i := 0; i < 20; i++ {
		resp, err := http.Get(serverURL + "/health")
		if err == nil {
			resp.Body.Close()
			return true
		}
		time.Sleep(500 * time.Millisecond)
	}
	return false
}

func registerAccount() (string, string) {
	username := fmt.Sprintf("p3test_%d", time.Now().UnixNano()%1000000)
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
	if loginResult.Token == "" {
		log.Fatal("login did not return a token")
	}
	return loginResult.Token, loginResult.AccountID
}

func createCharacter(token, speciesID, name string) string {
	body, _ := json.Marshal(map[string]string{"species": speciesID, "name": name})
	req, _ := http.NewRequest("POST", serverURL+"/api/characters", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("create character failed:", err)
	}
	respBody, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	var result map[string]interface{}
	json.Unmarshal(respBody, &result)
	id, _ := result["character_id"].(string)
	if id == "" {
		log.Fatalf("create character did not return a character_id: %s", string(respBody))
	}
	return id
}

func getTrainers() []map[string]interface{} {
	resp, err := http.Get(serverURL + "/api/trainers")
	if err != nil {
		log.Fatal("get trainers failed:", err)
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	var result []map[string]interface{}
	json.Unmarshal(body, &result)
	return result
}

func getCharacterSkills(token, charID string) map[string]interface{} {
	req, _ := http.NewRequest("GET", serverURL+"/api/characters/"+charID+"/skills", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("get skills failed:", err)
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	return result
}

func trainSkill(token, charID, trainerID, skillBoxID string) map[string]interface{} {
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
	if resp.StatusCode != 201 {
		log.Fatalf("train failed (HTTP %d): %s", resp.StatusCode, string(respBody))
	}
	var result map[string]interface{}
	json.Unmarshal(respBody, &result)
	return result
}
