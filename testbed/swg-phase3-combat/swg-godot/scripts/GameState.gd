# GameState.gd
# Global autoload singleton for sharing state between scenes.
# Add this as an Autoload named "GameState" in Project Settings.

extends Node

var token: String = ""
var account_id: String = ""
var selected_character_id: String = ""
var character_data: Dictionary = {}
