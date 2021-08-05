package ideas

import "math/rand"

// Identifier unlinke lowlevel id, can be exposed externally
type Identifier struct {
	id     int64
	source string
}

func newID() int64 {
	return rand.Int63()
}

func newIdentifier() Identifier {
	return Identifier{newID(), "https://identifier.suprafrontal.com/"}
}
