package effectutil

import "github.com/hajimehoshi/ebiten/v2"

type Effect interface {
	Type() string
	Update()
	Draw(dst *ebiten.Image)
	Finished() bool
}
