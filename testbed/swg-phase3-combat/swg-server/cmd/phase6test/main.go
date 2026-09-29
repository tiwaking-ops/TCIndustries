// Phase 6 integration test: wounds → medic heal, perform/watch → BF heal,
// tips, buffs (+relog persistence), gated revive, stim use.
//
// Exit criteria under test: "A combat player can accumulate Battle Fatigue and
// Wounds through Phase 3's combat loop and have them healed by a second player
// performing Section 23/24's mechanics."
//
// Ordering is exposure-budgeted: the revive-weakened fighter retreats out of
// aggro before any slow step, so retaliation cannot flake the assertions.
// Total runtime ~10 minutes.
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

func asMap(m protocol.WSMessage) map[string]interface{} {
	data, _ := json.Marshal(m.Data)
	var out map[string]interface{}
	json.Unmarshal(data, &out)
	return out
}

// Per-client tracked positions (dead reckoning; all moves via moveTo).
var ax, az = 20.0, 0.0
var bx, bz = 20.0, 0.0
var cx, cz = 20.0, 0.0

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
		before := countCorrections(wc)
		wc.send(protocol.MsgMove, protocol.MoveMsg{X: nx, Y: 5, Z: nz, Heading: 0})
		time.Sleep(600 * time.Millisecond)
		if countCorrections(wc) == before {
			cx, cz = nx, nz
			*px, *pz = cx, cz
		}
	}
	log.Fatalf("moveTo(%.0f, %.0f) did not converge", tx, tz)
}

func countCorrections(wc *wsClient) int {
	n := 0
	wc.mu.Lock()
	defer wc.mu.Unlock()
	for _, m := range wc.messages {
		if m.Type == protocol.MsgPositionCorrection {
			n++
		}
	}
	return n
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

func main() {
	fmt.Println("=== Phase 6 Integration Test (testbed fork, generic content) ===")
	fmt.Println("Wounds → heal, perform/watch → BF heal, tips, buffs, gated revive, stim")
	fmt.Println()

	if !waitForServer() {
		fmt.Println("FAIL: Server not reachable")
		os.Exit(1)
	}
	fmt.Println("Server is reachable.")

	// A fighter, B service (Medic + Entertainer Novice), C ungated stranger.
	_, tokenA := registerAccount()
	charA := createCharacter(tokenA, "human", "FighterChar")
	_, tokenB := registerAccount()
	charB := createCharacter(tokenB, "human", "ServiceChar")
	_, tokenC := registerAccount()
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

	// --- Test 1: fighter takes wounds + BF through the combat loop ---
	fmt.Println("\n--- Test 1: Wounds + BF via combat ---")
	moveTo(wcA, 12, 0, &ax, &az)
	fmt.Println("  Waiting for retaliation to incapacitate the fighter (real ticks)...")
	if _, ok := wcA.waitFor(protocol.MsgIncapacitated, 300*time.Second); !ok {
		fmt.Println("  [FAIL] fighter never incapacitated")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Fighter incapacitated by creature retaliation")

	// --- Test 2: revive gating (stranger rejected, medic accepted) ---
	fmt.Println("\n--- Test 2: Gated revive ---")
	moveTo(wcC, 12, 0, &cx, &cz)
	wcC.clearMessages()
	wcC.send(protocol.MsgRevive, protocol.ReviveMsg{TargetID: charA})
	if _, ok := wcC.waitFor(protocol.MsgError, 5*time.Second); !ok {
		fmt.Println("  [FAIL] ungated revive was not rejected")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Stranger revive rejected (requires Novice Medic)")
	moveTo(wcC, 30, 0, &cx, &cz) // stranger leaves; no further role
	moveTo(wcB, 12, 0, &bx, &bz)
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
	// Retreat the weak fighter out of aggro before slow steps.
	moveTo(wcA, 30, 0, &ax, &az)
	moveTo(wcB, 30, 0, &bx, &bz)

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

	// --- Test 4: perform / watch / BF heal + Mind drain ---
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
		if h, ok := lastHAM(wcB); ok {
			if mind, _ := h["mind_current"].(float64); mind < 1000 {
				mindDrained = true
			}
		}
	}
	if !bfDone {
		fmt.Println("  [FAIL] BF never reached 0 under performance")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Battle Fatigue healed to 0 by watching")
	if !mindDrained {
		fmt.Println("  [FAIL] performer Mind never drained")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Performer Mind drained by performing")

	// --- Test 5: tips ---
	fmt.Println("\n--- Test 5: Tips ---")
	wA0, wB0 := wallet(tokenA, charA), wallet(tokenB, charB)
	wcA.clearMessages()
	wcA.send(protocol.MsgTip, protocol.TipMsg{TargetID: charB, Amount: 100})
	if _, ok := wcA.waitFor(protocol.MsgTipReceipt, 5*time.Second); !ok {
		fmt.Println("  [FAIL] no tip receipt")
		os.Exit(1)
	}
	if w := wallet(tokenA, charA); w != wA0-100 {
		fmt.Printf("  [FAIL] tipper wallet wrong: %d\n", w)
		os.Exit(1)
	}
	if w := wallet(tokenB, charB); w != wB0+100 {
		fmt.Printf("  [FAIL] performer wallet wrong: %d\n", w)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Tip moved 100 credits tipper → performer")
	if xp := getXP(tokenB, charB, "entertaining"); xp < 25 {
		fmt.Printf("  [FAIL] no entertaining XP for tips (%d)\n", xp)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Entertaining XP awarded for tips")
	// Stop the performance (tip required active watching — done above first).
	wcA.send(protocol.MsgWatchStop, protocol.WatchMsg{})
	wcB.send(protocol.MsgPerformStop, protocol.PerformMsg{})
	time.Sleep(500 * time.Millisecond)

	// --- Test 6: buffs (+relog persistence) ---
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
	// Effective max = min(1000-wounds, 1000) + buff. Wounds are `left` from T3.
	wantMax := 1000 - left + amt
	if w := waitHAMMax(wcA, 10*time.Second); w != wantMax {
		fmt.Printf("  [FAIL] buffed max wrong: got %v want %d\n", w, wantMax)
		os.Exit(1)
	}
	fmt.Printf("  [PASS] Buff applied (+%d); effective Health max = %d\n", amt, wantMax)
	// Duplicate same-pool buff must be rejected.
	wcB.send(protocol.MsgApplyBuff, protocol.BuffMsg{TargetID: charA, Pool: "health"})
	if _, ok := wcB.waitFor(protocol.MsgError, 5*time.Second); !ok {
		fmt.Println("  [FAIL] duplicate buff was not rejected")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Duplicate same-pool buff rejected (no stacking)")
	// Relog: buff persists wall-clock across logout/login.
	wcA.conn.Close()
	time.Sleep(1000 * time.Millisecond)
	wcA = connectWS(tokenA)
	wcA.send(protocol.MsgEnterWorld, protocol.EnterWorldMsg{CharacterID: charA})
	if _, ok := wcA.waitFor(protocol.MsgWorldEnter, 5*time.Second); !ok {
		fmt.Println("  [FAIL] relog failed")
		os.Exit(1)
	}
	ax, az = 30, 0 // DB persisted (30,0) on disconnect
	buffs := getBuffs(tokenA, charA)
	if len(buffs) == 0 {
		fmt.Println("  [FAIL] buff gone after relog")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Buff persists across logout/login")

	// --- Test 7: stim craft + use ---
	fmt.Println("\n--- Test 7: Stim pack ---")
	sampleQuota(wcB, tokenB, charB, "organic", "cultured_organic", 6, &bx, &bz)
	sampleQuota(wcB, tokenB, charB, "chemical", "industrial_chemical", 4, &bx, &bz)
	moveTo(wcB, 30, 0, &bx, &bz)
	stimID := craftItem(tokenB, charB, "stim_pack",
		map[string]string{"bio": bioSrc(tokenB, charB), "binding": chemSrc(tokenB, charB)},
		"potency", 3)
	moveTo(wcA, 30, 0, &ax, &az)
	wcB.clearMessages()
	wcA.clearMessages()
	wcB.send(protocol.MsgUseStim, protocol.StimMsg{ItemID: stimID, TargetID: charA})
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
	if xp := getXP(tokenB, charB, "medical"); xp < 200 {
		fmt.Printf("  [FAIL] medical XP too low (%d)\n", xp)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Service loop complete (heal + revive + buff + stim XP)")

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

// --- helpers ---

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

func authed(method, path, token string, body interface{}) map[string]interface{} {
	var req *http.Request
	if body == nil {
		req, _ = http.NewRequest(method, serverURL+path, nil)
	} else {
		b, _ := json.Marshal(body)
		req, _ = http.NewRequest(method, serverURL+path, bytes.NewReader(b))
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("API %s failed: %v", path, err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		log.Fatalf("API %s HTTP %d: %s", path, resp.StatusCode, string(respBody))
	}
	var out map[string]interface{}
	json.Unmarshal(respBody, &out)
	return out
}

func wallet(token, charID string) int {
	out := authed("GET", "/api/characters/"+charID, token, nil)
	if v, ok := out["credits"].(float64); ok {
		return int(v)
	}
	log.Fatalf("wallet unreadable: %v", out)
	return 0
}

func getXP(token, charID, pool string) int {
	out := authed("GET", "/api/characters/"+charID+"/skills", token, nil)
	pools, _ := out["xp_pools"].(map[string]interface{})
	v, _ := pools[pool].(float64)
	return int(v)
}

func getBuffs(token, charID string) []interface{} {
	out := authed("GET", "/api/services/buffs?character_id="+charID, token, nil)
	list, _ := out["buffs"].([]interface{})
	return list
}

func doSurvey(wc *wsClient, tool string) (map[string]interface{}, bool) {
	wc.clearMessages()
	wc.send(protocol.MsgSurvey, protocol.SurveyMsg{Tool: tool})
	m, ok := wc.waitFor(protocol.MsgSurveyResult, 5*time.Second)
	if !ok {
		return nil, false
	}
	return asMap(m), true
}

func doSample(wc *wsClient) (map[string]interface{}, bool) {
	wc.send(protocol.MsgSample, map[string]interface{}{})
	m, ok := wc.waitFor(protocol.MsgSampleResult, 5*time.Second)
	if !ok {
		return nil, false
	}
	return asMap(m), true
}

func stackTotal(token, charID string, types []string) int {
	out := authed("GET", "/api/craft/resources?character_id="+charID, token, nil)
	total := 0
	want := map[string]bool{}
	for _, t := range types {
		want[t] = true
	}
	if list, ok := out["stacks"].([]interface{}); ok {
		for _, s := range list {
			sm, _ := s.(map[string]interface{})
			rt, _ := sm["ResourceType"].(string)
			if want[rt] {
				if q, ok := sm["Quantity"].(float64); ok {
					total += int(q)
				}
			}
		}
	}
	return total
}

// guaranteeCoords returns the fixed fast-cycle guarantee spawn coordinates
// for a resource type (test scaffolding documented in resource_db.go).
func guaranteeCoords(wantType string) (float64, float64) {
	switch wantType {
	case "cultured_organic":
		return 20, 25
	case "industrial_chemical":
		return 45, 15
	case "ferric_metal":
		return 30, 0
	case "structural_polymer":
		return 40, 0
	default:
		return 20, 0
	}
}

func sampleQuota(wc *wsClient, token, charID, tool, wantType string, qty int, px, pz *float64) {
	// Fill a quota of one explicit resource type: survey until the tool reports
	// that type (random geography may surface siblings first), walk to it, and
	// sample, re-surveying past despawns and recentering past empty radii.
	for outer := 0; outer < 12; outer++ {
		sr, ok := doSurvey(wc, tool)
		if !ok {
			fmt.Printf("  quota(%s): survey empty, recentering\n", wantType)
			moveTo(wc, 20, 0, px, pz)
			time.Sleep(11 * time.Second) // survey cooldown
			continue
		}
		found := sr["resource_type"].(string)
		if found != wantType {
			// Wrong sibling type is nearest (family search returns nearest of
			// either): walk to this want-type's guarantee coords instead. The
			// survey mechanism itself is proven by the successful searches;
			// guarantee coordinates are fixed test scaffolding.
			gx, gz := guaranteeCoords(wantType)
			fmt.Printf("  quota(%s): found %s instead, walking to guarantee (%.0f,%.0f)\n",
				wantType, found, gx, gz)
			moveTo(wc, gx, gz, px, pz)
			time.Sleep(11 * time.Second)
			continue
		}
		wp := sr["waypoint"].(map[string]interface{})
		moveTo(wc, wp["x"].(float64), wp["z"].(float64), px, pz)
		for i := 0; i < 60; i++ {
			if stackTotal(token, charID, []string{wantType}) >= qty {
				fmt.Printf("  quota(%s): met (%d)\n", wantType, qty)
				return
			}
			if _, ok := doSample(wc); !ok {
				break // despawned mid-quota: outer loop re-surveys
			}
			time.Sleep(300 * time.Millisecond)
		}
		if stackTotal(token, charID, []string{wantType}) >= qty {
			fmt.Printf("  quota(%s): met (%d)\n", wantType, qty)
			return
		}
		time.Sleep(11 * time.Second)
	}
	log.Fatalf("quota(%s): unmet after search", wantType)
}

func richestStack(token, charID string, types []string) string {
	out := authed("GET", "/api/craft/resources?character_id="+charID, token, nil)
	best, bestQ := "", -1
	want := map[string]bool{}
	for _, t := range types {
		want[t] = true
	}
	if list, ok := out["stacks"].([]interface{}); ok {
		for _, s := range list {
			sm, _ := s.(map[string]interface{})
			rt, _ := sm["ResourceType"].(string)
			if !want[rt] {
				continue
			}
			q := int(sm["Quantity"].(float64))
			if q > bestQ {
				best, bestQ = sm["SpawnID"].(string), q
			}
		}
	}
	if best == "" {
		log.Fatalf("no stack of types %v", types)
	}
	return best
}

func bioSrc(token, charID string) string {
	return richestStack(token, charID, []string{"cultured_organic", "fibrous_flora"})
}

func chemSrc(token, charID string) string {
	return richestStack(token, charID, []string{"industrial_chemical"})
}

func craftItem(token, charID, schemID string, slots map[string]string, prop string, points int) string {
	start := authed("POST", "/api/craft/sessions/start", token,
		map[string]string{"character_id": charID, "schematic_id": schemID})
	sid := start["session_id"].(string)
	for slot, spawn := range slots {
		authed("POST", "/api/craft/sessions/assign", token, map[string]string{
			"character_id": charID, "session_id": sid, "slot_id": slot, "spawn_id": spawn,
		})
	}
	authed("POST", "/api/craft/sessions/assemble", token,
		map[string]string{"character_id": charID, "session_id": sid})
	authed("POST", "/api/craft/sessions/experiment", token, map[string]interface{}{
		"character_id": charID, "session_id": sid, "property_id": prop, "points": points,
	})
	done := authed("POST", "/api/craft/sessions/finalize", token, map[string]string{
		"character_id": charID, "session_id": sid, "name": "Test Item",
	})
	id, _ := done["item_id"].(string)
	return id
}
