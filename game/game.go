package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/temidaradev/RpiZeroWEbiten/assets"
	"golang.org/x/exp/constraints"
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

	// arka plan döşemeleri çiz
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			dio := &ebiten.DrawImageOptions{}
			dio.GeoM.Translate(float64(i*50), float64(j*50))
			cam.Draw(assets.Tiles[1], dio, screen)
		}
	}

	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func Axis() (axisX, axisY float64) {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		axisY -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		axisY += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		axisX -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		axisX += 1
	}
	return axisX, axisY
}

// Clamp returns v clamped to [low, high].
func Clamp[T constraints.Ordered](v, low, high T) T {
	if v < low {
		return low
	}
	if v > high {
		return high
	}
	return v
}
