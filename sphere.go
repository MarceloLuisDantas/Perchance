package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sphere struct {
	Name     string
	Damage   int
	Hp       int
	Radius   float32
	Color    rl.Color
	Position rl.Vector2
	Velocity rl.Vector2
	Mass     float32
	Alive    bool
}

func NewSphere(
	name string, damage int, hp int,
	radius float32, color rl.Color,
	initial_pos rl.Vector2,
	initial_vel rl.Vector2,
) Sphere {
	s := Sphere{
		Name:     name,
		Damage:   damage,
		Hp:       hp,
		Radius:   radius,
		Color:    color,
		Position: initial_pos,
		Velocity: initial_vel,
		Mass:     1,
		Alive:    true,
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

	text_s := rl.MeasureText(strconv.Itoa(self.Hp), 30)
	rl.DrawText(
		strconv.Itoa(self.Hp),
		int32(self.Position.X-float32(text_s/2)),
		int32(self.Position.Y+self.Radius),
		30, rl.Black,
	)
}

func (self *Sphere) Move() {
	self.Position.X += self.Velocity.X
	self.Position.Y += self.Velocity.Y
}

func (self *Sphere) SetNewVelocity(vel rl.Vector2) {
	self.Velocity = vel
}
