package main

import (
	"fmt"
	"geometry-jumper/keyboard"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 400
	screenHeight = 400
)

type PlayerCharacter struct {
	keyboard *keyboard.KeyboardWrapper
	name     string
}

var (
	player *PlayerCharacter
	// this is weird. used primarily to make sure we capture key presses rather than key holds
	keyState = map[ebiten.Key]int{}
)

func (pc *PlayerCharacter) Update() error {
	pc.keyboard.Update()

	if pc.keyboard.KeyPushed(ebiten.KeySpace) {
		fmt.Print("you pushed space this cycle")
	}

	return nil
}

func update(screen *ebiten.Image) error {
	p := &personImageParts{image: personImage}
	screen.DrawImage(personImage, &ebiten.DrawImageOptions{
		ImageParts: p,
	})

	player.Update()
	ebitenutil.DebugPrint(screen, "Hello world!")
	return nil
}

func main() {
	var err error
	personImage, _, err = ebitenutil.NewImageFromFile("./resource/person.png", ebiten.FilterNearest)
	if err != nil {
		panic(err)
	}

	player = &PlayerCharacter{
		keyboard: keyboard.NewKeyboardWrapper(),
		name:     "Test",
	}
	ebiten.Run(update, screenWidth, screenHeight, 2, "Hello world!")
}
