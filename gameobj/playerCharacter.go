package gameobj

import (
	"fmt"
	"geometry-jumper/keyboard"

	"github.com/hajimehoshi/ebiten"
)

type PlayerCharacter struct {
	name            string
	Image           *ebiten.Image
	keyboardWrapper *keyboard.KeyboardWrapper
}

func NewPlayerCharacter(name string, image *ebiten.Image, keyboardWrapper *keyboard.KeyboardWrapper) *PlayerCharacter {
	var player = &PlayerCharacter{
		name:            "Test",
		Image:           image,
		keyboardWrapper: keyboardWrapper,
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
	return 20, 20, 60, 60
}

func (pc *PlayerCharacter) Src(i int) (x0, y0, x1, y1 int) {
	w, h := pc.Image.Size()
	return 0, 0, w, h
}
