package drawutil

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type DrawGridOptions struct {
	horizontalLineLength, verticalLineLength                 float64
	xMin, xMax, yMin, yMax                                   float64
	xMainInterval, xSubInterval, yMainInterval, ySubInterval float64
	xMainColor, xSubColor, yMainColor, ySubColor             color.Color
}

func DrawGrid(dst *ebiten.Image, opts *DrawGridOptions) {
	if opts.xMax == 0.0 {
		opts.xMax = opts.horizontalLineLength
	}
	if opts.yMax == 0.0 {
		opts.yMax = opts.verticalLineLength
	}
	if opts.xMainInterval == 0.0 {
		opts.xMainInterval = 100.0
	}
	if opts.xSubInterval == 0.0 {
		opts.xSubInterval = 50.0
	}
	if opts.yMainInterval == 0.0 {
		opts.yMainInterval = 100.0
	}
	if opts.ySubInterval == 0.0 {
		opts.ySubInterval = 50.0
	}
	if opts.xMainColor == nil {
		opts.xMainColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
	}
	if opts.xSubColor == nil {
		opts.xSubColor = color.RGBA{0xff, 0xff, 0xff, 0x40}
	}
	if opts.yMainColor == nil {
		opts.yMainColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
	}
	if opts.ySubColor == nil {
		opts.ySubColor = color.RGBA{0xff, 0xff, 0xff, 0x40}
	}

	for _, params := range []struct {
		vertical           bool
		min, max, interval float64
		color              color.Color
	}{
		{vertical: true, min: opts.xMin, max: opts.xMax, interval: opts.xSubInterval, color: opts.xSubColor},
		{vertical: true, min: opts.xMin, max: opts.xMax, interval: opts.xMainInterval, color: opts.xMainColor},
		{vertical: false, min: opts.yMin, max: opts.yMax, interval: opts.ySubInterval, color: opts.ySubColor},
		{vertical: false, min: opts.yMin, max: opts.yMax, interval: opts.yMainInterval, color: opts.yMainColor},
	} {
		if params.interval <= 0.0 || params.color == nil {
			continue
		}

		p := params.min
		for p < params.max {
			if params.vertical {
				ebitenutil.DrawLine(dst, p, 0, p, opts.verticalLineLength, params.color)
			} else {
				ebitenutil.DrawLine(dst, 0, p, opts.horizontalLineLength, p, params.color)
			}
			p += params.interval
		}
	}
}
