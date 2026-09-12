# World.gd
# The main game world scene. Connects to the server via WebSocket,
# enters the world with the selected character, and renders the character
# standing in a bare world with HAM display.
# GDD Phase 0 exit criteria: "see it standing in a bare world with correct starting HAM values."

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

func _ready() -> void:
	# Connect WebSocket signals
	ws_client.connected.connect(_on_ws_connected)
	ws_client.disconnected.connect(_on_ws_disconnected)
	ws_client.world_entered.connect(_on_world_entered)
	ws_client.entity_spawned.connect(_on_entity_spawned)
	ws_client.entity_moved.connect(_on_entity_moved)
	ws_client.entity_despawned.connect(_on_entity_despawned)
	ws_client.chat_received.connect(_on_chat_received)
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

func _on_ws_error(message: String) -> void:
	chat_log.append_text("[color=red]Error: " + message + "[/color]\n")

func _on_world_entered(data: Dictionary) -> void:
	# This is the Phase 0 verification: character standing in world with HAM values
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

func _on_entity_spawned(data: Dictionary) -> void:
	var entity_id: String = data.get("entity_id", "")
	var entity_name: String = data.get("name", "Unknown")
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
	
	chat_log.append_text(entity_name + " entered the area.\n")

func _on_entity_moved(data: Dictionary) -> void:
	var entity_id: String = data.get("entity_id", "")
	var pos: Dictionary = data.get("position", {})
	
	if other_players.has(entity_id):
		var node: Node3D = other_players[entity_id]
		var x := float(pos.get("x", 0))
		var y := float(pos.get("y", 0))
		var z := float(pos.get("z", 0))
		node.position = Vector3(x, y, z)

func _on_entity_despawned(data: Dictionary) -> void:
	var entity_id: String = data.get("entity_id", "")
	if other_players.has(entity_id):
		var node: Node3D = other_players[entity_id]
		node.queue_free()
		other_players.erase(entity_id)
		chat_log.append_text("A player left the area.\n")

func _on_chat_received(data: Dictionary) -> void:
	var sender: String = data.get("sender_name", "Unknown")
	var text: String = data.get("text", "")
	chat_log.append_text("[color=cyan]" + sender + ":[/color] " + text + "\n")

func _on_chat_submitted(text: String) -> void:
	if text.strip_edges().is_empty():
		return
	ws_client.send_chat("spatial", text)
	chat_input.clear()

func _process(_delta: float) -> void:
	# Update position label with current player position
	position_label.text = "Pos: (%.1f, %.1f, %.1f)" % [
		player_node.position.x,
		player_node.position.y,
		player_node.position.z
	]
