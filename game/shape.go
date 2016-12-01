package game

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
func getVelocityComponents(baseSpeed int, travelAngle float64) (xVelocity int, yVelocity int) {
	var travelAngleInRadians = degreesToRadians(travelAngle)

	xVelocity = int(float64(baseSpeed) + math.Cos(travelAngleInRadians))
	yVelocity = int(float64(baseSpeed) + math.Sin(travelAngleInRadians))
	return xVelocity, yVelocity
}

type BaseShape struct {
	Track       int
	Center      *coord
	BaseSpeed   int
	image       *ebiten.Image
	hitboxImage *ebiten.Image
	expired     bool
}

func NewBaseShape(track int, centerX int, baseSpeed int, image *ebiten.Image, hitboxImage *ebiten.Image) *BaseShape {
	var s = &BaseShape{
		Track: track,
		Center: &coord{
			x: centerX,
			y: TrackMappings[track],
		},
		BaseSpeed:   baseSpeed,
		image:       image,
		hitboxImage: hitboxImage,
		expired:     false,
	}
	return s
}

func (s *BaseShape) crossedLeftEdge() bool {
	var crossed bool
	w, _ := s.image.Size()
	if s.Center.x <= -(w / 2) {
		crossed = true
	} else {
		crossed = false
	}
	return crossed
}

func (s *BaseShape) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.hitboxImage, &ebiten.DrawImageOptions{
		ImageParts: s,
	})

	cm := ebiten.ColorM{}
	cm.Scale(0, 100, 0, 100)
	screen.DrawImage(s.image, &ebiten.DrawImageOptions{
		ColorM:     cm,
		ImageParts: s,
	})
}

func (s *BaseShape) Image() *ebiten.Image {
	return s.image
}

func (s *BaseShape) Len() int {
	return 1
}

func (s *BaseShape) Dst(i int) (x0, y0, x1, y1 int) {
	w, h := s.image.Size()
	halfHeight := h / 2
	halfWidth := w / 2
	return s.Center.x - halfHeight,
		s.Center.y - halfWidth,
		s.Center.x + halfHeight,
		s.Center.y + halfWidth
}

func (s *BaseShape) Src(i int) (x0, y0, x1, y1 int) {
	w, h := s.image.Size()
	return 0, 0, w, h
}

func (s *BaseShape) IsExpired() bool {
	return s.expired
}

func (s *BaseShape) CenterCoord() *coord {
	return s.Center
}

type Drawable interface {
	Draw(screen *ebiten.Image)
	Image() *ebiten.Image
	CenterCoord() *coord
	Update()
	Len() int
	Dst(int) (int, int, int, int)
	Src(int) (int, int, int, int)
	IsExpired() bool
}
