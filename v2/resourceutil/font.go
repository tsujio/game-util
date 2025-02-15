package resourceutil

import (
	"image/color"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func ForceLoadFont(fs fs.FS, fontName string) (large, medium, small text.Face) {
	f, err := fs.Open(fontName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	src, err := text.NewGoTextFaceSource(f)
	if err != nil {
		panic(err)
	}

	large = &text.GoTextFace{Source: src, Size: 36}
	medium = &text.GoTextFace{Source: src, Size: 24}
	small = &text.GoTextFace{Source: src, Size: 12}
	return
}

func DrawTextWithFace(dst *ebiten.Image, txt string, x, y float64, align text.Align, color color.Color, face text.Face, lineSpacingScale float64) {
	opts := &text.DrawOptions{}
	opts.PrimaryAlign = align
	opts.GeoM.Translate(x, y)
	opts.ColorScale.ScaleWithColor(color)
	_, h := text.Measure(txt, face, 0)
	opts.LineSpacing = h * lineSpacingScale
	text.Draw(dst, txt, face, opts)
}
