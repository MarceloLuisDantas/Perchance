package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 800, "Perchance")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	spheres := make([]Sphere, 2)
	spheres[0] = NewSphere(
		"Badingugilius", 9, 420, 50, rl.Red,
		rl.Vector2{X: 200, Y: 450},
		rl.Vector2{X: 5, Y: 0},
	)
	spheres[1] = NewSphere(
		"Smungulungus", 10, 400, 50, rl.Blue,
		rl.Vector2{X: 600, Y: 500},
		rl.Vector2{X: -10, Y: 0},
	)

	for !rl.WindowShouldClose() {
		if CheckCollisionSpheres(&spheres[0], &spheres[1]) {
			spheres[0].Hp -= spheres[1].Damage
			spheres[1].Hp -= spheres[0].Damage
		}

		spheres[0].Move()
		spheres[1].Move()

		rl.BeginDrawing()

		rl.ClearBackground(rl.White)
		rl.DrawRectangleLinesEx(
			rl.Rectangle{X: 20, Y: 200, Width: 760, Height: 580},
			5, rl.Black,
		)

		if spheres[0].Hp > 0 {
			spheres[0].Draw()
		}

		if spheres[1].Hp > 0 {
			spheres[1].Draw()
		}

		rl.EndDrawing()
	}
}
