package game

type Circle struct {
	*BaseShape
	// this is expected to be degrees.
	TravelAngle      float64
	DestinationTrack int
}

// default initializer for Circle. this sets TravelAngle to a default of 45 degrees
func NewCircle(base *BaseShape) *Circle {
	var c = &Circle{
		BaseShape:        base,
		TravelAngle:      DefaultCircleAngleOfDescent,
		DestinationTrack: SubsequentTracks[base.Track],
	}
	return c
}

// if you want a different angle of descent, use this initializer
func NewCircleNonStandardAngle(base *BaseShape, travelAngle float64) *Circle {
	var c = &Circle{
		BaseShape:        base,
		TravelAngle:      travelAngle,
		DestinationTrack: SubsequentTracks[base.Track],
	}
	return c
}

func (c *Circle) Update() {
	var xVelocity, yVelocity = getVelocityComponents(c.BaseSpeed, c.TravelAngle)

	if c.Track < c.DestinationTrack {
		yVelocity = yVelocity * -1
	}
	c.Center.x = c.Center.x - xVelocity
	c.Center.y = c.Center.y - yVelocity

	if (c.Track < c.DestinationTrack && c.Center.y >= TrackMappings[c.DestinationTrack]) ||
		(c.Track > c.DestinationTrack && c.Center.y <= TrackMappings[c.DestinationTrack]) {
		// then set the track to the destination
		c.Track = c.DestinationTrack
		// and snap the centerY to the new track
		c.Center.y = TrackMappings[c.Track]
		// and set our new destination to the one "after" our previous destination
		c.DestinationTrack = SubsequentTracks[c.Track]
	}

	c.crossedLeftEdge()
}
