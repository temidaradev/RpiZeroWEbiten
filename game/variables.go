package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/setanarut/kamera/v2"
)

var (
	w, h                                = 635., 475.
	camSpeed, zoomSpeedFactor, rotSpeed = 7.0, 1.02, 0.02
	targetX, targetY                    = w / 2, h / 2
	cam                                 = kamera.NewCamera(targetX, targetY, w, h)
	dio                                 = &ebiten.DrawImageOptions{}
)
