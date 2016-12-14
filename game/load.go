package game

import (
	"image/color"
	_ "image/png"
	"sync"
	"time"

	"github.com/ranchblt/geometry-jumper/resource"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/uber-go/zap"
)

func Load(lg zap.Logger) {
	logger = lg
	defer timeTrack(time.Now(), "Game.Load")
	var wg sync.WaitGroup

	wg.Add(5)

	go func() {
		defer wg.Done()
		// This is very fragile. initImages must come first!
		// I guess we could just call initImageMaps inside of initImages...?
		initImages()
		initImageMaps()
	}()

	go func() {
		defer wg.Done()

		err := loadAudio()
		handleErr(err)
	}()

	go func() {
		defer wg.Done()
		initColorMaps()
	}()

	go func() {
		defer wg.Done()
		initFont()
	}()

	go func() {
		defer wg.Done()
		initPatternCollection()
	}()

	wg.Wait()
}

func initImages() {
	pImage, err := openImage("Robot2.png")
	handleErr(err)

	PersonStandingImage, err = ebiten.NewImageFromImage(pImage, ebiten.FilterNearest)
	handleErr(err)

	pImage2, err := openImage("Robot2_jump.png")
	handleErr(err)

	PersonJumpingImage, err = ebiten.NewImageFromImage(pImage2, ebiten.FilterNearest)
	handleErr(err)

	platImage, err := openImage("platform.png")
	handleErr(err)

	PlatformImage, err = ebiten.NewImageFromImage(platImage, ebiten.FilterNearest)
	handleErr(err)

	sImage, err := openImage("square.png")
	handleErr(err)

	SquareImage, err = ebiten.NewImageFromImage(sImage, ebiten.FilterNearest)
	handleErr(err)

	tImage, err := openImage("triangle.png")
	handleErr(err)

	TriangleImage, err = ebiten.NewImageFromImage(tImage, ebiten.FilterNearest)
	handleErr(err)

	cImage, err := openImage("circle.png")
	handleErr(err)

	CircleImage, err = ebiten.NewImageFromImage(cImage, ebiten.FilterNearest)
	handleErr(err)

	UpperTrackLine, err = ebiten.NewImage(TrackLength, 1, ebiten.FilterNearest)
	UpperTrackLine.Fill(color.White)
	handleErr(err)

	LowerTrackLine, err = ebiten.NewImage(TrackLength, 1, ebiten.FilterNearest)
	LowerTrackLine.Fill(color.White)
	handleErr(err)

	UpperTrackOpts = &ebiten.DrawImageOptions{}
	UpperTrackOpts.GeoM.Translate(0, UpperTrackYAxis)

	LowerTrackOpts = &ebiten.DrawImageOptions{}
	LowerTrackOpts.GeoM.Translate(0, LowerTrackYAxis)

	titleImage, err := openImage("title.png")
	handleErr(err)

	TitleImage, err = ebiten.NewImageFromImage(titleImage, ebiten.FilterNearest)
	handleErr(err)

	end, err := openImage("end.png")
	handleErr(err)

	EndImage, err = ebiten.NewImageFromImage(end, ebiten.FilterNearest)
	handleErr(err)
}

func initImageMaps() {
	ShapeImageMap = map[int][]*ebiten.Image{
		TriangleType: []*ebiten.Image{TriangleImage},
		SquareType:   []*ebiten.Image{SquareImage},
		CircleType:   []*ebiten.Image{CircleImage},
	}
}

func initColorMaps() {
	DefaultSquareColorMap = ebiten.ColorM{}
	DefaultSquareColorMap.Scale(255, 0, 0, 100)

	DefaultCircleColorMap = ebiten.ColorM{}
	DefaultCircleColorMap.Scale(0, 255, 0, 100)

	DefaultTriangleColorMap = ebiten.ColorM{}
	DefaultTriangleColorMap.Scale(0, 0, 255, 100)

	ColorMappings = map[int]ebiten.ColorM{
		SquareType:   DefaultSquareColorMap,
		CircleType:   DefaultCircleColorMap,
		TriangleType: DefaultTriangleColorMap,
	}
}

func initFont() {
	fontAsset, err := resource.Asset("3Dventure.ttf")
	handleErr(err)

	Font, err = truetype.Parse(fontAsset)
	handleErr(err)
}

func initPatternCollection() {
	data, err := resource.Asset("patterns.json")
	handleErr(err)

	GamePatternCollection = PatternCollectionFromJSON(data)
}
