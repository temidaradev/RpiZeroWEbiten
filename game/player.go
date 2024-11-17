package game

import (
	"image/color"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
}

func (p *Player) Update() error {
	cam.LookAt(targetX, targetY)

	if inpututil.IsKeyJustPressed(ebiten.KeyL) {
		cam.LerpEnabled = !cam.LerpEnabled
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyK) {
		cam.ShakeEnabled = !cam.ShakeEnabled
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyT) {
		cam.AddTrauma(1.0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		cam.ZoomFactor *= 2
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		cam.ZoomFactor /= 2
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		targetX, targetY = rand.Float64()*w, rand.Float64()*h
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		targetX -= camSpeed / cam.ZoomFactor
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		targetX += camSpeed / cam.ZoomFactor
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		targetY -= camSpeed / cam.ZoomFactor
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		targetY += camSpeed / cam.ZoomFactor
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) { // zoom out
		cam.ZoomFactor /= zoomSpeedFactor
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) { // zoom in
		cam.ZoomFactor *= zoomSpeedFactor
	}

	if ebiten.IsKeyPressed(ebiten.KeyR) {
		cam.SetAngle(cam.Angle() + rotSpeed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyF) {
		cam.SetAngle(cam.Angle() - rotSpeed)
	}

	if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
		targetX, targetY = w/2, h/2
		cam.Reset()
	}
	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	s := assets.GopherIdle
	// if p.player.vx > 0 {
	// 	s = assets.GopherRight
	// } else if p.player.vx < 0 {
	// 	s = assets.GopherLeft
	// }

	cam.Draw(s, dio, screen)

	// Draw camera crosshair
	cx, cy := float32(w/2), float32(h/2)
	vector.StrokeLine(screen, cx-10, cy, cx+10, cy, 1, color.White, true)
	vector.StrokeLine(screen, cx, cy-10, cx, cy+10, 1, color.White, true)
}
