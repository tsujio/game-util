package touchutil

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type TouchPosition struct {
	X, Y int
}

type TouchContext struct {
	isTouchJustPressed        bool
	isTouchJustReleased       bool
	isMouseButtonJustPressed  bool
	isMouseButtonJustReleased bool
	isBeingTouched            bool
	touchIDs                  []ebiten.TouchID
	mainTouchID               *ebiten.TouchID
	touchPosition             TouchPosition
}

func CreateTouchContext() *TouchContext {
	return &TouchContext{
		touchIDs: []ebiten.TouchID{},
	}
}

func (c *TouchContext) Update() {
	c.touchIDs = inpututil.AppendJustPressedTouchIDs(c.touchIDs[:0])
	if c.mainTouchID == nil && len(c.touchIDs) > 0 {
		c.mainTouchID = &c.touchIDs[0]
		c.isTouchJustPressed = true
	} else {
		c.isTouchJustPressed = false
	}

	if c.mainTouchID != nil && inpututil.IsTouchJustReleased(*c.mainTouchID) {
		c.isTouchJustReleased = true
	} else {
		c.isTouchJustReleased = false
	}

	c.isMouseButtonJustPressed = inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)

	c.isMouseButtonJustReleased = inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)

	if c.IsJustTouched() {
		c.isBeingTouched = true
	}
	if c.IsJustReleased() {
		c.isBeingTouched = false
	}

	if c.isBeingTouched || c.IsJustReleased() {
		var x, y int
		if c.mainTouchID != nil {
			if c.isTouchJustReleased {
				x, y = inpututil.TouchPositionInPreviousTick(*c.mainTouchID)
			} else {
				x, y = ebiten.TouchPosition(*c.mainTouchID)
			}
		} else {
			x, y = ebiten.CursorPosition()
		}

		c.touchPosition = TouchPosition{
			X: x,
			Y: y,
		}
	} else {
		c.touchPosition = TouchPosition{}
	}

	if c.isTouchJustReleased {
		c.mainTouchID = nil
	}
}

func (c *TouchContext) IsJustTouched() bool {
	return c.isTouchJustPressed || c.isMouseButtonJustPressed
}

func (c *TouchContext) IsJustReleased() bool {
	return c.isTouchJustReleased || c.isMouseButtonJustReleased
}

func (c *TouchContext) IsBeingTouched() bool {
	return c.isBeingTouched
}

func (c *TouchContext) GetTouchPosition() TouchPosition {
	return c.touchPosition
}
