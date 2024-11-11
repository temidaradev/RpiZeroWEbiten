package assets

import (
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed *
var assets embed.FS

var GopherIdle = GetSingleImage("mainchar.png")
var GopherLeft = GetSingleImage("left.png")
var GopherRight = GetSingleImage("right.png")

func GetSingleImage(name string) *ebiten.Image {
	file, err := assets.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}
