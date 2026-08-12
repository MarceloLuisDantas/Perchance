package main

import (
	"Perchance/Components"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type BattleState int

const (
	WAITING BattleState = iota
	FIGHTING
	P1_WIN
	P2_WIN
	DRAW
)

type Battle struct {
	P1       Character
	P2       Character
	State    BattleState
	Arena    rl.Rectangle
	Count    int
	Sup      int
	MainMenu Components.Button
}

func NewBattle(p1, p2 Character) Battle {
	text_s := rl.MeasureText("MAIN MENU", 50)
	return Battle{
		P1:    p1,
		P2:    p2,
		State: WAITING,
		Arena: rl.Rectangle{
			X: 30, Y: 250, Width: 740, Height: 520,
		},
		Count: 3,
		Sup:   0,
		MainMenu: Components.NewButton(
			"MAIN MENU", 50, 400-text_s/2, 600, 310, 55, rl.Black,
		),
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

func (b *Battle) RenderBattle() {
	p1 := b.P1
	p2 := b.P2

	rl.DrawText(p1.Name, 30, 30, 30, rl.Black)
	rl.DrawCircle(90, 130, 60, p1.Color)
	RenderLifeBar(1, p1.MaxHp, p1.Hp)

	text_s := rl.MeasureText(p2.Name, 30)
	rl.DrawText(p2.Name, 770-text_s, 30, 30, rl.Black)
	rl.DrawCircle(710, 130, 60, p2.Color)
	RenderLifeBar(2, p2.MaxHp, p2.Hp)

	rl.DrawRectangleLinesEx(b.Arena, 5, rl.Black)
	if p1.Hp > 0 {
		rl.DrawCircle(
			int32(b.P1.Position.X),
			int32(p1.Position.Y),
			p1.Radius, p1.Color,
		)
	}
	if p2.Hp > 0 {
		rl.DrawCircle(
			int32(b.P2.Position.X),
			int32(p2.Position.Y),
			p2.Radius, p2.Color,
		)
	}
}

func (b *Battle) RenderWait() {
	b.RenderBattle()
	text_s := rl.MeasureText(strconv.Itoa(b.Count), 300)
	rl.DrawText(strconv.Itoa(b.Count), 420-text_s/2, 400, 200, rl.Black)
	b.Sup += 1
	if b.Sup == 60 {
		b.Count -= 1
		b.Sup = 0
		if b.Count == 0 {
			b.State = FIGHTING
		}
	}
}

func DrawOutlinedText(text string, posX, posY, fontSize int32, color rl.Color, outlineSize int32) {
	rl.DrawText(text, posX, posY-outlineSize, fontSize, rl.White)
	rl.DrawText(text, posX-outlineSize, posY-outlineSize, fontSize, rl.White)
	rl.DrawText(text, posX+outlineSize, posY-outlineSize, fontSize, rl.White)
	rl.DrawText(text, posX-outlineSize, posY+outlineSize, fontSize, rl.White)
	rl.DrawText(text, posX+outlineSize, posY+outlineSize, fontSize, rl.White)
	rl.DrawText(text, posX, posY+outlineSize, fontSize, rl.White)
	rl.DrawText(text, posX, posY, fontSize, color)
}

func (b *Battle) RenderEnd() {
	p1 := b.P1
	p2 := b.P2

	b.RenderBattle()
	var text string
	switch b.State {
	case DRAW:
		text = "DRAW"
	case P1_WIN:
		text = p1.Name
		text_s := rl.MeasureText("P1 WIN", 50)
		DrawOutlinedText("P1 WIN", 400-text_s/2, 500, 50, rl.Black, 5)
	case P2_WIN:
		text = p2.Name
		text_s := rl.MeasureText("P2 WIN", 50)
		DrawOutlinedText("P2 WIN", 400-text_s/2, 500, 50, rl.Black, 5)
	}
	text_s := rl.MeasureText(text, 100)
	DrawOutlinedText(text, 400-text_s/2, 380, 100, rl.Black, 5)

	b.MainMenu.Render()
}

func (b *Battle) Render() {
	switch b.State {
	case WAITING:
		b.RenderWait()
	case FIGHTING:
		b.RenderBattle()
	default:
		b.RenderEnd()
	}
}

func (b *Battle) CheckColissionArena(player Character) []bool {
	return []bool{
		(player.Position.Y-player.Radius <= b.Arena.Y),
		(player.Position.X-player.Radius <= b.Arena.X),
		(player.Position.X+player.Radius >= b.Arena.X+b.Arena.Width),
		(player.Position.Y+player.Radius >= b.Arena.Y+b.Arena.Height),
	}
}

func (b *Battle) Update() string {
	if b.State == FIGHTING {
		p1 := &b.P1
		p1_arena_c := b.CheckColissionArena(*p1)
		if p1_arena_c[0] || p1_arena_c[3] { // top bottom
			p1.Velocity.Y *= -1
		}
		if p1_arena_c[1] || p1_arena_c[2] { // left right
			p1.Velocity.X *= -1
		}

		p2 := &b.P2
		p2_arena_c := b.CheckColissionArena(*p2)
		if p2_arena_c[0] || p2_arena_c[3] { // top
			p2.Velocity.Y *= -1
		}
		if p2_arena_c[1] || p2_arena_c[2] { // left
			p2.Velocity.X *= -1
		}

		if rl.CheckCollisionCircles(p1.Position, p1.Radius, p2.Position, p2.Radius) {
			normalized_pos := rl.Vector2Subtract(p1.Position, p2.Position)
			distance := rl.Vector2Length(normalized_pos)
			normalized_pos = rl.Vector2Scale(normalized_pos, 1.0/distance)
			relative_vel := rl.Vector2Subtract(p1.Velocity, p2.Velocity)
			angular_vel := rl.Vector2DotProduct(relative_vel, normalized_pos)
			if angular_vel <= 0 {
				scalar_impulse := -(1 + 1.0) * angular_vel
				scalar_impulse /= float32(1.0/p1.Mass + 1.0/p2.Mass)
				impulse := rl.Vector2Scale(normalized_pos, scalar_impulse)
				p1.Velocity = rl.Vector2Add(p1.Velocity, rl.Vector2Scale(impulse, 1.0/p1.Mass))
				p2.Velocity = rl.Vector2Subtract(p2.Velocity, rl.Vector2Scale(impulse, 1.0/p2.Mass))

				p1.Hp -= p2.Damage
				p2.Hp -= p1.Damage

				if (p1.Hp <= 0) && (p2.Hp <= 0) {
					b.State = DRAW
				} else if p1.Hp <= 0 {
					b.State = P2_WIN
				} else if p2.Hp <= 0 {
					b.State = P1_WIN
				}

				if b.State != FIGHTING {
					return ""
				}
			}
		}

		p1.Position = p1.Position.Add(p1.Velocity)
		p2.Position = p2.Position.Add(p2.Velocity)
	} else if b.State == DRAW || b.State == P1_WIN || b.State == P2_WIN {
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), b.MainMenu.Body) &&
			rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
			return "MAIN_MENU"
		}
	}
	return ""
}
