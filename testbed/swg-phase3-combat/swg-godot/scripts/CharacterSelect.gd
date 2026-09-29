# CharacterSelect.gd
# Character selection and creation screen.
# Lists existing characters, allows creating new ones with species selection.
# GDD Section 6: Character creation with species, appearance, and HAM display.

extends Control

@onready var character_list: ItemList = $VBox/LeftPanel/CharacterList
@onready var create_btn: Button = $VBox/LeftPanel/CreateButton
@onready var enter_btn: Button = $VBox/LeftPanel/EnterButton
@onready var delete_btn: Button = $VBox/LeftPanel/DeleteButton
@onready var status_label: Label = $VBox/LeftPanel/StatusLabel

# Character creation panel
@onready var create_panel: Panel = $VBox/RightPanel
@onready var name_edit: LineEdit = $VBox/RightPanel/VBox/NameEdit
@onready var species_dropdown: OptionButton = $VBox/RightPanel/VBox/SpeciesDropdown
@onready var confirm_create_btn: Button = $VBox/RightPanel/VBox/ConfirmCreateButton
@onready var cancel_create_btn: Button = $VBox/RightPanel/VBox/CancelButton

@onready var api_client: SWGApiClient = $SWGApiClient

var characters: Array = []
var selected_index: int = -1
var species_data: Array = []

func _ready() -> void:
	create_btn.pressed.connect(_on_create_new)
	enter_btn.pressed.connect(_on_enter_world)
	cancel_create_btn.pressed.connect(_on_cancel_create)
	confirm_create_btn.pressed.connect(_on_confirm_create)
	character_list.item_selected.connect(_on_character_selected)
	
	api_client.characters_listed.connect(_on_characters_listed)
	api_client.species_received.connect(_on_species_received)
	api_client.character_created.connect(_on_character_created)
	api_client.character_creation_failed.connect(_on_creation_failed)
	api_client.character_received.connect(_on_character_data)
	
	# Propagate auth token from global state to this scene's API client
	api_client.token = GameState.token
	api_client.account_id = GameState.account_id
	
	create_panel.visible = false
	enter_btn.disabled = true
	
	# Load species list and characters
	api_client.get_species()
	api_client.list_characters()

func _on_species_received(species_list: Array) -> void:
	species_data = species_list
	species_dropdown.clear()
	for sp in species_list:
		species_dropdown.add_item(sp["display_name"])
		species_dropdown.set_item_metadata(species_dropdown.item_count - 1, sp["id"])

func _on_characters_listed(chars: Array) -> void:
	characters = chars
	character_list.clear()
	for c in chars:
		var label := "%s — %s (%s)" % [c["name"], c["species"], c["planet"]]
		character_list.add_item(label)
	if characters.size() > 0:
		character_list.select(0)
		_on_character_selected(0)

func _on_character_selected(index: int) -> void:
	selected_index = index
	enter_btn.disabled = false

func _on_create_new() -> void:
	create_panel.visible = true
	name_edit.text = ""
	name_edit.grab_focus()

func _on_cancel_create() -> void:
	create_panel.visible = false

func _on_confirm_create() -> void:
	var char_name := name_edit.text.strip_edges()
	if char_name.length() < 3:
		status_label.text = "Name must be 3-30 characters"
		return
	
	var species_idx := species_dropdown.selected
	if species_idx < 0:
		status_label.text = "Select a species"
		return
	
	var species_id: String = species_dropdown.get_item_metadata(species_idx)
	
	var appearance := {
		"body_type": "average",
		"skin_color": "default",
		"hair_style": "none",
		"hair_color": "none",
		"face_type": "default",
		"eye_color": "brown",
		"height": 1.0
	}
	
	status_label.text = "Creating character..."
	confirm_create_btn.disabled = true
	api_client.create_character(species_id, char_name, appearance)

func _on_character_created(char_id: String, name: String, species: String) -> void:
	status_label.text = "Character created: " + name
	confirm_create_btn.disabled = false
	create_panel.visible = false
	# Refresh character list
	api_client.list_characters()

func _on_creation_failed(error: String) -> void:
	status_label.text = "Creation failed: " + error
	confirm_create_btn.disabled = false

func _on_character_data(data: Dictionary) -> void:
	# Store character data for the world scene
	GameState.character_data = data

func _on_enter_world() -> void:
	if selected_index < 0 or selected_index >= characters.size():
		return
	
	var char_id: String = characters[selected_index]["id"]
	# Fetch full character data (with HAM) before entering world
	api_client.get_character(char_id)
	# Wait for character_received signal, then switch to world scene
	# The world scene will handle the WebSocket connection
	GameState.selected_character_id = char_id
	get_tree().change_scene_to_file("res://scenes/World.tscn")
