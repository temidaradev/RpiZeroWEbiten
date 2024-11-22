package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	X, Y  float64
	DIO   *ebiten.DrawImageOptions
	speed float64
}

func (p *Player) Update() {
	x, y := Axis()
	p.X += x * p.speed
	p.Y += y * p.speed
	//p.X = Clamp(p.X, -230, 1740)
	//p.Y = Clamp(p.Y, -120, 1610)
}
