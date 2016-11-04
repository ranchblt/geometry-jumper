package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 400
	screenHeight = 400
)

func update(screen *ebiten.Image) error {
	p := &personImageParts{image: personImage}
	screen.DrawImage(personImage, &ebiten.DrawImageOptions{
		ImageParts: p,
	})

	ebitenutil.DebugPrint(screen, "Hello world!")
	return nil
}

func main() {
	var err error
	personImage, _, err = ebitenutil.NewImageFromFile("./resource/person.png", ebiten.FilterNearest)
	if err != nil {
		panic(err)
	}

	ebiten.Run(update, screenWidth, screenHeight, 2, "Hello world!")
}
