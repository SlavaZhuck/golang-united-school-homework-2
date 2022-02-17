package square

import (
	"math"
)

//package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type sideNumber int

const (
	SidesTriangle sideNumber = 3
	SidesSquare   sideNumber = 4
	SidesCircle   sideNumber = 0
)

func CalcSquare(sideLen float64, sidesNum sideNumber) float64 {

	switch sidesNum {
	case SidesTriangle:
		return 0.5 * sideLen * sideLen * math.Sin(60.0*math.Pi/180.0)
	case SidesSquare:
		return sideLen * sideLen
	case SidesCircle:
		return sideLen * sideLen * math.Pi
	default:
		return 0
	}
}
