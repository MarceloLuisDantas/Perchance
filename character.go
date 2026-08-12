package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Character struct {
	Name     string
	Radius   float32
	Hp       int32
	Mass     float32
	MaxHp    int32
	Damage   int32
	Position rl.Vector2
	Velocity rl.Vector2
	Color    rl.Color
}

func NewCharacter(name string, radius float32, hp, damage int32, color rl.Color) Character {
	return Character{
		Name:   name,
		Radius: radius,
		Mass:   1,
		Hp:     hp,
		MaxHp:  hp,
		Damage: damage,
		Color:  color,
	}
}
