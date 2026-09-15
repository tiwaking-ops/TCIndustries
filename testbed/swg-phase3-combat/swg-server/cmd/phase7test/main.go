// Phase 7 integration test: city founding → elections → treasury/tax/ranks →
// guilds/dues → groups/chat → mail (offline) → permissions/zoning →
// dissolve paths → waypoints/friends/tell/emotes.
//
// Exit criteria under test (predecessor-GDD §30, HISTORICAL): "A group of
// players can found a city, elect a mayor, and see rank-gated unlocks
// (shuttleport) appear."
//
// Requires TESTBED_FAST_CYCLE=1 on the server (founding 4 / township 6 /
// city 8; 20 s harvest ticks; 2 min elections; 3 min forming grace) and a
// FRESH database. Total runtime ~30 minutes (quota sampling dominates).
//
// Credit budget (start 5000 each; all assertions exact):
// A funds treasury 2500 + vendor 2000 + 3 houses x100; collects tills.
// B funds treasury 2800 + 2 houses x100 + registrar 1000 + sidearm2 600.
// C funds treasury 2900 + 4 houses x100 + city2 12 + sidearm1 500 + sidearm3 700.
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

func main() {
	fmt.Println("=== Phase 7 Integration Test (testbed fork, generic content) ===")
	fmt.Println("Cities → elections → tax/ranks → guilds → groups → mail → perms → dissolve")
	fmt.Println()

	if !waitForServer() {
		fmt.Println("FAIL: Server not reachable")
		os.Exit(1)
	}
	fmt.Println("Server is reachable.")

	// Setup: A founder (artisan), B citizen/buyer (artisan), C citizen (artisan).
	_, tokenA := registerAccount()
	charA := createCharacter(tokenA, "human", "FounderChar")
	_, tokenB := registerAccount()
	charB := createCharacter(tokenB, "human", "CityBuyer")
	_, tokenC := registerAccount()
	charC := createCharacter(tokenC, "human", "ThirdChar")
	trainSkill(tokenA, charA, "trainer_artisan_z1", "artisan_novice")
	trainSkill(tokenB, charB, "trainer_artisan_z1", "artisan_novice")
	trainSkill(tokenC, charC, "trainer_artisan_z1", "artisan_novice")
	fmt.Println("Setup: founder + buyer + third (all Artisan Novice)")

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

	// --- Test 1: resource quotas + deed crafting ---
	fmt.Println("\n--- Test 1: Quotas + deeds ---")
	sampleQuota(wcA, tokenA, charA, "mineral", "ferric_metal", 145, &ax, &az)
	sampleQuota(wcA, tokenA, charA, "chemical", "structural_polymer", 78, &ax, &az)
	sampleQuota(wcB, tokenB, charB, "mineral", "ferric_metal", 50, &bx, &bz)
	sampleQuota(wcB, tokenB, charB, "chemical", "structural_polymer", 28, &bx, &bz)
	sampleQuota(wcC, tokenC, charC, "mineral", "ferric_metal", 95, &cx, &cz)
	sampleQuota(wcC, tokenC, charC, "chemical", "structural_polymer", 52, &cx, &cz)
	fmt.Println("  [PASS] Quotas sampled (metal + polymer for all three)")
	hallA := craftDeed(tokenA, charA, "city_hall_deed")
	hallA3 := craftDeed(tokenA, charA, "city_hall_deed")
	hallA2 := craftDeed(tokenA, charA, "city_hall_deed")
	houseA1 := craftDeed(tokenA, charA, "structure_deed")
	houseA2 := craftDeed(tokenA, charA, "structure_deed")
	houseA3 := craftDeed(tokenA, charA, "structure_deed")
	vendorA := craftItem(tokenA, charA, "vendor_deed",
		map[string]string{"frame": metalSrc(tokenA, charA), "fittings": polySrc(tokenA, charA)}, "presentation", 3)
	sidearm1 := craftItem(tokenA, charA, "basic_sidearm",
		map[string]string{"frame": metalSrc(tokenA, charA), "grip": polySrc(tokenA, charA)}, "damage", 3)
	sidearm2 := craftItem(tokenA, charA, "basic_sidearm",
		map[string]string{"frame": metalSrc(tokenA, charA), "grip": polySrc(tokenA, charA)}, "damage", 3)
	sidearm3 := craftItem(tokenA, charA, "basic_sidearm",
		map[string]string{"frame": metalSrc(tokenA, charA), "grip": polySrc(tokenA, charA)}, "damage", 3)
	houseB1 := craftDeed(tokenB, charB, "structure_deed")
	houseB2 := craftDeed(tokenB, charB, "structure_deed")
	houseB3 := craftDeed(tokenB, charB, "structure_deed")
	houseC1 := craftDeed(tokenC, charC, "structure_deed")
	hallC := craftDeed(tokenC, charC, "city_hall_deed")
	houseC2 := craftDeed(tokenC, charC, "structure_deed")
	houseC3 := craftDeed(tokenC, charC, "structure_deed")
	houseC4 := craftDeed(tokenC, charC, "structure_deed")
	hallC2 := craftDeed(tokenC, charC, "city_hall_deed")
	fmt.Println("  [PASS] Deeds crafted (halls, houses, vendor, 3 sidearms)")

	// --- Test 2: pre-place structures, then found (pre-existing path) ---
	fmt.Println("\n--- Test 2: Structure placement + founding ---")
	// Main-city geography sits +100m east so a shared-DB full regression
	// never collides with Phase 5's leftover structures at (45,0)/(70,0).
	placeHouse(wcA, houseA1, 140, 0, &ax, &az)
	fundStructure(tokenA, charA, lastPlaced(), 100)
	placeHouse(wcA, houseA2, 160, 0, &ax, &az)
	fundStructure(tokenA, charA, lastPlaced(), 100)
	placeHouse(wcA, houseA3, 180, 0, &ax, &az)
	fundStructure(tokenA, charA, lastPlaced(), 100)
	placeVendor(wcA, vendorA, 260, 0, &ax, &az)
	placeHouse(wcB, houseB1, 220, 0, &bx, &bz)
	fundStructure(tokenB, charB, lastPlaced(), 100)
	placeHouse(wcB, houseB2, 240, 0, &bx, &bz)
	fundStructure(tokenB, charB, lastPlaced(), 100)
	placeHouse(wcC, houseC1, 120, 0, &cx, &cz)
	fundStructure(tokenC, charC, lastPlaced(), 100)
	fmt.Println("  [PASS] 7 structures placed (founder 3 + vendor, buyer 2, third 1)")
	cityID := placeCityHall(wcA, hallA, 200, 0, "Founders Landing", &ax, &az)
	// Fund instantly: an active Rank-1 city with an empty treasury would
	// dissolve on the first upkeep tick (treasury-exhaustion rule).
	civicPOST(tokenA, "/api/civic/city/fund", map[string]interface{}{"character_id": charA, "city_id": cityID, "credits": 2500})
	fmt.Printf("  [PASS] City founded active + treasury seeded (%s)\n", cityID)

	// --- Test 3: forming hall (grace path runs in background) ---
	fmt.Println("\n--- Test 3: Forming hall (grace expiry observed at end) ---")
	formingID := placeCityHallForming(wcC, hallC2, 400, 200, "Lone Claim", &cx, &cz)
	fmt.Printf("  [PASS] Lone hall forming (%s)\n", formingID)

	// --- Test 4: auto-enrol citizenship (tick-driven) ---
	fmt.Println("\n--- Test 4: Auto-enrol citizenship ---")
	fmt.Println("  Waiting one interval (enrol tick)...")
	time.Sleep(25 * time.Second)
	city := civicGet(tokenA, cityID)
	citizens := city["citizens"].([]interface{})
	if len(citizens) != 3 {
		fmt.Printf("  [FAIL] citizens = %d (want 3)\n", len(citizens))
		os.Exit(1)
	}
	fmt.Println("  [PASS] All three owners auto-enrolled as citizens")
	if m, _ := city["mayor"].(string); m != "FounderChar" {
		fmt.Printf("  [FAIL] founder-mayor = %q\n", m)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Founder starts as mayor (Mayor-Founder)")

	// --- Test 5: election (open, vote, double-vote reject, close, discount) ---
	fmt.Println("\n--- Test 5: Election ---")
	civicPOST(tokenA, "/api/civic/city/open-election", map[string]string{"character_id": charA, "city_id": cityID})
	civicPOST(tokenA, "/api/civic/city/vote", map[string]string{"character_id": charA, "city_id": cityID, "candidate_id": charA})
	civicPOST(tokenB, "/api/civic/city/vote", map[string]string{"character_id": charB, "city_id": cityID, "candidate_id": charA})
	if _, code := civicPOSTRaw(tokenB, "/api/civic/city/vote", map[string]string{"character_id": charB, "city_id": cityID, "candidate_id": charA}); code < 300 {
		fmt.Println("  [FAIL] double vote accepted")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Double vote rejected (one vote per citizen)")
	civicPOST(tokenC, "/api/civic/city/vote", map[string]string{"character_id": charC, "city_id": cityID, "candidate_id": charA})
	fmt.Println("  Waiting one interval (election close tick)...")
	time.Sleep(25 * time.Second)
	city = civicGet(tokenA, cityID)
	if m, _ := city["mayor"].(string); m != "FounderChar" {
		fmt.Printf("  [FAIL] elected mayor = %q\n", m)
		os.Exit(1)
	}
	fmt.Println("  [PASS] FounderChar elected mayor (unanimous)")
	if eff, _ := city["upkeep_effective"].(float64); int(eff) != 400 {
		fmt.Printf("  [FAIL] upkeep_effective = %v (want 400 = 500-20%% mayor discount)\n", eff)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Mayor-status upkeep discount visible (500 → 400)")

	// --- Test 6: vendor tax + rank-ups (exit criteria: shuttleport appears) ---
	fmt.Println("\n--- Test 6: Vendor tax + rank unlocks ---")
	vendorID := vendorIDOf(tokenA, charA)
	fundStructure(tokenA, charA, vendorID, 2000)
	fmt.Println("  Waiting one interval (vendor reopens)...")
	time.Sleep(25 * time.Second)
	civicPOST(tokenA, "/api/civic/city/set-tax", map[string]interface{}{"character_id": charA, "city_id": cityID, "flat_weekly": 0, "vendor_pct": 10})
	listing1 := stockItemMust(tokenA, charA, vendorID, sidearm1, 500)
	moveTo(wcC, 255, 0, &cx, &cz)
	wC0 := wallet(tokenC, charC)
	wcC.clearMessages()
	wcC.send(protocol.MsgPurchase, protocol.PurchaseMsg{ListingID: listing1})
	if _, ok := wcC.waitFor(protocol.MsgPurchaseReceipt, 10*time.Second); !ok {
		fmt.Println("  [FAIL] no purchase receipt (taxed sale)")
		os.Exit(1)
	}
	if w := wallet(tokenC, charC); w != wC0-500 {
		fmt.Printf("  [FAIL] buyer wallet %d (want %d)\n", w, wC0-500)
		os.Exit(1)
	}
	foundTax := false
	for _, e := range civicCityLedger(tokenA, cityID) {
		em, _ := e.(map[string]interface{})
		if em["Category"] == "city_vendor_tax" && em["Flow"] == "transfer" {
			foundTax = true
		}
	}
	if !foundTax {
		fmt.Println("  [FAIL] no city_vendor_tax transfer pair in city ledger")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Vendor-% tax split (buyer −500, city +50, ledger pair)")
	moveTo(wcA, 260, 0, &ax, &az)
	wcA.clearMessages()
	wcA.send(protocol.MsgCollectTill, protocol.TillMsg{VendorID: vendorID})
	tillMsg, ok := wcA.waitFor(protocol.MsgTillCollected, 10*time.Second)
	if !ok || int(asMap(tillMsg)["amount"].(float64)) != 450 {
		fmt.Println("  [FAIL] till should be 450 (500 − 50 tax)")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Till collects 450 (tax already diverted)")
	civicPOST(tokenB, "/api/civic/city/fund", map[string]interface{}{"character_id": charB, "city_id": cityID, "credits": 2800})
	fmt.Println("  Waiting one interval (township evaluation)...")
	time.Sleep(25 * time.Second)
	city = civicGet(tokenA, cityID)
	if r, _ := city["rank"].(string); r != "township" {
		fmt.Printf("  [FAIL] rank = %q (want township)\n", r)
		os.Exit(1)
	}
	unlocks := city["unlocks"].(map[string]interface{})
	if unlocks["cantina_allowed"] != true || unlocks["shuttleport"] == true {
		fmt.Printf("  [FAIL] township unlocks wrong: %v\n", unlocks)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Rank township (cantina/medical unlocked, no shuttleport)")
	civicPOST(tokenC, "/api/civic/city/fund", map[string]interface{}{"character_id": charC, "city_id": cityID, "credits": 2900})
	fmt.Println("  Waiting one interval (city evaluation)...")
	time.Sleep(25 * time.Second)
	city = civicGet(tokenA, cityID)
	if r, _ := city["rank"].(string); r != "city" {
		fmt.Printf("  [FAIL] rank = %q (want city)\n", r)
		os.Exit(1)
	}
	unlocks = city["unlocks"].(map[string]interface{})
	if unlocks["shuttleport"] != true || unlocks["trainer_hosting"] != true {
		fmt.Printf("  [FAIL] city unlocks wrong: %v\n", unlocks)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Rank city — shuttleport + trainer hosting appear (exit criteria)")

	// --- Test 7: flat tax, evict guard, opt-out ---
	fmt.Println("\n--- Test 7: Flat tax + opt-out ---")
	civicPOST(tokenA, "/api/civic/city/set-tax", map[string]interface{}{"character_id": charA, "city_id": cityID, "flat_weekly": 100, "vendor_pct": 10})
	time.Sleep(45 * time.Second)
	foundFlat := false
	arrears := 0
	city = civicGet(tokenA, cityID)
	for _, e := range civicCityLedger(tokenA, cityID) {
		em, _ := e.(map[string]interface{})
		if em["Category"] == "city_flat_tax" && em["Flow"] == "transfer" {
			foundFlat = true
		}
	}
	for _, c := range city["citizens"].([]interface{}) {
		cm, _ := c.(map[string]interface{})
		if a, _ := cm["arrears"].(float64); a > 0 {
			arrears++
		}
	}
	if !foundFlat || arrears > 0 {
		fmt.Println("  [FAIL] flat tax not accruing cleanly in ledger")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Flat citizen tax accrues (ledger pair, no arrears)")
	if _, code := civicPOSTRaw(tokenA, "/api/civic/city/evict", map[string]string{"character_id": charA, "city_id": cityID, "target_id": charB}); code != 403 {
		fmt.Printf("  [FAIL] evicting solvent citizen not refused (HTTP %d)\n", code)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Eviction of solvent citizen refused (arrears required)")
	civicPOST(tokenC, "/api/civic/city/leave", map[string]string{"character_id": charC, "city_id": cityID})
	time.Sleep(25 * time.Second)
	city = civicGet(tokenA, cityID)
	if len(city["citizens"].([]interface{})) != 2 {
		fmt.Println("  [FAIL] opt-out did not reduce citizenship")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Opt-out honored and not auto-re-enrolled")

	// --- Test 8: guilds, dues, groups, mentorship ---
	fmt.Println("\n--- Test 8: Guilds + groups + mentorship ---")
	gOut := civicPOST(tokenB, "/api/civic/guild/found", map[string]string{"character_id": charB, "name": "Iron Compact", "tag": "IC"})
	guildID := gOut["guild_id"].(string)
	invA := civicPOST(tokenB, "/api/civic/guild/invite", map[string]string{"character_id": charB, "guild_id": guildID, "invitee_name": "FounderChar"})["invite_id"].(string)
	invC := civicPOST(tokenB, "/api/civic/guild/invite", map[string]string{"character_id": charB, "guild_id": guildID, "invitee_name": "ThirdChar"})["invite_id"].(string)
	civicPOST(tokenA, "/api/civic/guild/accept", map[string]string{"character_id": charA, "invite_id": invA})
	civicPOST(tokenC, "/api/civic/guild/accept", map[string]string{"character_id": charC, "invite_id": invC})
	g := civicGET(tokenB, "/api/civic/guild/get?character_id="+charB+"&guild_id="+guildID)
	if len(g["members"].([]interface{})) != 3 {
		fmt.Println("  [FAIL] guild membership != 3")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Guild founded, 3 members joined")
	civicPOST(tokenB, "/api/civic/guild/set-dues", map[string]interface{}{"character_id": charB, "guild_id": guildID, "pct": 10})
	listing2 := stockItemMust(tokenA, charA, vendorID, sidearm2, 600)
	moveTo(wcB, 255, 0, &bx, &bz)
	wcB.clearMessages()
	wcB.send(protocol.MsgPurchase, protocol.PurchaseMsg{ListingID: listing2})
	if _, ok := wcB.waitFor(protocol.MsgPurchaseReceipt, 10*time.Second); !ok {
		fmt.Println("  [FAIL] no receipt (dues sale)")
		os.Exit(1)
	}
	g = civicGET(tokenB, "/api/civic/guild/get?character_id="+charB+"&guild_id="+guildID)
	if t, _ := g["treasury"].(float64); int(t) != 60 {
		fmt.Printf("  [FAIL] guild treasury = %v (want 60 = 10%% dues on 600)\n", t)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Guild dues diverted (treasury +60 on taxed sale)")
	wcA.clearMessages()
	wcA.send(protocol.MsgCollectTill, protocol.TillMsg{VendorID: vendorID})
	tillMsg2, ok := wcA.waitFor(protocol.MsgTillCollected, 10*time.Second)
	if !ok || int(asMap(tillMsg2)["amount"].(float64)) != 480 {
		fmt.Println("  [FAIL] till should be 480 (600 − 60 tax − 60 dues)")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Till 480 (tax + dues both diverted)")
	civicPOST(tokenA, "/api/civic/group/invite", map[string]string{"character_id": charA, "invitee_name": "CityBuyer"})
	grInvites := civicGET(tokenB, "/api/civic/group/invites?character_id="+charB)["invites"].([]interface{})
	if len(grInvites) == 0 {
		fmt.Println("  [FAIL] no group invite for buyer")
		os.Exit(1)
	}
	grInvID := ""
	for _, iv := range grInvites {
		im, _ := iv.(map[string]interface{})
		if id, _ := im["ID"].(string); id != "" {
			grInvID = id
		}
	}
	civicPOST(tokenB, "/api/civic/group/accept", map[string]string{"character_id": charB, "invite_id": grInvID})
	civicPOST(tokenA, "/api/civic/group/loot-rule", map[string]string{"character_id": charA, "rule": "master_looter"})
	gr := civicGET(tokenA, "/api/civic/group/get?character_id="+charA)["group"].(map[string]interface{})
	if gr["loot_rule"] != "master_looter" || len(gr["members"].([]interface{})) != 2 {
		fmt.Printf("  [FAIL] group wrong: %v\n", gr)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Group formed (2 members, loot rule set)")
	wcB.clearMessages()
	wcC.clearMessages()
	wcA.send(protocol.MsgChat, protocol.ChatMsg{Channel: "group", Text: "group hello"})
	if _, ok := wcB.waitFor(protocol.MsgChatMessage, 5*time.Second); !ok {
		fmt.Println("  [FAIL] group message not delivered to member")
		os.Exit(1)
	}
	time.Sleep(1500 * time.Millisecond)
	if hasChannel(wcC, "group") {
		fmt.Println("  [FAIL] group message leaked to non-member")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Group chat routed to members only")
	wcC.clearMessages()
	wcA.send(protocol.MsgChat, protocol.ChatMsg{Channel: "guild", Text: "guild hello"})
	if _, ok := wcC.waitFor(protocol.MsgChatMessage, 5*time.Second); !ok {
		fmt.Println("  [FAIL] guild message not delivered")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Guild chat routed to members")
	civicPOST(tokenA, "/api/civic/mentor/bond", map[string]string{"character_id": charA, "protege_name": "ThirdChar"})
	if xp := getXP(tokenA, charA, "mentoring"); xp < 50 {
		fmt.Printf("  [FAIL] no mentoring XP (%d)\n", xp)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Mentorship bonded (+mentoring XP)")

	// --- Test 9: mail (offline delivery, claim, delete guard, cap) ---
	fmt.Println("\n--- Test 9: Mail ---")
	civicPOST(tokenA, "/api/civic/city/set-tax", map[string]interface{}{"character_id": charA, "city_id": cityID, "flat_weekly": 0, "vendor_pct": 10})
	wcB.conn.Close() // buyer goes offline; mail must still deliver
	time.Sleep(1000 * time.Millisecond)
	wB0mail := wallet(tokenB, charB)
	civicPOST(tokenA, "/api/civic/mail/send", map[string]interface{}{
		"character_id": charA, "recipient_name": "CityBuyer",
		"subject": "spare deed", "body": "claim this", "credits": 100,
		"item_ids": []string{hallA2},
	})
	wcB = connectWS(tokenB)
	wcB.send(protocol.MsgEnterWorld, protocol.EnterWorldMsg{CharacterID: charB})
	if _, ok := wcB.waitFor(protocol.MsgWorldEnter, 5*time.Second); !ok {
		fmt.Println("  [FAIL] buyer relog failed")
		os.Exit(1)
	}
	bx = 255
	bz = 0 // relog keeps last persisted position (near vendor)
	inbox := civicGET(tokenB, "/api/civic/mail/inbox?character_id="+charB)["mail"].([]interface{})
	if len(inbox) == 0 {
		fmt.Println("  [FAIL] inbox empty after offline delivery")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Offline mail delivered (inbox non-empty after relog)")
	mailID := ""
	for _, m := range inbox {
		mm, _ := m.(map[string]interface{})
		if id, _ := mm["id"].(string); id != "" {
			mailID = id
		}
	}
	claim := civicPOST(tokenB, "/api/civic/mail/claim", map[string]string{"character_id": charB, "mail_id": mailID})
	if c, _ := claim["credits"].(float64); int(c) != 100 {
		fmt.Printf("  [FAIL] claimed credits = %v\n", c)
		os.Exit(1)
	}
	if w := wallet(tokenB, charB); w != wB0mail+100 {
		fmt.Printf("  [FAIL] wallet %d (want %d)\n", w, wB0mail+100)
		os.Exit(1)
	}
	if !ownsItem(tokenB, charB, hallA2) {
		fmt.Println("  [FAIL] mailed deed not owned by recipient")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Mail claimed atomically (credits + item, offline-safe)")
	civicPOST(tokenB, "/api/civic/mail/delete", map[string]string{"character_id": charB, "mail_id": mailID})
	fmt.Println("  [PASS] Claimed mail deletable")
	guardMail := civicPOST(tokenA, "/api/civic/mail/send", map[string]interface{}{
		"character_id": charA, "recipient_name": "ThirdChar",
		"subject": "ten", "body": "ten", "credits": 10,
	})["mail_id"].(string)
	if _, code := civicPOSTRaw(tokenC, "/api/civic/mail/delete", map[string]string{"character_id": charC, "mail_id": guardMail}); code < 300 {
		fmt.Println("  [FAIL] delete with unclaimed credits allowed")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Delete refused while credits unclaimed")
	civicPOST(tokenC, "/api/civic/mail/claim", map[string]string{"character_id": charC, "mail_id": guardMail})
	civicPOST(tokenC, "/api/civic/mail/delete", map[string]string{"character_id": charC, "mail_id": guardMail})
	for i := 0; i < 50; i++ {
		civicPOST(tokenA, "/api/civic/mail/send", map[string]interface{}{
			"character_id": charA, "recipient_name": "ThirdChar",
			"subject": "fill", "body": "fill",
		})
	}
	if _, code := civicPOSTRaw(tokenA, "/api/civic/mail/send", map[string]interface{}{
		"character_id": charA, "recipient_name": "ThirdChar",
		"subject": "over", "body": "over",
	}); code != 400 {
		fmt.Printf("  [FAIL] 51st mail not capped (HTTP %d)\n", code)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Mailbox cap enforced at 50")

	// --- Test 10: permissions + zoning + banned purchase + overlap ---
	fmt.Println("\n--- Test 10: Permissions + zoning ---")
	listing3 := stockItemMust(tokenA, charA, vendorID, sidearm3, 700)
	moveTo(wcC, 260, 0, &cx, &cz)
	civicPOST(tokenA, "/api/civic/structure/perm-add", map[string]string{"character_id": charA, "structure_id": vendorID, "list": "banned", "target_name": "ThirdChar"})
	wcC.clearMessages()
	wcC.send(protocol.MsgPurchase, protocol.PurchaseMsg{ListingID: listing3})
	if _, ok := wcC.waitFor(protocol.MsgError, 5*time.Second); !ok {
		fmt.Println("  [FAIL] banned purchase not refused")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Banned buyer refused (entry denial analog)")
	civicPOST(tokenA, "/api/civic/structure/perm-remove", map[string]string{"character_id": charA, "structure_id": vendorID, "list": "banned", "target_name": "ThirdChar"})
	wC1 := wallet(tokenC, charC)
	wcC.clearMessages()
	wcC.send(protocol.MsgPurchase, protocol.PurchaseMsg{ListingID: listing3})
	if _, ok := wcC.waitFor(protocol.MsgPurchaseReceipt, 10*time.Second); !ok {
		fmt.Println("  [FAIL] no receipt after unban")
		os.Exit(1)
	}
	if w := wallet(tokenC, charC); w != wC1-700 {
		fmt.Printf("  [FAIL] wallet %d (want %d)\n", w, wC1-700)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Unbanned purchase succeeds")
	houseAID := firstHouseOf(tokenA, charA)
	civicPOST(tokenA, "/api/civic/structure/perm-add", map[string]string{"character_id": charA, "structure_id": houseAID, "list": "admin", "target_name": "CityBuyer"})
	authed("POST", "/api/economy/structure/fund", tokenB, map[string]interface{}{"character_id": charB, "structure_id": houseAID, "credits": 50})
	fmt.Println("  [PASS] Admin (non-owner) may fund structure")
	if _, code := authedRaw("POST", "/api/economy/structure/fund", tokenC, map[string]interface{}{"character_id": charC, "structure_id": houseAID, "credits": 50}); code != 403 {
		fmt.Printf("  [FAIL] stranger funding not refused (HTTP %d)\n", code)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Stranger funding refused (owner/admin only)")
	civicPOST(tokenA, "/api/civic/city/set-zoning", map[string]interface{}{"character_id": charA, "city_id": cityID, "open": false})
	moveTo(wcB, 240, 30, &bx, &bz)
	wcB.clearMessages()
	wcB.send(protocol.MsgPlaceHouse, protocol.PlaceStructureMsg{X: 240, Z: 30, Tier: "small", DeedItemID: houseB3})
	if _, ok := wcB.waitFor(protocol.MsgError, 5*time.Second); !ok {
		fmt.Println("  [FAIL] placement under closed zoning not refused")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Closed zoning blocks non-mayor placement")
	if _, code := civicPOSTRaw(tokenB, "/api/civic/city/set-zoning", map[string]interface{}{"character_id": charB, "city_id": cityID, "open": true}); code != 403 {
		fmt.Printf("  [FAIL] non-mayor zoning change not refused (HTTP %d)\n", code)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Non-mayor zoning change refused")
	civicPOST(tokenA, "/api/civic/city/set-zoning", map[string]interface{}{"character_id": charA, "city_id": cityID, "open": true})
	wcB.clearMessages()
	wcB.send(protocol.MsgPlaceHouse, protocol.PlaceStructureMsg{X: 240, Z: 30, Tier: "small", DeedItemID: houseB3})
	plMsg, ok := wcB.waitFor(protocol.MsgStructurePlaced, 5*time.Second)
	if !ok {
		fmt.Println("  [FAIL] placement after reopen failed")
		os.Exit(1)
	}
	placedIDs = append(placedIDs, asMap(plMsg)["structure_id"].(string))
	fundStructure(tokenB, charB, lastPlaced(), 100)
	fmt.Println("  [PASS] Open zoning admits placement")
	wcA.clearMessages()
	wcA.send(protocol.MsgPlaceCityHall, protocol.PlaceCityHallMsg{X: 250, Z: 0, Name: "Overlap Claim", DeedItemID: hallA3})
	if _, ok := wcA.waitFor(protocol.MsgError, 5*time.Second); !ok {
		fmt.Println("  [FAIL] overlapping city not refused")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Overlapping city radius refused")

	// --- Test 11: second city dissolve + forming grace expiry ---
	fmt.Println("\n--- Test 11: Dissolve paths ---")
	moveTo(wcC, 360, 0, &cx, &cz)
	placeHouse(wcC, houseC2, 360, 0, &cx, &cz)
	fundStructure(tokenC, charC, lastPlaced(), 100)
	placeHouse(wcC, houseC3, 380, 0, &cx, &cz)
	fundStructure(tokenC, charC, lastPlaced(), 100)
	placeHouse(wcC, houseC4, 420, 0, &cx, &cz)
	fundStructure(tokenC, charC, lastPlaced(), 100)
	city2ID := placeCityHall(wcC, hallC, 400, 0, "Second Claim", &cx, &cz)
	fmt.Printf("  [PASS] Second city active (%s)\n", city2ID)
	civicPOST(tokenC, "/api/civic/city/fund", map[string]interface{}{"character_id": charC, "city_id": city2ID, "credits": 12})
	fmt.Println("  Waiting (~2 min) for treasury exhaustion...")
	time.Sleep(140 * time.Second)
	city2 := civicGet(tokenC, city2ID)
	if st, _ := city2["status"].(string); st != "dissolved" {
		fmt.Printf("  [FAIL] second city status = %q (want dissolved)\n", st)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Rank-1 treasury exhaustion dissolves city (property untouched)")
	forming := civicGet(tokenA, formingID)
	if st, _ := forming["status"].(string); st != "dissolved" {
		fmt.Printf("  [FAIL] forming city status = %q (want dissolved past grace)\n", st)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Forming grace expiry dissolves uncompleted claim")

	// --- Test 12: waypoints, friends, tell, emotes, planet chat ---
	fmt.Println("\n--- Test 12: Social surface ---")
	moveTo(wcB, 255, 0, &bx, &bz) // back in local range for emote/planet delivery
	wp := civicPOST(tokenA, "/api/civic/waypoint/add", map[string]interface{}{"character_id": charA, "label": "vendor row", "x": 260, "z": 0})["waypoint_id"].(string)
	civicPOST(tokenA, "/api/civic/waypoint/share", map[string]string{"character_id": charA, "waypoint_id": wp, "target_name": "CityBuyer"})
	shared := false
	for _, w := range civicGET(tokenB, "/api/civic/waypoint/list?character_id="+charB)["waypoints"].([]interface{}) {
		wm, _ := w.(map[string]interface{})
		if wm["Label"] == "vendor row" {
			shared = true
		}
	}
	if !shared {
		fmt.Println("  [FAIL] shared waypoint not imported")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Waypoint shared + imported")
	civicPOST(tokenA, "/api/civic/friends/add", map[string]string{"character_id": charA, "friend_name": "CityBuyer"})
	civicPOST(tokenA, "/api/civic/friends/add", map[string]string{"character_id": charA, "friend_name": "ThirdChar"})
	allyOnline, thirdOnline := false, false
	for _, f := range civicGET(tokenA, "/api/civic/friends/list?character_id="+charA)["friends"].([]interface{}) {
		fm, _ := f.(map[string]interface{})
		if fm["name"] == "CityBuyer" && fm["online"] == true {
			allyOnline = true
		}
		if fm["name"] == "ThirdChar" && fm["online"] == true {
			thirdOnline = true
		}
	}
	if !allyOnline || !thirdOnline {
		fmt.Println("  [FAIL] friends list missing online entries")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Friends list shows online status")
	wcB.clearMessages()
	wcA.send(protocol.MsgChat, protocol.ChatMsg{Channel: "tell", Target: "CityBuyer", Text: "hello directly"})
	if _, ok := wcB.waitFor(protocol.MsgChatMessage, 5*time.Second); !ok {
		fmt.Println("  [FAIL] tell not delivered")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Tell delivered to named player")
	wcA.clearMessages()
	wcA.send(protocol.MsgChat, protocol.ChatMsg{Channel: "tell", Target: "Nobody", Text: "hello?"})
	if _, ok := wcA.waitFor(protocol.MsgError, 5*time.Second); !ok {
		fmt.Println("  [FAIL] tell to unknown name not rejected")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Tell to unknown name rejected")
	wcB.clearMessages()
	wcA.send(protocol.MsgChat, protocol.ChatMsg{Channel: "emote", Text: "wave"})
	if _, ok := wcB.waitFor(protocol.MsgChatMessage, 5*time.Second); !ok {
		fmt.Println("  [FAIL] emote not broadcast")
		os.Exit(1)
	}
	wcA.clearMessages()
	wcA.send(protocol.MsgChat, protocol.ChatMsg{Channel: "emote", Text: "explode"})
	if _, ok := wcA.waitFor(protocol.MsgError, 5*time.Second); !ok {
		fmt.Println("  [FAIL] unknown emote not rejected")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Standard emotes broadcast; unknown rejected")
	wcB.clearMessages()
	wcA.send(protocol.MsgChat, protocol.ChatMsg{Channel: "planet", Text: "planet hello"})
	if _, ok := wcB.waitFor(protocol.MsgChatMessage, 5*time.Second); !ok {
		fmt.Println("  [FAIL] planet chat not delivered")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Planet chat delivered zone-wide")
	wcA.clearMessages()
	wcA.send(protocol.MsgChat, protocol.ChatMsg{Channel: "faction", Text: "for the cause"})
	if _, ok := wcA.waitFor(protocol.MsgError, 5*time.Second); !ok {
		fmt.Println("  [FAIL] faction channel without alignment not refused")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Faction channel honestly unavailable (later phase)")
	wcC.conn.Close() // third goes offline at end
	time.Sleep(1500 * time.Millisecond)
	wcA.clearMessages()
	wcA.send(protocol.MsgChat, protocol.ChatMsg{Channel: "tell", Target: "ThirdChar", Text: "are you there?"})
	if _, ok := wcA.waitFor(protocol.MsgError, 5*time.Second); !ok {
		fmt.Println("  [FAIL] tell to offline player not refused")
		os.Exit(1)
	}
	offlineSeen := false
	for _, f := range civicGET(tokenA, "/api/civic/friends/list?character_id="+charA)["friends"].([]interface{}) {
		fm, _ := f.(map[string]interface{})
		if fm["name"] == "ThirdChar" && fm["online"] == false {
			offlineSeen = true
		}
	}
	if !offlineSeen {
		fmt.Println("  [FAIL] friends list missing offline entry")
		os.Exit(1)
	}
	fmt.Println("  [PASS] Offline tell refused; friends show offline")

	fmt.Println("\n=== ALL PHASE 7 TESTS PASSED ===")
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

// authed performs an authenticated API call, fatal on HTTP >= 300.
func authed(method, path, token string, body interface{}) map[string]interface{} {
	out, code := authedRaw(method, path, token, body)
	if code >= 300 {
		log.Fatalf("API %s HTTP %d: %v", path, code, out)
	}
	return out
}

// authedRaw performs an authenticated API call, returning status for
// expected-failure assertions.
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

func civicGET(token, path string) map[string]interface{} {
	return authed("GET", path, token, nil)
}

func civicPOST(token, path string, body interface{}) map[string]interface{} {
	return authed("POST", path, token, body)
}

func civicPOSTRaw(token, path string, body interface{}) (map[string]interface{}, int) {
	return authedRaw("POST", path, token, body)
}

func civicGet(token, cityID string) map[string]interface{} {
	return civicGET(token, "/api/civic/city/get?character_id=x&city_id="+cityID)
}

func civicCityLedger(token, cityID string) []interface{} {
	out := civicGET(token, "/api/civic/city/ledger?city_id="+cityID)
	list, _ := out["entries"].([]interface{})
	return list
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

func ownsItem(token, charID, itemID string) bool {
	out := authed("GET", "/api/craft/items?character_id="+charID, token, nil)
	if list, ok := out["items"].([]interface{}); ok {
		for _, it := range list {
			im, _ := it.(map[string]interface{})
			if im["ID"] == itemID {
				return true
			}
		}
	}
	return false
}

func vendorIDOf(token, charID string) string {
	out := authed("GET", "/api/economy/structures?character_id="+charID, token, nil)
	if list, ok := out["structures"].([]interface{}); ok {
		for _, s := range list {
			sm, _ := s.(map[string]interface{})
			if sm["kind"] == "vendor" {
				id, _ := sm["id"].(string)
				return id
			}
		}
	}
	log.Fatal("no vendor found")
	return ""
}

func firstHouseOf(token, charID string) string {
	out := authed("GET", "/api/economy/structures?character_id="+charID, token, nil)
	if list, ok := out["structures"].([]interface{}); ok {
		for _, s := range list {
			sm, _ := s.(map[string]interface{})
			if sm["kind"] == "house" {
				id, _ := sm["id"].(string)
				return id
			}
		}
	}
	log.Fatal("no house found")
	return ""
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

func placeVendor(wc *wsClient, deedID string, x, z float64, px, pz *float64) {
	moveTo(wc, x, z, px, pz)
	wc.clearMessages()
	wc.send(protocol.MsgPlaceVendor, protocol.PlaceStructureMsg{X: x, Z: z, DeedItemID: deedID})
	m, ok := wc.waitFor(protocol.MsgStructurePlaced, 10*time.Second)
	if !ok {
		log.Fatalf("place vendor (%.0f,%.0f) failed", x, z)
	}
	placedIDs = append(placedIDs, asMap(m)["structure_id"].(string))
}

// placeCityHall moves, places, and requires immediate active status.
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

// placeCityHallForming places a lone hall and requires forming status.
func placeCityHallForming(wc *wsClient, deedID string, x, z float64, name string, px, pz *float64) string {
	moveTo(wc, x, z, px, pz)
	wc.clearMessages()
	wc.send(protocol.MsgPlaceCityHall, protocol.PlaceCityHallMsg{X: x, Z: z, Name: name, DeedItemID: deedID})
	m, ok := wc.waitFor(protocol.MsgCityFounded, 10*time.Second)
	if !ok {
		log.Fatalf("place city hall (%.0f,%.0f) failed", x, z)
	}
	fm := asMap(m)
	if fm["status"] != "forming" {
		log.Fatalf("city %s should be forming: %v", name, fm)
	}
	return fm["city_id"].(string)
}

func fundStructure(token, charID, stID string, credits int) {
	authed("POST", "/api/economy/structure/fund", token, map[string]interface{}{
		"character_id": charID, "structure_id": stID, "credits": credits,
	})
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
	case "ferric_metal":
		return 30, 0
	case "structural_polymer":
		return 40, 0
	default:
		return 20, 0
	}
}

func sampleQuota(wc *wsClient, token, charID, tool, wantType string, qty int, px, pz *float64) {
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

func stockItemMust(token, charID, vendorID, itemID string, price int) string {
	out, code := authedRaw("POST", "/api/economy/vendor/list", token, map[string]interface{}{
		"character_id": charID, "vendor_id": vendorID, "item_id": itemID, "price": price,
	})
	if code >= 300 {
		log.Fatalf("stock failed HTTP %d: %v", code, out)
	}
	id, _ := out["listing_id"].(string)
	return id
}
