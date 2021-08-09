package ideas

import (
	"testing"
	"time"
)

func Test_basics(t *testing.T) {
	nid := newStrID()
	u := Unit{nid, "mg per Kg", "mg/kg"}
	if u.ID != nid {
		t.FailNow()
	}

	m := Measurment{nid, time.Now().UnixNano(), "42", &u}
	if m.ID != nid {
		t.FailNow()
	}
}
