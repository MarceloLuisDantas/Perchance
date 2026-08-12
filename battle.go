package main

import rl "github.com/gen2brain/raylib-go/raylib"

type BattleState int

const (
	WAITING BattleState = iota
	FIGHTING
	P1_WIN
	P2_WIN
	DRAW
)

type Battle struct {
	P1    Character
	P2    Character
	State BattleState
	Arena rl.Rectangle
}

func NewBattle(p1, p2 Character) Battle {
	return Battle{
		P1:    p1,
		P2:    p2,
		State: WAITING,
		Arena: rl.Rectangle{
			X: 30, Y: 250, Width: 740, Height: 520,
		},
	}
}

func RenderLifeBar(player, max, current int32) {
	border := rl.Rectangle{
		X: 30, Y: 200, Width: 200, Height: 40,
	}

	if player == 2 {
		border.X += 540
	}

	current_hp_percentage := (current * 100) / max
	current_bar := (current_hp_percentage * border.ToInt32().Width) / 100

	if player == 2 {
		rl.DrawRectangle(
			border.ToInt32().X+200-current_bar,
			border.ToInt32().Y,
			current_bar,
			border.ToInt32().Height,
			rl.Green,
		)
	} else {
		rl.DrawRectangle(
			border.ToInt32().X,
			border.ToInt32().Y,
			current_bar,
			border.ToInt32().Height,
			rl.Green,
		)
	}
	rl.DrawRectangleLinesEx(border, 4, rl.Black)
}

func (b *Battle) Render() {
	p1 := b.P1
	rl.DrawText(p1.Name, 30, 30, 30, rl.Black)
	rl.DrawCircle(90, 130, 60, p1.Color)
	RenderLifeBar(1, p1.MaxHp, p1.Hp)
	print(b.P1.Hp)
	b.P1.Hp -= 1

	p2 := b.P2
	text_s := rl.MeasureText(p2.Name, 30)
	rl.DrawText(p2.Name, 770-text_s, 30, 30, rl.Black)
	rl.DrawCircle(710, 130, 60, p2.Color)
	RenderLifeBar(2, p2.MaxHp, p2.Hp)
	b.P2.Hp -= 1

	rl.DrawRectangleLinesEx(b.Arena, 5, rl.Black)
}

func (b *Battle) Update() string {
	return "running"
}
