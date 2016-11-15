package gameobj

type Triangle struct {
	*BaseShape
	// this is expected to be degrees.
	TravelAngle      float64
	DestinationTrack int
	swapState        int
	midwayPoint      int
}

func NewTriangle(base *BaseShape) *Triangle {
	var t = &Triangle{
		BaseShape:        base,
		TravelAngle:      DefaultCircleAngleOfDescent,
		DestinationTrack: SubsequentTracks[base.Track],
		swapState:        TriangleBeforeSwap,
		midwayPoint:      int(base.CenterCoordinate.X / 2),
	}
	return t
}

func NewTriangleNonStandardAngle(base *BaseShape, travelAngle float64) *Triangle {
	var t = &Triangle{
		BaseShape:        base,
		TravelAngle:      DefaultCircleAngleOfDescent,
		DestinationTrack: SubsequentTracks[base.Track],
		swapState:        TriangleBeforeSwap,
		midwayPoint:      int(base.CenterCoordinate.X / 2),
	}
	return t
}

func (t *Triangle) Update() {
	if t.swapState == TriangleBeforeSwap || t.swapState == TriangleAfterSwap {
		// before and after swap, just slide along the track
		t.CenterCoordinate.X = t.CenterCoordinate.X - (t.BaseSpeed * t.SpeedModifier)

		if t.swapState == TriangleBeforeSwap && int(t.CenterCoordinate.X) <= t.midwayPoint {
			t.swapState = TriangleDuringSwap
		}
	} else {
		t.updateWithTrackSwitchingMovement()
	}
}

// this is the circle's up / down logic! wooo!
func (t *Triangle) updateWithTrackSwitchingMovement() {
	var xVelocity, yVelocity = getVelocityComponents(t.BaseSpeed, t.SpeedModifier, t.TravelAngle)

	if t.Track < t.DestinationTrack {
		yVelocity = yVelocity * -1
	}
	t.CenterCoordinate.X = t.CenterCoordinate.X - xVelocity
	t.CenterCoordinate.Y = t.CenterCoordinate.Y - yVelocity

	if (t.Track < t.DestinationTrack && t.CenterCoordinate.Y >= TrackMappings[t.DestinationTrack]) ||
		(t.Track > t.DestinationTrack && t.CenterCoordinate.Y <= TrackMappings[t.DestinationTrack]) {
		// then set the track to the destination
		t.Track = t.DestinationTrack
		// and snap the centerY to the new track
		t.CenterCoordinate.Y = TrackMappings[t.Track]
		t.swapState = TriangleAfterSwap
	}
}
