package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/exp/constraints"
)

func Axis() (axisX, axisY float64) {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		axisY -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		axisY += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		axisX -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
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
