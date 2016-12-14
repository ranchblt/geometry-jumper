package ranchblt

import (
	"bytes"
	"github.com/ranchblt/geometry-jumper/resource"
	"image"

	"github.com/hajimehoshi/ebiten"
)

type Logo struct {
	width  int
	height int
	image  *ebiten.Image
}

func (l *Logo) Len() int {
	return 1
}

func (l *Logo) Dst(i int) (x0, y0, x1, y1 int) {
	w, h := l.image.Size()
	return l.width/2 - w/5,
		l.height/2 - h/5,
		l.width/2 + w/5,
		l.height/2 + h/5
}

func (l *Logo) Src(i int) (x0, y0, x1, y1 int) {
	w, h := l.image.Size()
	return 0, 0, w, h
}

func (l *Logo) Draw(screen *ebiten.Image) {
	screen.DrawImage(l.image, &ebiten.DrawImageOptions{
		ImageParts: l,
	})
}

func NewLogoScreen(screenWidth, screenHeight int) *Logo {
	logoImage, err := openImage("ranchy.png")
	if err != nil {
		panic(err)
	}

	i, err := ebiten.NewImageFromImage(logoImage, ebiten.FilterLinear)
	if err != nil {
		panic(err)
	}

	return &Logo{
		width:  screenWidth,
		height: screenHeight,
		image:  i,
	}
}

func openImage(path string) (image.Image, error) {
	b, err := resource.Asset(path)
	if err != nil {
		return nil, err
	}

	image, _, err := image.Decode(bytes.NewReader(b))

	if err != nil {
		return nil, err
	}

	return image, nil
}
