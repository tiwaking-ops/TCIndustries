# World.gd
# The main game world scene. Connects to the server via WebSocket,
# enters the world with the selected character, and renders the character
# in the world with WASD movement and HAM display.
#
# Phase 1: Server-validated movement, spatial interest management, spatial chat.
# GDD Phase 1 exit criteria: "A player can walk around a persistent zone on
# one planet and see/chat with another connected player in real time."

extends Node3D

# World nodes
@onready var ground: MeshInstance3D = $Ground
@onready var sun: DirectionalLight3D = $Sun
@onready var camera: Camera3D = $Camera3D
@onready var player_node: Node3D = $Player
@onready var player_mesh: MeshInstance3D = $Player/MeshInstance3D

# UI overlay
@onready var ham_panel: Panel = $UI/HAMPanel
@onready var name_label: Label = $UI/HAMPanel/VBox/NameLabel
@onready var species_label: Label = $UI/HAMPanel/VBox/SpeciesLabel
@onready var planet_label: Label = $UI/HAMPanel/VBox/PlanetLabel
@onready var health_bar: ProgressBar = $UI/HAMPanel/VBox/HealthBar
@onready var action_bar: ProgressBar = $UI/HAMPanel/VBox/ActionBar
@onready var mind_bar: ProgressBar = $UI/HAMPanel/VBox/MindBar
@onready var ham_details: Label = $UI/HAMPanel/VBox/HAMDetails
@onready var position_label: Label = $UI/HAMPanel/VBox/PositionLabel

# Chat
@onready var chat_log: RichTextLabel = $UI/ChatPanel/ChatLog
@onready var chat_input: LineEdit = $UI/ChatPanel/ChatInput

# WebSocket client
@onready var ws_client: WSClient = $WSClient

# Other players (entity_id → Node3D)
var other_players: Dictionary = {}

# Phase 1: Movement state
# GDD Section 20.2.1: On Foot speed is 5-7 m/s.
const MOVE_SPEED := 6.0  # m/s — within the 5-7 m/s GDD range
const MOVE_SEND_INTERVAL := 0.066  # ~15Hz move updates (GDD recommends 10-20Hz)
const CAMERA_OFFSET := Vector3(0, 12, 10)  # Third-person camera offset

var move_timer: float = 0.0
var in_world: bool = false
var current_heading: float = 0.0

func _ready() -> void:
	# Connect WebSocket signals
	ws_client.connected.connect(_on_ws_connected)
	ws_client.disconnected.connect(_on_ws_disconnected)
	ws_client.world_entered.connect(_on_world_entered)
	ws_client.entity_spawned.connect(_on_entity_spawned)
	ws_client.entity_moved.connect(_on_entity_moved)
	ws_client.entity_despawned.connect(_on_entity_despawned)
	ws_client.chat_received.connect(_on_chat_received)
	ws_client.position_corrected.connect(_on_position_corrected)
	ws_client.error_received.connect(_on_ws_error)
	
	# Chat input
	chat_input.text_submitted.connect(_on_chat_submitted)
	
	# Connect to WebSocket server
	ws_client.connect_to_server(GameState.token)
	
	# Position camera to overview the spawn area
	camera.position = Vector3(3538, 15, -4790)
	camera.look_at(Vector3(3528, 5, -4804))

func _on_ws_connected() -> void:
	# Send enter_world with the selected character ID
	ws_client.send_enter_world(GameState.selected_character_id)

func _on_ws_disconnected() -> void:
	chat_log.append_text("[color=red]Disconnected from server.[/color]\n")
	in_world = false

func _on_ws_error(message: String) -> void:
	chat_log.append_text("[color=red]Error: " + message + "[/color]\n")

func _on_world_entered(data: Dictionary) -> void:
	var char_name: String = data.get("name", "Unknown")
	var species: String = data.get("species", "unknown")
	var planet: String = data.get("planet", "unknown")
	var pos: Dictionary = data.get("position", {})
	var ham: Dictionary = data.get("ham", {})
	
	# Update player position
	var x := float(pos.get("x", 3528))
	var y := float(pos.get("y", 5))
	var z := float(pos.get("z", -4804))
	player_node.position = Vector3(x, y, z)
	
	# Update UI
	name_label.text = "Name: " + char_name
	species_label.text = "Species: " + species
	planet_label.text = "Planet: " + planet
	
	# HAM bars
	var health: int = int(ham.get("health", 0))
	var action: int = int(ham.get("action", 0))
	var mind: int = int(ham.get("mind", 0))
	
	health_bar.max_value = max(health, 1)
	health_bar.value = health
	action_bar.max_value = max(action, 1)
	action_bar.value = action
	mind_bar.max_value = max(mind, 1)
	mind_bar.value = mind
	
	# HAM details (all 9 attributes)
	ham_details.text = "H:%d  A:%d  M:%d | Str:%d  Con:%d  Qui:%d  Sta:%d  Foc:%d  Wil:%d" % [
		health, action, mind,
		int(ham.get("strength", 0)),
		int(ham.get("constitution", 0)),
		int(ham.get("quickness", 0)),
		int(ham.get("stamina", 0)),
		int(ham.get("focus", 0)),
		int(ham.get("willpower", 0))
	]
	
	position_label.text = "Pos: (%.1f, %.1f, %.1f)" % [x, y, z]
	
	chat_log.append_text("[color=green]Entered world as " + char_name + " (" + species + ") on " + planet + "[/color]\n")
	chat_log.append_text("[color=yellow]WASD to move. Enter to chat.[/color]\n")
	
	in_world = true

# Phase 1: WASD movement with server validation
# GDD Section 20.2.1: On Foot 5-7 m/s. GDD Section 29.2.1: server-authoritative.
func _process(delta: float) -> void:
	if not in_world:
		return
	
	# WASD movement (local prediction — server validates)
	var input_dir := Vector2.ZERO
	if Input.is_key_pressed(KEY_W):
		input_dir.y -= 1
	if Input.is_key_pressed(KEY_S):
		input_dir.y += 1
	if Input.is_key_pressed(KEY_A):
		input_dir.x -= 1
	if Input.is_key_pressed(KEY_D):
		input_dir.x += 1
	
	if input_dir != Vector2.ZERO:
		input_dir = input_dir.normalized()
		
		# Calculate heading from input direction
		current_heading = rad_to_deg(atan2(input_dir.x, -input_dir.y))
		
		# Move player (XZ plane, Y is altitude)
		var move_vec := Vector3(input_dir.x, 0, input_dir.y) * MOVE_SPEED * delta
		player_node.position += move_vec
		
		# Rotate player mesh to face movement direction
		player_node.rotation.y = deg_to_rad(-current_heading + 180)
	
	# Send move updates at ~15Hz (GDD: 10-20Hz)
	move_timer += delta
	if move_timer >= MOVE_SEND_INTERVAL:
		move_timer = 0.0
		ws_client.send_move(
			player_node.position.x,
			player_node.position.y,
			player_node.position.z,
			current_heading
		)
	
	# Update position label
	position_label.text = "Pos: (%.1f, %.1f, %.1f)" % [
		player_node.position.x,
		player_node.position.y,
		player_node.position.z
	]
	
	# Camera follows player (third-person)
	var target_cam_pos := player_node.position + CAMERA_OFFSET
	camera.position = camera.position.lerp(target_cam_pos, delta * 5.0)
	camera.look_at(player_node.position)

# Phase 1: Handle server position correction (anti-cheat)
# GDD Section 29.2.1: server is authoritative. If the server rejects a move,
# the client must snap to the corrected position.
func _on_position_corrected(data: Dictionary) -> void:
	var pos: Dictionary = data.get("position", {})
	var reason: String = data.get("reason", "unknown")
	var x := float(pos.get("x", player_node.position.x))
	var y := float(pos.get("y", player_node.position.y))
	var z := float(pos.get("z", player_node.position.z))
	
	# Snap to server-authoritative position
	player_node.position = Vector3(x, y, z)
	
	chat_log.append_text("[color=orange]Position corrected: " + reason + "[/color]\n")

func _on_entity_spawned(data: Dictionary) -> void:
	var entity_id: String = data.get("entity_id", "")
	var entity_name: String = data.get("name", "Unknown")
	var species: String = data.get("species", "")
	var pos: Dictionary = data.get("position", {})
	
	if entity_id == GameState.selected_character_id:
		return # Don't spawn ourselves
	
	# Create a simple mesh for the other player
	var mesh := MeshInstance3D.new()
	var box := BoxMesh.new()
	box.size = Vector3(1, 2, 1)
	mesh.mesh = box
	
	var mat := StandardMaterial3D.new()
	mat.albedo_color = Color(0.8, 0.6, 0.4)
	mesh.material_override = mat
	
	var node := Node3D.new()
	node.name = "entity_" + entity_id
	node.add_child(mesh)
	
	var x := float(pos.get("x", 3528))
	var y := float(pos.get("y", 5))
	var z := float(pos.get("z", -4804))
	node.position = Vector3(x, y, z)
	
	# Add a label above the entity
	var label := Label3D.new()
	label.text = entity_name
	label.position = Vector3(0, 2, 0)
	label.font_size = 32
	node.add_child(label)
	
	add_child(node)
	other_players[entity_id] = node
	
	chat_log.append_text("[color=green]" + entity_name + " entered the area.[/color]\n")

func _on_entity_moved(data: Dictionary) -> void:
	var entity_id: String = data.get("entity_id", "")
	var pos: Dictionary = data.get("position", {})
	var heading: float = float(data.get("heading", 0))
	
	if other_players.has(entity_id):
		var node: Node3D = other_players[entity_id]
		var x := float(pos.get("x", 0))
		var y := float(pos.get("y", 0))
		var z := float(pos.get("z", 0))
		# Smoothly interpolate to the new position
		node.position = node.position.lerp(Vector3(x, y, z), 0.3)
		node.rotation.y = deg_to_rad(-heading + 180)

func _on_entity_despawned(data: Dictionary) -> void:
	var entity_id: String = data.get("entity_id", "")
	if other_players.has(entity_id):
		var node: Node3D = other_players[entity_id]
		node.queue_free()
		other_players.erase(entity_id)
		chat_log.append_text("[color=gray]A player left the area.[/color]\n")

func _on_chat_received(data: Dictionary) -> void:
	var sender: String = data.get("sender_name", "Unknown")
	var text: String = data.get("text", "")
	var channel: String = data.get("channel", "spatial")
	var prefix: String = ""
	if channel == "shout":
		prefix = "[shouts] "
	elif channel == "planet":
		prefix = "[planet] "
	chat_log.append_text("[color=cyan]" + sender + ":[/color] " + prefix + text + "\n")

func _on_chat_submitted(text: String) -> void:
	if text.strip_edges().is_empty():
		return
	
	# Parse chat command: /shout for shout, /planet for planet-wide, default spatial
	var channel := "spatial"
	var message := text
	
	if text.begins_with("/shout "):
		channel = "shout"
		message = text.substr(7)
	elif text.begins_with("/planet "):
		channel = "planet"
		message = text.substr(8)
	
	if message.strip_edges().is_empty():
		return
	
	ws_client.send_chat(channel, message)
	chat_input.clear()
