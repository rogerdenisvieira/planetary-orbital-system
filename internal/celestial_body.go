package internal

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	G     float64 = 6.67428e-11
	AU            = (149.6e6 * 1000) // 149.6 million km, in meters.
	SCALE         = 250 / AU
)

type CelestialBody struct {
	Name          string
	Mass          float64
	Radius        float32
	Color         color.Color
	Orbit         []CelestialBody
	IsSun         bool
	DistanceToSun float64

	VX, VY, PX, PY float64
}

func (body CelestialBody) draw(screen *ebiten.Image) {

	x := float32(body.PX*SCALE + 480/2)
	y := float32(body.PY*SCALE + 640/2)

	vector.DrawFilledCircle(screen, x, y, (body.Radius), body.Color, false)

}

func (body CelestialBody) getAttraction(otherBody CelestialBody) (forceX, forceY float64) {

	bodyPosX, bodyPosY := body.PX, body.PY
	otherPosX, otherPosY := otherBody.PX, otherBody.PY

	distanceX := otherPosX - bodyPosX
	distanceY := otherPosY - bodyPosY

	distance := math.Sqrt(math.Pow(distanceX, 2) + math.Pow(distanceY, 2))

	if distance == 0 {
		panic(fmt.Sprintf("Collision between objects %s and %s", body.Name, otherBody.Name))
	}

	force := G * body.Mass * otherBody.Mass / math.Pow(distance, 2)
	theta := math.Atan2(distanceX, distanceY)
	forceX = math.Cos(theta) * force
	forceY = math.Sin(theta) * force

	return forceX, forceY
}
