package game

import (
	"geometry-jumper/keyboard"

	"github.com/hajimehoshi/ebiten"
)

type PlayerCharacter struct {
	name             string
	image            *ebiten.Image
	imageJumping     *ebiten.Image
	keyboardWrapper  *keyboard.KeyboardWrapper
	CenterCoordinate *Coordinate
	jumping          bool
	maxHeightReached bool
	originalY        int
}

func NewPlayerCharacter(name string, image *ebiten.Image, jimage *ebiten.Image, keyboardWrapper *keyboard.KeyboardWrapper) *PlayerCharacter {
	var player = &PlayerCharacter{
		name:            "Test",
		image:           image,
		imageJumping:    jimage,
		keyboardWrapper: keyboardWrapper,
		CenterCoordinate: &Coordinate{
			X: PlayerX,
			Y: TrackMappings[LowerTrack],
		},
		jumping:   false,
		originalY: 0,
	}
	return player
}

func (pc *PlayerCharacter) Update() error {
	if pc.keyboardWrapper.KeyPushed(ebiten.KeySpace) {
		if !pc.jumping {
			pc.jumping = true
			pc.maxHeightReached = false
			pc.originalY = pc.CenterCoordinate.Y
		}
	}

	if pc.jumping {
		if pc.CenterCoordinate.Y >= pc.originalY-JumpHeight && !pc.maxHeightReached {
			pc.CenterCoordinate.Y -= JumpUpSpeed
		} else {
			pc.maxHeightReached = true
			pc.CenterCoordinate.Y += JumpDownSpeed

			if pc.CenterCoordinate.Y >= pc.originalY {
				pc.CenterCoordinate.Y = pc.originalY
				pc.jumping = false
			}
		}
	}

	return nil
}

func (pc *PlayerCharacter) Draw(screen *ebiten.Image) {
	if pc.jumping {
		screen.DrawImage(pc.imageJumping, &ebiten.DrawImageOptions{
			ImageParts: pc,
		})
	} else {
		screen.DrawImage(pc.image, &ebiten.DrawImageOptions{
			ImageParts: pc,
		})
	}
}

func (pc *PlayerCharacter) Image() *ebiten.Image {
	return pc.image
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
