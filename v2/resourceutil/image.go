package resourceutil

import (
	"image"
	_ "image/png"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
)

type imageLoader struct {
	img *ebiten.Image
}

func (l *imageLoader) Extract(x, y, width, height int) *ebiten.Image {
	return l.img.SubImage(image.Rect(x, y, x+width, y+height)).(*ebiten.Image)
}

func (l *imageLoader) ExtractList(x0, y0, width, height, xCount, yCount int) []*ebiten.Image {
	ret := make([]*ebiten.Image, 0)
	for xc := 0; xc < xCount; xc++ {
		for yc := 0; yc < yCount; yc++ {
			x := x0 + xc*width
			y := y0 + yc*height
			img := l.img.SubImage(image.Rect(x, y, x+width, y+height)).(*ebiten.Image)
			ret = append(ret, img)
		}
	}
	return ret
}

func NewImageLoader(fs fs.FS, path string) *imageLoader {
	f, err := fs.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return &imageLoader{img: ebiten.NewImageFromImage(img)}
}
