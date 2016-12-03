package game

import (
	"fmt"
	"geometry-jumper/keyboard"

	"geometry-jumper/collision"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
)

type PlayerCharacter struct {
	name             string
	image            *ebiten.Image
	imageJumping     *ebiten.Image
	keyboardWrapper  *keyboard.KeyboardWrapper
	Center           *coord
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
		Center: &coord{
			x: PlayerX,
			y: TrackMappings[LowerTrack],
		},
		jumping:   false,
		originalY: 0,
	}
	return player
}

func (pc *PlayerCharacter) Update() error {
	if jumpBytes == nil {
		select {
		case jumpBytes = <-jumpCh:
		default:
		}
	}

	if err := JumpSound.Update(); err != nil {
		return err
	}

	if pc.keyboardWrapper.KeyPushed(ebiten.KeySpace) {
		if !pc.jumping {
			pc.jumping = true
			jumpSoundPlayer, err := audio.NewPlayerFromBytes(JumpSound, jumpBytes)
			if err != nil {
				return err
			}
			jumpSoundPlayer.Play()
			pc.maxHeightReached = false
			pc.originalY = pc.Center.y
		}
	}

	if pc.jumping {
		if pc.Center.y >= pc.originalY-JumpHeight && !pc.maxHeightReached {
			pc.Center.y -= JumpUpSpeed
		} else {
			pc.maxHeightReached = true
			pc.Center.y += JumpDownSpeed

			if pc.Center.y >= pc.originalY {
				pc.Center.y = pc.originalY
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

func (pc *PlayerCharacter) CheckCollision(sc *ShapeCollection) {
	pcHitbox := collision.Hitbox{
		Image:  pc.Image(),
		Center: pc.Center,
	}

	for _, s := range sc.shapes {
		sHitBox := collision.Hitbox{
			Image:  s.Image(),
			Center: s.CenterCoord(),
		}
		if collision.IsColliding(&pcHitbox, &sHitBox) {
			fmt.Println("collision")
		}
	}
}

func (pc *PlayerCharacter) Image() *ebiten.Image {
	if pc.jumping {
		return pc.imageJumping
	}
	return pc.image
}

func (pc *PlayerCharacter) Len() int {
	return 1
}

func (pc *PlayerCharacter) Dst(i int) (x0, y0, x1, y1 int) {
	w, h := pc.image.Size()
	halfHeight := h / 2
	halfWidth := w / 2
	return pc.Center.x - halfHeight,
		pc.Center.y - halfWidth,
		pc.Center.x + halfHeight,
		pc.Center.y + halfWidth
}

func (pc *PlayerCharacter) Src(i int) (x0, y0, x1, y1 int) {
	w, h := pc.image.Size()
	return 0, 0, w, h
}
