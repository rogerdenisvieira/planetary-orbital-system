package main

import "github.com/hajimehoshi/ebiten/v2"

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// render the screen

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
