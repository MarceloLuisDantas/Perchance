package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Character struct {
	Name   string
	Radius float32
	Color  rl.Color
}

func NewCharacter(name string, radius float32, color rl.Color) Character {
	return Character{
		Name:   name,
		Radius: radius,
		Color:  color,
	}
}
