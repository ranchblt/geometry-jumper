package gameobj

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

type Circle struct {
	*BaseShape
	// this is expected to be degrees.
	TravelAngle      float64
	DestinationTrack int
	image            *ebiten.Image
}

// default initializer for Circle. this sets TravelAngle to a default of 45 degrees
func NewCircle(base *BaseShape, image *ebiten.Image, destinationTrack int) *Circle {
	var c = &Circle{
		BaseShape:        base,
		TravelAngle:      DefaultCircleAngleOfDescent,
		DestinationTrack: destinationTrack,
		image:            image,
	}
	return c
}

// if you want a different angle of descent, use this initializer
func NewCircleNonStandardAngle(base *BaseShape, image *ebiten.Image, destinationTrack int, travelAngle float64) *Circle {
	var c = &Circle{
		BaseShape:        base,
		TravelAngle:      travelAngle,
		DestinationTrack: destinationTrack,
		image:            image,
	}
	return c
}

// I'm sure this method can be streamlined somehow.
func (c *Circle) Update() {
	var xVelocity, yVelocity = c.getVelocityComponents()

	if c.Track < c.DestinationTrack {
		yVelocity = yVelocity * -1
	}
	c.CenterCoordinate.X = c.CenterCoordinate.X - xVelocity
	c.CenterCoordinate.Y = c.CenterCoordinate.Y - yVelocity

	if (c.Track < c.DestinationTrack && c.CenterCoordinate.Y >= TrackMappings[c.DestinationTrack]) ||
		(c.Track > c.DestinationTrack && c.CenterCoordinate.Y <= TrackMappings[c.DestinationTrack]) {
		// then set the track to the destination
		c.Track = c.DestinationTrack
		// and snap the centerY to the new track
		c.CenterCoordinate.Y = TrackMappings[c.Track]
		// and set our new destination to the one "after" our previous destination
		c.DestinationTrack = SubsequentTracks[c.Track]
	}
}

// unpublished methods are sweet!
func (c *Circle) getVelocityComponents() (xVelocity float64, yVelocity float64) {
	var travelAngleInRadians = degreesToRadians(c.TravelAngle)

	xVelocity = c.BaseSpeed * c.SpeedModifier * math.Cos(travelAngleInRadians)
	yVelocity = c.BaseSpeed * c.SpeedModifier * math.Sin(travelAngleInRadians)
	return xVelocity, yVelocity
}

func (s *Circle) Len() int {
	return 1
}

func (c *Circle) Dst(i int) (x0, y0, x1, y1 int) {
	w, h := c.image.Size()
	halfHeight := float64(h / 2)
	halfWidth := float64(w / 2)
	return int(c.CenterCoordinate.X - halfHeight),
		int(c.CenterCoordinate.Y - halfWidth),
		int(c.CenterCoordinate.X + halfHeight),
		int(c.CenterCoordinate.Y + halfWidth)
}

func (c *Circle) Src(i int) (x0, y0, x1, y1 int) {
	w, h := c.image.Size()
	return 0, 0, w, h
}

func (c *Circle) Image() *ebiten.Image {
	return c.image
}
