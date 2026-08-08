package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 800, "Perchance")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	var spheres []Sphere
	spheres = append(spheres, NewSphere(
		"Badingugilius", 9, 420, 50, rl.Red,
		rl.Vector2{X: 200, Y: 450},
		rl.Vector2{X: 5, Y: 0},
	))

	spheres = append(spheres, NewSphere(
		"Smungulungus", 10, 400, 50, rl.Blue,
		rl.Vector2{X: 600, Y: 500},
		rl.Vector2{X: -10, Y: 0},
	))

	for !rl.WindowShouldClose() {
		for _, s := range spheres {
			CheckCollisionSphereMap(&s)
		}

		for i := 0; i < len(spheres)-2; i++ {
			s1 := &spheres[i]
			for j := i + 1; j < len(spheres)-1; j++ {
				s2 := &spheres[j]
				if CheckCollisionSpheres(s1, s2) {
					s1.Hp -= s2.Damage
					s2.Hp -= s1.Damage
				}
			}
		}

		for i := range spheres {
			if spheres[i].Hp <= 0 {
				spheres = append(spheres[:i], spheres[i+1:]...)
			}
			spheres[i].Move()
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.White)
		rl.DrawRectangleLinesEx(
			rl.Rectangle{
				X: 20, Y: 200, Width: 760, Height: 580,
			},
			5, rl.Black,
		)

		for i := range spheres {
			spheres[i].Draw()
		}

		rl.EndDrawing()
	}
}
