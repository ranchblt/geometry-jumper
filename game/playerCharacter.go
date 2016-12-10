package game

import (
	"fmt"
	"geometry-jumper/keyboard"
	"image"

	"geometry-jumper/collision"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
)

type PlayerCharacter struct {
	name             string
	image            *ebiten.Image
	imageJumping     *ebiten.Image
	rgbaImage        *image.RGBA
	rgbaJumpImage    *image.RGBA
	keyboardWrapper  *keyboard.KeyboardWrapper
	Center           *coord
	Collided         bool
	jumping          bool
	maxHeightReached bool
	originalY        int
	score            int
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
	if JumpBytes == nil {
		select {
		case JumpBytes = <-JumpCh:
		default:
		}
	}

	if JumpSound != nil {
		if err := JumpSound.Update(); err != nil {
			return err
		}

	}

	if pc.keyboardWrapper.KeyPushed(ebiten.KeySpace) {
		if !pc.jumping {
			pc.jumping = true
			playJumpSound()
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

func (pc *PlayerCharacter) Score() int {
	return pc.score
}

func (pc *PlayerCharacter) CheckScore(sc *ShapeCollection) {
	for _, s := range sc.shapes {
		x0, _, _, _ := s.Dst(1)
		if x0 < pc.Center.X() && !s.Scored() {
			s.SetScore(true)
			pc.score++
		}
	}
}

func (pc *PlayerCharacter) CheckCollision(sc *ShapeCollection) {
	pcHitbox := collision.Hitbox{
		Image:  pc.RgbaImage(),
		Center: pc.Center,
	}

	for _, s := range sc.shapes {
		sHitBox := collision.Hitbox{
			Image:  s.RgbaImage(),
			Center: s.CenterCoord(),
		}
		if collision.IsColliding(&pcHitbox, &sHitBox) {
			if Debug {
				go fmt.Println("collision")
			}
			pc.Collided = true
		}
	}
}

func (pc *PlayerCharacter) Image() *ebiten.Image {
	if pc.jumping {
		return pc.imageJumping
	}
	return pc.image
}

func (pc *PlayerCharacter) RgbaImage() *image.RGBA {
	if pc.jumping {
		if pc.rgbaJumpImage == nil {
			pc.rgbaJumpImage = toRGBA(pc.imageJumping)
		}
		return pc.rgbaJumpImage
	}
	if pc.rgbaImage == nil {
		pc.rgbaImage = toRGBA(pc.image)
	}
	return pc.rgbaImage
}

func (pc *PlayerCharacter) Len() int {
	return 1
}

func (pc *PlayerCharacter) Dst(i int) (x0, y0, x1, y1 int) {
	w, h := pc.image.Size()
	h = h / 2
	w = w / 2
	return pc.Center.x - h,
		pc.Center.y - w,
		pc.Center.x + h,
		pc.Center.y + w
}

func (pc *PlayerCharacter) Src(i int) (x0, y0, x1, y1 int) {
	w, h := pc.image.Size()
	return 0, 0, w, h
}

// made into function so can goroutine
func playJumpSound() {
	jumpSoundPlayer, _ := audio.NewPlayerFromBytes(JumpSound, JumpBytes)
	jumpSoundPlayer.Play()
}
