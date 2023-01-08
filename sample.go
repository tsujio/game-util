package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/tsujio/game-util/touchutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	touchContext       *touchutil.TouchContext
	touchStartPosition *touchutil.TouchPosition
}

func (g *Game) Update() error {
	g.touchContext.Update()

	if g.touchContext.IsJustTouched() {
		pos := g.touchContext.GetTouchPosition()
		g.touchStartPosition = &pos
	}

	if g.touchContext.IsJustReleased() {
		g.touchStartPosition = nil
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	if g.touchStartPosition != nil {
		pos := g.touchContext.GetTouchPosition()
		ebitenutil.DrawLine(screen,
			float64(g.touchStartPosition.X), float64(g.touchStartPosition.Y),
			float64(pos.X), float64(pos.Y),
			color.White)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("%.1f FPS", ebiten.ActualFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Sample")

	game := &Game{
		touchContext: touchutil.CreateTouchContext(),
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
