package drawutil

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func DrawImageAt(dst *ebiten.Image, img *ebiten.Image, x, y float64, opts *ebiten.DrawImageOptions) {
	g := ebiten.GeoM{}

	w, h := img.Size()
	g.Translate(-float64(w)/2, -float64(h)/2)

	g.Concat(opts.GeoM)

	g.Translate(x, y)

	opts.GeoM = g

	dst.DrawImage(img, opts)
}
