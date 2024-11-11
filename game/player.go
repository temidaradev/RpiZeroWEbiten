package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/RpiZeroWEbiten/assets"
)

type Char struct {
	x  int
	y  int
	vx int
	vy int
}

const (
	groundY = 395
	unit    = 10
)

func (c *Char) tryJump() {
	if c.y == groundY*unit {
		c.vy = -10 * unit
	}
}

func (c *Char) update() {
	c.x += c.vx
	c.y += c.vy

	if c.y > groundY*unit {
		c.y = groundY * unit
	}
	if c.vx > 0 {
		c.vx -= 2
	} else if c.vx < 0 {
		c.vx += 2
	}
	if c.vy < 20*unit {
		c.vy += 8
	}
}

type Player struct {
	player *Char
}

func (p *Player) Update() error {
	if p.player == nil {
		p.player = &Char{x: 50 * unit, y: groundY * unit}
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.player.vx = -5 * unit
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.player.vx = 5 * unit
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.player.tryJump()
	}

	p.player.update()
	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	s := assets.GopherIdle
	if p.player.vx > 0 {
		s = assets.GopherRight
	} else if p.player.vx < 0 {
		s = assets.GopherLeft
	}

	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Scale(0.3, 0.3)
	op2.GeoM.Translate(float64(p.player.x)/unit, float64(p.player.y)/unit)
	screen.DrawImage(s, op2)
}
