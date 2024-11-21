package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/temidaradev/RpiZeroWEbiten/assets"
)

type Player struct {
	X, Y  float64
	DIO   *ebiten.DrawImageOptions
	speed float64
}

func (p *Player) Update() error {
	x, y := Axis()

	p.X += x * p.speed
	p.Y += y * p.speed

	p.DIO.GeoM.Reset()

	p.DIO.GeoM.Translate(
		p.X-float64(assets.GopherIdle.Bounds().Dx()/2),
		p.Y-float64(assets.GopherIdle.Bounds().Dy()/2),
	)

	p.X = Clamp(p.X, -230, 1740)
	p.Y = Clamp(p.Y, -120, 1610)

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	cam.Draw(assets.GopherIdle, p.DIO, screen)

	// Crosshair
	cx, cy := float32(w/2), float32(h/2)
	vector.StrokeLine(screen, cx-10, cy, cx+10, cy, 1, color.White, true)
	vector.StrokeLine(screen, cx, cy-10, cx, cy+10, 1, color.White, true)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("X: %+v", p.X), 5, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Y: %+v", p.Y), 5, 15)
}
