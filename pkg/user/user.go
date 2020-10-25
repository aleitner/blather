package user

import (
	"github.com/aleitner/spacialPhone/pkg/coordinates"
)

type User struct {
	id       int64
	address  string
	position *coordinates.Coordinate
}
