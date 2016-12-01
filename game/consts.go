package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
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

	ScreenWidth  = 400
	ScreenHeight = 400
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

	JumpSound *audio.Context
	jumpCh    = make(chan []byte)
	jumpBytes []byte
)
