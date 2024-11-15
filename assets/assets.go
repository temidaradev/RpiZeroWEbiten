package assets

import (
	"embed"
	_ "image/png"

	"github.com/temidaradev/esset"
)

//go:embed *
var assets embed.FS

var GopherIdle = esset.GetAsset(assets, "gopher/mainchar.png")
var GopherLeft = esset.GetAsset(assets, "gopher/left.png")
var GopherRight = esset.GetAsset(assets, "gopher/right.png")

var Tiles = esset.GetMultiAssets(assets, "tiles/*.png")
