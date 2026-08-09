package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type State int

const (
	FIGHTING State = iota
	DRAW
	WON
)

type GameState int

const (
	MAIN_MENU GameState = iota
)

const (
	HEIGHT = 800
	WIDTH  = 800
)

func main() {
	// state := MAIN_MENU

	rl.InitWindow(WIDTH, HEIGHT, "Perchance")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	main_menu := NewMainMenu()

	for !rl.WindowShouldClose() {
		next, bt := main_menu.Update()
		if next {
			println(bt)
		}

		rl.BeginDrawing()

		main_menu.Render()
		rl.ClearBackground(rl.White)

		rl.EndDrawing()

		// if state == WON {
		// 	var winner string
		// 	for _, s := range spheres {
		// 		if s.Alive {
		// 			winner = s.Name
		// 			break
		// 		}
		// 	}
		// 	text_s := rl.MeasureText(winner, 100)
		// 	rl.DrawText(
		// 		winner,
		// 		int32(400-float32(text_s/2)),
		// 		int32(400),
		// 		100, rl.Black,
		// 	)
		// } else if state == DRAW {
		// 	text_s := rl.MeasureText("EMPATE", 100)
		// 	rl.DrawText(
		// 		"EMPATE",
		// 		int32(400-float32(text_s/2)),
		// 		int32(400),
		// 		100, rl.Black,
		// 	)
		// }
	}
}
