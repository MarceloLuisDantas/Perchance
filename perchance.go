package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type State int

const (
	FIGHTING State = iota
	DRAW
	WON
)

func main() {
	state := FIGHTING

	rl.InitWindow(800, 800, "Perchance")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	spheres := make([]Sphere, 0)
	spheres = append(spheres, NewSphere(
		"Badingugilius", 200, 400, 50, rl.Red,
		rl.Vector2{
			X: float32(200 + rand.Int31n(150)),
			Y: float32(450 + rand.Int31n(150))},
		rl.Vector2{X: 5, Y: 0},
	))

	spheres = append(spheres, NewSphere(
		"Smungulungus", 200, 400, 50, rl.Blue,
		rl.Vector2{
			X: 600 + rand.Float32(),
			Y: 500 + rand.Float32(),
		},
		rl.Vector2{X: -10, Y: 0},
	))

	for !rl.WindowShouldClose() {
		if state == FIGHTING {
			RunPhysics(spheres)
		}
		state = RunLogic(spheres)

		rl.BeginDrawing()

		rl.ClearBackground(rl.White)
		rl.DrawRectangleLinesEx(
			rl.Rectangle{X: 20, Y: 200, Width: 760, Height: 580},
			5, rl.Black,
		)

		for i := range spheres {
			if spheres[i].Alive {
				spheres[i].Draw()
			}
		}

		if state == WON {
			var winner string
			for _, s := range spheres {
				if s.Alive {
					winner = s.Name
					break
				}
			}
			text_s := rl.MeasureText(winner, 100)
			rl.DrawText(
				winner,
				int32(400-float32(text_s/2)),
				int32(400),
				100, rl.Black,
			)
		} else if state == DRAW {
			text_s := rl.MeasureText("EMPATE", 100)
			rl.DrawText(
				"EMPATE",
				int32(400-float32(text_s/2)),
				int32(400),
				100, rl.Black,
			)
		}

		rl.EndDrawing()
	}
}
