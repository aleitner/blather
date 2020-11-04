package coordinates

import (
	"math"

	call "github.com/aleitner/blather/pkg/protobuf"
)

type Coordinate struct {
	X float64
	Y float64
	Z float64
}

func (c Coordinate) ToGRPC() *call.Coordinates {
	return &call.Coordinates{
		X: c.X,
		Y: c.Y,
		Z: c.Z,
	}
}

func (c Coordinate) Distance(object Coordinate) float64 {
	return math.Sqrt(math.Pow(object.X-c.X, 2) + math.Pow(object.Y-c.Y, 2) + math.Pow(object.Z-c.Z, 2))
}
