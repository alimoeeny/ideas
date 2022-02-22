package ideas

import (
	"fmt"
	"math/rand"
)

// Identifier unlinke lowlevel id, can be exposed externally
type Identifier struct {
	ID     string `json:"id,omitempty"`
	Source string `json:"source,omitempty"`
}

func newID() int64 {
	return rand.Int63()
}

func newStrID() string {
	return fmt.Sprintf("%d", newID())
}

func newIdentifier() Identifier {
	return Identifier{newStrID(), "https://identifier.suprafrontal.com/"}
}
