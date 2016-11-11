package main

import (
	"bytes"
	"errors"
	"geometry-jumper/gameobj"
	"geometry-jumper/keyboard"
	"geometry-jumper/resource"

	"fmt"
	"image"

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

	screen.DrawImage(player.Image, &ebiten.DrawImageOptions{
		ImageParts: player,
	})

	player.Update()

	ebitenutil.DebugPrint(screen, "Hello world!")

	if keyboardWrapper.KeyPushed(ebiten.KeyEscape) {
		return errors.New("User wanted to quit") //Best way to do this?
	}

	return nil
}

func main() {
	shapeCollection = gameobj.NewShapeCollection()

	pImage, err := openImage("person.png")
	handleErr(err)

	personImage, err := ebiten.NewImageFromImage(pImage, ebiten.FilterNearest)
	handleErr(err)

	sImage, err := openImage("square.png")
	handleErr(err)

	squareImage, err := ebiten.NewImageFromImage(sImage, ebiten.FilterNearest)
	handleErr(err)

	tImage, err := openImage("triangle.png")
	handleErr(err)

	triangleImage, err := ebiten.NewImageFromImage(tImage, ebiten.FilterNearest)
	handleErr(err)

	cImage, err := openImage("circle.png")
	handleErr(err)

	circleImage, err := ebiten.NewImageFromImage(cImage, ebiten.FilterNearest)
	handleErr(err)

	circle := gameobj.NewCircle(gameobj.NewBaseShape(gameobj.UpperTrack, gameobj.RightSide, 1, 1), circleImage)
	square := gameobj.NewSquare(gameobj.NewBaseShape(gameobj.LowerTrack, gameobj.RightSide, 1, 1), squareImage)
	triangle := gameobj.NewTriangle(gameobj.NewBaseShape(gameobj.LowerTrack, gameobj.RightSide, 2, 1), triangleImage)
	shapeCollection.Add(circle)
	shapeCollection.Add(square)
	shapeCollection.Add(triangle)

	player = &PlayerCharacter{
		name:  "Test",
		Image: personImage,
	}

	fmt.Printf("Starting up game. Version %s, Build %s", Version, Build)

	ebiten.Run(update, screenWidth, screenHeight, 2, "Hello world!")
}

func openImage(path string) (image.Image, error) {
	b, err := resource.Asset(path)
	if err != nil {
		return nil, err
	}

	image, _, err := image.Decode(bytes.NewReader(b))

	if err != nil {
		return nil, err
	}

	return image, nil
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}
