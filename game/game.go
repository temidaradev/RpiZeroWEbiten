package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/RpiZeroWEbiten/assets"
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

var playerOffsetX = float64(assets.GopherIdle.Bounds().Dx() / 2)
var playerOffsetY = float64(assets.GopherIdle.Bounds().Dy() / 2)

func init() {
	fmt.Println(playerOffsetX, playerOffsetY)
}

// var low float64 = 0
// var high float64 = 1500

func (g *Game) Update() error {
	g.player.Update()
	cam.LookAt(g.player.X, g.player.Y)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	MakeBackground(screen)
	g.player.DIO.GeoM.Translate(g.player.X-playerOffsetX, g.player.Y-playerOffsetY)
	cam.Draw(assets.GopherIdle, g.player.DIO, screen)
	g.player.DIO.GeoM.Reset()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
