package loggingutil

import (
	"github.com/tsujio/game-logging-server/client"
	"github.com/tsujio/game-util/touchutil"
)

func RegisterScoreToRankingAsync(gameName string, playerID string, playID string, score int) <-chan []client.GameScore {
	ch := make(chan []client.GameScore, 1)

	go func(c chan<- []client.GameScore) {
		client.RegisterScore(gameName, playerID, playID, score)
		if ranking, err := client.GetScoreList(gameName); err == nil {
			c <- ranking
		}
		close(c)
	}(ch)

	return ch
}

type touchRecord struct {
	Ticks        uint64 `json:"ticks"`
	JustTouched  bool   `json:"just_touched"`
	JustReleased bool   `json:"just_released"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
}

var touchBuffer []touchRecord

func SendTouchLog(gameName string, playerID string, playID string, ticks uint64, touchContext *touchutil.TouchContext) {
	if touchContext.IsBeingTouched() || touchContext.IsJustReleased() {
		pos := touchContext.GetTouchPosition()
		touchBuffer = append(touchBuffer, touchRecord{
			Ticks:        ticks,
			JustTouched:  touchContext.IsJustTouched(),
			JustReleased: touchContext.IsJustReleased(),
			X:            pos.X,
			Y:            pos.Y,
		})
	}
	if len(touchBuffer) > 0 {
		lastTicks := touchBuffer[len(touchBuffer)-1].Ticks
		if len(touchBuffer) >= 60 ||
			lastTicks > ticks ||
			ticks-lastTicks > 60 {
			SendLog(gameName, playerID, playID, map[string]interface{}{
				"touches": touchBuffer,
			})
			touchBuffer = nil
		}
	}
}

func SendLog(gameName string, playerID string, playID string, payload map[string]interface{}) {
	p := map[string]interface{}{
		"player_id": playerID,
		"play_id":   playID,
	}
	for k, v := range payload {
		p[k] = v
	}

	client.LogAsync(gameName, p)
}
