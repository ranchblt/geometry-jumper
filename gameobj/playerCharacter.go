package gameobj

import (
	"fmt"
	"geometry-jumper/keyboard"

	"github.com/hajimehoshi/ebiten"
)

type PlayerCharacter struct {
	name             string
	image            *ebiten.Image
	keyboardWrapper  *keyboard.KeyboardWrapper
	CenterCoordinate *Coordinate
}

func NewPlayerCharacter(name string, image *ebiten.Image, keyboardWrapper *keyboard.KeyboardWrapper) *PlayerCharacter {
	var player = &PlayerCharacter{
		name:            "Test",
		image:           image,
		keyboardWrapper: keyboardWrapper,
		CenterCoordinate: &Coordinate{
			X: LeftSide,
			Y: TrackMappings[LowerTrack],
		},
	}
	return player
}

func (pc *PlayerCharacter) Update() error {
	if pc.keyboardWrapper.KeyPushed(ebiten.KeySpace) {
		fmt.Print("you pushed space")
	}
	return nil
}

func (pc *PlayerCharacter) Len() int {
	return 1
}

func (pc *PlayerCharacter) Dst(i int) (x0, y0, x1, y1 int) {
	w, h := pc.image.Size()
	halfHeight := h / 2
	halfWidth := w / 2
	return pc.CenterCoordinate.X - halfHeight,
		pc.CenterCoordinate.Y - halfWidth,
		pc.CenterCoordinate.X + halfHeight,
		pc.CenterCoordinate.Y + halfWidth
}

func (pc *PlayerCharacter) Src(i int) (x0, y0, x1, y1 int) {
	w, h := pc.image.Size()
	return 0, 0, w, h
}

func (pc *PlayerCharacter) Image() *ebiten.Image {
	return pc.image
}

func (pc *PlayerCharacter) Draw(screen *ebiten.Image) {
	screen.DrawImage(pc.image, &ebiten.DrawImageOptions{
		ImageParts: pc,
	})
}
