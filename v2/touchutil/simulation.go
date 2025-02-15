package touchutil

import (
	"github.com/tsujio/game-util/mathutil"
)

type TouchSimulation struct {
	index          int
	actions        []any
	actionDuration int
}

type touchAction struct{}

type releaseAction struct{}

type waitAction struct {
	duration int
}

func (t *TouchSimulation) Touch() *TouchSimulation {
	t.actions = append(t.actions, touchAction{})
	return t
}

func (t *TouchSimulation) Release() *TouchSimulation {
	t.actions = append(t.actions, releaseAction{})
	return t
}

func (t *TouchSimulation) Wait(duration int) *TouchSimulation {
	t.actions = append(t.actions, waitAction{duration: duration})
	return t
}

func (t *TouchSimulation) Next() []Touch {
	t.actionDuration++

	action := t.actions[t.index%len(t.actions)]
	switch a := action.(type) {
	case touchAction:
		t.index++
		t.actionDuration = 0
		return []Touch{&simulatedTouch{isJustTouched: true}}
	case releaseAction:
		t.index++
		t.actionDuration = 0
		return []Touch{&simulatedTouch{isJustReleased: true}}
	case waitAction:
		index := t.index
		if t.actionDuration >= a.duration {
			t.index++
			t.actionDuration = 0
		}
		for i := index; i >= 0; i-- {
			switch t.actions[i%len(t.actions)].(type) {
			case touchAction:
				return []Touch{&simulatedTouch{}}
			case releaseAction:
				return nil
			}
		}
		return nil
	default:
		panic("error")
	}
}

func NewSimulation() *TouchSimulation {
	return &TouchSimulation{}
}

type simulatedTouch struct {
	isJustTouched  bool
	isJustReleased bool
}

func (t *simulatedTouch) ID() TouchID {
	return TouchID{}
}

func (t *simulatedTouch) Update() {}

func (t *simulatedTouch) IsJustTouched() bool {
	return t.isJustTouched
}

func (t *simulatedTouch) IsJustReleased() bool {
	return t.isJustReleased
}

func (t *simulatedTouch) Position() *mathutil.Vector2D {
	return nil
}

func (t *simulatedTouch) PreviousPosition() *mathutil.Vector2D {
	return nil
}

func (t *simulatedTouch) isReleased() bool {
	return t.isJustReleased
}
