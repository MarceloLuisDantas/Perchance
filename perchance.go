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
	rl.SetTargetFPS(30)

	main_menu := NewMainMenu()
	char_select := NewCharacterSelection()
	var battle Battle

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
			action := char_select.Update()
			switch action {
			case "BACK":
				state = MAIN_MENU
			case "BATTLE":
				p1 := char_select.Chars[char_select.Player1]
				p2 := char_select.Chars[char_select.Player2]
				battle = NewBattle(p1, p2)
				state = BATTLE
			}
		case BATTLE:
			println(battle.Update())
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.White)
		switch state {
		case MAIN_MENU:
			main_menu.Render()
		case CHAR_SELECT:
			char_select.Render()
		case BATTLE:
			battle.Render()
		}

		rl.EndDrawing()
	}

}
