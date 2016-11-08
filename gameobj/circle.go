package gameobj

import "math"

type Circle struct {
	*BaseShape
	// this is expected to be degrees.
	travelAngle float64
}

// default initializer for Circle. this sets travelAngle to a default of 45 degrees
func NewCircle(track int, centerX float64, centerY float64, baseSpeed float64, speedModifier float64) *Circle {
	var c = &Circle{
		BaseShape:   NewBaseShape(track, centerX, centerY, baseSpeed, speedModifier),
		travelAngle: DefaultCircleAngleOfDescent,
	}
	return c
}

// if you want a different angle of descent, use this initializer
func NewCircleNonStandardAngle(track int, centerX float64, centerY float64, baseSpeed float64, speedModifier float64, travelAngle float64) *Circle {
	var c = &Circle{
		BaseShape:   NewBaseShape(track, centerX, centerY, baseSpeed, speedModifier),
		travelAngle: travelAngle,
	}
	return c
}

// I'm sure this method can be streamlined somehow.
func (c *Circle) Update() {
	var xVelocity, yVelocity = c.getVelocityComponents()
	c.centerX = c.centerX - xVelocity

	if c.track == UpperTrack {
		// upper track means we're moving down (going from upper track to lower)
		c.centerY += yVelocity

		// if the center of the circle reached the lower Y axis, flip the track to lower so we reverse directions
		if c.centerY >= LowerTrackYAxis {
			c.centerY = LowerTrackYAxis
			c.track = LowerTrack
		}
	} else {
		// otherwise we're moving up (going from lower track to upper)
		c.centerY -= yVelocity

		// if the center of the circle reached the upper Y axis, flip the track to upper so we reverse directions
		if c.centerY <= UpperTrackYAxis {
			c.centerY = UpperTrackYAxis
			c.track = UpperTrack
		}
	}
}

// unpublished methods are sweet!
func (c *Circle) getVelocityComponents() (xVelocity float64, yVelocity float64) {
	var travelAngleInRadians = degreesToRadians(c.travelAngle)

	xVelocity = c.baseSpeed * c.speedModifier * math.Cos(travelAngleInRadians)
	yVelocity = c.baseSpeed * c.speedModifier * math.Sin(travelAngleInRadians)
	return xVelocity, yVelocity
}
