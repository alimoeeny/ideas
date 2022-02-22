package ideas

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"
)

// Identifier unlinke lowlevel id, can be exposed externally
type Identifier struct {
	ID     string `json:"id,omitempty"`
	Source string `json:"source,omitempty"`
}

const (
	IDENTIFIER_SEPARATOR_SEQUENCE = "IDENTIFIER_SEPARATOR_SEQUENCE"
)

func (id *Identifier) String() string {
	s := id.ID + IDENTIFIER_SEPARATOR_SEQUENCE + id.Source
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func (id *Identifier) FromEncodedString(encodedString string) error {
	s, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return err
	}
	bits := strings.Split(string(s), IDENTIFIER_SEPARATOR_SEQUENCE)
	if len(bits) != 2 {
		return fmt.Errorf("this does not look like an Identifier: %s", string(s))
	}
	id.ID = bits[0]
	id.Source = bits[1]
	return nil
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
