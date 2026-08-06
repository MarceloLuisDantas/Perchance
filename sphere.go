package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Sphere struct {
	Name     string
	Radius   float32
	Color    rl.Color
	Position rl.Vector2
	Velocity rl.Vector2
	Mass     float32
}

func NewSphere(
	name string, radius float32, color rl.Color,
	initial_pos rl.Vector2, initial_vel rl.Vector2,
) Sphere {
	s := Sphere{
		Name:     name,
		Radius:   radius,
		Color:    color,
		Position: initial_pos,
		Velocity: initial_vel,
		Mass:     1,
	}
	return s
}

func (self *Sphere) Draw() {
	rl.DrawCircle(
		int32(self.Position.X),
		int32(self.Position.Y),
		self.Radius,
		self.Color,
	)
}

func (self *Sphere) Move() {
	self.Position.X += self.Velocity.X
	self.Position.Y += self.Velocity.Y
}

func (self *Sphere) SetNewVelocity(vel rl.Vector2) {
	self.Velocity = vel
}
