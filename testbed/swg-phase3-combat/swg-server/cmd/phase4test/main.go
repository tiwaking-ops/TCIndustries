// Phase 4 integration test: survey → sample → harvest → craft-with-
// experimentation → equip, plus corpse harvest and spawn lifecycle.
//
// Exit criteria under test: "A player can survey, harvest a resource, craft an
// item from a schematic with experimentation, and equip a crafted weapon/armor
// with correctly-computed stats."
//
// Requires TESTBED_FAST_CYCLE=1 on the server (compressed lifecycles/ticks;
// documented test scaffolding). Total runtime ~8 minutes.
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

func main() {
	fmt.Println("=== Phase 4 Integration Test (testbed fork, generic content) ===")
	fmt.Println("Survey → sample → harvest → craft → equip, corpse harvest, lifecycle")
	fmt.Println()

	if !waitForServer() {
		fmt.Println("FAIL: Server not reachable")
		os.Exit(1)
	}
	fmt.Println("Server is reachable.")

	accountID, token := registerAccount()
	_ = accountID
	charID := createCharacter(token, "human", "CraftTestChar")
	trainSkill(token, charID, "trainer_artisan_z1", "artisan_novice")
	trainSkill(token, charID, "trainer_brawler_z1", "brawler_novice")
	fmt.Println("Setup: character + Novice Artisan + Novice Brawler")

	wc := connectWS(token)
	wc.send(protocol.MsgEnterWorld, protocol.EnterWorldMsg{CharacterID: charID})
	if _, ok := wc.waitFor(protocol.MsgWorldEnter, 5*time.Second); !ok {
		fmt.Println("  [FAIL] did not receive world_enter")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Entered world")
	// Snapshot creature IDs at entry: visibility is edge-triggered, so spawns
	// seen here will NOT re-arrive later while we stay in range.
	time.Sleep(500 * time.Millisecond)
	liveCreatures := visibleCreatures(wc)
	if len(liveCreatures) == 0 {
		fmt.Println("  [FAIL] no creature spawns visible at entry")
		os.Exit(1)
	}
	fmt.Printf("  Tracking %d creature instances from entry\n", len(liveCreatures))

	// --- Test 1: schematic registry ---
	fmt.Println("\n--- Test 1: Schematic registry ---")
	schems := getSchematics()
	for _, want := range []string{"basic_sidearm", "basic_plate", "harvester_deed"} {
		found := false
		for _, s := range schems {
			if sm, ok := s.(map[string]interface{}); ok && sm["id"] == want {
				found = true
			}
		}
		if !found {
			fmt.Printf("  [FAIL] schematic %s absent from registry\n", want)
			os.Exit(1)
		}
	}
	fmt.Println("  [PASS] Schematic registry holds the Phase 4 trio (+ later additions)")

	// --- Test 2: survey finds a spawn + waypoint ---
	fmt.Println("\n--- Test 2: Survey ---")
	wc.clearMessages()
	wc.send(protocol.MsgSurvey, protocol.SurveyMsg{Tool: "mineral"})
	surveyMsg, ok := wc.waitFor(protocol.MsgSurveyResult, 5*time.Second)
	if !ok {
		fmt.Println("  [FAIL] no survey_result")
		os.Exit(1)
	}
	sr := asMap(surveyMsg)
	metalSpawn := sr["spawn_id"].(string)
	wp := sr["waypoint"].(map[string]interface{})
	mwx, mwz := wp["x"].(float64), wp["z"].(float64)
	fmt.Printf("  [PASS] Survey found %s (%s) conc %v at waypoint (%.0f, %.0f)\n",
		metalSpawn, sr["resource_type"], sr["concentration"], mwx, mwz)

	// --- Test 3: move to waypoint + sample into stack ---
	fmt.Println("\n--- Test 3: Sample ---")
	moveTo(wc, mwx, mwz)
	wc.clearMessages()
	wc.send(protocol.MsgSample, map[string]interface{}{})
	sampleMsg, ok := wc.waitFor(protocol.MsgSampleResult, 5*time.Second)
	if !ok {
		fmt.Println("  [FAIL] no sample_result")
		os.Exit(1)
	}
	sm := asMap(sampleMsg)
	units := int(sm["units"].(float64))
	if units < 1 {
		fmt.Println("  [FAIL] zero units sampled")
		os.Exit(1)
	}
	fmt.Printf("  [PASS] Sampled %d units of %s\n", units, sm["resource_type"])
	// Top up metal to deed quota (8) by repeated sampling.
	for stackQty(token, charID, metalSpawn) < 8 {
		wc.send(protocol.MsgSample, map[string]interface{}{})
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Printf("  Metal stack topped to %d units\n", stackQty(token, charID, metalSpawn))

	// --- Test 4: second resource (polymer) via chemical tool ---
	fmt.Println("\n--- Test 4: Second resource via survey ---")
	time.Sleep(11 * time.Second) // survey cooldown (GDD-given 10 s)
	wc.clearMessages()
	wc.send(protocol.MsgSurvey, protocol.SurveyMsg{Tool: "chemical"})
	surveyMsg2, ok := wc.waitFor(protocol.MsgSurveyResult, 5*time.Second)
	if !ok {
		fmt.Println("  [FAIL] no survey_result for chemical tool")
		os.Exit(1)
	}
	sr2 := asMap(surveyMsg2)
	polySpawn := sr2["spawn_id"].(string)
	wp2 := sr2["waypoint"].(map[string]interface{})
	moveTo(wc, wp2["x"].(float64), wp2["z"].(float64))
	for stackQty(token, charID, polySpawn) < 9 {
		wc.send(protocol.MsgSample, map[string]interface{}{})
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Printf("  [PASS] Polymer stack at %d units\n", stackQty(token, charID, polySpawn))

	// --- Test 5: craft harvester deed → place → fund cycle → extract ---
	fmt.Println("\n--- Test 5: Deed, harvester, hopper ---")
	deedID := craftItem(token, charID, "harvester_deed",
		map[string]string{"housing": metalSpawn, "fittings": polySpawn}, "efficiency", 3)
	fmt.Printf("  Deed crafted: %s\n", deedID)
	moveTo(wc, 30, 0)
	wc.clearMessages()
	wc.send(protocol.MsgPlaceHarvester, protocol.PlaceHarvesterMsg{
		X: 30, Z: 0, Kind: "personal", DeedItemID: deedID,
	})
	placedMsg, ok := wc.waitFor(protocol.MsgHarvesterPlaced, 5*time.Second)
	if !ok {
		fmt.Println("  [FAIL] harvester not placed")
		os.Exit(1)
	}
	fmt.Printf("  [PASS] Harvester placed (%s)\n", asMap(placedMsg)["harvester_id"])
	// Unfunded tick: hopper must stay empty (shutdown state).
	fmt.Println("  Waiting one harvest interval unfunded (expect shutdown)...")
	time.Sleep(25 * time.Second)
	wc.clearMessages()
	wc.send(protocol.MsgEmptyHopper, map[string]interface{}{})
	emptyMsg, ok := wc.waitFor(protocol.MsgHopperEmptied, 5*time.Second)
	if !ok {
		fmt.Println("  [FAIL] no hopper_emptied response")
		os.Exit(1)
	}
	if int(asMap(emptyMsg)["units"].(float64)) != 0 {
		fmt.Println("  [FAIL] unfunded harvester extracted units")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Unfunded harvester extracts nothing (shutdown)")
	fundHarvester(token, charID, asMap(placedMsg)["harvester_id"].(string), 1000)
	fmt.Println("  Waiting one harvest interval funded...")
	time.Sleep(25 * time.Second)
	wc.clearMessages()
	wc.send(protocol.MsgEmptyHopper, map[string]interface{}{})
	emptyMsg2, ok := wc.waitFor(protocol.MsgHopperEmptied, 5*time.Second)
	if !ok {
		fmt.Println("  [FAIL] no hopper_emptied response after funding")
		os.Exit(1)
	}
	got := int(asMap(emptyMsg2)["units"].(float64))
	if got <= 0 {
		fmt.Println("  [FAIL] funded harvester extracted nothing")
		os.Exit(1)
	}
	fmt.Printf("  [PASS] Funded harvester extracted %d units to stack\n", got)

	// --- Test 6: craft sidearm with experimentation ---
	fmt.Println("\n--- Test 6: Craft sidearm with experimentation ---")
	metalSrc := richestStack(token, charID, []string{"ferric_metal", "conductive_alloy"})
	polySrc := richestStack(token, charID, []string{"structural_polymer", "fibrous_flora"})
	sidearmID, sidearmStats := craftItemFull(token, charID, "basic_sidearm",
		map[string]string{"frame": metalSrc, "grip": polySrc}, "damage", 3, "Test Sidearm")
	dmin, dmax := sidearmStats["damage_min"], sidearmStats["damage_max"]
	if dmin < 20 || dmax <= dmin {
		fmt.Printf("  [FAIL] sidearm stats wrong: min=%v max=%v\n", dmin, dmax)
		os.Exit(1)
	}
	fmt.Printf("  [PASS] Sidearm crafted: damage %.0f–%.0f (base 20–40 + quality + experiment)\n", dmin, dmax)
	if getXP(token, charID, "crafting") < 500 {
		fmt.Println("  [FAIL] no crafting XP awarded")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Crafting XP awarded")

	// --- Test 7: equip sidearm, crafted damage lands in combat ---
	fmt.Println("\n--- Test 7: Equip + stat-delta in combat ---")
	equipItem(token, charID, sidearmID)
	moveTo(wc, 12, 0)
	// Collect the first Hit result; unarmed max is 17, so >= 18 proves the item range.
	proved := false
	deadline := time.Now().Add(30 * time.Second)
	for time.Now().Before(deadline) && !proved {
		wc.send(protocol.MsgCombatAction, protocol.CombatActionMsg{TargetID: liveCreatures[0]})
		time.Sleep(500 * time.Millisecond)
		wc.mu.Lock()
		for _, m := range wc.messages {
			if m.Type == protocol.MsgCombatResult {
				data, _ := json.Marshal(m.Data)
				var cr protocol.CombatResultMsg
				json.Unmarshal(data, &cr)
				if cr.Hit && cr.Damage >= 18 {
					fmt.Printf("  [PASS] Crafted sidearm hit for %d (unarmed max 17 — equip matters)\n", cr.Damage)
					proved = true
				}
			}
			if m.Type == protocol.MsgCreatureDeath {
				dm := asMap(m)
				if id, _ := dm["instance_id"].(string); id == liveCreatures[0] {
					liveCreatures = liveCreatures[1:]
				}
			}
		}
		wc.mu.Unlock()
	}
	if !proved {
		fmt.Println("  [FAIL] no crafted-range hit observed in 30s")
		os.Exit(1)
	}

	// --- Test 8: kill + corpse harvest + craft/equip plate ---
	fmt.Println("\n--- Test 8: Corpse harvest + armor ---")
	if len(liveCreatures) == 0 {
		fmt.Println("  [FAIL] no live creatures left for harvest test")
		os.Exit(1)
	}
	killedID := killCreature(wc, liveCreatures[0])
	if killedID == "" {
		fmt.Println("  [FAIL] could not kill a creature for harvest")
		os.Exit(1)
	}
	fmt.Println("  Creature killed; harvesting corpse...")
	wc.clearMessages()
	wc.send(protocol.MsgHarvestCorpse, map[string]interface{}{})
	chMsg, ok := wc.waitFor(protocol.MsgCorpseHarvested, 5*time.Second)
	if !ok {
		fmt.Println("  [FAIL] no corpse_harvested")
		os.Exit(1)
	}
	chm := asMap(chMsg)
	if int(chm["units"].(float64)) < 2 {
		fmt.Println("  [FAIL] corpse yield too small")
		os.Exit(1)
	}
	fmt.Printf("  [PASS] Corpse harvested: %d units %s (quality %v)\n",
		int(chm["units"].(float64)), chm["resource_type"], chm["quality"])
	// Top up organics to the plate lining quota (5): kill + harvest more.
	for liveIdx := 1; organicTotal(token, charID) < 5 && liveIdx < len(liveCreatures); liveIdx++ {
		fmt.Printf("  Quota loop: organic=%d, trying %s\n", organicTotal(token, charID), liveCreatures[liveIdx])
		killed := killCreature(wc, liveCreatures[liveIdx])
		fmt.Printf("  Quota loop: kill result %q\n", killed)
		if killed == "" {
			continue
		}
		wc.clearMessages()
		wc.send(protocol.MsgHarvestCorpse, map[string]interface{}{})
		if hm, ok := wc.waitFor(protocol.MsgCorpseHarvested, 5*time.Second); ok {
			fmt.Printf("  Quota loop: harvested %v\n", asMap(hm)["units"])
		} else {
			fmt.Println("  Quota loop: harvest got no response")
		}
	}
	if organicTotal(token, charID) < 5 {
		fmt.Println("  [FAIL] organic quota unmet for plate lining")
		os.Exit(1)
	}
	// Retreat out of aggro while crafting plate.
	moveTo(wc, 30, 0)
	plateID, plateStats := craftItemFull(token, charID, "basic_plate",
		map[string]string{"plating": metalSrc, "lining": richestStack(token, charID, []string{"cultured_organic", "fibrous_flora"})},
		"protection", 3, "Test Plate")
	if plateStats["protection_max"] < 5 {
		fmt.Println("  [FAIL] plate stats wrong")
		os.Exit(1)
	}
	equipItem(token, charID, plateID)
	fmt.Printf("  [PASS] Plate crafted (protection max %.0f) and equipped\n", plateStats["protection_max"])

	// --- Test 9: spawn lifecycle relocation ---
	fmt.Println("\n--- Test 9: Spawn lifecycle relocation ---")
	before := liveSpawnIDs(token, charID)
	fmt.Printf("  Tracking %d non-guarantee spawns; waiting out fast lifespans...\n", len(before))
	time.Sleep(210 * time.Second)
	after := liveSpawnIDs(token, charID)
	stillThere := 0
	for id := range before {
		if after[id] {
			stillThere++
		}
	}
	if stillThere == len(before) || len(after) == 0 {
		fmt.Printf("  [FAIL] no relocation observed (before=%d after=%d overlap=%d)\n",
			len(before), len(after), stillThere)
		os.Exit(1)
	}
	fmt.Printf("  [PASS] Lifecycle relocation: %d before → %d after (%d overlap)\n",
		len(before), len(after), stillThere)

	fmt.Println("\n=== ALL PHASE 4 TESTS PASSED ===")
	wc.conn.Close()
}

// --- helpers ---

func moveTo(wc *wsClient, tx, tz float64) {
	// Closed-loop stepwise walk at ~7 m/s (under the 10.5 validation ceiling):
	// each step is accepted only if no NEW position_correction arrives for it;
	// rejected steps are retried (dt grows, so they eventually pass).
	cx, cz := posX, posZ
	time.Sleep(1200 * time.Millisecond) // settle LastMoveTime after prior actions
	for i := 0; i < 200; i++ {
		dx, dz := tx-cx, tz-cz
		dist := math.Sqrt(dx*dx + dz*dz)
		if dist < 1.0 {
			posX, posZ = tx, tz
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
			posX, posZ = cx, cz
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

// killCreature attacks one live instance until ITS death message arrives (stale
// deaths for other instances are ignored). Retaliation may incapacitate the
// tester mid-hunt: on incap, it clones, walks back, and resumes.
func killCreature(wc *wsClient, targetID string) string {
	wc.clearMessages()
	deadline := time.Now().Add(120 * time.Second)
	for time.Now().Before(deadline) {
		wc.send(protocol.MsgCombatAction, protocol.CombatActionMsg{TargetID: targetID})
		time.Sleep(800 * time.Millisecond)
		// Scan fresh buffer only.
		dead, incapped := "", false
		wc.mu.Lock()
		for _, m := range wc.messages {
			switch m.Type {
			case protocol.MsgCreatureDeath:
				if id, _ := asMap(m)["instance_id"].(string); id == targetID {
					dead = id
				}
			case protocol.MsgIncapacitated:
				incapped = true
			case protocol.MsgError:
				if em, _ := asMap(m)["message"].(string); em == "incapacitated, cannot act" {
					incapped = true
				}
			}
		}
		wc.mu.Unlock()
		if dead != "" {
			return dead
		}
		if incapped {
			fmt.Println("  (incapacitated mid-hunt — cloning and resuming)")
			wc.clearMessages()
			wc.send(protocol.MsgClone, map[string]interface{}{})
			if _, ok := wc.waitFor(protocol.MsgCloned, 5*time.Second); !ok {
				return ""
			}
			posX, posZ = 20, 0 // clone drops us at the zone spawn
			moveTo(wc, 12, 0)
			wc.clearMessages()
		}
	}
	return ""
}

// Tracked client position (dead reckoning from spawn; tests always move via moveTo).
var posX, posZ = 20.0, 0.0

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

func craftAuthed(method, path string, token string, body interface{}) map[string]interface{} {
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

func getSchematics() []interface{} {
	resp, err := http.Get(serverURL + "/api/craft/schematics")
	if err != nil {
		log.Fatal("schematics failed:", err)
	}
	defer resp.Body.Close()
	var out map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&out)
	list, _ := out["schematics"].([]interface{})
	return list
}

func stackQty(token, charID, spawnID string) int {
	out := craftAuthed("GET", "/api/craft/resources?character_id="+charID, token, nil)
	total := 0
	if list, ok := out["stacks"].([]interface{}); ok {
		for _, s := range list {
			if sm, ok := s.(map[string]interface{}); ok && sm["SpawnID"] == spawnID {
				if q, ok := sm["Quantity"].(float64); ok {
					total += int(q)
				}
			}
		}
	}
	return total
}

func richestStack(token, charID string, types []string) string {	out := craftAuthed("GET", "/api/craft/resources?character_id="+charID, token, nil)
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

// organicTotal sums owned organic/flora stacks (plate lining quota helper).
func organicTotal(token, charID string) int {
	out := craftAuthed("GET", "/api/craft/resources?character_id="+charID, token, nil)
	total := 0
	if list, ok := out["stacks"].([]interface{}); ok {
		for _, s := range list {
			sm, _ := s.(map[string]interface{})
			rt, _ := sm["ResourceType"].(string)
			if rt == "cultured_organic" || rt == "fibrous_flora" {
				if q, ok := sm["Quantity"].(float64); ok {
					total += int(q)
				}
			}
		}
	}
	return total
}

func craftItem(token, charID, schemID string, slots map[string]string, prop string, points int) string {
	id, _ := craftItemFull(token, charID, schemID, slots, prop, points, "Test Item")
	return id
}

func craftItemFull(token, charID, schemID string, slots map[string]string, prop string, points int, name string) (string, map[string]float64) {
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
		"character_id": charID, "session_id": sid, "name": name,
	})
	stats := map[string]float64{}
	if sm, ok := done["stats"].(map[string]interface{}); ok {
		for k, v := range sm {
			if f, ok := v.(float64); ok {
				stats[k] = f
			}
		}
	}
	id, _ := done["item_id"].(string)
	return id, stats
}

func equipItem(token, charID, itemID string) {
	craftAuthed("POST", "/api/craft/equip", token,
		map[string]string{"character_id": charID, "item_id": itemID})
}

func fundHarvester(token, charID, hvid string, credits int) {
	craftAuthed("POST", "/api/craft/harvester/fund", token, map[string]interface{}{
		"character_id": charID, "harvester_id": hvid, "credits": credits,
	})
}

func getXP(token, charID, pool string) int {
	req, _ := http.NewRequest("GET", serverURL+"/api/characters/"+charID+"/skills", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("get skills failed:", err)
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	pools, _ := result["xp_pools"].(map[string]interface{})
	v, _ := pools[pool].(float64)
	return int(v)
}

func visibleCreatures(wc *wsClient) []string {
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

func liveSpawnIDs(token, charID string) map[string]bool {
	out := craftAuthed("GET", "/api/craft/spawns?character_id="+charID, token, nil)
	ids := map[string]bool{}
	if list, ok := out["spawns"].([]interface{}); ok {
		for _, s := range list {
			if sm, ok := s.(map[string]interface{}); ok {
				id, _ := sm["id"].(string)
				if len(id) >= 9 && id[:9] == "testspawn" {
					continue
				}
				ids[id] = true
			}
		}
	}
	return ids
}
