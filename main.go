package main

import (
	"geometry-jumper/keyboard"

	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 400
	screenHeight = 400
)

var (
	player          *PlayerCharacter
	keyboardWrapper = keyboard.NewKeyboardWrapper()
)

func (pc *PlayerCharacter) Update() error {
	if keyboardWrapper.KeyPushed(ebiten.KeySpace) {
		fmt.Print("you pushed space")
	}
	return nil
}

func update(screen *ebiten.Image) error {
	screen.DrawImage(player.Image, &ebiten.DrawImageOptions{
		ImageParts: player,
	})

	player.Update()
	keyboardWrapper.Update()
	ebitenutil.DebugPrint(screen, "Hello world!")
	return nil
}

func main() {
	personImage, _, err := ebitenutil.NewImageFromFile("./resource/person.png", ebiten.FilterNearest)
	if err != nil {
		panic(err)
	}

	player = &PlayerCharacter{
		name:  "Test",
		Image: personImage,
	}
	ebiten.Run(update, screenWidth, screenHeight, 2, "Hello world!")
}
