package internal

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 800
)

var (
	InitialXPosition float32 = 100
	InitialYPosition float32 = 100
	Earth            CelestialBody
	Sun              CelestialBody
	Bodies           []CelestialBody
)

type Game struct {
}

func (g *Game) Init() {
	Earth = CelestialBody{Name: "Terra", Radius: 16, Mass: math.Pow((5.9742 * 10), 24), Color: color.RGBA{0, 0, 128, 0}, Orbit: nil, IsSun: false, DistanceToSun: 100e30, VX: 0, VY: (29.783 * 1000), PX: (-1 * AU), PY: 0}
	Sun = CelestialBody{Name: "Sol", Radius: 30, Mass: math.Pow((1.98892 * 10), 30), Color: color.RGBA{0, 128, 128, 0}, Orbit: nil, IsSun: true, DistanceToSun: 0, VX: 0, VY: 0, PX: 0, PY: 0}

	Bodies = append(Bodies, Sun, Earth)

}

func (g *Game) Update() error {
	// Bodies := make([]CelestialBody, 1)

	for _, body := range Bodies {
		body.updatePosition(Bodies)
		fmt.Println(body)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// render the screen
	for _, body := range Bodies {
		body.draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
