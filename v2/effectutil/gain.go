package effectutil

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type GainEffect struct {
	ticks, maxTicks uint64
	x, y, y0        float64
	opts            GainEffectOptions
}

type GainEffectOptions struct {
	Gain  int
	Face  text.Face
	Color color.Color
}

func NewGainEffect(x, y float64, span uint64, opts *GainEffectOptions) *GainEffect {
	_opts := *opts

	if _opts.Color == nil {
		_opts.Color = color.RGBA{0xff, 0xe0, 0, 0xff}
	}

	return &GainEffect{
		maxTicks: span,
		x:        x,
		y:        y,
		y0:       y,
		opts:     _opts,
	}
}

func (e *GainEffect) Type() string {
	return "gain"
}

func (e *GainEffect) Update() {
	e.ticks++
	e.y = e.y0 - 30*math.Sin(float64(e.ticks)/float64(e.maxTicks)*math.Pi)
}

func (e *GainEffect) Draw(dst *ebiten.Image) {
	t := fmt.Sprintf("%+d", e.opts.Gain)
	opts := &text.DrawOptions{}
	opts.GeoM.Translate(e.x, e.y)
	opts.ColorScale.ScaleWithColor(e.opts.Color)
	text.Draw(dst, t, e.opts.Face, opts)
}

func (e *GainEffect) Finished() bool {
	return e.ticks > e.maxTicks
}
