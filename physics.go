package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func CheckCollisionSphereMap(s *Sphere) {
	if s.Position.X-s.Radius <= 20 || s.Position.X+s.Radius >= 780 {
		s.Velocity.X *= -1
	}

	if s.Position.Y-s.Radius <= 200 || s.Position.Y+s.Radius >= 780 {
		s.Velocity.Y *= -1
	}
}

func CheckCollisionSpheres(s1, s2 *Sphere) bool {
	collided := false
	if rl.CheckCollisionCircles(s1.Position, s1.Radius, s2.Position, s2.Radius) {
		collided = true
		normalized_pos := rl.Vector2Subtract(s1.Position, s2.Position)
		distance := rl.Vector2Length(normalized_pos)
		normalized_pos = rl.Vector2Scale(normalized_pos, 1.0/distance)
		relative_vel := rl.Vector2Subtract(s1.Velocity, s2.Velocity)
		angular_vel := rl.Vector2DotProduct(relative_vel, normalized_pos)
		if angular_vel > 0 {
			return collided
		}

		scalar_impulse := -(1 + 1.0) * angular_vel
		scalar_impulse /= (1.0/s1.Mass + 1.0/s2.Mass)
		impulse := rl.Vector2Scale(normalized_pos, scalar_impulse)
		s1.Velocity = rl.Vector2Add(s1.Velocity, rl.Vector2Scale(impulse, 1.0/s1.Mass))
		s2.Velocity = rl.Vector2Subtract(s2.Velocity, rl.Vector2Scale(impulse, 1.0/s2.Mass))
	}

	return collided
}

func RunPhysics(spheres []Sphere) {
	for i := range spheres {
		CheckCollisionSphereMap(&spheres[i])
	}
	for i := 0; i < len(spheres)-1; i++ {
		for j := i + 1; j < len(spheres); j++ {
			if !spheres[i].Alive || !spheres[j].Alive {
				continue
			}

			if CheckCollisionSpheres(&spheres[i], &spheres[j]) {
				spheres[i].Hp -= spheres[j].Damage
				spheres[j].Hp -= spheres[i].Damage
			}
		}
	}

	for i := range spheres {
		spheres[i].Move()
	}
}
