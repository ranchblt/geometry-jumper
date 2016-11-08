package gameobj

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

type Circle struct {
	*BaseShape
	// this is expected to be degrees.
	travelAngle float64
	image       *ebiten.Image
}

// default initializer for Circle. this sets travelAngle to a default of 45 degrees
func NewCircle(base *BaseShape, image *ebiten.Image) *Circle {
	var c = &Circle{
		BaseShape:   base,
		travelAngle: DefaultCircleAngleOfDescent,
		image:       image,
	}
	return c
}

// if you want a different angle of descent, use this initializer
func NewCircleNonStandardAngle(base *BaseShape, image *ebiten.Image, travelAngle float64) *Circle {
	var c = &Circle{
		BaseShape:   base,
		travelAngle: travelAngle,
		image:       image,
	}
	return c
}

// I'm sure this method can be streamlined somehow.
func (c *Circle) Update() {
	var xVelocity, yVelocity = c.getVelocityComponents()
	c.CenterX = c.CenterX - xVelocity

	if c.Track == UpperTrack {
		// upper track means we're moving down (going from upper track to lower)
		c.CenterY += yVelocity

		// if the center of the circle reached the lower Y axis, flip the track to lower so we reverse directions
		if c.CenterY >= LowerTrackYAxis {
			c.CenterY = LowerTrackYAxis
			c.Track = LowerTrack
		}
	} else {
		// otherwise we're moving up (going from lower track to upper)
		c.CenterY -= yVelocity

		// if the center of the circle reached the upper Y axis, flip the track to upper so we reverse directions
		if c.CenterY <= UpperTrackYAxis {
			c.CenterY = UpperTrackYAxis
			c.Track = UpperTrack
		}
	}
}

// unpublished methods are sweet!
func (c *Circle) getVelocityComponents() (xVelocity float64, yVelocity float64) {
	var travelAngleInRadians = degreesToRadians(c.travelAngle)

	xVelocity = c.BaseSpeed * c.SpeedModifier * math.Cos(travelAngleInRadians)
	yVelocity = c.BaseSpeed * c.SpeedModifier * math.Sin(travelAngleInRadians)
	return xVelocity, yVelocity
}
