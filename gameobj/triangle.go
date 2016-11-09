package gameobj

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

type Triangle struct {
	*BaseShape
	// this is expected to be degrees.
	TravelAngle      float64
	DestinationTrack int
	swapState        int
	Image            *ebiten.Image
}

func NewTriangle(base *BaseShape, image *ebiten.Image, destinationTrack int) *Triangle {
	var t = &Triangle{
		BaseShape:        base,
		TravelAngle:      DefaultCircleAngleOfDescent,
		DestinationTrack: destinationTrack,
		swapState:        TriangleBeforeSwap,
		Image:            image,
	}
	return t
}

func (r *Triangle) Update() {
	//var xVelocity, yVelocity = r.getVelocityComponents()

}

// unpublished methods are sweet!
func (r *Triangle) getVelocityComponents() (xVelocity float64, yVelocity float64) {
	var travelAngleInRadians = degreesToRadians(r.TravelAngle)

	xVelocity = r.BaseSpeed * r.SpeedModifier * math.Cos(travelAngleInRadians)
	yVelocity = r.BaseSpeed * r.SpeedModifier * math.Sin(travelAngleInRadians)
	return xVelocity, yVelocity
}
