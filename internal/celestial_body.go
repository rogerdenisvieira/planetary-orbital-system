package internal

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	G        float64 = 6.67428e-11
	AU               = (149.6e6 * 1000) // 149.6 million km, in meters.
	SCALE            = 250 / AU
	TIMESTEP         = 3600 * 24 // 1 day
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

	x := float32(body.PX*SCALE + 800/2)
	y := float32(body.PY*SCALE + 800/2)

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

func (body CelestialBody) updatePosition(otherBodies []CelestialBody) {
	var total_fx, total_fy = 0, 0

	for _, otherBody := range otherBodies {

		if body.Name == otherBody.Name {
			continue
		}

		fx, fy := body.getAttraction(otherBody)

		total_fx += int(fx)
		total_fy += int(fy)

	}

	body.VX += float64(total_fx) / body.Mass * TIMESTEP
	body.VY += float64(total_fy) / body.Mass * TIMESTEP

	body.PX += body.VX * TIMESTEP
	body.PY += body.VY * TIMESTEP
}
