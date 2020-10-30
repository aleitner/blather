package coordinates

import (
	"math"

	call "github.com/aleitner/spacialPhone/internal/protobuf"
)

type Coordinate struct {
	x float64
	y float64
	z float64
}

func (c Coordinate) ToGRPC() *call.Coordinates {
	return &call.Coordinates{
		X: c.x,
		Y: c.y,
		Z: c.z,
	}
}

func (c Coordinate) Distance(object Coordinate) float64 {
	return math.Sqrt(math.Pow(object.x-c.x, 2) + math.Pow(object.y-c.y, 2) + math.Pow(object.z-c.z, 2))
}
