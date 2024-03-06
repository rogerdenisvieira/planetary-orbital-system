package internal

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	G        float64 = 6.67428e-11
	AU               = 149.6e6 * 1000 // 149.6 million km, in meters.
	Scale    float64 = 250 / AU
	Timestep         = 3600 * 24 // 1 day
)

type Point struct {
	X float64
	Y float64
}

type CelestialBody struct {
	Name   string
	Mass   float64
	Radius float64
	Color  color.Color
	Orbit  []Point
	IsSun  bool

	VelX, VelY, PosX, PosY float64
}

func NewCelestialBody(name string, posX float64, posY float64, radius float64, color color.Color, mass float64, velY float64) *CelestialBody {
	return &CelestialBody{
		Name:   name,
		PosX:   posX,
		PosY:   posY,
		Radius: radius,
		Color:  color,
		Mass:   mass,
		VelY:   velY,
		// VelX:   0,
		IsSun: false,
	}
}

func (body CelestialBody) draw(screen *ebiten.Image) {

	x := body.PosX*Scale + ScreenWidth/2
	y := body.PosY*Scale + ScreenHeight/2

	// if len(body.Orbit) > 2 {

	// 	for _, point := range body.Orbit {
	// 		x, y = point.X, point.Y

	// 		x = x*Scale + ScreenWidth/2
	// 		y = y*Scale + ScreenHeight/2

	vector.DrawFilledRect(screen, float32(x), float32(y), 1, 1, color.RGBA{255, 255, 255, 0}, false)
	// 	}

	// }

	vector.DrawFilledCircle(screen, float32(x), float32(y), float32(body.Radius), body.Color, false)

}

func (body CelestialBody) getAttraction(otherBody CelestialBody) (forceX, forceY float64) {

	otherBodyPosX, otherBodyPosY := otherBody.PosX, otherBody.PosY

	distanceX := otherBodyPosX - body.PosX
	distanceY := otherBodyPosY - body.PosY

	distance := math.Sqrt(math.Pow(distanceX, 2) + math.Pow(distanceY, 2))

	if distance == 0 {
		panic(fmt.Sprintf("Collision between objects %s and %s", body.Name, otherBody.Name))
	}

	// if body.IsSun {
	// 	body.DistanceToSun = distance
	// }

	force := G * body.Mass * otherBody.Mass / math.Pow(distance, 2)
	theta := math.Atan2(distanceY, distanceX)
	forceX = math.Cos(theta) * force
	forceY = math.Sin(theta) * force

	return forceX, forceY
}

func (body *CelestialBody) updatePosition(otherBodies []*CelestialBody) {
	var total_fx, total_fy = 0.0, 0.0

	for _, otherBody := range otherBodies {

		if body.Name == otherBody.Name {
			continue
		}

		fx, fy := body.getAttraction(*otherBody)

		total_fx += fx
		total_fy += fy

	}

	body.VelX += total_fx / body.Mass * Timestep
	body.VelY += total_fy / body.Mass * Timestep

	body.PosX += body.VelX * Timestep
	body.PosY += body.VelY * Timestep
	body.Orbit = append(body.Orbit, Point{body.PosX, body.PosY})

	log.Printf("Name: %s X: %f Y: %f", body.Name, body.PosX, body.PosY)
}
