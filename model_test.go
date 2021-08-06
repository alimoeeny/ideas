package ideas

import (
	"testing"
	"time"
)

func Test_basics(t *testing.T) {
	nid := newID()
	u := Unit{nid, "mg per Kg", "mg/kg"}
	if u.id != nid {
		t.FailNow()
	}

	m := Measurment{nid, time.Now().UnixNano(), "42", &u}
	if m.id != nid {
		t.FailNow()
	}
}
