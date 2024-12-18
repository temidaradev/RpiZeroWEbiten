package game

import (
	"github.com/setanarut/kamera/v2"
)

var (
	w, h             = 635., 475.
	targetX, targetY = w / 2, h / 2
	cam              = kamera.NewCamera(targetX, targetY, w, h)
)
