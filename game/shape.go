package game

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten"
)

// hi there, future Tom or Mike. I'm not sure if this is how ebiten works, but the negatives on centerX calcs is
// to move left along the screen, and the seemingly inverted calculations for moving up / down are because screens typically
// have the zero at the top, not the bottom.

// helper function for gameobj for shapes to convert from degrees to radians
func degreesToRadians(degreeValue float64) float64 {
	return (degreeValue * math.Pi) / 180
}

// helper function to get the velocity components for a given velocity and travelAngle
func getVelocityComponents(baseSpeed int, travelAngle float64) (xVelocity int, yVelocity int) {
	var travelAngleInRadians = degreesToRadians(travelAngle)

	xVelocity = int(float64(baseSpeed) * math.Cos(travelAngleInRadians))
	yVelocity = int(float64(baseSpeed) * math.Sin(travelAngleInRadians))
	return xVelocity, yVelocity
}

type BaseShape struct {
	Track       int
	Center      *coord
	BaseSpeed   int
	image       *ebiten.Image
	rgbaImage   *image.RGBA
	hitboxImage *ebiten.Image
	expired     bool
	colorMap    ebiten.ColorM
	scored      bool
}

func NewBaseShape(track int, centerX int, baseSpeed int, image *ebiten.Image, colorMap ebiten.ColorM) *BaseShape {
	var s = &BaseShape{
		Track: track,
		Center: &coord{
			x: centerX,
			y: TrackMappings[track],
		},
		BaseSpeed: baseSpeed,
		image:     image,
		expired:   false,
		colorMap:  colorMap,
	}
	return s
}

func (s *BaseShape) crossedLeftEdge() {
	w, _ := s.image.Size()
	if s.Center.x <= -(w / 2) {
		s.expired = true
	} else {
		s.expired = false
	}
}

func (s *BaseShape) Draw(screen *ebiten.Image) {
	screen.DrawImage(s.image, &ebiten.DrawImageOptions{
		ColorM:     s.colorMap,
		ImageParts: s,
	})
}

func (s *BaseShape) Image() *ebiten.Image {
	return s.image
}

func (s *BaseShape) RgbaImage() *image.RGBA {
	if s.rgbaImage == nil {
		s.rgbaImage = toRGBA(s.image)
	}
	return s.rgbaImage
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

func (s *BaseShape) Scored() bool {
	return s.scored
}

func (s *BaseShape) SetScore(b bool) {
	s.scored = b
}

type Drawable interface {
	Draw(screen *ebiten.Image)
	Image() *ebiten.Image
	RgbaImage() *image.RGBA
	CenterCoord() *coord
	Update()
	Len() int
	Dst(int) (int, int, int, int)
	Src(int) (int, int, int, int)
	IsExpired() bool
	Scored() bool
	SetScore(bool)
}
