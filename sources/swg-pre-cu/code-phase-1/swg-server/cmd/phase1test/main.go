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

// Phase 1 integration test: server-validated movement, spatial interest
// management, and spatial chat.
//
// GDD Phase 1 exit criteria: "A player can walk around a persistent zone
// on one planet and see/chat with another connected player in real time."

const serverURL = "http://localhost:8080"
const wsURL = "ws://localhost:8080/ws"

type wsClient struct {
	conn *websocket.Conn
	// Messages received, keyed by type
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

func (wc *wsClient) countMessages(msgType string) int {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	count := 0
	for _, m := range wc.messages {
		if m.Type == msgType {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println("=== SWG Pre-CU Phase 1 Integration Test ===")
	fmt.Println("Server-validated movement + spatial interest management + spatial chat")
	fmt.Println()

	if !waitForServer() {
		fmt.Println("FAIL: Server not reachable")
		os.Exit(1)
	}
	fmt.Println("Server is reachable.")

	// --- Setup: register account, create two characters ---
	fmt.Println("\n--- Setup: Register account + create two characters ---")
	accountID, token := registerAccount()
	fmt.Printf("  Account: %s\n", accountID[:8])

	// Create two characters at spawn
	charA := createCharacter(token, "human", "TestAlpha")
	charB := createCharacter(token, "bothan", "TestBeta")
	if charA == "" || charB == "" {
		fmt.Println("FAIL: Failed to create characters")
		os.Exit(1)
	}
	fmt.Printf("  Character A: %s (%s)\n", "TestAlpha", charA[:8])
	fmt.Printf("  Character B: %s (%s)\n", "TestBeta", charB[:8])

	// --- Test 1: Both players enter world at spawn and see each other ---
	fmt.Println("\n--- Test 1: Proximity-based entity spawn ---")
	wcA := connectWS(token)
	wcB := connectWS(token)
	time.Sleep(200 * time.Millisecond)

	// A enters world first
	wcA.send(protocol.MsgEnterWorld, protocol.EnterWorldMsg{CharacterID: charA})
	worldEnterA, ok := wcA.waitFor(protocol.MsgWorldEnter, 3*time.Second)
	if !ok {
		fmt.Println("  [FAIL] A did not receive world_enter")
		os.Exit(1)
	}
	fmt.Printf("  A entered world at (%.1f, %.1f, %.1f)\n",
		worldEnterA.Data.(map[string]interface{})["position"].(map[string]interface{})["x"].(float64),
		worldEnterA.Data.(map[string]interface{})["position"].(map[string]interface{})["y"].(float64),
		worldEnterA.Data.(map[string]interface{})["position"].(map[string]interface{})["z"].(float64),
	)

	// B enters world — A should receive entity_spawn for B (both at spawn, within 128m)
	wcB.send(protocol.MsgEnterWorld, protocol.EnterWorldMsg{CharacterID: charB})
	_, ok = wcB.waitFor(protocol.MsgWorldEnter, 3*time.Second)
	if !ok {
		fmt.Println("  [FAIL] B did not receive world_enter")
		os.Exit(1)
	}

	// A should see B spawn
	spawnMsg, ok := wcA.waitFor(protocol.MsgEntitySpawn, 3*time.Second)
	if !ok {
		fmt.Println("  [FAIL] A did not receive entity_spawn for B")
		os.Exit(1)
	}
	spawnData := spawnMsg.Data.(map[string]interface{})
	fmt.Printf("  A received entity_spawn: %s (%s)\n", spawnData["name"], spawnData["species"])
	fmt.Println("  [PASS] Both players see each other at spawn")

	// --- Test 2: Movement broadcast to nearby player ---
	fmt.Println("\n--- Test 2: Movement broadcast to nearby player ---")
	// A moves slightly (1m, well within speed limit)
	wcA.send(protocol.MsgMove, protocol.MoveMsg{X: 3529, Y: 5, Z: -4804, Heading: 90})

	moveMsg, ok := wcB.waitFor(protocol.MsgEntityMove, 3*time.Second)
	if !ok {
		fmt.Println("  [FAIL] B did not receive entity_move for A")
		os.Exit(1)
	}
	moveData := moveMsg.Data.(map[string]interface{})
	fmt.Printf("  B received entity_move for A to (%.1f, %.1f, %.1f)\n",
		moveData["position"].(map[string]interface{})["x"].(float64),
		moveData["position"].(map[string]interface{})["y"].(float64),
		moveData["position"].(map[string]interface{})["z"].(float64),
	)
	fmt.Println("  [PASS] Movement broadcast to nearby player")

	// --- Test 3: Spatial chat (/say) delivered within 20m ---
	fmt.Println("\n--- Test 3: Spatial chat (/say) within 20m ---")
	// A and B are ~1m apart, well within 20m
	chatText := "Hello from TestAlpha!"
	wcA.send(protocol.MsgChat, protocol.ChatMsg{Channel: "spatial", Text: chatText})

	chatMsg, ok := wcB.waitFor(protocol.MsgChatMessage, 3*time.Second)
	if !ok {
		fmt.Println("  [FAIL] B did not receive spatial chat from A")
		os.Exit(1)
	}
	chatData := chatMsg.Data.(map[string]interface{})
	receivedText := chatData["text"].(string)
	if receivedText != chatText {
		fmt.Printf("  [FAIL] Chat text mismatch: got %q, expected %q\n", receivedText, chatText)
		os.Exit(1)
	}
	fmt.Printf("  B received chat: \"%s\" from %s\n", receivedText, chatData["sender_name"])
	fmt.Println("  [PASS] Spatial chat delivered within 20m")

	// --- Test 4: Move B far away (beyond 128m interest radius) ---
	fmt.Println("\n--- Test 4: Entity despawn beyond interest radius (128m) ---")
	// Move B to 3700, 5, -4804 (172m from A at 3529)
	// We need to move B gradually to avoid speed rejection
	// Actually, let's move B in one step — it might get rejected.
	// Instead, let's just move B far away and see if A gets a despawn.
	// But the server will reject the teleport... Let's move B step by step.

	// Move B in small steps to ~200m away (total ~200m, each step ~7m, so ~30 steps)
	// Actually, for the test let's just do it in one step and see if it works.
	// The speed validation allows 7*1.5 = 10.5 m/s, and dt since last move is ~1s,
	// so max distance is ~11.5m. We need to move in steps of ~10m.

	startX := 3528.0
	targetX := 3740.0                      // ~212m away, beyond 128m interest radius
	step := 5.0                            // 5m per step, realistic on-foot speed (GDD 20.2.1: 5-7 m/s)
	stepInterval := 500 * time.Millisecond // 500ms → 5m/0.5s = 10 m/s (within 7*1.5=10.5 tolerance)

	for x := startX; x <= targetX; x += step {
		wcB.send(protocol.MsgMove, protocol.MoveMsg{X: x, Y: 5, Z: -4804, Heading: 0})
		time.Sleep(stepInterval)
	}
	// Final position
	wcB.send(protocol.MsgMove, protocol.MoveMsg{X: targetX, Y: 5, Z: -4804, Heading: 0})
	time.Sleep(500 * time.Millisecond)

	// A should have received entity_despawn for B
	despawnCount := wcA.countMessages(protocol.MsgEntityDespawn)
	if despawnCount > 0 {
		fmt.Printf("  A received %d entity_despawn message(s) for B\n", despawnCount)
		fmt.Println("  [PASS] Entity despawned beyond interest radius")
	} else {
		fmt.Println("  [FAIL] A did not receive entity_despawn when B moved away")
		os.Exit(1)
	}

	// --- Test 5: Spatial chat NOT delivered beyond 20m ---
	fmt.Println("\n--- Test 5: Spatial chat NOT delivered beyond 20m ---")
	// Clear B's messages
	wcB.mu.Lock()
	wcB.messages = wcB.messages[:0]
	wcB.mu.Unlock()

	// A sends /say — B is ~212m away, should NOT receive it
	wcA.send(protocol.MsgChat, protocol.ChatMsg{Channel: "spatial", Text: "Can you hear me?"})
	time.Sleep(1 * time.Second)

	chatCount := wcB.countMessages(protocol.MsgChatMessage)
	if chatCount == 0 {
		fmt.Println("  B did NOT receive spatial chat (correct — out of 20m range)")
		fmt.Println("  [PASS] Spatial chat not delivered beyond 20m")
	} else {
		fmt.Printf("  [FAIL] B received %d spatial chat messages (should be 0)\n", chatCount)
		os.Exit(1)
	}

	// --- Test 6: Teleport rejection (anti-cheat) ---
	fmt.Println("\n--- Test 6: Teleport rejection (movement validation) ---")
	// Clear A's messages
	wcA.mu.Lock()
	wcA.messages = wcA.messages[:0]
	wcA.mu.Unlock()

	// A tries to teleport 1000m away (should be rejected)
	wcA.send(protocol.MsgMove, protocol.MoveMsg{X: 4529, Y: 5, Z: -4804, Heading: 0})

	correction, ok := wcA.waitFor(protocol.MsgPositionCorrection, 3*time.Second)
	if !ok {
		fmt.Println("  [FAIL] A did not receive position_correction for teleport attempt")
		os.Exit(1)
	}
	corrData := correction.Data.(map[string]interface{})
	corrReason := corrData["reason"].(string)
	fmt.Printf("  A received position_correction: %s\n", corrReason)
	fmt.Println("  [PASS] Teleport rejected with position correction")

	// --- Test 7: Move B back into range, entity_spawn ---
	fmt.Println("\n--- Test 7: Entity respawn when moving back into range ---")
	// Clear A's messages
	wcA.mu.Lock()
	wcA.messages = wcA.messages[:0]
	wcA.mu.Unlock()

	// Move B back toward spawn in steps
	for x := targetX; x >= startX; x -= step {
		wcB.send(protocol.MsgMove, protocol.MoveMsg{X: x, Y: 5, Z: -4804, Heading: 180})
		time.Sleep(stepInterval)
	}
	wcB.send(protocol.MsgMove, protocol.MoveMsg{X: startX, Y: 5, Z: -4804, Heading: 180})
	time.Sleep(500 * time.Millisecond)

	// A should have received entity_spawn for B again
	spawnCount := wcA.countMessages(protocol.MsgEntitySpawn)
	if spawnCount > 0 {
		fmt.Printf("  A received %d entity_spawn message(s) for B\n", spawnCount)
		fmt.Println("  [PASS] Entity respawned when moving back into range")
	} else {
		fmt.Println("  [FAIL] A did not receive entity_spawn when B moved back")
		os.Exit(1)
	}

	// --- Summary ---
	fmt.Println("\n=== Phase 1 Test Summary ===")
	fmt.Println("  Test 1: Proximity entity spawn     [PASS]")
	fmt.Println("  Test 2: Movement broadcast         [PASS]")
	fmt.Println("  Test 3: Spatial chat within 20m    [PASS]")
	fmt.Println("  Test 4: Despawn beyond 128m        [PASS]")
	fmt.Println("  Test 5: Chat NOT beyond 20m        [PASS]")
	fmt.Println("  Test 6: Teleport rejection         [PASS]")
	fmt.Println("  Test 7: Respawn back in range      [PASS]")
	fmt.Println()
	fmt.Println("=== ALL TESTS PASSED — Phase 1 exit criteria met ===")
	fmt.Println("A player can walk around a persistent zone on one planet and")
	fmt.Println("see/chat with another connected player in real time.")

	wcA.conn.Close()
	wcB.conn.Close()
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
		log.Fatalf("create character %s failed (HTTP %d): %s", name, resp.StatusCode, string(respBody))
	}
	var result protocol.CreateCharacterResponse
	json.Unmarshal(respBody, &result)
	return result.CharacterID
}

func connectWS(token string) *wsClient {
	dialer := websocket.Dialer{HandshakeTimeout: 5 * time.Second}
	url := fmt.Sprintf("%s?token=%s", wsURL, token)
	conn, _, err := dialer.Dial(url, nil)
	if err != nil {
		log.Fatal("WebSocket dial failed:", err)
	}
	wc := &wsClient{conn: conn}
	go wc.readLoop()
	return wc
}
