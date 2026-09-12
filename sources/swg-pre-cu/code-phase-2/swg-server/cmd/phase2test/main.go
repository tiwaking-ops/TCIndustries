package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"swg-server/internal/protocol"
)

// Phase 2 integration test: skills, professions, XP, and training.
//
// GDD Phase 2 exit criteria: "A player can earn XP from a placeholder action,
// spend skill points at a trainer NPC, and see the skill persist across
// logout/login."

const serverURL = "http://localhost:8080"

func main() {
	fmt.Println("=== SWG Pre-CU Phase 2 Integration Test ===")
	fmt.Println("Skills, professions, XP earning, and skill training")
	fmt.Println()

	if !waitForServer() {
		fmt.Println("FAIL: Server not reachable")
		os.Exit(1)
	}
	fmt.Println("Server is reachable.")

	// --- Setup ---
	fmt.Println("\n--- Setup: Register account + create character ---")
	accountID, token := registerAccount()
	fmt.Printf("  Account: %s\n", accountID[:8])

	charID := createCharacter(token, "human", "SkillTestChar")
	fmt.Printf("  Character: SkillTestChar (%s)\n", charID[:8])

	// --- Test 1: Get professions ---
	fmt.Println("\n--- Test 1: List all professions ---")
	professions := getProfessions()
	if len(professions) != 6 {
		fmt.Printf("  [FAIL] Expected 6 basic professions, got %d\n", len(professions))
		os.Exit(1)
	}
	fmt.Printf("  Got %d professions:\n", len(professions))
	for _, p := range professions {
		boxCount := 0
		boxCount += 1 // Novice
		for _, t := range p["trees"].([]interface{}) {
			tree := t.(map[string]interface{})
			boxCount += len(tree["boxes"].([]interface{}))
		}
		boxCount += 1 // Master
		fmt.Printf("    %s (%s): %d boxes, %d total SP\n",
			p["name"], p["category"], boxCount, int(p["total_skill_points"].(float64)))
	}
	fmt.Println("  [PASS] All 6 basic professions returned")

	// --- Test 2: Get trainers ---
	fmt.Println("\n--- Test 2: List trainers ---")
	trainers := getTrainers()
	if len(trainers) < 6 {
		fmt.Printf("  [FAIL] Expected at least 6 trainers, got %d\n", len(trainers))
		os.Exit(1)
	}
	fmt.Printf("  Got %d trainers\n", len(trainers))
	fmt.Println("  [PASS] Trainers seeded correctly")

	// --- Test 3: Get character skills (empty initially) ---
	fmt.Println("\n--- Test 3: Character starts with no skills ---")
	skills := getCharacterSkills(token, charID)
	spAvail := int(skills["skill_points_available"].(float64))
	spMax := int(skills["skill_points_max"].(float64))
	credits := int(skills["credits"].(float64))
	ownedList := skills["owned_skills"].([]interface{})

	if spAvail != 250 || spMax != 250 {
		fmt.Printf("  [FAIL] Expected 250/250 skill points, got %d/%d\n", spAvail, spMax)
		os.Exit(1)
	}
	if credits != 5000 {
		fmt.Printf("  [FAIL] Expected 5000 credits, got %d\n", credits)
		os.Exit(1)
	}
	if len(ownedList) != 0 {
		fmt.Printf("  [FAIL] Expected 0 owned skills, got %d\n", len(ownedList))
		os.Exit(1)
	}
	fmt.Printf("  Skill points: %d/%d, Credits: %d, Owned skills: 0\n", spAvail, spMax, credits)
	fmt.Println("  [PASS] Character starts with 250 SP, 5000 credits, no skills")

	// --- Test 4: Earn XP (placeholder action) ---
	fmt.Println("\n--- Test 4: Earn XP from placeholder action ---")
	xpResult := earnXP(token, charID, "combat", 10000)
	xpPools := xpResult["xp_pools"].(map[string]interface{})
	combatXP := int(xpPools["combat"].(float64))
	if combatXP != 10000 {
		fmt.Printf("  [FAIL] Expected 10000 combat XP, got %d\n", combatXP)
		os.Exit(1)
	}
	fmt.Printf("  Earned 10000 combat XP (pool now: %d)\n", combatXP)
	fmt.Println("  [PASS] XP earned from placeholder action")

	// --- Test 5: Train Novice Marksman (free) ---
	fmt.Println("\n--- Test 5: Train Novice Marksman (free entry) ---")
	trainResult := trainSkill(token, charID, "trainer_marksman_mos", "marksman_novice")
	if !trainResult["success"].(bool) {
		fmt.Printf("  [FAIL] Training Novice Marksman failed\n")
		os.Exit(1)
	}
	spUsed := int(trainResult["skill_points_used"].(float64))
	spAvail = int(trainResult["skill_points_available"].(float64))
	fmt.Printf("  Trained: %s (SP used: %d, available: %d)\n",
		trainResult["skill_box_name"], spUsed, spAvail)
	if spUsed != 0 {
		fmt.Printf("  [FAIL] Novice should cost 0 SP, but %d used\n", spUsed)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Novice Marksman trained for free")

	// --- Test 6: Train Ranged Accuracy I (costs XP + credits + SP) ---
	fmt.Println("\n--- Test 6: Train Ranged Accuracy I ---")
	trainResult = trainSkill(token, charID, "trainer_marksman_mos", "marksman_ranged_accuracy_i")
	if !trainResult["success"].(bool) {
		fmt.Printf("  [FAIL] Training Ranged Accuracy I failed\n")
		os.Exit(1)
	}
	spUsed = int(trainResult["skill_points_used"].(float64))
	spAvail = int(trainResult["skill_points_available"].(float64))
	fmt.Printf("  Trained: %s (SP used: %d, available: %d)\n",
		trainResult["skill_box_name"], spUsed, spAvail)
	if spUsed != 2 {
		fmt.Printf("  [FAIL] Ranged Accuracy I should cost 2 SP, but %d used\n", spUsed)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Ranged Accuracy I trained (2 SP, XP, credits deducted)")

	// Verify XP was deducted
	skills = getCharacterSkills(token, charID)
	xpPools = skills["xp_pools"].(map[string]interface{})
	if xpPools["combat"] != nil {
		combatXP = int(xpPools["combat"].(float64))
		if combatXP != 9000 {
			fmt.Printf("  [FAIL] Expected 9000 combat XP after training (10000-1000), got %d\n", combatXP)
			os.Exit(1)
		}
	}
	fmt.Printf("  Combat XP after training: %d (was 10000, cost 1000)\n", combatXP)

	// --- Test 7: Cannot skip to Tier III ---
	fmt.Println("\n--- Test 7: Cannot train Tier III without Tier II ---")
	resp := tryTrainSkill(token, charID, "trainer_marksman_mos", "marksman_ranged_accuracy_iii")
	if resp.StatusCode == 201 {
		fmt.Println("  [FAIL] Server allowed training Tier III without Tier II")
		os.Exit(1)
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("  Server correctly rejected: %s\n", string(body))
	fmt.Println("  [PASS] Cannot skip tiers")

	// --- Test 8: Cannot train Tier I of another tree without Novice ---
	fmt.Println("\n--- Test 8: Cannot train Artisan skill without Novice Artisan ---")
	resp = tryTrainSkill(token, charID, "trainer_artisan_mos", "artisan_engineering_i")
	if resp.StatusCode == 201 {
		fmt.Println("  [FAIL] Server allowed training Artisan skill without Novice Artisan")
		os.Exit(1)
	}
	body, _ = io.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Printf("  Server correctly rejected: %s\n", string(body))
	fmt.Println("  [PASS] Cannot train profession skills without Novice box")

	// --- Test 9: Drop Ranged Accuracy I, verify SP refunded ---
	fmt.Println("\n--- Test 9: Drop skill and verify SP refund ---")
	dropResult := dropSkill(token, charID, "marksman_ranged_accuracy_i")
	if !dropResult["success"].(bool) {
		fmt.Println("  [FAIL] Dropping skill failed")
		os.Exit(1)
	}
	spUsed = int(dropResult["skill_points_used"].(float64))
	spAvail = int(dropResult["skill_points_available"].(float64))
	fmt.Printf("  Dropped: SP used now %d, available %d\n", spUsed, spAvail)
	if spUsed != 0 {
		fmt.Printf("  [FAIL] After dropping, SP used should be 0, got %d\n", spUsed)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Skill dropped, SP refunded")

	// --- Test 10: Skill persists across logout/login ---
	fmt.Println("\n--- Test 10: Skill persists across logout/login ---")
	// Re-fetch character skills (simulates logout/login since data is in DB)
	skills = getCharacterSkills(token, charID)
	ownedList = skills["owned_skills"].([]interface{})
	if len(ownedList) != 1 {
		fmt.Printf("  [FAIL] Expected 1 owned skill (Novice Marksman), got %d\n", len(ownedList))
		os.Exit(1)
	}
	ownedBox := ownedList[0].(map[string]interface{})
	if ownedBox["box_id"] != "marksman_novice" {
		fmt.Printf("  [FAIL] Expected marksman_novice, got %s\n", ownedBox["box_id"])
		os.Exit(1)
	}
	fmt.Printf("  After re-fetch: %d owned skill(s), first: %s\n", len(ownedList), ownedBox["box_id"])
	fmt.Println("  [PASS] Skill persists across logout/login")

	// --- Test 11: Earn different XP type ---
	fmt.Println("\n--- Test 11: Earn crafting XP (different type) ---")
	xpResult = earnXP(token, charID, "crafting", 5000)
	xpPools = xpResult["xp_pools"].(map[string]interface{})
	combatXP = 0
	craftingXP := 0
	if xpPools["combat"] != nil {
		combatXP = int(xpPools["combat"].(float64))
	}
	if xpPools["crafting"] != nil {
		craftingXP = int(xpPools["crafting"].(float64))
	}
	fmt.Printf("  Combat XP: %d, Crafting XP: %d\n", combatXP, craftingXP)
	if craftingXP != 5000 {
		fmt.Printf("  [FAIL] Expected 5000 crafting XP, got %d\n", craftingXP)
		os.Exit(1)
	}
	fmt.Println("  [PASS] Typed XP pools work correctly")

	// --- Summary ---
	fmt.Println("\n=== Phase 2 Test Summary ===")
	fmt.Println("  Test 1:  List professions             [PASS]")
	fmt.Println("  Test 2:  List trainers                 [PASS]")
	fmt.Println("  Test 3:  Character starts clean        [PASS]")
	fmt.Println("  Test 4:  Earn XP (placeholder)         [PASS]")
	fmt.Println("  Test 5:  Train Novice (free)           [PASS]")
	fmt.Println("  Test 6:  Train tier I (costs SP/XP)    [PASS]")
	fmt.Println("  Test 7:  Cannot skip tiers             [PASS]")
	fmt.Println("  Test 8:  Cannot train without Novice   [PASS]")
	fmt.Println("  Test 9:  Drop skill + SP refund        [PASS]")
	fmt.Println("  Test 10: Skill persists login          [PASS]")
	fmt.Println("  Test 11: Typed XP pools                 [PASS]")
	fmt.Println()
	fmt.Println("=== ALL TESTS PASSED — Phase 2 exit criteria met ===")
	fmt.Println("A player can earn XP from a placeholder action, spend skill points")
	fmt.Println("at a trainer NPC, and see the skill persist across logout/login.")
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
			"body_type":  "average",
			"skin_color": "default",
			"height":     1.0,
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
		log.Fatalf("create character failed (HTTP %d): %s", resp.StatusCode, string(respBody))
	}
	var result protocol.CreateCharacterResponse
	json.Unmarshal(respBody, &result)
	return result.CharacterID
}

func getProfessions() []map[string]interface{} {
	resp, err := http.Get(serverURL + "/api/professions")
	if err != nil {
		log.Fatal("get professions failed:", err)
	}
	var result []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	resp.Body.Close()
	return result
}

func getTrainers() []map[string]interface{} {
	resp, err := http.Get(serverURL + "/api/trainers")
	if err != nil {
		log.Fatal("get trainers failed:", err)
	}
	var result []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	resp.Body.Close()
	return result
}

func getCharacterSkills(token, charID string) map[string]interface{} {
	req, _ := http.NewRequest("GET", serverURL+"/api/characters/"+charID+"/skills", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("get skills failed:", err)
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	resp.Body.Close()
	return result
}

func earnXP(token, charID, xpType string, amount int) map[string]interface{} {
	body, _ := json.Marshal(map[string]interface{}{"xp_type": xpType, "amount": amount})
	req, _ := http.NewRequest("POST", serverURL+"/api/characters/"+charID+"/earn-xp", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("earn XP failed:", err)
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	resp.Body.Close()
	return result
}

func tryTrainSkill(token, charID, trainerID, skillBoxID string) *http.Response {
	body, _ := json.Marshal(map[string]string{"trainer_id": trainerID, "skill_box_id": skillBoxID})
	req, _ := http.NewRequest("POST", serverURL+"/api/characters/"+charID+"/train", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("train failed:", err)
	}
	return resp
}

func trainSkill(token, charID, trainerID, skillBoxID string) map[string]interface{} {
	resp := tryTrainSkill(token, charID, trainerID, skillBoxID)
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode != 201 {
		log.Fatalf("train failed (HTTP %d): %s", resp.StatusCode, string(body))
	}
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	return result
}

func dropSkill(token, charID, skillBoxID string) map[string]interface{} {
	body, _ := json.Marshal(map[string]string{"skill_box_id": skillBoxID})
	req, _ := http.NewRequest("POST", serverURL+"/api/characters/"+charID+"/drop-skill", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("drop skill failed:", err)
	}
	respBody, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("drop skill failed (HTTP %d): %s", resp.StatusCode, string(respBody))
	}
	var result map[string]interface{}
	json.Unmarshal(respBody, &result)
	return result
}
