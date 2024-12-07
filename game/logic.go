package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/exp/constraints"
)

const (
	groundY = 395
	unit    = 10
)

func Jump(axisX, axisY float64) {
	if axisY == groundY*unit {
		axisY = -10 * unit
	}

	if axisY > groundY*unit {
		axisY = groundY * unit
	}
	if axisY < 20*unit {
		axisY += 8
	}
}

func Axis() (axisX, axisY float64) {
	// if ebiten.IsKeyPressed(ebiten.KeyW) {
	// 	axisY -= 1
	// }
	// if ebiten.IsKeyPressed(ebiten.KeyS) {
	// 	axisY += 1
	// }
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if axisY == groundY*unit {
			axisY = -10 * unit
		}

		if axisY > groundY*unit {
			axisY = groundY * unit
		}
		if axisY < 20*unit {
			axisY -= 8
		}
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
