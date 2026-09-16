// Phase 8 integration test: alignment + flagging, PvP gating, overt duels
// to incap with death penalties (transfer + condition + points), ranks,
// faction bases (place, window, siege, destroy), city PvP permission,
// faction chat, roster, switches, and city leaning.
//
// Exit criteria under test (predecessor-GDD §30, HISTORICAL): "Two Overt,
// opposing-faction players can engage in valid PvP combat with correct death
// penalties (Section 9.4.2), and faction points/ranks accrue correctly."
//
// Requires TESTBED_FAST_CYCLE=1 on the server (rank thresholds 100/200/300;
// 30 s covert delay; 20 s harvest ticks) and a FRESH database for standalone
// runs. Shared-DB full-suite note: the mini-city sits east (center 515) to
// avoid Phase 5 leftovers (45/70) and Phase 7 cities (center 100/200);
// the siege base sits at (320,0) outside all active cities. Total runtime
// ~35 minutes (quotas + three duels + walk-backs dominate).
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

func seenType(wc *wsClient, msgType string) bool {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	for _, m := range wc.messages {
		if m.Type == msgType {
			return true
		}
	}
	return false
}

func hasChannel(wc *wsClient, channel string) bool {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	for _, m := range wc.messages {
		if m.Type == protocol.MsgChatMessage && asMap(m)["channel"] == channel {
			return true
		}
	}
	return false
}

// Per-client tracked positions (dead reckoning; all moves via moveTo).
var ax, az = 20.0, 0.0
var bx, bz = 20.0, 0.0
var cx, cz = 20.0, 0.0
var dx, dz = 20.0, 0.0

func moveTo(wc *wsClient, tx, tz float64, px, pz *float64) {
	cx, cz := *px, *pz
	time.Sleep(1200 * time.Millisecond)
	for i := 0; i < 200; i++ {
		ddx, ddz := tx-cx, tz-cz
		dist := math.Sqrt(ddx*ddx + ddz*ddz)
		if dist < 1.0 {
			*px, *pz = tx, tz
			return
		}
		step := 4.0
		if dist < step {
			step = dist
		}
		nx, nz := cx+ddx/dist*step, cz+ddz/dist*step
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

func main() {
	fmt.Println("=== Phase 8 Integration Test (testbed fork, generic content) ===")
	fmt.Println("Flagging → PvP duels → death penalties → ranks → bases → city permission → social")
	fmt.Println()

	if !waitForServer() {
		fmt.Println("FAIL: Server not reachable")
		os.Exit(1)
	}
	fmt.Println("Server is reachable.")

	// Setup: A + B duelists (artisan + brawler), C ally (no skills), D neutral.
	_, tokenA := registerAccount()
	charA := createCharacter(tokenA, "human", "OvertA")
	_, tokenB := registerAccount()
	charB := createCharacter(tokenB, "human", "OvertB")
	_, tokenC := registerAccount()
	charC := createCharacter(tokenC, "human", "AllyC")
	_, tokenD := registerAccount()
	charD := createCharacter(tokenD, "human", "NeutralD")
	trainSkill(tokenA, charA, "trainer_artisan_z1", "artisan_novice")
	trainSkill(tokenA, charA, "trainer_brawler_z1", "brawler_novice")
	trainSkill(tokenB, charB, "trainer_artisan_z1", "artisan_novice")
	trainSkill(tokenB, charB, "trainer_brawler_z1", "brawler_novice")
	fmt.Println("Setup: two duelists + ally + neutral")

	wcA := connectWS(tokenA)
	wcB := connectWS(tokenB)
	wcC := connectWS(tokenC)
	wcD := connectWS(tokenD)
	wcA.send(protocol.MsgEnterWorld, protocol.EnterWorldMsg{CharacterID: charA})
	wcB.send(protocol.MsgEnterWorld, protocol.EnterWorldMsg{CharacterID: charB})
	wcC.send(protocol.MsgEnterWorld, protocol.EnterWorldMsg{CharacterID: charC})
	wcD.send(protocol.MsgEnterWorld, protocol.EnterWorldMsg{CharacterID: charD})
	for i, wc := range []*wsClient{wcA, wcB, wcC, wcD} {
		if _, ok := wc.waitFor(protocol.MsgWorldEnter, 5*time.Second); !ok {
			fmt.Printf("  [FAIL] client %d did not enter world\n", i)
			os.Exit(1)
		}
	}
	fmt.Println("  [PASS] All four clients entered world")

	// --- Test 1: quotas + gear ---
	fmt.Println("\n--- Test 1: Quotas + gear ---")
	// A crafts five deeds (3 houses + hall + base: 5×15M+8P) + sidearm (10+5).
	sampleQuota(wcA, tokenA, charA, "mineral", "ferric_metal", 90, &ax, &az)
	sampleQuota(wcA, tokenA, charA, "chemical", "structural_polymer", 50, &ax, &az)
	sampleQuota(wcB, tokenB, charB, "mineral", "ferric_metal", 25, &bx, &bz)
	sampleQuota(wcB, tokenB, charB, "chemical", "structural_polymer", 14, &bx, &bz)
	// Deed slots bind single stacks: guarantee the richest stack covers the
	// largest slot (A frame 15 / fittings 8; B frame 10 / grip 5).
	ensureStack(wcA, tokenA, charA, "mineral", "ferric_metal", 15, &ax, &az)
	ensureStack(wcA, tokenA, charA, "chemical", "structural_polymer", 8, &ax, &az)
	ensureStack(wcB, tokenB, charB, "mineral", "ferric_metal", 10, &bx, &bz)
	ensureStack(wcB, tokenB, charB, "chemical", "structural_polymer", 5, &bx, &bz)
	fmt.Println("  [PASS] Quotas sampled (single-stack needs covered)")
	houseA1 := craftDeed(tokenA, charA, "structure_deed")
	houseA2 := craftDeed(tokenA, charA, "structure_deed")
	houseA3 := craftDeed(tokenA, charA, "structure_deed")
	hallA := craftDeed(tokenA, charA, "city_hall_deed")
	baseDeed := craftDeed(tokenA, charA, "base_deed")
	swordA := craftItem(tokenA, charA, "basic_sidearm",
		map[string]string{"frame": metalSrc(tokenA, charA), "grip": polySrc(tokenA, charA)}, "damage", 3)
	swordB := craftItem(tokenB, charB, "basic_sidearm",
		map[string]string{"frame": metalSrc(tokenB, charB), "grip": polySrc(tokenB, charB)}, "damage", 3)
	equipItem(tokenA, charA, swordA)
	equipItem(tokenB, charB, swordB)
	fmt.Println("  [PASS] Deeds + sidearms crafted and equipped")

	// --- Test 2: mini-city + alignment + overt ---
	fmt.Println("\n--- Test 2: Mini-city + alignment ---")
	placeHouse(wcA, houseA1, 440, 0, &ax, &az)
	fundStructure(tokenA, charA, lastPlaced(), 100)
	placeHouse(wcA, houseA2, 465, 0, &ax, &az)
	fundStructure(tokenA, charA, lastPlaced(), 100)
	placeHouse(wcA, houseA3, 490, 0, &ax, &az)
	fundStructure(tokenA, charA, lastPlaced(), 100)
	cityID := placeCityHall(wcA, hallA, 515, 0, "Duel Ground", &ax, &az)
	factionPOST(tokenA, "/api/civic/city/fund", map[string]interface{}{"character_id": charA, "city_id": cityID, "credits": 500})
	fmt.Printf("  [PASS] Mini-city active (%s)\n", cityID)
	fmt.Println("  Waiting one interval (enrol tick)...")
	time.Sleep(25 * time.Second)
	factionPOST(tokenA, "/api/faction/declare", map[string]string{"character_id": charA, "alignment": "alignment_a"})
	factionPOST(tokenB, "/api/faction/declare", map[string]string{"character_id": charB, "alignment": "alignment_b"})
	factionPOST(tokenC, "/api/faction/declare", map[string]string{"character_id": charC, "alignment": "alignment_a"})
	fmt.Println("  [PASS] Alignments declared (a/b/a, D neutral)")
	if _, code := factionPOSTRaw(tokenA, "/api/faction/declare", map[string]string{"character_id": charA, "alignment": "alignment_a"}); code < 300 {
		fmt.Println("  [FAIL] re-declaring same alignment accepted")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Same-alignment re-declare refused")
	goOvert(wcA)
	goOvert(wcB)
	fmt.Println("  [PASS] Duelists overt")
	moveTo(wcA, 480, 0, &ax, &az)
	moveTo(wcB, 480, 0, &bx, &bz)

	// --- Test 3: PvP gating rejects ---
	fmt.Println("\n--- Test 3: PvP gating ---")
	attackExpectError(wcA, charD, "neutral target")   // A→D: neutral never valid
	attackExpectError(wcA, charC, "covert victim")    // A→C: victim must be overt
	attackExpectError(wcD, charA, "neutral attacker") // D→A: no enemy alignment
	fmt.Println("  [PASS] Neutral/covert gating rejects")
	goCovert(wcB) // B overt→covert (delay long passed since T2)
	wcB.clearMessages()
	wcB.send(protocol.MsgCombatAction, protocol.CombatActionMsg{TargetID: charA})
	flagMsg, ok := wcB.waitFor(protocol.MsgFlagChanged, 5*time.Second)
	if !ok || asMap(flagMsg)["overt"] != true {
		fmt.Println("  [FAIL] covert attack did not auto-flag overt")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Covert attack auto-flags overt (both sentences reconciled)")

	// --- Test 4: duel #1 → incap → clone penalties ---
	fmt.Println("\n--- Test 4: Duel #1 + death penalties ---")
	wA0, wB0 := wallet(tokenA, charA), wallet(tokenB, charB)
	duelToIncap(wcA, wcB, charB)
	fmt.Println("  [PASS] Overt duel to incapacitation (real resolution)")
	wcB.clearMessages()
	wcB.send(protocol.MsgClone, map[string]interface{}{})
	if _, ok := wcB.waitFor(protocol.MsgCloned, 10*time.Second); !ok {
		fmt.Println("  [FAIL] no clone after PvP incap")
		os.Exit(1)
	}
	bx, bz = 20, 0 // clone bind: zone spawn
	drop := wB0 * 10 / 100
	if w := wallet(tokenB, charB); w != wB0-drop {
		fmt.Printf("  [FAIL] victim wallet %d (want %d)\n", w, wB0-drop)
		os.Exit(1)
	}
	if w := wallet(tokenA, charA); w != wA0+drop {
		fmt.Printf("  [FAIL] killer wallet %d (want %d)\n", w, wA0+drop)
		os.Exit(1)
	}
	foundDrop := false
	for _, e := range econLedger(tokenB, charB) {
		em, _ := e.(map[string]interface{})
		if em["Category"] == "pvp_death_drop" && em["Flow"] == "transfer" {
			if amt, _ := em["Amount"].(float64); int(amt) == -drop {
				foundDrop = true
			}
		}
	}
	if !foundDrop {
		fmt.Println("  [FAIL] no pvp_death_drop transfer pair in ledger")
		os.Exit(1)
	}
	fmt.Printf("  [PASS] 10%% credit transfer victim→killer (%d) + ledger pair\n", drop)
	if c := itemCondition(tokenB, charB, swordB); c != 90 {
		fmt.Printf("  [FAIL] equipped condition = %d (want 90)\n", c)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Equipped condition −10 (floor-0 minimal surface)")
	if w := waitWounds(wcB, 10*time.Second); w <= 0 {
		fmt.Println("  [FAIL] no wounds after PvP death")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Wounds/BF applied on PvP death")
	st := factionStatus(tokenA, charA)
	if p, _ := st["points"].(float64); int(p) != 100 {
		fmt.Printf("  [FAIL] killer points = %v (want 100)\n", p)
		os.Exit(1)
	}
	if r, _ := st["rank"].(string); r != "sergeant" {
		fmt.Printf("  [FAIL] killer rank = %q (want sergeant)\n", r)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Killer +100 points → Sergeant (fast-cycle threshold)")
	moveTo(wcB, 480, 0, &bx, &bz) // walk back from clone bind

	// --- Test 5: city PvP permission gate ---
	fmt.Println("\n--- Test 5: City PvP permission ---")
	factionPOST(tokenA, "/api/faction/city/set-pvp", map[string]interface{}{"character_id": charA, "city_id": cityID, "allowed": false})
	attackExpectMessage(wcA, charB, "denied by city", "city-disallowed")
	fmt.Println("  [PASS] City-disallowed attack refused")
	if _, code := factionPOSTRaw(tokenB, "/api/faction/city/set-pvp", map[string]interface{}{"character_id": charB, "city_id": cityID, "allowed": true}); code != 403 {
		fmt.Printf("  [FAIL] non-mayor PvP toggle not refused (HTTP %d)\n", code)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Non-mayor PvP toggle refused")
	factionPOST(tokenA, "/api/faction/city/set-pvp", map[string]interface{}{"character_id": charA, "city_id": cityID, "allowed": true})
	wcA.clearMessages()
	wcA.send(protocol.MsgCombatAction, protocol.CombatActionMsg{TargetID: charB})
	if _, ok := wcA.waitFor(protocol.MsgCombatResult, 10*time.Second); !ok {
		fmt.Println("  [FAIL] no combat result after re-allow")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Re-allowed attack resolves (gate is live, not cosmetic)")

	// --- Test 6: duels #2–3 → Major → Colonel ---
	fmt.Println("\n--- Test 6: Duels #2–3 → ranks ---")
	duelToIncap(wcA, wcB, charB)
	cloneAndReturn(wcB, tokenB, charB, &bx, &bz)
	st = factionStatus(tokenA, charA)
	if p, _ := st["points"].(float64); int(p) != 200 {
		fmt.Printf("  [FAIL] killer points = %v (want 200)\n", p)
		os.Exit(1)
	}
	if r, _ := st["rank"].(string); r != "major" {
		fmt.Printf("  [FAIL] killer rank = %q (want major)\n", r)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Second kill → 200 points → Major")
	duelToIncap(wcA, wcB, charB)
	cloneAndReturn(wcB, tokenB, charB, &bx, &bz)
	st = factionStatus(tokenA, charA)
	if p, _ := st["points"].(float64); int(p) != 300 {
		fmt.Printf("  [FAIL] killer points = %v (want 300)\n", p)
		os.Exit(1)
	}
	if r, _ := st["rank"].(string); r != "colonel" {
		fmt.Printf("  [FAIL] killer rank = %q (want colonel)\n", r)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Third kill → 300 points → Colonel (MVP cap)")
	if p, _ := factionStatus(tokenB, charB)["points"].(float64); int(p) != 0 {
		fmt.Println("  [FAIL] victim earned points")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Victims earn nothing")

	// --- Test 7: guild + base + siege ---
	fmt.Println("\n--- Test 7: Guild + base + siege ---")
	gOut := factionPOST(tokenA, "/api/civic/guild/found", map[string]string{"character_id": charA, "name": "Storm Compact", "tag": "SC"})
	guildID := gOut["guild_id"].(string)
	invC := factionPOST(tokenA, "/api/civic/guild/invite", map[string]string{"character_id": charA, "guild_id": guildID, "invitee_name": "AllyC"})["invite_id"].(string)
	factionPOST(tokenC, "/api/civic/guild/accept", map[string]string{"character_id": charC, "invite_id": invC})
	factionPOST(tokenA, "/api/faction/guild/set-alignment", map[string]string{"character_id": charA, "guild_id": guildID, "alignment": "alignment_a"})
	fmt.Println("  [PASS] War guild founded (a-aligned, 2 members)")
	if _, code := factionPOSTRaw(tokenB, "/api/faction/guild/set-alignment", map[string]string{"character_id": charB, "guild_id": guildID, "alignment": "alignment_b"}); code != 403 {
		fmt.Printf("  [FAIL] non-leader guild alignment change not refused (HTTP %d)\n", code)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Non-leader guild alignment change refused")
	moveTo(wcA, 320, 0, &ax, &az)
	baseID := placeBase(wcA, baseDeed, guildID, 320, 0, &ax, &az)
	fmt.Printf("  [PASS] Base placed in wilderness (%s, Major-gated)\n", baseID)
	nowSec := time.Now().Unix()
	factionPOST(tokenA, "/api/faction/base/window", map[string]interface{}{"character_id": charA, "base_id": baseID, "window_start": nowSec + 300, "window_end": nowSec + 900})
	moveTo(wcB, 320, 0, &bx, &bz)
	wcB.clearMessages()
	wcB.send(protocol.MsgCombatAction, protocol.CombatActionMsg{TargetID: baseID})
	wm, ok := wcB.waitFor(protocol.MsgError, 5*time.Second)
	if !ok {
		fmt.Println("  [FAIL] out-of-window siege not refused")
		os.Exit(1)
	}
	if msg, _ := asMap(wm)["message"].(string); !contains(msg, "window") {
		fmt.Printf("  [FAIL] wrong refusal reason: %q\n", msg)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Out-of-window siege refused (precise reason)")
	factionPOST(tokenA, "/api/faction/base/window", map[string]interface{}{"character_id": charA, "base_id": baseID, "window_start": nowSec, "window_end": nowSec + 600})
	wcB.clearMessages()
	wcB.send(protocol.MsgCombatAction, protocol.CombatActionMsg{TargetID: baseID})
	if _, ok := wcB.waitFor(protocol.MsgCombatResult, 10*time.Second); !ok {
		fmt.Println("  [FAIL] no combat result vs base")
		os.Exit(1)
	}
	if hp := baseHP(tokenA, baseID); hp >= 3000 {
		fmt.Printf("  [FAIL] base HP not reduced (%d)\n", hp)
		os.Exit(1)
	}
	fmt.Println("  [PASS] In-window siege hits register (HP drops)")
	siegeToDestroyed(wcB, baseID)
	fmt.Println("  [PASS] Base destroyed (3000 HP, real resolution)")
	stB := factionStatus(tokenB, charB)
	if p, _ := stB["points"].(float64); int(p) != 250 {
		fmt.Printf("  [FAIL] besieger points = %v (want 250 destroy award)\n", p)
		os.Exit(1)
	}
	if r, _ := stB["rank"].(string); r != "major" {
		fmt.Printf("  [FAIL] besieger rank = %q (want major at 250)\n", r)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Destroy award +250 → Major")
	wcB.clearMessages()
	wcB.send(protocol.MsgCombatAction, protocol.CombatActionMsg{TargetID: baseID})
	dm, ok := wcB.waitFor(protocol.MsgError, 5*time.Second)
	if !ok {
		fmt.Println("  [FAIL] destroyed-base attack not refused")
		os.Exit(1)
	}
	if msg, _ := asMap(dm)["message"].(string); !contains(msg, "destroyed") {
		fmt.Printf("  [FAIL] wrong refusal reason: %q\n", msg)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Destroyed base refuses further attacks (precise reason)")

	// --- Test 8: faction chat + roster + leaning ---
	fmt.Println("\n--- Test 8: Faction social ---")
	wcC.clearMessages()
	wcB.clearMessages()
	wcA.send(protocol.MsgChat, protocol.ChatMsg{Channel: "faction", Text: "for the alignment"})
	if _, ok := wcC.waitFor(protocol.MsgChatMessage, 5*time.Second); !ok {
		fmt.Println("  [FAIL] faction chat not delivered same-side")
		os.Exit(1)
	}
	time.Sleep(1500 * time.Millisecond)
	if hasChannel(wcB, "faction") {
		fmt.Println("  [FAIL] faction chat leaked to enemy side")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Faction chat isolated to alignment")
	wcD.clearMessages()
	wcD.send(protocol.MsgChat, protocol.ChatMsg{Channel: "faction", Text: "anyone?"})
	if _, ok := wcD.waitFor(protocol.MsgError, 5*time.Second); !ok {
		fmt.Println("  [FAIL] neutral faction chat not refused")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Neutral faction chat refused (no alignment)")
	roster := factionGET(tokenA, "/api/faction/roster?character_id="+charA)["overt"].([]interface{})
	if len(roster) != 1 || roster[0] != "OvertA" {
		fmt.Printf("  [FAIL] overt roster = %v (want [OvertA])\n", roster)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Overt roster same-side only (covert ally excluded)")
	city := factionGET(tokenA, "/api/civic/city/get?character_id=x&city_id="+cityID)
	leaning, _ := city["leaning"].(map[string]interface{})
	if a, _ := leaning["alignment_a"].(float64); int(a) < 1 {
		fmt.Printf("  [FAIL] leaning = %v\n", leaning)
		os.Exit(1)
	}
	fmt.Printf("  [PASS] City leaning displayed (display-only): %v\n", leaning)

	// --- Test 9: covert delay + switches ---
	fmt.Println("\n--- Test 9: Covert delay + switches ---")
	goOvert(wcC)
	wcC.clearMessages()
	wcC.send(protocol.MsgGoCovert, protocol.FlagMsg{})
	if _, ok := wcC.waitFor(protocol.MsgError, 5*time.Second); !ok {
		fmt.Println("  [FAIL] fresh-overt covert request not refused")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Covert delay enforced (anti-combat-log)")
	fmt.Println("  Waiting out the fast-cycle delay...")
	time.Sleep(35 * time.Second)
	wcC.clearMessages()
	wcC.send(protocol.MsgGoCovert, protocol.FlagMsg{})
	if _, ok := wcC.waitFor(protocol.MsgFlagChanged, 5*time.Second); !ok {
		fmt.Println("  [FAIL] covert after delay failed")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Covert succeeds after delay")
	factionPOST(tokenA, "/api/faction/declare", map[string]string{"character_id": charA, "alignment": "neutral"})
	st = factionStatus(tokenA, charA)
	if p, _ := st["points"].(float64); int(p) != 0 {
		fmt.Printf("  [FAIL] points after leaving = %v (want reset 0)\n", p)
		os.Exit(1)
	}
	if o, _ := st["overt"].(bool); o {
		fmt.Println("  [FAIL] overt flag survived leaving")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Leaving resets points and clears overt")
	if _, code := factionPOSTRaw(tokenA, "/api/faction/declare", map[string]string{"character_id": charA, "alignment": "alignment_a"}); code != 403 {
		fmt.Printf("  [FAIL] cooldown re-declare not refused (HTTP %d)\n", code)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Cooldown re-declare refused (30-day rule, live refusal)")
	wcD.clearMessages()
	wcD.send(protocol.MsgGoOvert, protocol.FlagMsg{})
	if _, ok := wcD.waitFor(protocol.MsgError, 5*time.Second); !ok {
		fmt.Println("  [FAIL] neutral overt not refused")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Neutral cannot go overt")
	factionPOST(tokenD, "/api/faction/declare", map[string]string{"character_id": charD, "alignment": "alignment_b"})
	factionPOST(tokenD, "/api/faction/declare", map[string]string{"character_id": charD, "alignment": "neutral"})
	fmt.Println("  [PASS] Neutral declare + leave cycle (first declaration free)")

	fmt.Println("\n=== ALL PHASE 8 TESTS PASSED ===")
	wcA.conn.Close()
	wcB.conn.Close()
	wcC.conn.Close()
	wcD.conn.Close()
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
	out, code := authedRaw(method, path, token, body)
	if code >= 300 {
		log.Fatalf("API %s HTTP %d: %v", path, code, out)
	}
	return out
}

func authedRaw(method, path, token string, body interface{}) (map[string]interface{}, int) {
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
	var out map[string]interface{}
	json.Unmarshal(respBody, &out)
	return out, resp.StatusCode
}

func factionGET(token, path string) map[string]interface{} {
	return authed("GET", path, token, nil)
}

func factionPOST(token, path string, body interface{}) map[string]interface{} {
	return authed("POST", path, token, body)
}

func factionPOSTRaw(token, path string, body interface{}) (map[string]interface{}, int) {
	return authedRaw("POST", path, token, body)
}

func factionStatus(token, charID string) map[string]interface{} {
	return factionGET(token, "/api/faction/status?character_id="+charID)
}

func wallet(token, charID string) int {
	out := authed("GET", "/api/characters/"+charID, token, nil)
	if v, ok := out["credits"].(float64); ok {
		return int(v)
	}
	log.Fatalf("wallet unreadable: %v", out)
	return 0
}

func econLedger(token, charID string) []interface{} {
	out := authed("GET", "/api/economy/ledger?character_id="+charID, token, nil)
	list, _ := out["entries"].([]interface{})
	return list
}

func itemCondition(token, charID, itemID string) int {
	out := authed("GET", "/api/craft/items?character_id="+charID, token, nil)
	if list, ok := out["items"].([]interface{}); ok {
		for _, it := range list {
			im, _ := it.(map[string]interface{})
			if im["ID"] == itemID {
				c, _ := im["Condition"].(float64)
				return int(c)
			}
		}
	}
	log.Fatalf("item %s not found", itemID)
	return -1
}

func waitWounds(wc *wsClient, timeout time.Duration) int {
	deadline := time.Now().Add(timeout)
	best := -1
	for time.Now().Before(deadline) {
		wc.mu.Lock()
		for _, m := range wc.messages {
			if m.Type == protocol.MsgHAMUpdate {
				if w, ok := asMap(m)["wounds_health"].(float64); ok {
					best = int(w)
				}
			}
		}
		wc.mu.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
	return best
}

// placedIDs tracks WS placement receipts in order.
var placedIDs []string

func lastPlaced() string {
	if len(placedIDs) == 0 {
		log.Fatal("no placements yet")
	}
	return placedIDs[len(placedIDs)-1]
}

func placeHouse(wc *wsClient, deedID string, x, z float64, px, pz *float64) {
	moveTo(wc, x, z, px, pz)
	wc.clearMessages()
	wc.send(protocol.MsgPlaceHouse, protocol.PlaceStructureMsg{X: x, Z: z, Tier: "small", DeedItemID: deedID})
	m, ok := wc.waitFor(protocol.MsgStructurePlaced, 10*time.Second)
	if !ok {
		log.Fatalf("place house (%.0f,%.0f) failed", x, z)
	}
	placedIDs = append(placedIDs, asMap(m)["structure_id"].(string))
}

func placeCityHall(wc *wsClient, deedID string, x, z float64, name string, px, pz *float64) string {
	moveTo(wc, x, z, px, pz)
	wc.clearMessages()
	wc.send(protocol.MsgPlaceCityHall, protocol.PlaceCityHallMsg{X: x, Z: z, Name: name, DeedItemID: deedID})
	m, ok := wc.waitFor(protocol.MsgCityFounded, 10*time.Second)
	if !ok {
		log.Fatalf("place city hall (%.0f,%.0f) failed", x, z)
	}
	fm := asMap(m)
	if fm["status"] != "active" {
		log.Fatalf("city %s not immediately active: %v", name, fm)
	}
	return fm["city_id"].(string)
}

func placeBase(wc *wsClient, deedID, guildID string, x, z float64, px, pz *float64) string {
	moveTo(wc, x, z, px, pz)
	wc.clearMessages()
	wc.send(protocol.MsgPlaceBase, protocol.PlaceBaseMsg{X: x, Z: z, GuildID: guildID, DeedItemID: deedID})
	m, ok := wc.waitFor(protocol.MsgBasePlaced, 10*time.Second)
	if !ok {
		// Surface the server's refusal reason for forensics.
		wc.mu.Lock()
		for _, em := range wc.messages {
			if em.Type == protocol.MsgError {
				log.Printf("place base refused: %v", asMap(em)["message"])
			}
		}
		wc.mu.Unlock()
		log.Fatalf("place base (%.0f,%.0f) failed", x, z)
	}
	return asMap(m)["base_id"].(string)
}

func fundStructure(token, charID, stID string, credits int) {
	authed("POST", "/api/economy/structure/fund", token, map[string]interface{}{
		"character_id": charID, "structure_id": stID, "credits": credits,
	})
}

func goOvert(wc *wsClient) {
	wc.clearMessages()
	wc.send(protocol.MsgGoOvert, protocol.FlagMsg{})
	if _, ok := wc.waitFor(protocol.MsgFlagChanged, 5*time.Second); !ok {
		log.Fatal("go_overt failed")
	}
}

func goCovert(wc *wsClient) {
	wc.clearMessages()
	wc.send(protocol.MsgGoCovert, protocol.FlagMsg{})
	if _, ok := wc.waitFor(protocol.MsgFlagChanged, 5*time.Second); !ok {
		log.Fatal("go_covert failed")
	}
}

// attackExpectError sends one attack and requires a server error.
func attackExpectError(att *wsClient, targetID, why string) {
	att.clearMessages()
	att.send(protocol.MsgCombatAction, protocol.CombatActionMsg{TargetID: targetID})
	if _, ok := att.waitFor(protocol.MsgError, 5*time.Second); !ok {
		log.Fatalf("attack (%s) not refused", why)
	}
}

// attackExpectMessage sends one attack and requires an error containing text.
func attackExpectMessage(att *wsClient, targetID, text, why string) {
	att.clearMessages()
	att.send(protocol.MsgCombatAction, protocol.CombatActionMsg{TargetID: targetID})
	m, ok := att.waitFor(protocol.MsgError, 5*time.Second)
	if !ok {
		log.Fatalf("attack (%s) not refused", why)
	}
	if msg, _ := asMap(m)["message"].(string); !contains(msg, text) {
		log.Fatalf("attack (%s) wrong reason: %q", why, msg)
	}
}

func contains(s, sub string) bool {
	if len(sub) == 0 {
		return true
	}
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

func baseHP(token, baseID string) int {
	out := factionGET(token, "/api/faction/base/info?base_id="+baseID)
	hp, _ := out["hp"].(float64)
	return int(hp)
}

// duelToIncap spams attacks until the victim reports incapacitation.
func duelToIncap(att, vic *wsClient, vicChar string) {
	vic.clearMessages()
	deadline := time.Now().Add(300 * time.Second)
	for time.Now().Before(deadline) {
		att.send(protocol.MsgCombatAction, protocol.CombatActionMsg{TargetID: vicChar})
		time.Sleep(800 * time.Millisecond)
		if seenType(vic, protocol.MsgIncapacitated) {
			return
		}
	}
	log.Fatal("duel never reached incapacitation")
}

// cloneAndReturn clones the victim and walks them back to the duel ground.
func cloneAndReturn(vic *wsClient, token, vicChar string, px, pz *float64) {
	vic.clearMessages()
	vic.send(protocol.MsgClone, map[string]interface{}{})
	if _, ok := vic.waitFor(protocol.MsgCloned, 10*time.Second); !ok {
		log.Fatal("no clone after PvP incap")
	}
	*px, *pz = 20, 0 // clone bind: zone spawn
	moveTo(vic, 480, 0, px, pz)
}

// siegeToDestroyed spams base attacks until the destroyed broadcast.
func siegeToDestroyed(att *wsClient, baseID string) {
	deadline := time.Now().Add(400 * time.Second)
	for time.Now().Before(deadline) {
		att.send(protocol.MsgCombatAction, protocol.CombatActionMsg{TargetID: baseID})
		time.Sleep(800 * time.Millisecond)
		if seenType(att, protocol.MsgBaseDestroyed) {
			return
		}
	}
	log.Fatal("siege never destroyed the base")
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

func guaranteeCoords(wantType string) (float64, float64) {
	switch wantType {
	case "ferric_metal":
		return 30, 0
	case "structural_polymer":
		return 40, 0
	default:
		return 20, 0
	}
}

func sampleQuota(wc *wsClient, token, charID, tool, wantType string, qty int, px, pz *float64) {
	// Anchor on the guarantee spawn first: surveying from its coordinates
	// returns it (nearest), so units accumulate in ONE stack and deed slots
	// (which bind a single stack) never starve on split totals.
	gx, gz := guaranteeCoords(wantType)
	moveTo(wc, gx, gz, px, pz)
	for outer := 0; outer < 16; outer++ {
		if stackTotal(token, charID, []string{wantType}) >= qty {
			fmt.Printf("  quota(%s): met (%d)\n", wantType, qty)
			return
		}
		sr, ok := doSurvey(wc, tool)
		if !ok {
			fmt.Printf("  quota(%s): survey empty, recentering\n", wantType)
			moveTo(wc, 20, 0, px, pz)
			time.Sleep(11 * time.Second)
			continue
		}
		found := sr["resource_type"].(string)
		if found != wantType {
			gx, gz := guaranteeCoords(wantType)
			fmt.Printf("  quota(%s): found %s instead, walking to guarantee (%.0f,%.0f)\n",
				wantType, found, gx, gz)
			moveTo(wc, gx, gz, px, pz)
			time.Sleep(11 * time.Second)
			continue
		}
		wp := sr["waypoint"].(map[string]interface{})
		moveTo(wc, wp["x"].(float64), wp["z"].(float64), px, pz)
		for i := 0; i < 80; i++ {
			if stackTotal(token, charID, []string{wantType}) >= qty {
				fmt.Printf("  quota(%s): met (%d)\n", wantType, qty)
				return
			}
			if _, ok := doSample(wc); !ok {
				break
			}
			time.Sleep(300 * time.Millisecond)
		}
		time.Sleep(11 * time.Second)
	}
	log.Fatalf("quota(%s): unmet after search", wantType)
}

func richestStack(token, charID string, types []string) string {
	id, q := richestStackID(token, charID, types)
	if id == "" || q <= 0 {
		log.Fatalf("no stack of types %v", types)
	}
	return id
}

// richestStackID returns the spawn ID and quantity of the largest stack of
// the wanted types (deed slots bind ONE stack, so totals are not enough).
func richestStackID(token, charID string, types []string) (string, int) {
	out := authed("GET", "/api/craft/resources?character_id="+charID, token, nil)
	best, bestQ := "", 0
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
			if sid, _ := sm["SpawnID"].(string); q > bestQ {
				best, bestQ = sid, q
			}
		}
	}
	return best, bestQ
}

// ensureStack tops up until the RICHEST stack of wantType holds ≥ need units.
// Sampling grows whatever spawn the survey returns; standing still at a
// wantType spawn, the nearest-survey is stable, so bursts accumulate into one
// stack (despawn churn is retried by the outer loop).
func ensureStack(wc *wsClient, token, charID, tool, wantType string, need int, px, pz *float64) {
	for i := 0; i < 40; i++ {
		if _, q := richestStackID(token, charID, []string{wantType}); q >= need {
			fmt.Printf("  stack(%s): richest %d (need %d)\n", wantType, q, need)
			return
		}
		sr, ok := doSurvey(wc, tool)
		if !ok {
			time.Sleep(11 * time.Second)
			continue
		}
		if found, _ := sr["resource_type"].(string); found != wantType {
			gx, gz := guaranteeCoords(wantType)
			moveTo(wc, gx, gz, px, pz)
			time.Sleep(11 * time.Second)
			continue
		}
		wp := sr["waypoint"].(map[string]interface{})
		moveTo(wc, wp["x"].(float64), wp["z"].(float64), px, pz)
		for j := 0; j < 20; j++ {
			if _, ok := doSample(wc); !ok {
				break
			}
			time.Sleep(300 * time.Millisecond)
			if _, q := richestStackID(token, charID, []string{wantType}); q >= need {
				fmt.Printf("  stack(%s): richest %d (need %d)\n", wantType, q, need)
				return
			}
		}
		time.Sleep(11 * time.Second)
	}
	log.Fatalf("ensureStack(%s,%d): unmet", wantType, need)
}

func metalSrc(token, charID string) string {
	return richestStack(token, charID, []string{"ferric_metal", "conductive_alloy"})
}

func polySrc(token, charID string) string {
	return richestStack(token, charID, []string{"structural_polymer", "fibrous_flora"})
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

func craftDeed(token, charID, schemID string) string {
	return craftItem(token, charID, schemID,
		map[string]string{"frame": metalSrc(token, charID), "fittings": polySrc(token, charID)}, "capacity", 3)
}

func equipItem(token, charID, itemID string) {
	authed("POST", "/api/craft/equip", token, map[string]string{
		"character_id": charID, "item_id": itemID,
	})
}
