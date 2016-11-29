package main

import (
	"errors"
	"geometry-jumper/gameobj"
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
	player          *gameobj.PlayerCharacter
	keyboardWrapper = keyboard.NewKeyboardWrapper()
	shapeCollection *gameobj.ShapeCollection
)

// Version is autoset from the build script
var Version string

// Build is autoset from the build script
var Build string

func update(screen *ebiten.Image) error {
	if gameobj.Debug {
		screen.DrawImage(gameobj.UpperTrackLine, gameobj.UpperTrackOpts)
		screen.DrawImage(gameobj.LowerTrackLine, gameobj.LowerTrackOpts)
	}

	keyboardWrapper.Update()
	shapeCollection.Update()
	shapeCollection.Draw(screen)

	player.Update()
	player.Draw(screen)

	ebitenutil.DebugPrint(screen, "Hello world!")

	if keyboardWrapper.KeyPushed(ebiten.KeyEscape) {
		return errors.New("User wanted to quit") //Best way to do this?
	}

	return nil
}

func main() {
	gameobj.InitImages()
	gameobj.InitImageMaps()

	shapeCollection = gameobj.NewShapeCollection()

	shapeCollection.SpawnRandomShape()
	shapeCollection.SpawnRandomShape()

	player = gameobj.NewPlayerCharacter("Test", gameobj.PersonStandingImage, gameobj.PersonJumpingImage, keyboardWrapper)

	fmt.Printf("Starting up game. Version %s, Build %s", Version, Build)

	ebiten.Run(update, screenWidth, screenHeight, 2, "Hello world!")
}
