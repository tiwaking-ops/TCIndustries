# SWGApiClient.gd
# Handles all HTTP API communication with the SWG server.
# Endpoints: register, login, species list, character CRUD.
# GDD Section 29.2.4: non-real-time operations use request/response API.

extends Node
class_name SWGApiClient

signal register_success(account_id: String)
signal register_failed(error: String)
signal login_success(account_id: String, token: String)
signal login_failed(error: String)
signal species_received(species_list: Array)
signal character_created(character_id: String, name: String, species: String)
signal character_creation_failed(error: String)
signal characters_listed(characters: Array)
signal character_received(character_data: Dictionary)

const SERVER_URL := "http://localhost:8080"
var token: String = ""
var account_id: String = ""

func _ready() -> void:
	# Load server URL from env or config if needed
	pass

# POST /api/register
func register(username: String, password: String) -> void:
	var body := {"username": username, "password": password}
	var http := HTTPRequest.new()
	add_child(http)
	http.request_completed.connect(_on_register_completed)
	http.request(SERVER_URL + "/api/register", ["Content-Type: application/json"], HTTPClient.METHOD_POST, JSON.stringify(body))

func _on_register_completed(result: int, response_code: int, headers: PackedStringArray, body: PackedByteArray) -> void:
	var data = JSON.parse_string(body.get_string_from_utf8())
	if response_code == 201:
		account_id = data["account_id"]
		register_success.emit(data["account_id"])
	else:
		var err = data["error"] if data else "Unknown error"
		register_failed.emit(err)

# POST /api/login
func login(username: String, password: String) -> void:
	var body := {"username": username, "password": password}
	var http := HTTPRequest.new()
	add_child(http)
	http.request_completed.connect(_on_login_completed)
	http.request(SERVER_URL + "/api/login", ["Content-Type: application/json"], HTTPClient.METHOD_POST, JSON.stringify(body))

func _on_login_completed(result: int, response_code: int, headers: PackedStringArray, body: PackedByteArray) -> void:
	var data = JSON.parse_string(body.get_string_from_utf8())
	if response_code == 200:
		account_id = data["account_id"]
		token = data["token"]
		login_success.emit(data["account_id"], data["token"])
	else:
		var err = data["error"] if data else "Unknown error"
		login_failed.emit(err)

# GET /api/species
func get_species() -> void:
	var http := HTTPRequest.new()
	add_child(http)
	http.request_completed.connect(_on_species_completed)
	http.request(SERVER_URL + "/api/species")

func _on_species_completed(result: int, response_code: int, headers: PackedStringArray, body: PackedByteArray) -> void:
	var data = JSON.parse_string(body.get_string_from_utf8())
	if response_code == 200:
		species_received.emit(data["species"])

# POST /api/characters
func create_character(species_id: String, char_name: String, appearance: Dictionary) -> void:
	var body := {
		"species": species_id,
		"name": char_name,
		"appearance": appearance
	}
	var http := HTTPRequest.new()
	add_child(http)
	http.request_completed.connect(_on_create_character_completed)
	http.request(
		SERVER_URL + "/api/characters",
		["Content-Type: application/json", "Authorization: Bearer " + token],
		HTTPClient.METHOD_POST,
		JSON.stringify(body)
	)

func _on_create_character_completed(result: int, response_code: int, headers: PackedStringArray, body: PackedByteArray) -> void:
	var data = JSON.parse_string(body.get_string_from_utf8())
	if response_code == 201:
		character_created.emit(data["character_id"], data["name"], data["species"])
	else:
		var err = data["error"] if data else "Unknown error"
		character_creation_failed.emit(err)

# GET /api/characters
func list_characters() -> void:
	var http := HTTPRequest.new()
	add_child(http)
	http.request_completed.connect(_on_list_characters_completed)
	http.request(
		SERVER_URL + "/api/characters",
		["Authorization: Bearer " + token]
	)

func _on_list_characters_completed(result: int, response_code: int, headers: PackedStringArray, body: PackedByteArray) -> void:
	var data = JSON.parse_string(body.get_string_from_utf8())
	if response_code == 200:
		characters_listed.emit(data["characters"])

# GET /api/characters/{id}
func get_character(character_id: String) -> void:
	var http := HTTPRequest.new()
	add_child(http)
	http.request_completed.connect(_on_get_character_completed)
	http.request(
		SERVER_URL + "/api/characters/" + character_id,
		["Authorization: Bearer " + token]
	)

func _on_get_character_completed(result: int, response_code: int, headers: PackedStringArray, body: PackedByteArray) -> void:
	var data = JSON.parse_string(body.get_string_from_utf8())
	if response_code == 200:
		character_received.emit(data)
