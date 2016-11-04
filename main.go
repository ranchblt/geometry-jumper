package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
        "fmt"
)

const (
	screenWidth  = 400
	screenHeight = 400
)

type PlayerCharacter struct {
        name string
}

var (
        player *PlayerCharacter
        // this is weird. used primarily to make sure we capture key presses rather than key holds
        keyState = map[ebiten.Key]int{}
)


func (pc *PlayerCharacter) keyboardInput() error {


        if !ebiten.IsKeyPressed(ebiten.KeySpace) {
            keyState[ebiten.KeySpace] = 0
                
        } else {
                keyState[ebiten.KeySpace]++
                // when this is 1, it's just after we pushed a key, which means we should actually do something with it.
                // should probably pull this out into its own thing, I think? we might be able to loop over all of the KeySpace
                // and do it implicitly. not sure. 
                if keyState[ebiten.KeySpace] == 1 {
                        fmt.Print("you pushed space")
                }
        }
        return nil
}

func update(screen *ebiten.Image) error {
	p := &personImageParts{image: personImage}
	screen.DrawImage(personImage, &ebiten.DrawImageOptions{
		ImageParts: p,
	})

        player.keyboardInput()
	ebitenutil.DebugPrint(screen, "Hello world!")
	return nil
}


func main() {
	var err error
	personImage, _, err = ebitenutil.NewImageFromFile("./resource/person.png", ebiten.FilterNearest)
	if err != nil {
		panic(err)
	}

        player = &PlayerCharacter{ name : "Test" }
	ebiten.Run(update, screenWidth, screenHeight, 2, "Hello world!")
}
