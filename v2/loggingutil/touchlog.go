package loggingutil

import (
	"github.com/tsujio/game-rhinoceros/touchutil"
)

type touch struct {
	ID           string  `json:"id"`
	JustTouched  bool    `json:"just_touched"`
	JustReleased bool    `json:"just_released"`
	X            float64 `json:"x"`
	Y            float64 `json:"y"`
}

type touchRecord struct {
	Ticks   uint64  `json:"ticks"`
	Touches []touch `json:"touches"`
}

var touchBuffer []touchRecord

func SendTouchLog(gameName string, playerID, sessionID, playID string, ticks uint64, touches []touchutil.Touch) {
	if len(touches) > 0 {
		record := touchRecord{Ticks: ticks}
		for i := range touches {
			t := touches[i]
			pos := t.Position()
			record.Touches = append(record.Touches, touch{
				ID:           t.ID().Serialize(),
				JustTouched:  t.IsJustTouched(),
				JustReleased: t.IsJustReleased(),
				X:            pos.X,
				Y:            pos.Y,
			})
		}
		touchBuffer = append(touchBuffer, record)
	}

	if len(touchBuffer) > 0 {
		lastTicks := touchBuffer[len(touchBuffer)-1].Ticks
		if len(touchBuffer) >= 60 ||
			lastTicks > ticks ||
			ticks-lastTicks > 60 {
			SendLog(gameName, playerID, sessionID, playID, &touchPayload{
				touches: touchBuffer,
			})
			touchBuffer = nil
		}
	}
}
