package main

import (
	"Perchance/Components"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MainMenu struct {
	Texts  []Components.Text
	Button []Components.Button
}

func NewMainMenu() MainMenu {
	mm := MainMenu{
		Texts:  make([]Components.Text, 1),
		Button: make([]Components.Button, 2),
	}

	text_s := rl.MeasureText("PerChance", 100)
	mm.Texts[0] = Components.NewText(
		"PerChance", 100,
		rl.Vector2{
			X: float32(WIDTH/2 - text_s/2),
			Y: float32(HEIGHT / 3),
		},
		rl.Black,
	)

	text_s = rl.MeasureText("START", 50)
	println(text_s)
	mm.Button[0] = Components.NewButton(
		"START", 50, WIDTH/5, HEIGHT/2, 200, 55, rl.Black,
	)

	mm.Button[1] = Components.NewButton(
		"END", 50, WIDTH/2, HEIGHT/2, 200, 55, rl.Black,
	)

	return mm
}

func (self *MainMenu) Render() {
	for _, text := range self.Texts {
		text.Render()
	}

	for _, button := range self.Button {
		button.Render()
	}
}

func (self *MainMenu) Update() (bool, string) {
	mouse := rl.GetMousePosition()
	for _, b := range self.Button {
		if (mouse.X >= b.Body.X) && (mouse.X <= b.Body.X+b.Body.Width) && (mouse.Y >= b.Body.Y) && (mouse.Y <= b.Body.Y+b.Body.Height) {
			if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
				return true, b.Label
			}
		}
	}
	return false, ""
}
