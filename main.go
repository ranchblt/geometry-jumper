package main

import (
	"errors"
	"flag"
	"image"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"

	"geometry-jumper/game"
	"geometry-jumper/keyboard"
	"geometry-jumper/menu"
	"geometry-jumper/ranchblt"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/uber-go/zap"
)

var (
	player          *game.PlayerCharacter
	keyboardWrapper = keyboard.NewKeyboardWrapper()
	shapeCollection *game.ShapeCollection
	logoScreen      *ranchblt.Logo
	showLogo        = true
	showMenu        = true
	mainMenu        menu.Menu
	endMenu         menu.Menu
	slowDownCount   int64
	logger          zap.Logger
)

// Version is autoset from the build script
var Version string

// Build is autoset from the build script
var Build string

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func gameLoop(screen *ebiten.Image) error {
	if ebiten.IsRunningSlowly() {
		slowDownCount++
		logger.Debug("Running slow",
			zap.Int64("Amount", slowDownCount),
		)
		return nil
	}

	keyboardWrapper.Update()

	if showLogo && !game.Debug {
		logoScreen.Draw(screen)
		return nil
	}

	if showMenu {
		mainMenu.Update()
		mainMenu.Draw(screen)
		if keyboardWrapper.KeyPushed(ebiten.KeyEnter) {
			if strings.ToLower(mainMenu.Selected()) == "start" {
				showMenu = false
			} else if strings.ToLower(mainMenu.Selected()) == "exit" {
				return errors.New("User wanted to quit")
			}
		}
		return nil
	}

	if game.Debug {
		screen.DrawImage(game.UpperTrackLine, game.UpperTrackOpts)
		screen.DrawImage(game.LowerTrackLine, game.LowerTrackOpts)
	}

	if !player.Collided {
		shapeCollection.Update()
		player.Update()
	} else {
		shapeCollection.Stop = true
		endMenu.Update()
		endMenu.Draw(screen)
		screen.DrawImage(getScoreImage(player.Score()), &ebiten.DrawImageOptions{})
		if keyboardWrapper.KeyPushed(ebiten.KeyEnter) {
			if strings.ToLower(endMenu.Selected()) == "restart" {
				shapeCollection = game.NewShapeCollection()
				player = game.NewPlayerCharacter("Test", game.PersonStandingImage, game.PersonJumpingImage, keyboardWrapper)
			} else if strings.ToLower(endMenu.Selected()) == "exit" {
				return errors.New("User wanted to quit")
			}
		}
	}

	shapeCollection.Draw(screen)
	player.Draw(screen)

	player.CheckCollision(shapeCollection)
	player.CheckScore(shapeCollection)

	//ebitenutil.DebugPrint(screen, "Hello world!")

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

	logger = zap.New(zap.NewTextEncoder(zap.TextNoTime()), zap.DebugLevel)

	game.Load(logger)

	options := []*menu.Option{}
	options = append(options, &menu.Option{
		Text: "Start",
	})
	options = append(options, &menu.Option{
		Text: "Exit",
	})

	mainMenu = &menu.Regular{
		BackgroundImage: game.TitleImage,
		Height:          game.ScreenHeight,
		Width:           game.ScreenWidth,
		KeyboardWrapper: keyboardWrapper,
		Options:         options,
		Font:            game.Font,
	}

	options2 := []*menu.Option{}
	options2 = append(options2, &menu.Option{
		Text: "Restart",
	})
	options2 = append(options2, &menu.Option{
		Text: "Exit",
	})

	endMenu = &menu.Regular{
		BackgroundImage: game.EndImage,
		Height:          game.ScreenHeight,
		Width:           game.ScreenWidth,
		KeyboardWrapper: keyboardWrapper,
		Options:         options2,
		Font:            game.Font,
	}

	shapeCollection = game.NewShapeCollection()

	player = game.NewPlayerCharacter("Test", game.PersonStandingImage, game.PersonJumpingImage, keyboardWrapper)

	logoScreen = ranchblt.NewLogoScreen(game.ScreenWidth, game.ScreenHeight)

	logger.Info("Starting up game",
		zap.String("Version", Version),
		zap.String("Build", Build),
	)

	go logoTimer()
	ebiten.Run(gameLoop, game.ScreenWidth, game.ScreenHeight, 2, "Geom Jump")
}

func logoTimer() {
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	showLogo = false
}

func getScoreImage(score int) *ebiten.Image {
	const size = 24
	const dpi = 72

	textImage, _ := ebiten.NewImage(game.ScreenWidth, game.ScreenHeight, ebiten.FilterNearest)
	dst := image.NewRGBA(image.Rect(0, 0, game.ScreenWidth, game.ScreenHeight))

	d := &font.Drawer{
		Dst: dst,
		Src: image.White,
		Face: truetype.NewFace(game.Font, &truetype.Options{
			Size:    size,
			DPI:     dpi,
			Hinting: font.HintingFull,
		}),
	}

	st := "Score: " + strconv.Itoa(score)

	s := font.MeasureString(d.Face, st)
	d.Dot = fixed.P(game.ScreenWidth/2-s.Round()/2, game.ScreenHeight-100)
	d.DrawString(st)

	textImage.ReplacePixels(dst.Pix)
	return textImage
}
