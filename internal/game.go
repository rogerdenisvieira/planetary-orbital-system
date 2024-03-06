package internal

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 1280
	ScreenHeight = 1280
)

var (
	InitialXPosition float32 = 100
	InitialYPosition float32 = 100
	Bodies           []*CelestialBody
)

type Game struct {
}

func (g *Game) Init() {

	Sun := NewCelestialBody(
		"Sol",
		0,
		0,
		30,
		color.RGBA{255, 255, 0, 0},
		1.98892*math.Pow(10, 30),
		0)

	Sun.IsSun = true

	Earth := NewCelestialBody(
		"Terra",
		(-1 * AU),
		400,
		16,
		color.RGBA{0, 0, 255, 0},
		5.9742*math.Pow(10, 24),
		(29.783 * 1000))

	Mars := NewCelestialBody(
		"Marte",
		(-1.524 * AU),
		0,
		12,
		color.RGBA{255, 0, 0, 0},
		6.39*math.Pow(10, 23),
		(24.077 * 1000))

	Mercury := NewCelestialBody(
		"Mercúrio",
		(0.387 * AU),
		0,
		8,
		color.RGBA{64, 64, 64, 0},
		3.30*math.Pow(10, 23),
		(-35.02 * 1000))

	Venus := NewCelestialBody(
		"Vênus",
		(0.723 * AU),
		0,
		14,
		color.RGBA{255, 255, 255, 0},
		4.8685*math.Pow(10, 24),
		(-35.02 * 1000))

	Bodies = append(Bodies, Sun, Earth, Mars, Mercury, Venus)

}

func (g *Game) Update() error {
	// Bodies := make([]CelestialBody, 1)

	for _, body := range Bodies {
		body.updatePosition(Bodies)
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
