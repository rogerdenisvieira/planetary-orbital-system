package internal

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

var (
	InitialXPosition float32 = 100
	InitialYPosition float32 = 100
	Earth            CelestialBody
)

type Game struct {
}

func (g *Game) Init() {
	Earth = CelestialBody{Name: "Terra", Radius: 50, Mass: 10e20, Color: color.RGBA{0, 0, 128, 0}, Orbit: nil, IsSun: false, DistanceToSun: 100e30, VX: 0, VY: 0, PX: 0, PY: 0}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// render the screen
	Earth.draw(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
