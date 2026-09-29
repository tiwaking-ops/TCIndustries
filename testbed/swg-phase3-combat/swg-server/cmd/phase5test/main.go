// Phase 5 integration test: ledger, houses, vendors, bazaar search, atomic
// purchase (two clients), maintenance lapse, telemetry.
//
// Exit criteria under test: "A player can place a house, place a vendor, list a
// crafted item, and have a second player find and buy it via Bazaar search —
// with correct atomic credit/item transfer."
//
// Requires TESTBED_FAST_CYCLE=1 on the server. Total runtime ~9 minutes.
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

// syncDBPos forces one persisted-position update at the current spot: the
// server persists every 5 s while moving, so a 6 s settle plus a 1 m nudge
// guarantees the DB position matches arrival (HTTP handlers read DB position).
func syncDBPos(wc *wsClient, px, pz *float64) {
	time.Sleep(6 * time.Second)
	before := countCorrections(wc)
	wc.send(protocol.MsgMove, protocol.MoveMsg{X: *px + 1, Y: 5, Z: *pz, Heading: 0})
	time.Sleep(1500 * time.Millisecond)
	if countCorrections(wc) == before {
		*px = *px + 1
	}
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

func main() {
	fmt.Println("=== Phase 5 Integration Test (testbed fork, generic content) ===")
	fmt.Println("Ledger, houses, vendors, bazaar search, atomic purchase, lapse, telemetry")
	fmt.Println()

	if !waitForServer() {
		fmt.Println("FAIL: Server not reachable")
		os.Exit(1)
	}
	fmt.Println("Server is reachable.")

	// Setup A (seller/crafter) and B (buyer).
	_, tokenA := registerAccount()
	charA := createCharacter(tokenA, "human", "SellerChar")
	_, tokenB := registerAccount()
	charB := createCharacter(tokenB, "human", "BuyerChar")
	trainSkill(tokenA, charA, "trainer_artisan_z1", "artisan_novice")
	trainSkill(tokenA, charA, "trainer_marksman_z1", "marksman_novice")
	earnXP(tokenA, charA, "combat", 10000)
	trainSkill(tokenA, charA, "trainer_marksman_z1", "marksman_ranged_accuracy_i") // 50cr sink
	fmt.Println("Setup: seller + buyer; seller holds Artisan/Marksman Novice + Tier I")

	wcA := connectWS(tokenA)
	wcB := connectWS(tokenB)
	wcA.send(protocol.MsgEnterWorld, protocol.EnterWorldMsg{CharacterID: charA})
	wcB.send(protocol.MsgEnterWorld, protocol.EnterWorldMsg{CharacterID: charB})
	if _, ok := wcA.waitFor(protocol.MsgWorldEnter, 5*time.Second); !ok {
		fmt.Println("  [FAIL] A did not enter world")
		os.Exit(1)
	}
	if _, ok := wcB.waitFor(protocol.MsgWorldEnter, 5*time.Second); !ok {
		fmt.Println("  [FAIL] B did not enter world")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Both clients entered world")

	// --- Test 1: ledger records the training sink ---
	fmt.Println("\n--- Test 1: Ledger ---")
	entries := getLedger(tokenA, charA)
	foundSink := false
	for _, e := range entries {
		em, _ := e.(map[string]interface{})
		if em["Category"] == "training_cost" && em["Flow"] == "sink" {
			foundSink = true
		}
	}
	if !foundSink {
		fmt.Println("  [FAIL] no training_cost sink in ledger")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Ledger records the training-cost sink")

	// --- Test 2: drop faucet on kill ---
	fmt.Println("\n--- Test 2: Creature-drop faucet ---")
	trainSkill(tokenB, charB, "trainer_brawler_z1", "brawler_novice")
	wB0drop := wallet(tokenB, charB)
	moveTo(wcB, 14, 0, &bx, &bz)
	time.Sleep(500 * time.Millisecond)
	killIDs := entryCreatures(wcB)
	if len(killIDs) == 0 {
		fmt.Println("  [FAIL] no creatures visible for drop test")
		os.Exit(1)
	}
	killedDrop := ""
	deadline := time.Now().Add(150 * time.Second)
	wcB.clearMessages()
	for time.Now().Before(deadline) && killedDrop == "" {
		wcB.send(protocol.MsgCombatAction, protocol.CombatActionMsg{TargetID: killIDs[0]})
		time.Sleep(800 * time.Millisecond)
		incapped := false
		wcB.mu.Lock()
		for _, m := range wcB.messages {
			switch m.Type {
			case protocol.MsgCreatureDeath:
				if id, _ := asMap(m)["instance_id"].(string); id == killIDs[0] {
					killedDrop = id
				}
			case protocol.MsgIncapacitated:
				incapped = true
			case protocol.MsgError:
				if em, _ := asMap(m)["message"].(string); em == "incapacitated, cannot act" {
					incapped = true
				}
			}
		}
		wcB.mu.Unlock()
		if killedDrop == "" && incapped {
			fmt.Println("  (incapacitated mid-hunt — cloning and resuming)")
			wcB.clearMessages()
			wcB.send(protocol.MsgClone, map[string]interface{}{})
			if _, ok := wcB.waitFor(protocol.MsgCloned, 5*time.Second); !ok {
				break
			}
			bx, bz = 20, 0
			moveTo(wcB, 14, 0, &bx, &bz)
			wcB.clearMessages()
		}
	}
	if killedDrop == "" {
		fmt.Println("  [FAIL] drop-test kill failed")
		os.Exit(1)
	}
	if w := wallet(tokenB, charB); w <= wB0drop {
		fmt.Printf("  [FAIL] wallet unchanged after kill (%d)\n", w)
		os.Exit(1)
	} else {
		fmt.Printf("  [PASS] Kill paid credits (%d → %d)\n", wB0drop, w)
	}
	foundFaucet := false
	for _, e := range getLedger(tokenB, charB) {
		em, _ := e.(map[string]interface{})
		if em["Category"] == "creature_drop" && em["Flow"] == "faucet" {
			foundFaucet = true
		}
	}
	if !foundFaucet {
		fmt.Println("  [FAIL] no creature_drop faucet in ledger")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Ledger records the creature-drop faucet")
	moveTo(wcB, 20, 0, &bx, &bz) // back out of aggro for the trade phase

	// --- Test 3: gather quotas (house 15M+8P, vendor 15M+8P, 2× sidearm 10M+5P) ---
	fmt.Println("\n--- Test 3: Stockpile resources ---")
	sampleQuota(wcA, tokenA, charA, "mineral", 50, &ax, &az)
	sampleQuota(wcA, tokenA, charA, "chemical", 26, &ax, &az)
	fmt.Println("  [PASS] Resource quotas sampled (metal + polymer)")

	// --- Test 4: house + vendor placement ---
	fmt.Println("\n--- Test 4: House + vendor placement ---")
	houseDeed := craftItem(tokenA, charA, "structure_deed",
		map[string]string{"frame": metalSrc(tokenA, charA), "fittings": polySrc(tokenA, charA)},
		"capacity", 3)
	moveTo(wcA, 45, 0, &ax, &az)
	wcA.clearMessages()
	wcA.send(protocol.MsgPlaceHouse, protocol.PlaceStructureMsg{X: 45, Z: 0, Tier: "small", DeedItemID: houseDeed})
	phMsg, ok := wcA.waitFor(protocol.MsgStructurePlaced, 5*time.Second)
	if !ok {
		fmt.Println("  [FAIL] house not placed")
		os.Exit(1)
	}
	houseID := asMap(phMsg)["structure_id"].(string)
	fmt.Printf("  [PASS] House placed (%s)\n", houseID)
	vendorDeed := craftItem(tokenA, charA, "vendor_deed",
		map[string]string{"frame": metalSrc(tokenA, charA), "fittings": polySrc(tokenA, charA)},
		"presentation", 3)
	moveTo(wcA, 70, 0, &ax, &az)
	wcA.clearMessages()
	wcA.send(protocol.MsgPlaceVendor, protocol.PlaceStructureMsg{X: 70, Z: 0, DeedItemID: vendorDeed})
	pvMsg, ok := wcA.waitFor(protocol.MsgStructurePlaced, 5*time.Second)
	if !ok {
		fmt.Println("  [FAIL] vendor not placed")
		os.Exit(1)
	}
	vendorID := asMap(pvMsg)["structure_id"].(string)
	fmt.Printf("  [PASS] Vendor placed (%s)\n", vendorID)
	// Unfunded vendor lapses closed on the first tick: stocking must fail.
	fmt.Println("  Waiting one interval (vendor starts unfunded → closed)...")
	time.Sleep(25 * time.Second)
	if _, ok := stockItem(tokenA, charA, vendorID, "bogus", 100); ok {
		fmt.Println("  [FAIL] stocking a closed vendor succeeded")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Closed vendor refuses stocking")
	fundStructure(tokenA, charA, vendorID, 1000)
	time.Sleep(25 * time.Second)
	sidearm := craftItem(tokenA, charA, "basic_sidearm",
		map[string]string{"frame": metalSrc(tokenA, charA), "grip": polySrc(tokenA, charA)},
		"damage", 3)
	listing := stockItemMust(tokenA, charA, vendorID, sidearm, 500)
	fmt.Printf("  [PASS] Vendor funded, opened, sidearm listed (%s @ 500)\n", listing)

	// --- Test 5: bazaar search finds it ---
	fmt.Println("\n--- Test 5: Bazaar search ---")
	// Search requires terminal proximity, checked against the DB-persisted
	// position (5 s persist cadence): walk to the terminal, then nudge once
	// after 6 s so the persisted position catches up to arrival.
	moveTo(wcB, 22, 0, &bx, &bz)
	syncDBPos(wcB, &bx, &bz)
	results := bazaarSearch(tokenB, charB, "slot=weapon&max_price=600")
	found := false
	for _, r := range results {
		rm, _ := r.(map[string]interface{})
		if rm["listing_id"] == listing {
			found = true
			fmt.Printf("  Found: %s @ %v by %v at %v (%.0f, %.0f)\n",
				rm["schematic"], rm["price"], rm["seller"], rm["zone"], rm["vendor_x"], rm["vendor_z"])
		}
	}
	if !found {
		fmt.Println("  [FAIL] listing absent from bazaar search")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Bazaar search finds the listing with location")

	// --- Test 6: atomic purchase + double-buy race ---
	fmt.Println("\n--- Test 6: Atomic purchase ---")
	moveTo(wcB, 68, 0, &bx, &bz) // within purchase range of the vendor at (70,0)
	wB0 := wallet(tokenB, charB)
	wcB.clearMessages()
	wcB.send(protocol.MsgPurchase, protocol.PurchaseMsg{ListingID: listing})
	prMsg, ok := wcB.waitFor(protocol.MsgPurchaseReceipt, 5*time.Second)
	if !ok {
		fmt.Println("  [FAIL] no purchase receipt")
		os.Exit(1)
	}
	if asMap(prMsg)["item_id"] == "" {
		fmt.Println("  [FAIL] receipt missing item")
		os.Exit(1)
	}
	if w := wallet(tokenB, charB); w != wB0-500 {
		fmt.Printf("  [FAIL] buyer wallet wrong: %d (expected %d)\n", w, wB0-500)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Purchase atomic: buyer −500, receipt issued")
	// Till collect by owner in person.
	moveTo(wcA, 70, 0, &ax, &az)
	wcA.clearMessages()
	wcA.send(protocol.MsgCollectTill, protocol.TillMsg{VendorID: vendorID})
	tillMsg, ok := wcA.waitFor(protocol.MsgTillCollected, 5*time.Second)
	if !ok || int(asMap(tillMsg)["amount"].(float64)) != 500 {
		fmt.Println("  [FAIL] till collection wrong")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Till collected in person (+500)")
	// Race: second sidearm listed once, bought twice at once → exactly one wins.
	sidearm2 := craftItem(tokenA, charA, "basic_sidearm",
		map[string]string{"frame": metalSrc(tokenA, charA), "grip": polySrc(tokenA, charA)},
		"damage", 3)
	listing2 := stockItemMust(tokenA, charA, vendorID, sidearm2, 600)
	wB1 := wallet(tokenB, charB)
	moveTo(wcB, 70, 0, &bx, &bz)
	wcB.clearMessages()
	wcB.send(protocol.MsgPurchase, protocol.PurchaseMsg{ListingID: listing2})
	wcB.send(protocol.MsgPurchase, protocol.PurchaseMsg{ListingID: listing2})
	time.Sleep(2 * time.Second)
	receipts, errors := 0, 0
	wcB.mu.Lock()
	for _, m := range wcB.messages {
		if m.Type == protocol.MsgPurchaseReceipt {
			receipts++
		}
		if m.Type == protocol.MsgError {
			errors++
		}
	}
	wcB.mu.Unlock()
	if receipts != 1 || errors < 1 {
		fmt.Printf("  [FAIL] race resolved wrong: receipts=%d errors=%d\n", receipts, errors)
		os.Exit(1)
	}
	if w := wallet(tokenB, charB); w != wB1-600 {
		fmt.Printf("  [FAIL] double-charge: wallet %d (expected %d)\n", w, wB1-600)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Double-purchase race: exactly one winner, loser never charged")

	// --- Test 7: ledger transfer pair + market record ---
	fmt.Println("\n--- Test 7: Ledger transfer pair ---")
	entriesB := getLedger(tokenB, charB)
	entriesA := getLedger(tokenA, charA)
	hasDebit, hasCredit := false, false
	for _, e := range entriesB {
		em, _ := e.(map[string]interface{})
		if em["Category"] == "vendor_sale" && em["Flow"] == "transfer" {
			if amt, _ := em["Amount"].(float64); amt == -500 {
				hasDebit = true
			}
		}
	}
	for _, e := range entriesA {
		em, _ := e.(map[string]interface{})
		if em["Category"] == "vendor_sale" && em["Flow"] == "transfer" {
			if amt, _ := em["Amount"].(float64); amt == 500 {
				hasCredit = true
			}
		}
	}
	if !hasDebit || !hasCredit {
		fmt.Println("  [FAIL] transfer pair incomplete in ledger")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Ledger holds the buyer-debit/seller-credit transfer pair")

	// --- Test 8: house lapse → condemned (fast-cycle grace) ---
	fmt.Println("\n--- Test 8: House lapse → condemned ---")
	fmt.Println("  House pool has been empty since placement; waiting out grace...")
	time.Sleep(200 * time.Second)
	st := structureStatus(tokenA, charA, houseID)
	if st != "condemned" && st != "destroyed" {
		fmt.Printf("  [FAIL] house status %q (expected condemned/destroyed)\n", st)
		os.Exit(1)
	}
	fmt.Printf("  [PASS] Unfunded house reached %q\n", st)

	// --- Test 9: telemetry snapshot ---
	fmt.Println("\n--- Test 9: Telemetry snapshot ---")
	snap := getSnapshot(tokenA)
	if snap["Circulation"] == nil {
		fmt.Println("  [FAIL] no snapshot")
		os.Exit(1)
	}
	fmt.Printf("  [PASS] Snapshot: circulation=%v faucet_30d=%v sink_30d=%v net=%v%%\n",
		snap["Circulation"], snap["Faucet30d"], snap["Sink30d"], snap["NetPct"])

	fmt.Println("\n=== ALL PHASE 5 TESTS PASSED ===")
	wcA.conn.Close()
	wcB.conn.Close()
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

func earnXP(token, charID, xpType string, amount int) {
	body, _ := json.Marshal(map[string]interface{}{"xp_type": xpType, "amount": amount})
	req, _ := http.NewRequest("POST", serverURL+"/api/characters/"+charID+"/earn-xp", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("earn XP failed:", err)
	}
	resp.Body.Close()
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

func econAuthed(method, path, token string, body interface{}) map[string]interface{} {
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
		log.Fatalf("economy API %s failed: %v", path, err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		log.Fatalf("economy API %s HTTP %d: %s", path, resp.StatusCode, string(respBody))
	}
	var out map[string]interface{}
	json.Unmarshal(respBody, &out)
	return out
}

func wallet(token, charID string) int {
	req, _ := http.NewRequest("GET", serverURL+"/api/characters/"+charID, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("get character failed:", err)
	}
	defer resp.Body.Close()
	var out map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&out)
	if v, ok := out["credits"].(float64); ok {
		return int(v)
	}
	log.Fatalf("wallet unreadable: %v", out)
	return 0
}

func craftAuthed(method, path, token string, body interface{}) map[string]interface{} {
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
		log.Fatalf("craft API %s failed: %v", path, err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		log.Fatalf("craft API %s HTTP %d: %s", path, resp.StatusCode, string(respBody))
	}
	var out map[string]interface{}
	json.Unmarshal(respBody, &out)
	return out
}

func doSurvey(wc *wsClient, tool string) map[string]interface{} {
	wc.clearMessages()
	wc.send(protocol.MsgSurvey, protocol.SurveyMsg{Tool: tool})
	m, ok := wc.waitFor(protocol.MsgSurveyResult, 5*time.Second)
	if !ok {
		log.Fatal("no survey_result")
	}
	return asMap(m)
}

func doSample(wc *wsClient) map[string]interface{} {
	wc.send(protocol.MsgSample, map[string]interface{}{})
	m, ok := wc.waitFor(protocol.MsgSampleResult, 5*time.Second)
	if !ok {
		log.Fatal("no sample_result")
	}
	return asMap(m)
}

func stackTotal(token, charID string, types []string) int {
	out := craftAuthed("GET", "/api/craft/resources?character_id="+charID, token, nil)
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

// sampleQuota walks to the surveyed waypoint and samples until the type total
// across all stacks reaches qty.
func sampleQuota(wc *wsClient, token, charID, tool string, qty int, px, pz *float64) {
	sr := doSurvey(wc, tool)
	wp := sr["waypoint"].(map[string]interface{})
	moveTo(wc, wp["x"].(float64), wp["z"].(float64), px, pz)
	rtype := sr["resource_type"].(string)
	for i := 0; i < 60 && stackTotal(token, charID, []string{rtype}) < qty; i++ {
		doSample(wc)
		time.Sleep(300 * time.Millisecond)
	}
}

func metalSrc(token, charID string) string {
	return richestStack(token, charID, []string{"ferric_metal", "conductive_alloy"})
}

func polySrc(token, charID string) string {
	return richestStack(token, charID, []string{"structural_polymer", "fibrous_flora"})
}

func richestStack(token, charID string, types []string) string {
	out := craftAuthed("GET", "/api/craft/resources?character_id="+charID, token, nil)
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

func craftItem(token, charID, schemID string, slots map[string]string, prop string, points int) string {
	start := craftAuthed("POST", "/api/craft/sessions/start", token,
		map[string]string{"character_id": charID, "schematic_id": schemID})
	sid := start["session_id"].(string)
	for slot, spawn := range slots {
		craftAuthed("POST", "/api/craft/sessions/assign", token, map[string]string{
			"character_id": charID, "session_id": sid, "slot_id": slot, "spawn_id": spawn,
		})
	}
	craftAuthed("POST", "/api/craft/sessions/assemble", token,
		map[string]string{"character_id": charID, "session_id": sid})
	craftAuthed("POST", "/api/craft/sessions/experiment", token, map[string]interface{}{
		"character_id": charID, "session_id": sid, "property_id": prop, "points": points,
	})
	done := craftAuthed("POST", "/api/craft/sessions/finalize", token, map[string]string{
		"character_id": charID, "session_id": sid, "name": "Test Item",
	})
	id, _ := done["item_id"].(string)
	return id
}

func stockItem(token, charID, vendorID, itemID string, price int) (string, bool) {
	body, _ := json.Marshal(map[string]interface{}{
		"character_id": charID, "vendor_id": vendorID, "item_id": itemID, "price": price,
	})
	req, _ := http.NewRequest("POST", serverURL+"/api/economy/vendor/list", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("stock failed:", err)
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		return "", false
	}
	var out map[string]interface{}
	json.Unmarshal(respBody, &out)
	id, _ := out["listing_id"].(string)
	return id, true
}

func stockItemMust(token, charID, vendorID, itemID string, price int) string {
	id, ok := stockItem(token, charID, vendorID, itemID, price)
	if !ok {
		log.Fatal("stock failed")
	}
	return id
}

func fundStructure(token, charID, stID string, credits int) {
	econAuthed("POST", "/api/economy/structure/fund", token, map[string]interface{}{
		"character_id": charID, "structure_id": stID, "credits": credits,
	})
}

func bazaarSearch(token, charID, query string) []interface{} {
	out := econAuthed("GET", "/api/economy/bazaar/search?character_id="+charID+"&"+query, token, nil)
	list, _ := out["results"].([]interface{})
	return list
}

func getLedger(token, charID string) []interface{} {
	out := econAuthed("GET", "/api/economy/ledger?character_id="+charID, token, nil)
	list, _ := out["entries"].([]interface{})
	return list
}

func structureStatus(token, charID, stID string) string {
	out := econAuthed("GET", "/api/economy/structures?character_id="+charID, token, nil)
	if list, ok := out["structures"].([]interface{}); ok {
		for _, s := range list {
			sm, _ := s.(map[string]interface{})
			if sm["id"] == stID {
				st, _ := sm["status"].(string)
				return st
			}
		}
	}
	return ""
}

func getSnapshot(token string) map[string]interface{} {
	out := econAuthed("GET", "/api/economy/snapshot?x=1", token, nil)
	snap, _ := out["snapshot"].(map[string]interface{})
	return snap
}

// entryCreatures collects creature entity IDs from spawn messages (visibility
// is edge-triggered: snapshot right after world entry).
func entryCreatures(wc *wsClient) []string {
	ids := []string{}
	seen := map[string]bool{}
	wc.mu.Lock()
	defer wc.mu.Unlock()
	for _, m := range wc.messages {
		if m.Type != protocol.MsgEntitySpawn {
			continue
		}
		data, _ := json.Marshal(m.Data)
		var sp protocol.EntitySpawnMsg
		json.Unmarshal(data, &sp)
		if sp.EntityType == "creature" && !seen[sp.EntityID] {
			seen[sp.EntityID] = true
			ids = append(ids, sp.EntityID)
		}
	}
	return ids
}
