package touchutil

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tsujio/game-util/mathutil"
)

type mouseButtonPress struct {
	mouseButton  ebiten.MouseButton
	pos, prevPos *mathutil.Vector2D
	released     bool
}

func (m *mouseButtonPress) ID() TouchID {
	return TouchID{touchType: touchTypeMouseButtonPress, id: int(m.mouseButton)}
}

func (m *mouseButtonPress) Update() {
	if m.pos != nil {
		m.prevPos = m.pos.Clone()
	}

	x, y := ebiten.CursorPosition()
	m.pos = mathutil.NewVector2D(float64(x), float64(y))

	if !m.released {
		m.released = m.IsJustReleased()
	}
}

func (m *mouseButtonPress) IsJustTouched() bool {
	return inpututil.IsMouseButtonJustPressed(m.mouseButton)
}

func (m *mouseButtonPress) IsJustReleased() bool {
	return inpututil.IsMouseButtonJustReleased(m.mouseButton)
}

func (m *mouseButtonPress) Position() *mathutil.Vector2D {
	return m.pos
}

func (m *mouseButtonPress) PreviousPosition() *mathutil.Vector2D {
	return m.prevPos
}

func (m *mouseButtonPress) isReleased() bool {
	return m.released
}
