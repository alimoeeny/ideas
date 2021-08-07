package ideas

import (
	"fmt"
	"math/rand"
)

// Identifier unlinke lowlevel id, can be exposed externally
type Identifier struct {
	id     int64
	source string
}

func newID() int64 {
	return rand.Int63()
}

func newStrID() string {
	return fmt.Sprintf("%d", newID())
}

func newIdentifier() Identifier {
	return Identifier{newID(), "https://identifier.suprafrontal.com/"}
}
