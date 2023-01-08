package drawutil

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/tsujio/game-logging-server/client"
	"github.com/tsujio/game-util/resourceutil"
)

const (
	defaultScreenWidth  = 640
	defaultScreenHeight = 480
)

type DrawRankingOption struct {
	BackgroundColor, FontColor color.Color
	TitleFont, BodyFont        *resourceutil.Font
	ScreenWidth, ScreenHeight  int
	PlayerID                   string
}

func DrawRanking(dst *ebiten.Image, ranking []client.GameScore, opt *DrawRankingOption) {
	if opt == nil {
		opt = &DrawRankingOption{}
	}
	if opt.BackgroundColor == nil {
		opt.BackgroundColor = color.RGBA{0, 0, 0, 0xa0}
	}
	if opt.FontColor == nil {
		opt.FontColor = color.White
	}
	if opt.ScreenWidth == 0 {
		opt.ScreenWidth = defaultScreenWidth
	}
	if opt.ScreenHeight == 0 {
		opt.ScreenHeight = defaultScreenHeight
	}

	ebitenutil.DrawRect(
		dst,
		20,
		40,
		float64(opt.ScreenWidth)-20*2,
		float64(opt.ScreenHeight)-40*2,
		opt.BackgroundColor,
	)

	rankingText := "RANKING"
	text.Draw(
		dst,
		rankingText,
		opt.TitleFont.Face,
		opt.ScreenWidth/2-len(rankingText)*int(opt.TitleFont.FaceOptions.Size)/2,
		110,
		opt.FontColor,
	)

	ranking = ranking[:5]

	rankIn := false
	for _, r := range ranking {
		if opt.PlayerID != "" && r.PlayerID == opt.PlayerID {
			rankIn = true
		}
	}

	hText := "   SCORE    DATE   "
	if rankIn {
		hText += "     "
	}
	text.Draw(
		dst,
		hText,
		opt.BodyFont.Face,
		opt.ScreenWidth/2-len(hText)*int(opt.BodyFont.FaceOptions.Size)/2,
		170,
		opt.FontColor,
	)

	for i, r := range ranking {
		rank := 1
		for _, s := range ranking[:i] {
			if r.Score < s.Score {
				rank++
			}
		}
		t := fmt.Sprintf("%d. %5d %s", rank, r.Score, r.Timestamp.Local().Format("2006.01.02"))
		if rankIn {
			if r.PlayerID == opt.PlayerID {
				t += " YOU!"
			} else {
				t += "     "
			}
		}
		text.Draw(
			dst,
			t,
			opt.BodyFont.Face,
			opt.ScreenWidth/2-len(t)*int(opt.BodyFont.FaceOptions.Size)/2,
			170+(i+1)*int(opt.BodyFont.FaceOptions.Size*2),
			opt.FontColor,
		)
	}
}
