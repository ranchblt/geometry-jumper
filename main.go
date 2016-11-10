package main

import (
	"errors"
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
	shapes          *Shape
)

type Shape struct {
	shapes []gameobj.Drawable
	num    int
}

func (s *Shape) Update() {
	for _, d := range s.shapes {
		d.Update()
	}
}

func (s *Shape) Add(g gameobj.Drawable) {
	s.shapes = append(s.shapes, g)
}

func (s *Shape) Draw(screen *ebiten.Image) {
	for _, d := range s.shapes {
		screen.DrawImage(d.Image(), &ebiten.DrawImageOptions{
			ImageParts: d,
		})
	}
}

func update(screen *ebiten.Image) error {
	keyboardWrapper.Update()
	shapes.Update()
	shapes.Draw(screen)

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
	shapes = &Shape{
		shapes: []gameobj.Drawable{},
	}
	personImage, _, err := ebitenutil.NewImageFromFile("./resource/person.png", ebiten.FilterNearest)
	if err != nil {
		panic(err)
	}

	circle := gameobj.NewCircle(gameobj.NewBaseShape(gameobj.UpperTrack, gameobj.RightSide, 1, 1), personImage)
	square := gameobj.NewSquare(gameobj.NewBaseShape(gameobj.LowerTrack, gameobj.RightSide, 1, 1), personImage)
	triangle := gameobj.NewTriangle(gameobj.NewBaseShape(gameobj.LowerTrack, gameobj.RightSide, 2, 1), personImage)
	shapes.Add(circle)
	shapes.Add(square)
	shapes.Add(triangle)

	player = &PlayerCharacter{
		name:  "Test",
		Image: personImage,
	}
	ebiten.Run(update, screenWidth, screenHeight, 2, "Hello world!")
}
