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
	Scale          float64
	ScaleX, ScaleY float64
	Rotate         float64
	BasePosition   DrawImagePosition
}

func DrawImage(dst *ebiten.Image, img *ebiten.Image, x, y float64, opt *DrawImageOption) {
	o := &ebiten.DrawImageOptions{}

	switch opt.BasePosition {
	case DrawImagePositionCenter:
		w, h := img.Size()
		r := math.Sqrt(math.Pow(float64(w), 2)+math.Pow(float64(h), 2)) / 2
		rad := math.Atan2(float64(h), float64(w))
		o.GeoM.Translate(-r*math.Cos(rad), -r*math.Sin(rad))
	}

	if opt.ScaleX != 0 || opt.ScaleY != 0 {
		o.GeoM.Scale(opt.ScaleX, opt.ScaleY)
	} else if opt.Scale != 0 {
		o.GeoM.Scale(opt.Scale, opt.Scale)
	}
	o.GeoM.Rotate(opt.Rotate)
	o.GeoM.Translate(x, y)

	dst.DrawImage(img, o)
}
