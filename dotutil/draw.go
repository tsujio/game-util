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
	Scale        float64
	Rotate       float64
	BasePosition DrawImagePosition
}

func DrawImage(dst *ebiten.Image, img *ebiten.Image, x, y float64, opt *DrawImageOption) {
	o := &ebiten.DrawImageOptions{}
	if opt.Scale != 0 {
		o.GeoM.Scale(opt.Scale, opt.Scale)
	}
	o.GeoM.Rotate(opt.Rotate)
	o.GeoM.Translate(x, y)

	switch opt.BasePosition {
	case DrawImagePositionCenter:
		w, h := img.Size()
		r := math.Sqrt(math.Pow(float64(w)*opt.Scale, 2)+math.Pow(float64(h)*opt.Scale, 2)) / 2
		rad := math.Atan2(float64(h), float64(w))
		o.GeoM.Translate(-r*math.Cos(opt.Rotate+rad), -r*math.Sin(opt.Rotate+rad))
	}

	dst.DrawImage(img, o)
}
