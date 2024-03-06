package main

import (
	game "rogerdenisvieira/planetary-orbital-system/internal"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Planetary Orbital System")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	g := game.Game{}
	g.Init()

	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}
