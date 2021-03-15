package userid

import (
	"strconv"
)

type ID string

func (id ID) String() string {
	return string(id)
}

func FromInt(id int) ID {
	return ID(strconv.Itoa(id))
}
