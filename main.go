package main

import (
	"errors"
	"flag"
	"geometry-jumper/game"
	"geometry-jumper/keyboard"
	"os"
	"runtime/pprof"

	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
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

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func gameLoop(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		if game.Debug {
			go fmt.Println("slow")
		}
		return nil
	}

	if game.Debug {
		screen.DrawImage(game.UpperTrackLine, game.UpperTrackOpts)
		screen.DrawImage(game.LowerTrackLine, game.LowerTrackOpts)
	}

	keyboardWrapper.Update()
	shapeCollection.Update()
	shapeCollection.Draw(screen)

	player.Update()
	player.Draw(screen)

	player.CheckCollision(shapeCollection)

	ebitenutil.DebugPrint(screen, "Hello world!")

	if keyboardWrapper.KeyPushed(ebiten.KeyEscape) {
		return errors.New("User wanted to quit") //Best way to do this?
	}

	return nil
}

func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)

		if err != nil {
			panic(err)
		}

		pprof.StartCPUProfile(f)

		defer pprof.StopCPUProfile()

	}

	game.Load()

	square := game.NewSpawnDefaultSpeed(game.SquareType, game.LowerTrack, 2000)
	triangle := game.NewSpawnDefaultSpeed(game.TriangleType, game.UpperTrack, 2000)

	pattern := game.NewPattern([]*game.Spawn{square, triangle})
	patternCollection := &game.PatternCollection{
		Patterns: map[int][]*game.Pattern{
			game.LowDifficulty: []*game.Pattern{pattern},
		},
	}

	shapeCollection = game.NewShapeCollection(patternCollection)
	shapeCollection.UnlockColorSwap()
	player = game.NewPlayerCharacter("Test", game.PersonStandingImage, game.PersonJumpingImage, keyboardWrapper)

	fmt.Printf("Starting up game. Version %s, Build %s", Version, Build)

	ebiten.Run(gameLoop, game.ScreenWidth, game.ScreenHeight, 2, "Hello world!")
}
