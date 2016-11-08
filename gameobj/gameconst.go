package gameobj

const (
	// Track constants
	UpperTrack = 1
	LowerTrack = 2

	// space (in pixels, I guess?) between the two tracks
	TrackGap        = 100
	UpperTrackYAxis = 50
	LowerTrackYAxis = 150

	// default angle values IN DEGREES!!!! (go math requires radians but degrees make more sense...)
	DefaultCircleAngleOfDescent   float64 = 45
	DefaultTriangleAngleOfDescent float64 = 45
)
