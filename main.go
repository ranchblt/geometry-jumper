package main

import (
	"errors"
	"geometry-jumper/game"
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
	player          *game.PlayerCharacter
	keyboardWrapper = keyboard.NewKeyboardWrapper()
	shapeCollection *game.ShapeCollection
)

// Version is autoset from the build script
var Version string

// Build is autoset from the build script
var Build string

func update(screen *ebiten.Image) error {
	if game.Debug {
		screen.DrawImage(game.UpperTrackLine, game.UpperTrackOpts)
		screen.DrawImage(game.LowerTrackLine, game.LowerTrackOpts)
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
	game.Load()

	square := game.NewSpawnDefaultSpeed(game.SquareType, game.LowerTrack, 5)
	triangle := game.NewSpawnDefaultSpeed(game.TriangleType, game.UpperTrack, 5)

	pattern := game.NewPattern([]*game.Spawn{square, triangle})
	patternCollection := &game.PatternCollection{
		Patterns: map[int][]*game.Pattern{
			game.LowDifficulty: []*game.Pattern{pattern},
		},
	}

	shapeCollection = game.NewShapeCollection(patternCollection)

	player = game.NewPlayerCharacter("Test", game.PersonStandingImage, game.PersonJumpingImage, keyboardWrapper)

	fmt.Printf("Starting up game. Version %s, Build %s", Version, Build)

	ebiten.Run(update, screenWidth, screenHeight, 2, "Hello world!")
}
