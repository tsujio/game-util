package loggingutil

import "github.com/tsujio/game-logging-server/client"

type InitializePayload struct {
	RandomSeed int64
}

type StartGamePayload struct {
}

type GameOverPayload struct {
	Score int
}

type touchPayload struct {
	touches []touchRecord
}

func SendLog(gameName string, playerID, sessionID, playID string, payload any) {
	p := map[string]any{
		"player_id":  playerID,
		"session_id": sessionID,
		"play_id":    playID,
	}

	switch pl := payload.(type) {
	case *InitializePayload:
		p["action"] = "initialize"
		p["seed"] = pl.RandomSeed
	case *StartGamePayload:
		p["action"] = "start_game"
	case *GameOverPayload:
		p["action"] = "game_over"
		p["score"] = pl.Score
	case *touchPayload:
		p["action"] = "touch"
		p["touches"] = pl.touches
	default:
		panic("Invalid payload type")
	}

	client.LogAsync(gameName, p)
}
