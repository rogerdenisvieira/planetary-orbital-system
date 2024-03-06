package main

import (
	game "rogerdenisvieira/planetary-orbital-system/internal"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Planetary Orbital System")

	g := game.Game{}
	g.Init()

	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}
