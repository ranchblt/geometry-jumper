package gameobj

import (
	"bytes"
	"image"
	"image/color"

	"geometry-jumper/resource"

	"github.com/hajimehoshi/ebiten"
)

const (
	// can we specify this via a property? mike pls.
	Debug = true

	// when we're on the last shape in a pattern, use this constant to flag the end.
	EndOfPatternSpawnDelay = -1

	DefaultSpeed = 4

	// 'difficulty' constants for patterns?
	LowDifficulty    = 1
	MediumDifficulty = 2
	HighDifficulty   = 3

	// Track constants
	UpperTrack = 1
	LowerTrack = 2

	// shape constants
	TriangleType = 1
	CircleType   = 2
	SquareType   = 3

	// triangle movement state constants
	TriangleBeforeSwap = 1
	TriangleDuringSwap = 2
	TriangleAfterSwap  = 3

	// space (in pixels, I guess?) between the two tracks
	UpperTrackYAxis = 150
	LowerTrackYAxis = 250

	// x constants for sides and player
	PlayerX   = 60
	LeftSide  = 50
	RightSide = 375
	// this should be screen width probably, not a constant 400
	TrackLength = 400

	// default angle values IN DEGREES!!!! (go math requires radians but degrees make more sense...)
	DefaultCircleAngleOfDescent   float64 = 45
	DefaultTriangleAngleOfDescent float64 = 45

	JumpHeight    = LowerTrackYAxis - UpperTrackYAxis
	JumpUpSpeed   = 5
	JumpDownSpeed = 3

	StartingUpperSpeedLimit = 4
	MinimumSpeed            = 4
)

var (
	// slice of all shape types
	ShapeTypes = [...]int{TriangleType, CircleType, SquareType}

	// track mappings so you can use the track ID to get the track's position on the y axis
	TrackMappings = map[int]int{
		UpperTrack: UpperTrackYAxis,
		LowerTrack: LowerTrackYAxis,
	}

	// subsequent track shows us what track comes after the one we're currently on.
	SubsequentTracks = map[int]int{
		UpperTrack: LowerTrack,
		LowerTrack: UpperTrack,
	}
	PersonStandingImage *ebiten.Image
	PersonJumpingImage  *ebiten.Image
	SquareImage         *ebiten.Image
	SquareBorder        *ebiten.Image
	TriangleImage       *ebiten.Image
	CircleImage         *ebiten.Image
	UpperTrackLine      *ebiten.Image
	LowerTrackLine      *ebiten.Image

	UpperTrackOpts *ebiten.DrawImageOptions
	LowerTrackOpts *ebiten.DrawImageOptions

	ShapeImageMap  map[int][]*ebiten.Image
	HitboxImageMap map[int]*ebiten.Image
)

func InitImages() {
	pImage, err := openImage("person-standing.png")
	handleErr(err)

	PersonStandingImage, err = ebiten.NewImageFromImage(pImage, ebiten.FilterNearest)
	handleErr(err)

	pImage2, err := openImage("person-jumping.png")
	handleErr(err)

	PersonJumpingImage, err = ebiten.NewImageFromImage(pImage2, ebiten.FilterNearest)
	handleErr(err)

	sImage, err := openImage("square.png")
	handleErr(err)

	SquareImage, err = ebiten.NewImageFromImage(sImage, ebiten.FilterNearest)
	handleErr(err)

	squareWidth, squareHeight := SquareImage.Size()
	// this is wrong. need to figure out how to do hollow shapes
	SquareBorder, err = ebiten.NewImage(squareWidth, squareHeight, ebiten.FilterNearest)
	SquareBorder.Fill(color.White)

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
}

func InitImageMaps() {
	ShapeImageMap = map[int][]*ebiten.Image{
		TriangleType: []*ebiten.Image{TriangleImage},
		SquareType:   []*ebiten.Image{SquareImage},
		CircleType:   []*ebiten.Image{CircleImage},
	}

	HitboxImageMap = map[int]*ebiten.Image{
		// todo: real values
		TriangleType: SquareBorder,
		SquareType:   SquareBorder,
		CircleType:   SquareBorder,
	}
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
