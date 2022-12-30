package drawutil

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var drawPatternCanvas *ebiten.Image

func init() {
	drawPatternCanvas = ebiten.NewImage(100, 100)
}

type CreatePatternImageOption struct {
	Color       color.Color
	ColorMap    map[int]color.Color
	DotSize     float64
	DotInterval float64
}

func CreatePatternImage(pattern [][]int, opt *CreatePatternImageOption) *ebiten.Image {
	dotSize := opt.DotSize
	if dotSize == 0 {
		dotSize = 15
	}

	canvasWidth := int(float64(len(pattern[0]))*(dotSize+opt.DotInterval) - opt.DotInterval)
	canvasHeight := int(float64(len(pattern))*(dotSize+opt.DotInterval) - opt.DotInterval)
	canvas := ebiten.NewImage(canvasWidth, canvasHeight)

	DrawPattern(canvas, pattern, 0, 0, &DrawPatternOption{
		Color:       opt.Color,
		ColorMap:    opt.ColorMap,
		DotSize:     dotSize,
		DotInterval: opt.DotInterval,
	})

	return canvas
}

type PatternPosition int

const (
	PatternPositionTopLeft PatternPosition = iota
	PatternPositionCenter
)

type DrawPatternOption struct {
	Color        color.Color
	ColorMap     map[int]color.Color
	DotSize      float64
	DotInterval  float64
	Rotate       float64
	BasePosition PatternPosition
}

func DrawPattern(dst *ebiten.Image, pattern [][]int, x, y float64, option *DrawPatternOption) {
	var opt DrawPatternOption
	if option != nil {
		opt = *option
	}
	if opt.Color == nil {
		opt.Color = color.White
	}
	if opt.DotSize == 0 {
		opt.DotSize = 15
	}

	canvasWidth := int(float64(len(pattern[0]))*(opt.DotSize+opt.DotInterval) - opt.DotInterval)
	canvasHeight := int(float64(len(pattern))*(opt.DotSize+opt.DotInterval) - opt.DotInterval)
	if w, h := drawPatternCanvas.Size(); w < canvasWidth || h < canvasHeight {
		drawPatternCanvas = ebiten.NewImage(canvasWidth*2, canvasHeight*2)
	}
	canvas := drawPatternCanvas.SubImage(image.Rect(0, 0, canvasWidth, canvasHeight)).(*ebiten.Image)
	canvas.Clear()

	var cmap map[int]color.Color
	if opt.ColorMap != nil {
		cmap = opt.ColorMap
	} else {
		cmap = map[int]color.Color{1: opt.Color}
	}

	for v, c := range cmap {
		d := emptyImage.SubImage(image.Rect(0, 0, 1, 1)).(*ebiten.Image)
		d.Fill(c)
		for i := 0; i < len(pattern); i++ {
			for j := 0; j < len(pattern[i]); j++ {
				if pattern[i][j] == v {
					xij := float64(j) * (opt.DotSize + opt.DotInterval)
					yij := float64(i) * (opt.DotSize + opt.DotInterval)

					o := &ebiten.DrawImageOptions{}
					o.GeoM.Scale(opt.DotSize, opt.DotSize)
					o.GeoM.Translate(xij, yij)

					canvas.DrawImage(d, o)
				}
			}
		}
	}

	var pos DrawImagePosition
	switch opt.BasePosition {
	case PatternPositionTopLeft:
		pos = DrawImagePositionTopLeft
	case PatternPositionCenter:
		pos = DrawImagePositionCenter
	default:
		panic("Invalid position")
	}

	DrawImage(dst, canvas, x, y, &DrawImageOption{
		Rotate:       opt.Rotate,
		BasePosition: pos,
	})
}
