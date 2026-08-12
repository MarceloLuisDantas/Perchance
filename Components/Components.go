package Components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	Label string
	Font  int32
	Body  rl.Rectangle
	Color rl.Color
}

func NewButton(
	label string, font int32, x int32, y int32,
	width int32, height int32, color rl.Color,
) Button {
	return Button{
		Label: label,
		Font:  font,
		Body: rl.Rectangle{
			X:      float32(x),
			Y:      float32(y),
			Width:  float32(width),
			Height: float32(height),
		},
		Color: color,
	}
}

func (b *Button) Render() {
	mouse := rl.GetMousePosition()

	text_s := rl.MeasureText(b.Label, b.Font)
	rl.DrawText(
		b.Label,
		b.Body.ToInt32().X+(int32(b.Body.Width)-text_s)/2,
		b.Body.ToInt32().Y+5,
		b.Font, rl.Black,
	)

	if (mouse.X >= b.Body.X) && (mouse.X <= b.Body.X+b.Body.Width) && (mouse.Y >= b.Body.Y) && (mouse.Y <= b.Body.Y+b.Body.Height) {
		rl.DrawRectangleLinesEx(b.Body, 7, b.Color)
	} else {
		rl.DrawRectangleLinesEx(b.Body, 5, b.Color)
	}
}

type BorderLessButton struct {
	Label string
	Font  int32
	Body  rl.Rectangle
	Color rl.Color
}

func NewBorderlessButton(
	label string, font int32, x int32, y int32,
	width int32, height int32, color rl.Color,
) BorderLessButton {
	return BorderLessButton{
		Label: label,
		Font:  font,
		Body: rl.Rectangle{
			X:      float32(x),
			Y:      float32(y),
			Width:  float32(width),
			Height: float32(height),
		},
		Color: color,
	}
}

func (b *BorderLessButton) Render() {
	mouse := rl.GetMousePosition()

	text_s := rl.MeasureText(b.Label, b.Font)
	if (mouse.X >= b.Body.X) && (mouse.X <= b.Body.X+b.Body.Width) && (mouse.Y >= b.Body.Y) && (mouse.Y <= b.Body.Y+b.Body.Height) {
		rl.DrawText(
			b.Label,
			b.Body.ToInt32().X+(int32(b.Body.Width)-text_s)/2,
			b.Body.ToInt32().Y+5,
			b.Font, rl.Gray,
		)
	} else {
		rl.DrawText(
			b.Label,
			b.Body.ToInt32().X+(int32(b.Body.Width)-text_s)/2,
			b.Body.ToInt32().Y+5,
			b.Font, rl.Black,
		)
	}
	// rl.DrawRectangleLinesEx(b.Body, 5, b.Color)
}

type Text struct {
	Label string
	Font  int32
	Start rl.Vector2
	Color rl.Color
}

func NewText(label string, font int32, start rl.Vector2, color rl.Color) Text {
	return Text{
		Label: label,
		Font:  font,
		Start: start,
		Color: color,
	}
}

func (t *Text) Render() {
	rl.DrawText(
		t.Label, int32(t.Start.X), int32(t.Start.Y), t.Font, t.Color,
	)
}
