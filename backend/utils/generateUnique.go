package utils

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func GenUniqID() ulid.ULID {
	t := time.Unix(1000000, 0)
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id
}
