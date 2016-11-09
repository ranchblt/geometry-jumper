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
	circle          *gameobj.Circle
	square          *gameobj.Square
)

func update(screen *ebiten.Image) error {
	screen.DrawImage(player.Image, &ebiten.DrawImageOptions{
		ImageParts: player,
	})

	screen.DrawImage(square.Image, &ebiten.DrawImageOptions{
		ImageParts: square,
	})

	screen.DrawImage(circle.Image, &ebiten.DrawImageOptions{
		ImageParts: circle,
	})

	player.Update()
	square.Update()
	circle.Update()
	keyboardWrapper.Update()
	ebitenutil.DebugPrint(screen, "Hello world!")
	return nil
}

func main() {
	personImage, _, err := ebitenutil.NewImageFromFile("./resource/person.png", ebiten.FilterNearest)
	if err != nil {
		panic(err)
	}

	circle = gameobj.NewCircle(gameobj.NewBaseShape(gameobj.UpperTrack, gameobj.RightSide, .35, 1), personImage, gameobj.SubsequentTracks[gameobj.UpperTrack])
	square = gameobj.NewSquare(gameobj.NewBaseShape(gameobj.LowerTrack, gameobj.RightSide, .10, 1), personImage)
	player = &PlayerCharacter{
		name:  "Test",
		Image: personImage,
	}
	ebiten.Run(update, screenWidth, screenHeight, 2, "Hello world!")
}
