package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"

	"swg-server/internal/protocol"
)

// testclient is a CLI integration test that verifies Phase 0 exit criteria:
// "A player can create an account, create a character with species/appearance,
// and see it standing in a bare world with correct starting HAM values."
//
// It tests ALL 9 species and verifies HAM values match the GDD spec.

const serverURL = "http://localhost:8080"
const wsURL = "ws://localhost:8080/ws"

// expectedHAM is a hard-coded table of expected HAM values for each species,
// derived directly from the GDD (Section 6.2.1 + 6.2.2).
// This is INDEPENDENT of the server's ComputeHAM function so the test
// genuinely validates the implementation against the spec.
var expectedHAM = map[string]protocol.HAMState{
	"human":        {Health: 1000, Strength: 500, Constitution: 500, Action: 1000, Quickness: 500, Stamina: 500, Mind: 1000, Focus: 500, Willpower: 500},
	"bothan":       {Health: 1000, Strength: 490, Constitution: 500, Action: 1000, Quickness: 510, Stamina: 500, Mind: 1000, Focus: 500, Willpower: 500},
	"rodian":       {Health: 1000, Strength: 500, Constitution: 490, Action: 1000, Quickness: 510, Stamina: 500, Mind: 1000, Focus: 500, Willpower: 500},
	"trandoshan":   {Health: 1000, Strength: 520, Constitution: 500, Action: 1000, Quickness: 490, Stamina: 500, Mind: 1000, Focus: 500, Willpower: 500},
	"twilek":       {Health: 1000, Strength: 500, Constitution: 500, Action: 1000, Quickness: 500, Stamina: 490, Mind: 1000, Focus: 500, Willpower: 510},
	"wookiee":      {Health: 1100, Strength: 550, Constitution: 500, Action: 1000, Quickness: 450, Stamina: 500, Mind: 1000, Focus: 500, Willpower: 500},
	"zabrak":       {Health: 1000, Strength: 500, Constitution: 500, Action: 1000, Quickness: 500, Stamina: 510, Mind: 1000, Focus: 500, Willpower: 490},
	"mon_calamari": {Health: 1000, Strength: 500, Constitution: 490, Action: 1000, Quickness: 500, Stamina: 500, Mind: 1000, Focus: 510, Willpower: 500},
	"sullustan":    {Health: 1000, Strength: 500, Constitution: 500, Action: 1000, Quickness: 510, Stamina: 490, Mind: 1000, Focus: 500, Willpower: 500},
}

type testResult struct {
	species string
	passed  bool
	ham     protocol.HAMState
	detail  string
}

func main() {
	fmt.Println("=== SWG Pre-CU Phase 0 Integration Test ===")
	fmt.Println()

	// Wait for server to be ready
	if !waitForServer() {
		fmt.Println("FAIL: Server not reachable at", serverURL)
		os.Exit(1)
	}
	fmt.Println("Server is reachable.")
	fmt.Println()

	// Step 1: Register an account
	fmt.Println("--- Step 1: Register Account ---")
	accountID, token := registerAccount()
	fmt.Printf("  Account ID: %s\n", accountID)
	fmt.Printf("  Token: %s\n", token)
	fmt.Println()

	// Step 2: Get species list
	fmt.Println("--- Step 2: Get Species List ---")
	speciesList := getSpecies()
	fmt.Printf("  Found %d species:\n", len(speciesList))
	for _, sp := range speciesList {
		fmt.Printf("    - %s (%s)\n", sp.DisplayName, sp.ID)
	}
	fmt.Println()

	// Step 3: Create characters for ALL 9 species and verify HAM
	fmt.Println("--- Step 3: Create Characters (all 9 species) & Verify HAM ---")
	results := make([]testResult, 0, 9)

	allSpecies := []string{
		"human", "bothan", "rodian", "trandoshan", "twilek",
		"wookiee", "zabrak", "mon_calamari", "sullustan",
	}

	for i, sp := range allSpecies {
		name := fmt.Sprintf("Test%c_%s", 'A'+i, sp)
		charID := createCharacter(token, sp, name)
		charData := getCharacter(token, charID)

		receivedHAM := charData.HAM

		// Compare against hard-coded expected values from the GDD
		expected := expectedHAM[sp]
		passed := true
		detail := ""
		if receivedHAM.Health != expected.Health {
			passed = false
			detail += fmt.Sprintf("Health: got %d, expected %d; ", receivedHAM.Health, expected.Health)
		}
		if receivedHAM.Strength != expected.Strength {
			passed = false
			detail += fmt.Sprintf("Strength: got %d, expected %d; ", receivedHAM.Strength, expected.Strength)
		}
		if receivedHAM.Constitution != expected.Constitution {
			passed = false
			detail += fmt.Sprintf("Constitution: got %d, expected %d; ", receivedHAM.Constitution, expected.Constitution)
		}
		if receivedHAM.Action != expected.Action {
			passed = false
			detail += fmt.Sprintf("Action: got %d, expected %d; ", receivedHAM.Action, expected.Action)
		}
		if receivedHAM.Quickness != expected.Quickness {
			passed = false
			detail += fmt.Sprintf("Quickness: got %d, expected %d; ", receivedHAM.Quickness, expected.Quickness)
		}
		if receivedHAM.Stamina != expected.Stamina {
			passed = false
			detail += fmt.Sprintf("Stamina: got %d, expected %d; ", receivedHAM.Stamina, expected.Stamina)
		}
		if receivedHAM.Mind != expected.Mind {
			passed = false
			detail += fmt.Sprintf("Mind: got %d, expected %d; ", receivedHAM.Mind, expected.Mind)
		}
		if receivedHAM.Focus != expected.Focus {
			passed = false
			detail += fmt.Sprintf("Focus: got %d, expected %d; ", receivedHAM.Focus, expected.Focus)
		}
		if receivedHAM.Willpower != expected.Willpower {
			passed = false
			detail += fmt.Sprintf("Willpower: got %d, expected %d; ", receivedHAM.Willpower, expected.Willpower)
		}

		if passed {
			detail = "All HAM values match GDD spec"
		}

		results = append(results, testResult{
			species: sp,
			passed:  passed,
			ham:     receivedHAM,
			detail:  detail,
		})

		status := "PASS"
		if !passed {
			status = "FAIL"
		}
		fmt.Printf("  [%s] %s (H:%d S:%d C:%d A:%d Q:%d St:%d M:%d F:%d W:%d)\n",
			status, sp,
			receivedHAM.Health, receivedHAM.Strength, receivedHAM.Constitution,
			receivedHAM.Action, receivedHAM.Quickness, receivedHAM.Stamina,
			receivedHAM.Mind, receivedHAM.Focus, receivedHAM.Willpower)
		if !passed {
			fmt.Printf("         %s\n", detail)
		}
	}
	fmt.Println()

	// Step 4: List characters
	fmt.Println("--- Step 4: List Characters ---")
	chars := listCharacters(token)
	fmt.Printf("  Found %d characters for this account:\n", len(chars))
	for _, c := range chars {
		fmt.Printf("    - %s (%s, %s) on %s\n", c.Name, c.Species, c.ID[:8], c.Planet)
	}
	fmt.Println()

	// Step 5: WebSocket enter_world — the critical Phase 0 test
	fmt.Println("--- Step 5: WebSocket Enter World ---")
	wsResult := false
	if len(chars) == 0 {
		fmt.Println("  [SKIP] No characters to test WebSocket with")
		wsResult = false
	} else {
		wsResult = testWebSocket(token, chars[0].ID)
	}
	if wsResult {
		fmt.Println("  [PASS] WebSocket world_enter received with correct character state")
	} else {
		fmt.Println("  [FAIL] WebSocket enter_world failed")
	}
	fmt.Println()

	// Summary
	fmt.Println("=== Test Summary ===")
	passCount := 0
	for _, r := range results {
		if r.passed {
			passCount++
		}
	}
	fmt.Printf("  Species HAM tests: %d/%d passed\n", passCount, len(results))
	fmt.Printf("  WebSocket test:    %v\n", wsResult)
	fmt.Println()

	allPassed := passCount == len(results) && wsResult
	if allPassed {
		fmt.Println("=== ALL TESTS PASSED — Phase 0 exit criteria met ===")
		fmt.Println("A player can create an account, create a character with species/appearance,")
		fmt.Println("and see it standing in a bare world with correct starting HAM values.")
	} else {
		fmt.Println("=== SOME TESTS FAILED ===")
		os.Exit(1)
	}
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
	req := protocol.RegisterRequest{Username: username, Password: "testpass123"}

	body, _ := json.Marshal(req)
	resp, err := http.Post(serverURL+"/api/register", "application/json", bytes.NewReader(body))
	if err != nil {
		panic(fmt.Sprintf("register failed: %v", err))
	}
	respBody, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	var result protocol.RegisterResponse
	json.Unmarshal(respBody, &result)
	fmt.Printf("  Register response: %s\n", string(respBody))

	// Now login to get the token
	loginReq := protocol.LoginRequest{Username: username, Password: "testpass123"}
	body, _ = json.Marshal(loginReq)
	resp2, err := http.Post(serverURL+"/api/login", "application/json", bytes.NewReader(body))
	if err != nil {
		panic(fmt.Sprintf("login failed: %v", err))
	}
	loginBody, _ := io.ReadAll(resp2.Body)
	resp2.Body.Close()

	var loginResult protocol.LoginResponse
	json.Unmarshal(loginBody, &loginResult)
	fmt.Printf("  Login response: %s\n", string(loginBody))

	return loginResult.AccountID, loginResult.Token
}

func getSpecies() []protocol.SpeciesInfo {
	resp, err := http.Get(serverURL + "/api/species")
	if err != nil {
		panic(fmt.Sprintf("get species failed: %v", err))
	}
	defer resp.Body.Close()

	var result protocol.SpeciesListResponse
	json.NewDecoder(resp.Body).Decode(&result)
	return result.Species
}

func createCharacter(token, speciesID, name string) string {
	req := protocol.CreateCharacterRequest{
		Species: speciesID,
		Name:    name,
		Appearance: protocol.CharacterAppearanceReq{
			BodyType:  "average",
			SkinColor: "default",
			Height:    1.0,
		},
	}

	body, _ := json.Marshal(req)
	req2, _ := http.NewRequest("POST", serverURL+"/api/characters", bytes.NewReader(body))
	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req2)
	if err != nil {
		panic(fmt.Sprintf("create character failed: %v", err))
	}
	defer resp.Body.Close()

	var result protocol.CreateCharacterResponse
	json.NewDecoder(resp.Body).Decode(&result)
	return result.CharacterID
}

func getCharacter(token, charID string) protocol.GetCharacterResponse {
	req, _ := http.NewRequest("GET", serverURL+"/api/characters/"+charID, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(fmt.Sprintf("get character failed: %v", err))
	}
	defer resp.Body.Close()

	var result protocol.GetCharacterResponse
	json.NewDecoder(resp.Body).Decode(&result)
	return result
}

func listCharacters(token string) []protocol.CharacterSummary {
	req, _ := http.NewRequest("GET", serverURL+"/api/characters", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(fmt.Sprintf("list characters failed: %v", err))
	}
	defer resp.Body.Close()

	var result protocol.ListCharactersResponse
	json.NewDecoder(resp.Body).Decode(&result)
	return result.Characters
}

func testWebSocket(token, charID string) bool {
	// Connect to WebSocket
	dialer := websocket.Dialer{HandshakeTimeout: 5 * time.Second}
	url := fmt.Sprintf("%s?token=%s", wsURL, token)
	conn, _, err := dialer.Dial(url, nil)
	if err != nil {
		fmt.Printf("  WebSocket dial failed: %v\n", err)
		return false
	}
	defer conn.Close()

	// Send enter_world message
	enterMsg := protocol.WSMessage{
		Type: protocol.MsgEnterWorld,
		Data: protocol.EnterWorldMsg{CharacterID: charID},
	}
	msgBytes, _ := json.Marshal(enterMsg)
	if err := conn.WriteMessage(websocket.TextMessage, msgBytes); err != nil {
		fmt.Printf("  WebSocket write failed: %v\n", err)
		return false
	}

	// Read response (world_enter)
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	_, msg, err := conn.ReadMessage()
	if err != nil {
		fmt.Printf("  WebSocket read failed: %v\n", err)
		return false
	}

	var response protocol.WSMessage
	if err := json.Unmarshal(msg, &response); err != nil {
		fmt.Printf("  WebSocket unmarshal failed: %v\n", err)
		return false
	}

	if response.Type != protocol.MsgWorldEnter {
		fmt.Printf("  Expected %s, got %s\n", protocol.MsgWorldEnter, response.Type)
		return false
	}

	// Parse the world_enter data
	dataBytes, _ := json.Marshal(response.Data)
	var worldEnter protocol.WorldEnterMsg
	if err := json.Unmarshal(dataBytes, &worldEnter); err != nil {
		fmt.Printf("  Failed to parse world_enter data: %v\n", err)
		return false
	}

	// Verify the character data
	if worldEnter.CharacterID != charID {
		fmt.Printf("  Character ID mismatch: got %s, expected %s\n", worldEnter.CharacterID, charID)
		return false
	}
	if worldEnter.Planet != "tatooine" {
		fmt.Printf("  Planet mismatch: got %s, expected tatooine\n", worldEnter.Planet)
		return false
	}
	if worldEnter.HAM.Health == 0 {
		fmt.Println("  HAM Health is 0 — not computed")
		return false
	}

	fmt.Printf("  Character: %s (%s)\n", worldEnter.Name, worldEnter.Species)
	fmt.Printf("  Position: (%.1f, %.1f, %.1f) on %s\n", worldEnter.Position.X, worldEnter.Position.Y, worldEnter.Position.Z, worldEnter.Planet)
	fmt.Printf("  HAM: H:%d A:%d M:%d | Str:%d Con:%d Qui:%d Sta:%d Foc:%d Wil:%d\n",
		worldEnter.HAM.Health, worldEnter.HAM.Action, worldEnter.HAM.Mind,
		worldEnter.HAM.Strength, worldEnter.HAM.Constitution, worldEnter.HAM.Quickness,
		worldEnter.HAM.Stamina, worldEnter.HAM.Focus, worldEnter.HAM.Willpower)

	return true
}

// Silence unused import warning
var _ = io.EOF
