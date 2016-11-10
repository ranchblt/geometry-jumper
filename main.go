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
	drawables       []gameobj.Drawable
)

func update(screen *ebiten.Image) error {
	for _, d := range drawables {
		screen.DrawImage(d.Image(), &ebiten.DrawImageOptions{
			ImageParts: d,
		})
		d.Update()
	}

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

	circle := gameobj.NewCircle(gameobj.NewBaseShape(gameobj.UpperTrack, gameobj.RightSide, .35, 1), personImage)
	square := gameobj.NewSquare(gameobj.NewBaseShape(gameobj.LowerTrack, gameobj.RightSide, .10, 1), personImage)
	triangle := gameobj.NewTriangle(gameobj.NewBaseShape(gameobj.LowerTrack, gameobj.RightSide, 1, 1), personImage)
	drawables = append(drawables, circle)
	drawables = append(drawables, square)
	drawables = append(drawables, triangle)

	player = &PlayerCharacter{
		name:  "Test",
		Image: personImage,
	}
	ebiten.Run(update, screenWidth, screenHeight, 2, "Hello world!")
}
