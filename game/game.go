package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/setanarut/kamera/v2"
)

var (
	w, h                                = 635., 475.
	camSpeed, zoomSpeedFactor, rotSpeed = 10.0, 1.02, 0.02
	targetX, targetY                    = w / 2, h / 2
	cam                                 = kamera.NewCamera(targetX, targetY, w, h)
)

type Game struct {
	player Player
}

func NewGame() *Game {
	g := &Game{}
	cam.LerpEnabled = true
	cam.ShakeEnabled = true
	return g
}

func (g *Game) Update() error {
	g.player.Update()
	cam.LookAt(targetX, targetY)
	cam.Center()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
