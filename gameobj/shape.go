package gameobj

import "math"

// hi there, future Tom or Mike. I'm not sure if this is how ebiten works, but the negatives on centerX calcs is
// to move left along the screen, and the seemingly inverted calculations for moving up / down are because screens typically
// have the zero at the top, not the bottom.

// helper function for gameobj for shapes to convert from degrees to radians
func degreesToRadians(degreeValue float64) float64 {
	return degreeValue * math.Pi / 180
}

type BaseShape struct {
	Track         int
	CenterX       float64
	CenterY       float64
	BaseSpeed     float64
	SpeedModifier float64
}
