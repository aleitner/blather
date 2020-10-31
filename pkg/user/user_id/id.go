package user_id

import (
	"context"
	"fmt"
	"strconv"

	"google.golang.org/grpc/metadata"
)

type ID int

func (id ID) toInt() int {
	return int(id)
}

func (id ID) String() string {
	return strconv.Itoa(id.toInt())
}

func IDFromString(id string) (ID, error) {
	n, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return 0, nil
	}

	return ID(n), nil
}

func IDFromMetaData(ctx context.Context) (ID, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0, fmt.Errorf("Failed to retrieve incoming context")
	}

	IDAsString := md.Get("client-id")

	if len(IDAsString) <= 0 {
		return 0, fmt.Errorf("Failed to retrieve incoming id")
	}

	contactID, err := IDFromString(IDAsString[0])
	if err != nil {
		return 0, fmt.Errorf("Failed to parse incoming id: %s", err.Error())
	}

	return contactID, nil
}
