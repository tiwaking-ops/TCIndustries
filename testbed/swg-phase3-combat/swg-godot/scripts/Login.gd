# Login.gd
# Login/Register screen controller.
# Allows the player to register a new account or log in to an existing one.

extends Control

@onready var username_edit: LineEdit = $VBox/UsernameEdit
@onready var password_edit: LineEdit = $VBox/PasswordEdit
@onready var status_label: Label = $VBox/StatusLabel
@onready var login_btn: Button = $VBox/HBox/LoginButton
@onready var register_btn: Button = $VBox/HBox/RegisterButton
@onready var api_client: SWGApiClient = $SWGApiClient

func _ready() -> void:
	login_btn.pressed.connect(_on_login)
	register_btn.pressed.connect(_on_register)
	api_client.login_success.connect(_on_login_success)
	api_client.login_failed.connect(_on_login_failed)
	api_client.register_success.connect(_on_register_success)
	api_client.register_failed.connect(_on_register_failed)

func _on_login() -> void:
	var username := username_edit.text.strip_edges()
	var password := password_edit.text
	if username.is_empty() or password.is_empty():
		status_label.text = "Enter username and password"
		return
	status_label.text = "Logging in..."
	login_btn.disabled = true
	register_btn.disabled = true
	api_client.login(username, password)

func _on_register() -> void:
	var username := username_edit.text.strip_edges()
	var password := password_edit.text
	if username.length() < 3:
		status_label.text = "Username must be 3-20 characters"
		return
	if password.length() < 6:
		status_label.text = "Password must be at least 6 characters"
		return
	status_label.text = "Registering..."
	login_btn.disabled = true
	register_btn.disabled = true
	api_client.register(username, password)

func _on_login_success(account_id: String, token: String) -> void:
	status_label.text = "Login successful!"
	# Store token in autoload/global for other scenes
	GameState.token = token
	GameState.account_id = account_id
	get_tree().change_scene_to_file("res://scenes/CharacterSelect.tscn")

func _on_login_failed(error: String) -> void:
	status_label.text = "Login failed: " + error
	login_btn.disabled = false
	register_btn.disabled = false

func _on_register_success(account_id: String) -> void:
	# Auto-login after registration
	status_label.text = "Account created! Logging in..."
	api_client.login(username_edit.text.strip_edges(), password_edit.text)

func _on_register_failed(error: String) -> void:
	status_label.text = "Registration failed: " + error
	login_btn.disabled = false
	register_btn.disabled = false
