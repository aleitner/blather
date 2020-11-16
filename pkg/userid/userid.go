package userid

import (
	"strconv"
)

type ID int

func (id ID) toInt() int {
	return int(id)
}

func (id ID) String() string {
	return strconv.Itoa(id.toInt())
}

func FromString(id string) (ID, error) {
	n, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, nil
	}

	return ID(n), nil
}
