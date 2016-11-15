package gameobj

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

// hi there, future Tom or Mike. I'm not sure if this is how ebiten works, but the negatives on centerX calcs is
// to move left along the screen, and the seemingly inverted calculations for moving up / down are because screens typically
// have the zero at the top, not the bottom.

// helper function for gameobj for shapes to convert from degrees to radians
func degreesToRadians(degreeValue float64) float64 {
	return degreeValue * math.Pi / 180
}

// helper function to get the velocity components for a given velocity and travelAngle
func getVelocityComponents(baseSpeed int, speedModifier int, travelAngle float64) (xVelocity int, yVelocity int) {
	var travelAngleInRadians = degreesToRadians(travelAngle)

	xVelocity = int(float64(baseSpeed)*float64(speedModifier) + math.Cos(travelAngleInRadians))
	yVelocity = int(float64(baseSpeed)*float64(speedModifier) + math.Sin(travelAngleInRadians))
	return xVelocity, yVelocity
}

type BaseShape struct {
	Track            int
	CenterCoordinate *Coordinate
	BaseSpeed        int
	SpeedModifier    int
	image            *ebiten.Image
}

func NewBaseShape(track int, centerX int, baseSpeed int, speedModifier int, image *ebiten.Image) *BaseShape {
	var s = &BaseShape{
		Track: track,
		CenterCoordinate: &Coordinate{
			X: centerX,
			Y: TrackMappings[track],
		},
		BaseSpeed:     baseSpeed,
		SpeedModifier: speedModifier,
		image:         image,
	}
	return s
}

func (s *BaseShape) Len() int {
	return 1
}

func (s *BaseShape) Dst(i int) (x0, y0, x1, y1 int) {
	w, h := s.image.Size()
	halfHeight := h / 2
	halfWidth := w / 2
	return s.CenterCoordinate.X - halfHeight,
		s.CenterCoordinate.Y - halfWidth,
		s.CenterCoordinate.X + halfHeight,
		s.CenterCoordinate.Y + halfWidth
}

func (s *BaseShape) Src(i int) (x0, y0, x1, y1 int) {
	w, h := s.image.Size()
	return 0, 0, w, h
}

func (s *BaseShape) Image() *ebiten.Image {
	return s.image
}

type Drawable interface {
	Image() *ebiten.Image
	Update()
	Len() int
	Dst(int) (int, int, int, int)
	Src(int) (int, int, int, int)
}
