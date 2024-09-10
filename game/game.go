package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

func NewGame() *Game {
	g := &Game{}
	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
