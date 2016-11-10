package gameobj

import "github.com/hajimehoshi/ebiten"

type Triangle struct {
	*BaseShape
	// this is expected to be degrees.
	TravelAngle      float64
	DestinationTrack int
	swapState        int
	midwayPoint      int
	image            *ebiten.Image
}

func NewTriangle(base *BaseShape, image *ebiten.Image) *Triangle {
	var t = &Triangle{
		BaseShape:        base,
		TravelAngle:      DefaultCircleAngleOfDescent,
		DestinationTrack: SubsequentTracks[base.Track],
		swapState:        TriangleBeforeSwap,
		midwayPoint:      int(base.CenterCoordinate.X / 2),
		image:            image,
	}
	return t
}

func NewTriangleNonStandardAngle(base *BaseShape, image *ebiten.Image, travelAngle float64) *Triangle {
	var t = &Triangle{
		BaseShape:        base,
		TravelAngle:      DefaultCircleAngleOfDescent,
		DestinationTrack: SubsequentTracks[base.Track],
		swapState:        TriangleBeforeSwap,
		midwayPoint:      int(base.CenterCoordinate.X / 2),
		image:            image,
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

func (t *Triangle) Len() int {
	return 1
}

func (t *Triangle) Dst(i int) (x0, y0, x1, y1 int) {
	w, h := t.image.Size()
	halfHeight := float64(h / 2)
	halfWidth := float64(w / 2)
	return int(t.CenterCoordinate.X - halfHeight),
		int(t.CenterCoordinate.Y - halfWidth),
		int(t.CenterCoordinate.X + halfHeight),
		int(t.CenterCoordinate.Y + halfWidth)
}

func (t *Triangle) Src(i int) (x0, y0, x1, y1 int) {
	w, h := t.image.Size()
	return 0, 0, w, h
}

func (t *Triangle) Image() *ebiten.Image {
	return t.image
}
