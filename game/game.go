package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player *Player
}

func NewGame() *Game {
	g := &Game{}
	cam.LerpEnabled = true
	cam.ShakeEnabled = true

	g.player = &Player{
		X:     0,
		Y:     0,
		DIO:   &ebiten.DrawImageOptions{},
		speed: 20,
	}
	return g
}

var low float64 = 0
var high float64 = 1500

func (g *Game) Update() error {

	g.player.Update()

	camPosX := Clamp(g.player.X, low, high)
	camPosY := Clamp(g.player.Y, low, high)

	cam.LookAt(camPosX, camPosY)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	MakeBackground(screen)
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
