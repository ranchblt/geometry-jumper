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

	shapeImageMap := map[int]*ebiten.Image{
		gameobj.TriangleType: gameobj.TriangleImage,
		gameobj.SquareType:   gameobj.SquareImage,
		gameobj.CircleType:   gameobj.CircleImage,
	}

	shapeCollection = gameobj.NewShapeCollection(shapeImageMap)

	//circle := gameobj.NewCircle(gameobj.NewBaseShape(gameobj.UpperTrack, gameobj.RightSide, 1, circleImage))
	//square := gameobj.NewSquare(gameobj.NewBaseShape(gameobj.LowerTrack, gameobj.RightSide, 1, squareImage))
	//triangle := gameobj.NewTriangle(gameobj.NewBaseShape(gameobj.LowerTrack, gameobj.RightSide, 2, triangleImage))
	//shapeCollection.Add(circle)
	//shapeCollection.Add(square)
	//shapeCollection.Add(triangle)
	shapeCollection.SpawnRandomShape()
	shapeCollection.IncreaseSpeedModifier()
	shapeCollection.IncreaseSpeedModifier()
	shapeCollection.IncreaseSpeedModifier()
	shapeCollection.IncreaseSpeedModifier()
	shapeCollection.SpawnRandomShape()

	player = gameobj.NewPlayerCharacter("Test", gameobj.PersonStandingImage, gameobj.PersonJumpingImage, keyboardWrapper)

	fmt.Printf("Starting up game. Version %s, Build %s", Version, Build)

	ebiten.Run(update, screenWidth, screenHeight, 2, "Hello world!")
}
