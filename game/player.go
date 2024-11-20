package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/temidaradev/RpiZeroWEbiten/assets"
)

// type Char struct {
// 	x  float64
// 	y  float64
// 	vx float64
// 	vy float64
// }

// const (
// 	groundY = 395
// 	unit    = 16
// )

// func (c *Char) tryJump() {
// 	if c.y == groundY*unit {
// 		c.vy = -10 * unit
// 	}
// }

// func (c *Char) update() {
// 	c.x += c.vx
// 	c.y += c.vy

// 	if c.vx > 0 {
// 		c.vx -= 4
// 	} else if c.vx < 0 {
// 		c.vx += 4
// 	}
// 	if c.vy > 0 {
// 		c.vy -= 4
// 	} else if c.vy < 0 {
// 		c.vy += 4
// 	}
// }

type Player struct {
	// player *Char
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

	//sides
	if p.X < -230 {
		p.X = -230
	}
	if p.X > 1740 {
		p.X = 1740
	}

	//corners
	if p.Y < -120 {
		p.Y = -120
	}
	if p.Y > 1610 {
		p.Y = 1610
	}
	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	cam.Draw(assets.GopherIdle, p.DIO, screen)

	// Draw camera crosshair
	cx, cy := float32(w/2), float32(h/2)
	vector.StrokeLine(screen, cx-10, cy, cx+10, cy, 1, color.White, true)
	vector.StrokeLine(screen, cx, cy-10, cx, cy+10, 1, color.White, true)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("X: %+v", p.X), 5, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Y: %+v", p.Y), 5, 15)
}
