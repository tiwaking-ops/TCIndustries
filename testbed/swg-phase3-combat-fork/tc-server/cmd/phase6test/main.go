// Phase 6 integration test: wounds → medic heal, perform/watch → BF heal,
// tips, buffs (+relog persistence), gated revive, stim use.
//
// Exit criteria under test: "A combat player can accumulate Battle Fatigue and
// Wounds through Phase 3's combat loop and have them healed by a second player
// performing Section 23/24's mechanics."
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"tc-server/internal/protocol"
)

var serverURL = envOr("SERVER_URL", "http://localhost:8080")
var wsURL = func() string {
	if v := os.Getenv("SERVER_URL"); v != "" {
		return strings.Replace(v, "http", "ws", 1) + "/ws"
	}
	return "ws://localhost:8080/ws"
}()

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

func asMap(m protocol.WSMessage) map[string]interface{} {
	data, _ := json.Marshal(m.Data)
	var out map[string]interface{}
	json.Unmarshal(data, &out)
	return out
}

func moveTo(wc *wsClient, tx, tz float64, px, pz *float64) {
	cx, cz := *px, *pz
	time.Sleep(1200 * time.Millisecond)
	for i := 0; i < 200; i++ {
		dx, dz := tx-cx, tz-cz
		dist := math.Sqrt(dx*dx + dz*dz)
		if dist < 1.0 {
			*px, *pz = tx, tz
			return
		}
		step := 4.0
		if dist < step {
			step = dist
		}
		nx, nz := cx+dx/dist*step, cz+dz/dist*step
		wc.send(protocol.MsgMove, protocol.MoveMsg{X: nx, Y: 5, Z: nz, Heading: 0})
		time.Sleep(600 * time.Millisecond)
		cx, cz = nx, nz
		*px, *pz = cx, cz
	}
}

func lastHAM(wc *wsClient) (map[string]interface{}, bool) {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	for i := len(wc.messages) - 1; i >= 0; i-- {
		if wc.messages[i].Type == protocol.MsgHAMUpdate {
			return asMap(wc.messages[i]), true
		}
	}
	return nil, false
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	fmt.Println("=== Phase 6 Integration Test (testbed fork, generic content) ===")
	fmt.Println("Wounds → heal, perform/watch → BF heal, tips, buffs, gated revive, stim")
	fmt.Println()

	if !waitForServer() {
		fmt.Println("FAIL: Server not reachable")
		os.Exit(1)
	}
	fmt.Println("Server is reachable.")

	tokenA, _ := registerAccount()
	charA := createCharacter(tokenA, "human", "FighterChar")
	tokenB, _ := registerAccount()
	charB := createCharacter(tokenB, "human", "ServiceChar")
	tokenC, _ := registerAccount()
	charC := createCharacter(tokenC, "human", "StrangerChar")
	trainSkill(tokenA, charA, "trainer_brawler_z1", "brawler_novice")
	trainSkill(tokenB, charB, "trainer_medic_z1", "medic_novice")
	trainSkill(tokenB, charB, "trainer_entertainer_z1", "entertainer_novice")
	fmt.Println("Setup: fighter + service (Medic/Entertainer) + stranger")

	wcA := connectWS(tokenA)
	wcB := connectWS(tokenB)
	wcC := connectWS(tokenC)
	wcA.send(protocol.MsgEnterWorld, protocol.EnterWorldMsg{CharacterID: charA})
	wcB.send(protocol.MsgEnterWorld, protocol.EnterWorldMsg{CharacterID: charB})
	wcC.send(protocol.MsgEnterWorld, protocol.EnterWorldMsg{CharacterID: charC})
	for i, wc := range []*wsClient{wcA, wcB, wcC} {
		if _, ok := wc.waitFor(protocol.MsgWorldEnter, 5*time.Second); !ok {
			fmt.Printf("  [FAIL] client %d did not enter world\n", i)
			os.Exit(1)
		}
	}
	fmt.Println("  [PASS] All three clients entered world")

	ax, az := 3528.0, -4804.0
	bx, bz := 3528.0, -4804.0
	cx, cz := 3528.0, -4804.0

	// --- Test 1: fighter takes wounds + BF via combat ---
	fmt.Println("\n--- Test 1: Wounds + BF via combat ---")
	moveTo(wcA, 3560, -4790, &ax, &az)
	fmt.Println("  Waiting for retaliation to incapacitate the fighter (real ticks)...")
	if _, ok := wcA.waitFor(protocol.MsgIncapacitated, 300*time.Second); !ok {
		fmt.Println("  [FAIL] fighter never incapacitated")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Fighter incapacitated by creature retaliation")

	// --- Test 2: revive gating ---
	fmt.Println("\n--- Test 2: Gated revive ---")
	moveTo(wcC, 3560, -4790, &cx, &cz)
	wcC.clearMessages()
	wcC.send(protocol.MsgRevive, protocol.ReviveMsg{TargetID: charA})
	if _, ok := wcC.waitFor(protocol.MsgError, 5*time.Second); !ok {
		fmt.Println("  [FAIL] ungated revive was not rejected")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Stranger revive rejected (requires Novice Medic)")
	moveTo(wcC, 3510, -4800, &cx, &cz)
	moveTo(wcB, 3560, -4790, &bx, &bz)
	wcB.clearMessages()
	wcA.clearMessages()
	wcB.send(protocol.MsgRevive, protocol.ReviveMsg{TargetID: charA})
	hamMsg, ok := wcA.waitFor(protocol.MsgHAMUpdate, 10*time.Second)
	if !ok {
		fmt.Println("  [FAIL] no ham_update after medic revive")
		os.Exit(1)
	}
	ham := asMap(hamMsg)
	if ham["incapacitated"] == true {
		fmt.Println("  [FAIL] still incapacitated after revive")
		os.Exit(1)
	}
	wounds0 := int(ham["wounds_health"].(float64))
	bf0 := ham["battle_fatigue_pct"].(float64)
	if wounds0 <= 0 || bf0 <= 0 {
		fmt.Printf("  [FAIL] revive yielded no wounds/BF (w=%d bf=%v)\n", wounds0, bf0)
		os.Exit(1)
	}
	fmt.Printf("  [PASS] Medic revive: standing with wounds=%d BF=%.1f\n", wounds0, bf0)
	moveTo(wcA, 3520, -4810, &ax, &az)
	moveTo(wcB, 3520, -4810, &bx, &bz)

	// --- Test 3: wound healing + cooldown ---
	fmt.Println("\n--- Test 3: Wound healing ---")
	wcB.clearMessages()
	wcA.clearMessages()
	wcB.send(protocol.MsgHealWounds, protocol.TargetMsg{TargetID: charA})
	healMsg, ok := wcB.waitFor(protocol.MsgWoundHealed, 5*time.Second)
	if !ok {
		fmt.Println("  [FAIL] no wound_healed receipt")
		os.Exit(1)
	}
	hm := asMap(healMsg)
	healed := int(hm["amount"].(float64))
	left := int(hm["wounds_left"].(float64))
	if healed <= 0 || left != wounds0-healed {
		fmt.Printf("  [FAIL] heal math wrong: healed=%d left=%d (was %d)\n", healed, left, wounds0)
		os.Exit(1)
	}
	fmt.Printf("  [PASS] Wounds %d → %d\n", wounds0, left)
	wcB.send(protocol.MsgHealWounds, protocol.TargetMsg{TargetID: charA})
	if _, ok := wcB.waitFor(protocol.MsgError, 5*time.Second); !ok {
		fmt.Println("  [FAIL] immediate re-heal was not cooled down")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Immediate re-heal rejected (cooldown window)")
	if xp := getXP(tokenB, charB, "medical"); xp < 50 {
		fmt.Printf("  [FAIL] no medical XP for healing (%d)\n", xp)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Medical XP awarded for healing")

	// --- Test 4: perform / watch / BF heal ---
	fmt.Println("\n--- Test 4: Performance → BF heal ---")
	wcB.clearMessages()
	wcA.clearMessages()
	wcB.send(protocol.MsgPerformStart, protocol.PerformMsg{Kind: "music"})
	if _, ok := wcB.waitFor(protocol.MsgPerformanceStarted, 5*time.Second); !ok {
		fmt.Println("  [FAIL] performance did not start")
		os.Exit(1)
	}
	wcA.send(protocol.MsgWatchStart, protocol.WatchMsg{PerformerID: charB})
	if _, ok := wcA.waitFor(protocol.MsgWatchStarted, 5*time.Second); !ok {
		fmt.Println("  [FAIL] watch did not start")
		os.Exit(1)
	}
	fmt.Println("  Performance running, watcher paired; waiting for BF → 0...")
	bfDone, mindDrained := false, false
	deadline := time.Now().Add(120 * time.Second)
	for time.Now().Before(deadline) && (!bfDone || !mindDrained) {
		time.Sleep(3 * time.Second)
		if h, ok := lastHAM(wcA); ok {
			if bf, _ := h["battle_fatigue_pct"].(float64); bf == 0 {
				bfDone = true
			}
		}
		if hm, ok := lastHAM(wcB); ok {
			if mind, _ := hm["mind_current"].(float64); mind < 500 {
				mindDrained = true
			}
		}
	}
	if !bfDone || !mindDrained {
		fmt.Printf("  [FAIL] timeout: bfDone=%v mindDrained=%v\n", bfDone, mindDrained)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Battle Fatigue healed to 0 by watching")
	fmt.Println("  [PASS] Performer Mind drained by performing")

	// --- Test 5: Tips ---
	fmt.Println("\n--- Test 5: Tips ---")
	balBefore := getCredits(tokenB, charB)
	wcA.send(protocol.MsgTip, protocol.TipMsg{TargetID: charB, Amount: 100})
	if _, ok := wcA.waitFor(protocol.MsgTipReceipt, 5*time.Second); !ok {
		fmt.Println("  [FAIL] no tip receipt")
		os.Exit(1)
	}
	if balAfter := getCredits(tokenB, charB); balAfter != balBefore+100 {
		fmt.Printf("  [FAIL] tip credits wrong: %d → %d\n", balBefore, balAfter)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Tip moved 100 credits tipper → performer")
	if xp := getXP(tokenB, charB, "entertainer"); xp < 25 {
		fmt.Printf("  [FAIL] no entertainer XP for tips (%d)\n", xp)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Entertaining XP awarded for tips")

	wcA.send(protocol.MsgWatchStop, protocol.WatchMsg{})
	wcB.send(protocol.MsgPerformStop, protocol.PerformMsg{})
	time.Sleep(500 * time.Millisecond)

	// --- Test 6: Buffs ---
	fmt.Println("\n--- Test 6: Buffs ---")
	wcA.clearMessages()
	wcB.send(protocol.MsgApplyBuff, protocol.BuffMsg{TargetID: charA, Pool: "health"})
	buffMsg, ok := wcB.waitFor(protocol.MsgBuffApplied, 5*time.Second)
	if !ok {
		fmt.Println("  [FAIL] no buff_applied receipt")
		os.Exit(1)
	}
	amt := int(asMap(buffMsg)["amount"].(float64))
	if amt < 500 {
		fmt.Printf("  [FAIL] buff magnitude below GDD band: %d\n", amt)
		os.Exit(1)
	}
	wantMax := 1000 - left + amt
	if w := waitHAMMax(wcA, 10*time.Second); w != wantMax {
		fmt.Printf("  [FAIL] buffed max wrong: got %v want %d\n", w, wantMax)
		os.Exit(1)
	}
	fmt.Printf("  [PASS] Buff applied (+%d); effective Health max = %d\n", amt, wantMax)

	wcB.send(protocol.MsgApplyBuff, protocol.BuffMsg{TargetID: charA, Pool: "health"})
	if _, ok := wcB.waitFor(protocol.MsgError, 5*time.Second); !ok {
		fmt.Println("  [FAIL] duplicate buff was not rejected")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Duplicate same-pool buff rejected (no stacking)")

	// Relog persistence
	wcA.conn.Close()
	time.Sleep(1000 * time.Millisecond)
	wcA = connectWS(tokenA)
	wcA.send(protocol.MsgEnterWorld, protocol.EnterWorldMsg{CharacterID: charA})
	if _, ok := wcA.waitFor(protocol.MsgWorldEnter, 5*time.Second); !ok {
		fmt.Println("  [FAIL] relog failed")
		os.Exit(1)
	}
	buffs := getBuffs(tokenA, charA)
	if len(buffs) == 0 {
		fmt.Println("  [FAIL] buff gone after relog")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Buff persists across logout/login")

	// --- Test 7: Stim pack ---
	fmt.Println("\n--- Test 7: Stim pack ---")
	wcB.clearMessages()
	wcA.clearMessages()
	wcB.send(protocol.MsgUseStim, protocol.StimMsg{ItemID: "stim_pack_1", TargetID: charA})
	stimMsg, ok := wcB.waitFor(protocol.MsgStimUsed, 5*time.Second)
	if !ok {
		fmt.Println("  [FAIL] no stim_used receipt")
		os.Exit(1)
	}
	leftAfter := waitWounds(wcA, 10*time.Second)
	if leftAfter >= left {
		fmt.Printf("  [FAIL] stim did not reduce wounds (%d → %d)\n", left, leftAfter)
		os.Exit(1)
	}
	fmt.Printf("  [PASS] Stim used: wounds %d → %d, charges left %v\n",
		left, leftAfter, asMap(stimMsg)["charges_left"])
	fmt.Println("  [PASS] Service loop complete (heal + revive + buff + stim)")

	fmt.Println("\n=== ALL PHASE 6 TESTS PASSED ===")
	wcA.conn.Close()
	wcB.conn.Close()
	wcC.conn.Close()
}

func waitHAMMax(wc *wsClient, timeout time.Duration) int {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if h, ok := lastHAM(wc); ok {
			if m, ok := h["health_max"].(float64); ok && m > 1000 {
				return int(m)
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
	if h, ok := lastHAM(wc); ok {
		if m, ok := h["health_max"].(float64); ok {
			return int(m)
		}
	}
	return -1
}

func waitWounds(wc *wsClient, timeout time.Duration) int {
	deadline := time.Now().Add(timeout)
	best := -1
	for time.Now().Before(deadline) {
		if h, ok := lastHAM(wc); ok {
			if w, ok := h["wounds_health"].(float64); ok {
				best = int(w)
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
	return best
}

func connectWS(token string) *wsClient {
	conn, _, err := websocket.DefaultDialer.Dial(wsURL+"?token="+token, nil)
	if err != nil {
		log.Fatal("ws dial failed:", err)
	}
	wc := &wsClient{conn: conn}
	go wc.readLoop()
	return wc
}

func waitForServer() bool {
	for i := 0; i < 30; i++ {
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
	username := fmt.Sprintf("p6test_%d", time.Now().UnixNano()%1000000)
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

func getCharacter(token, charID string) map[string]interface{} {
	req, _ := http.NewRequest("GET", serverURL+"/api/characters/"+charID, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("get character failed:", err)
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	return result
}

func getCredits(token, charID string) int {
	ch := getCharacter(token, charID)
	c, _ := ch["credits"].(float64)
	return int(c)
}

func getXP(token, charID, pool string) int {
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
	pools, _ := result["xp_pools"].(map[string]interface{})
	v, _ := pools[pool].(float64)
	return int(v)
}

func getBuffs(token, charID string) []interface{} {
	req, _ := http.NewRequest("GET", serverURL+"/api/services/buffs?character_id="+charID, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("get buffs failed:", err)
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	list, _ := result["buffs"].([]interface{})
	return list
}
