package ideas

import "testing"

func Test_Factsheet(t *testing.T) {
	id := newStrID()
	fs := NewDictionaryFactsheet(id)
	if fs.CurrentValue("ali") != nil {
		t.FailNow()
	}
	fs.SetValue("ali", "ali")
	if fs.CurrentValue("ali") != "ali" {
		t.FailNow()
	}
	if fs.ID() != id {
		t.Errorf("ID should be %s, but is %s", id, fs.ID())
	}
}
