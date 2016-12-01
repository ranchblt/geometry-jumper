package game

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
		midwayPoint:      int(base.Center.X / 2),
	}
	return t
}

func NewTriangleNonStandardAngle(base *BaseShape, travelAngle float64) *Triangle {
	var t = &Triangle{
		BaseShape:        base,
		TravelAngle:      DefaultCircleAngleOfDescent,
		DestinationTrack: SubsequentTracks[base.Track],
		swapState:        TriangleBeforeSwap,
		midwayPoint:      int(base.Center.X / 2),
	}
	return t
}

func (t *Triangle) Update() {
	if t.swapState == TriangleBeforeSwap || t.swapState == TriangleAfterSwap {
		// before and after swap, just slide along the track
		t.Center.X = t.Center.X - t.BaseSpeed

		if t.swapState == TriangleBeforeSwap && int(t.Center.X) <= t.midwayPoint {
			t.swapState = TriangleDuringSwap
		}
	} else {
		t.updateWithTrackSwitchingMovement()
	}

	if t.crossedLeftEdge() {
		t.expired = true
	}
}

// this is the circle's up / down logic! wooo!
func (t *Triangle) updateWithTrackSwitchingMovement() {
	var xVelocity, yVelocity = getVelocityComponents(t.BaseSpeed, t.TravelAngle)

	if t.Track < t.DestinationTrack {
		yVelocity = yVelocity * -1
	}
	t.Center.X = t.Center.X - xVelocity
	t.Center.Y = t.Center.Y - yVelocity

	if (t.Track < t.DestinationTrack && t.Center.Y >= TrackMappings[t.DestinationTrack]) ||
		(t.Track > t.DestinationTrack && t.Center.Y <= TrackMappings[t.DestinationTrack]) {
		// then set the track to the destination
		t.Track = t.DestinationTrack
		// and snap the centerY to the new track
		t.Center.Y = TrackMappings[t.Track]
		t.swapState = TriangleAfterSwap
	}
}