package touchutil

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tsujio/game-util/mathutil"
)

type screenTouch struct {
	touchID      ebiten.TouchID
	pos, prevPos *mathutil.Vector2D
	released     bool
}

func (s *screenTouch) ID() TouchID {
	return TouchID{touchType: touchTypeScreenTouch, id: int(s.touchID)}
}

func (s *screenTouch) Update() {
	if s.pos != nil {
		s.prevPos = s.pos.Clone()
	}

	var x, y int
	if s.IsJustReleased() {
		x, y = inpututil.TouchPositionInPreviousTick(s.touchID)
	} else {
		x, y = ebiten.TouchPosition(s.touchID)
	}
	s.pos = mathutil.NewVector2D(float64(x), float64(y))

	if !s.released {
		s.released = s.IsJustReleased()
	}
}

func (s *screenTouch) IsJustTouched() bool {
	for _, touchID := range justPressedTouchedIDs {
		if touchID == s.touchID {
			return true
		}
	}
	return false
}

func (s *screenTouch) IsJustReleased() bool {
	return inpututil.IsTouchJustReleased(s.touchID)
}

func (s *screenTouch) Position() *mathutil.Vector2D {
	return s.pos
}

func (s *screenTouch) PreviousPosition() *mathutil.Vector2D {
	return s.prevPos
}

func (s *screenTouch) isReleased() bool {
	return s.released
}
