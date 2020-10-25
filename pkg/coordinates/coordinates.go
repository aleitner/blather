package coordinates

import "math"

type Coordinate struct {
	x float64
	y float64
	z float64
}

func (c Coordinate) Distance(object Coordinate) float64 {
	return math.Sqrt(math.Pow(object.x-c.x, 2) + math.Pow(object.y-c.y, 2) + math.Pow(object.z-c.z, 2))
}
