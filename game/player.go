package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/RpiZeroWEbiten/assets"
)

type Char struct {
	x  float64
	y  float64
	vx float64
	vy float64
}

const (
	groundY = 395
	unit    = 16
)

func (c *Char) tryJump() {
	if c.y == groundY*unit {
		c.vy = -10 * unit
	}
}

func (c *Char) update() {
	c.x += c.vx
	c.y += c.vy

	if c.vx > 0 {
		c.vx -= 4
	} else if c.vx < 0 {
		c.vx += 4
	}
	if c.vy > 0 {
		c.vy -= 4
	} else if c.vy < 0 {
		c.vy += 4
	}
}

type Player struct {
	player *Char
}

func (p *Player) Update() error {
	if p.player == nil {
		p.player = &Char{x: w, y: h}
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.player.vx = -5 * unit
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.player.vx = 5 * unit
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.player.vy = -5 * unit
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.player.vy = 5 * unit
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
	cam.Draw(s, op2, screen)
}
