package ideas

import "testing"

func Test_Identifier_Encoding(t *testing.T) {
	id := newIdentifier()
	s := id.String()
	var idBack Identifier
	idBack.FromEncodedString(s)
	if id.ID != idBack.ID || id.Source != idBack.Source {
		t.FailNow()
	}
}
