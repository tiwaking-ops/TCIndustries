# WSClient.gd
# WebSocket client for real-time world communication.
# GDD Section 29.2.4: real-time state sync over persistent connection.
# Handles: enter_world, entity spawn/move/despawn, chat.

extends Node
class_name WSClient

signal connected()
signal disconnected()
signal world_entered(data: Dictionary)
signal entity_spawned(data: Dictionary)
signal entity_moved(data: Dictionary)
signal entity_despawned(data: Dictionary)
signal chat_received(data: Dictionary)
signal error_received(message: String)

const WS_URL := "ws://localhost:8080/ws"

var socket: WebSocketPeer
var is_connected_flag := false
var token: String = ""

func _ready() -> void:
	socket = WebSocketPeer.new()

func connect_to_server(auth_token: String) -> void:
	token = auth_token
	var url := WS_URL + "?token=" + token
	var err := socket.connect_to_url(url)
	if err != OK:
		error_received.emit("Failed to connect to WebSocket: " + str(err))
		return

func _process(_delta: float) -> void:
	if socket == null:
		return
	
	socket.poll()
	var state := socket.get_ready_state()
	
	if state == WebSocketPeer.STATE_OPEN:
		if not is_connected_flag:
			is_connected_flag = true
			connected.emit()
		
		# Read all available messages
		while socket.get_available_packet_count() > 0:
			var packet := socket.get_packet()
			var message := packet.get_string_from_utf8()
			_handle_message(message)
	
	elif state == WebSocketPeer.STATE_CLOSED:
		if is_connected_flag:
			is_connected_flag = false
			disconnected.emit()

func _handle_message(message: String) -> void:
	var data = JSON.parse_string(message)
	if data == null:
		return
	
	var msg_type: String = data.get("type", "")
	var msg_data = data.get("data", {})
	
	match msg_type:
		"world_enter":
			world_entered.emit(msg_data)
		"entity_spawn":
			entity_spawned.emit(msg_data)
		"entity_move":
			entity_moved.emit(msg_data)
		"entity_despawn":
			entity_despawned.emit(msg_data)
		"chat_message":
			chat_received.emit(msg_data)
		"error":
			error_received.emit(msg_data.get("message", "Unknown error"))
		_:
			print("Unknown message type: ", msg_type)

# Send enter_world request
func send_enter_world(character_id: String) -> void:
	var msg := {
		"type": "enter_world",
		"data": {"character_id": character_id}
	}
	_send(msg)

# Send movement update (client-predicted position)
func send_move(x: float, y: float, z: float, heading: float) -> void:
	var msg := {
		"type": "move",
		"data": {"x": x, "y": y, "z": z, "heading": heading}
	}
	_send(msg)

# Send spatial chat message
func send_chat(channel: String, text: String) -> void:
	var msg := {
		"type": "chat",
		"data": {"channel": channel, "text": text}
	}
	_send(msg)

func _send(msg: Dictionary) -> void:
	if socket.get_ready_state() == WebSocketPeer.STATE_OPEN:
		socket.send_text(JSON.stringify(msg))
