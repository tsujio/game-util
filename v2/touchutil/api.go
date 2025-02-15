package touchutil

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/tsujio/game-util/mathutil"
)

type touchType int

const (
	touchTypeMouseButtonPress = iota
	touchTypeScreenTouch
)

type TouchID struct {
	touchType touchType
	id        int
}

func (t TouchID) Serialize() string {
	return fmt.Sprintf("%d-%d", t.touchType, t.id)
}

type Touch interface {
	ID() TouchID
	Update()
	IsJustTouched() bool
	IsJustReleased() bool
	Position() *mathutil.Vector2D
	PreviousPosition() *mathutil.Vector2D
	isReleased() bool
}

var (
	justPressedTouchedIDs = make([]ebiten.TouchID, 0)
)

func UpdateTouches(touches []Touch) []Touch {
	if len(touches) > 0 {
		_touches := touches[:0]
		for _, t := range touches {
			if !t.isReleased() {
				_touches = append(_touches, t)
			}
		}
		touches = _touches
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		touches = append(touches, &mouseButtonPress{
			mouseButton: ebiten.MouseButtonLeft,
		})
	}

	justPressedTouchedIDs = inpututil.AppendJustPressedTouchIDs(justPressedTouchedIDs[:0])
	for _, touchID := range justPressedTouchedIDs {
		touches = append(touches, &screenTouch{
			touchID: touchID,
		})
	}

	for _, t := range touches {
		t.Update()
	}

	return touches
}

func AnyTouchesJustTouched(touches []Touch) bool {
	for _, t := range touches {
		if t.IsJustTouched() {
			return true
		}
	}
	return false
}

func AnyTouchesActive(touches []Touch) bool {
	for _, t := range touches {
		if !t.isReleased() {
			return true
		}
	}
	return false
}

func AllTouchesJustReleased(touches []Touch) bool {
	if len(touches) == 0 {
		return false
	}
	for _, t := range touches {
		if !t.IsJustReleased() {
			return false
		}
	}
	return true
}
