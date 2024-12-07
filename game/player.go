package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/RpiZeroWEbiten/assets"
)

type Player struct {
	X, Y  float64
	DIO   *ebiten.DrawImageOptions
	speed float64
}

func (p *Player) Draw(screen *ebiten.Image) {
	MakeBackground(screen)
	p.DIO.GeoM.Scale(0.5, 0.5)
	p.DIO.GeoM.Translate(p.X-playerOffsetX+50, p.Y-playerOffsetY+50)
	cam.Draw(assets.GopherIdle, p.DIO, screen)
	p.DIO.GeoM.Reset()
}

func (p *Player) Update() {
	x, y := Axis()
	p.X += x * p.speed
	p.Y += y * p.speed

	//p.X = Clamp(p.X, -230, 1740)
	//p.Y = Clamp(p.Y, -120, 1610)
}
