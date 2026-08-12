package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameState int

const (
	MAIN_MENU GameState = iota
	CHAR_SELECT
	BATTLE
)

const (
	HEIGHT = 800
	WIDTH  = 800
)

func main() {

	rl.InitWindow(WIDTH, HEIGHT, "Perchance")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	main_menu := NewMainMenu()
	char_select := NewCharacterSelection()

	state := MAIN_MENU
game_loop:
	for !rl.WindowShouldClose() {
		switch state {
		case MAIN_MENU:
			action := main_menu.Update()
			if action == "EXIT" {
				break game_loop
			} else if action == "START" {
				state = CHAR_SELECT
			}
		case CHAR_SELECT:
			_ = char_select.Update()
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.White)
		switch state {
		case MAIN_MENU:
			main_menu.Render()
		case CHAR_SELECT:
			char_select.Render()
		}

		rl.EndDrawing()
	}

}
