package main

import (
	"geometry-jumper/gameobj"
	"geometry-jumper/keyboard"

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
	square          *gameobj.Square
)

func update(screen *ebiten.Image) error {
	screen.DrawImage(player.Image, &ebiten.DrawImageOptions{
		ImageParts: player,
	})

	screen.DrawImage(square.Image, &ebiten.DrawImageOptions{
		ImageParts: square,
	})

	player.Update()
	square.Update()
	keyboardWrapper.Update()
	ebitenutil.DebugPrint(screen, "Hello world!")
	return nil
}

func main() {
	personImage, _, err := ebitenutil.NewImageFromFile("./resource/person.png", ebiten.FilterNearest)
	if err != nil {
		panic(err)
	}

	square = gameobj.NewSquare(&gameobj.BaseShape{
		Track:         1,
		CenterX:       30,
		CenterY:       30,
		BaseSpeed:     -.05,
		SpeedModifier: 1,
	}, personImage)

	player = &PlayerCharacter{
		name:  "Test",
		Image: personImage,
	}
	ebiten.Run(update, screenWidth, screenHeight, 2, "Hello world!")
}
