package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Battle() State {
	state := FIGHTING

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

	for {
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

		if state != FIGHTING {
			return state
		}
		rl.EndDrawing()
	}
}
