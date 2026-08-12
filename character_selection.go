package main

import (
	"Perchance/Components"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CharacterSelection struct {
	Chars     []Character
	Texts     []Components.Text
	BLButtons []Components.BorderLessButton
	Player1   int
	Player2   int
}

func NewCharacterSelection() CharacterSelection {
	cs := CharacterSelection{
		Chars:     make([]Character, 4),
		Texts:     make([]Components.Text, 1),
		BLButtons: make([]Components.BorderLessButton, 6),
		Player1:   0,
		Player2:   0,
	}

	cs.Chars[0] = NewCharacter("LineyStiney", 50, 500, 10, rl.Red)
	cs.Chars[1] = NewCharacter("SpikeyPikey", 40, 500, 10, rl.Black)
	cs.Chars[2] = NewCharacter("SpookyDooky", 60, 500, 10, rl.Green)
	cs.Chars[3] = NewCharacter("SickySticky", 40, 500, 10, rl.Blue)

	text_s := rl.MeasureText("SELECT YOUR WOOKY", 50)
	cs.Texts[0] = Components.NewText(
		"SELECT YOUR WOOKY", 50, rl.Vector2{X: float32(400 - text_s/2), Y: 100}, rl.Black,
	)

	text_s = rl.MeasureText("<", 70) + 20
	cs.BLButtons[0] = Components.NewBorderlessButton("<", 70, 110, 287, text_s, 70, rl.Black)
	cs.BLButtons[1] = Components.NewBorderlessButton(">", 70, 300, 287, text_s, 70, rl.Black)
	cs.BLButtons[2] = Components.NewBorderlessButton("<", 70, 460, 287, text_s, 70, rl.Black)
	cs.BLButtons[3] = Components.NewBorderlessButton(">", 70, 650, 287, text_s, 70, rl.Black)

	text_s = rl.MeasureText("BATTLE", 70)
	cs.BLButtons[4] = Components.NewBorderlessButton("BATTLE", 70, 400-(text_s+30)/2, 600, text_s+30, 75, rl.Black)

	cs.BLButtons[5] = Components.NewBorderlessButton("<<", 40, 20, 20, 50, 50, rl.Black)

	return cs
}

func (cs *CharacterSelection) Update() string {
	for i, b := range cs.BLButtons {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), b.Body) &&
			rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			switch i {
			case 0: // p1 prev
				print("Prev to ", cs.Player1, " is: ")
				if cs.Player1 == 0 {
					cs.Player1 = len(cs.Chars) - 1
				} else {
					cs.Player1 -= 1
				}
			case 1: // p1 next
				print("Next to ", cs.Player1, " is: ")
				if cs.Player1 == len(cs.Chars)-1 {
					cs.Player1 = 0
				} else {
					cs.Player1 += 1
				}
				println(cs.Player1)
			case 2: // p2 prev
				print("Prev to ", cs.Player2, " is: ")
				if cs.Player2 == 0 {
					cs.Player2 = len(cs.Chars) - 1
				} else {
					cs.Player2 -= 1
				}
				println(cs.Player1)
			case 3: // p2 next
				print("Next to ", cs.Player2, " is: ")
				if cs.Player2 == len(cs.Chars)-1 {
					cs.Player2 = 0
				} else {
					cs.Player2 += 1
				}
				println(cs.Player1)
			case 4: // battle
				return "BATTLE"
			case 5: // back
				return "BACK"
			}
		}
	}
	return ""
}

func (cs *CharacterSelection) CharSelectionSquare() {
	rl.DrawRectangleLinesEx(
		rl.Rectangle{
			X: 150, Y: 250, Width: 150, Height: 150,
		},
		4, rl.Black,
	)
	p1 := cs.Chars[cs.Player1]
	text_s := rl.MeasureText(p1.Name, 30)
	rl.DrawText(p1.Name, 225-text_s/2, 200, 30, rl.Black)
	rl.DrawCircle(225, 325, 60, p1.Color)
	println(p1.Color.A, p1.Color.B, p1.Color.G)

	radius := fmt.Sprintf("Radius: %2.f", p1.Radius)
	text_s = rl.MeasureText(radius, 30)
	rl.DrawText(radius, 150, 410, 30, rl.Black)

	rl.DrawRectangleLinesEx(
		rl.Rectangle{
			X: 500, Y: 250, Width: 150, Height: 150,
		},
		4, rl.Black,
	)
	p2 := cs.Chars[cs.Player2]
	text_s = rl.MeasureText(p2.Name, 30)
	rl.DrawText(p2.Name, 575-text_s/2, 200, 30, rl.Black)
	rl.DrawCircle(575, 325, 60, p2.Color)

	radius = fmt.Sprintf("Radius: %2.f", p2.Radius)
	text_s = rl.MeasureText(radius, 30)
	rl.DrawText(radius, 500, 410, 30, rl.Black)
}

func (cs *CharacterSelection) Render() {
	cs.CharSelectionSquare()

	for _, e := range cs.Texts {
		e.Render()
	}

	for _, e := range cs.BLButtons {
		e.Render()
	}
}
