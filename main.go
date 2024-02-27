package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("E-biome")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
