package dotutil

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type DrawImagePosition int

const (
	DrawImagePositionTopLeft DrawImagePosition = iota
	DrawImagePositionCenter
)

type DrawImageOption struct {
	Rotate       float64
	BasePosition DrawImagePosition
}

func DrawImage(dst *ebiten.Image, img *ebiten.Image, x, y float64, opt *DrawImageOption) {
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Rotate(opt.Rotate)
	o.GeoM.Translate(x, y)

	switch opt.BasePosition {
	case DrawImagePositionCenter:
		w, h := img.Size()
		r := math.Sqrt(math.Pow(float64(w), 2)+math.Pow(float64(h), 2)) / 2
		rad := math.Atan2(float64(h), float64(w))
		o.GeoM.Translate(-r*math.Cos(opt.Rotate+rad), -r*math.Sin(opt.Rotate+rad))
	}

	dst.DrawImage(img, o)
}
