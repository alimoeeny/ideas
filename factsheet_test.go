package ideas

import "testing"

func Test_Factsheet(t *testing.T) {
	fs := NewDictionaryFactsheet()
	if fs.CurrentValue("ali") != nil {
		t.FailNow()
	}
	fs.SetValue("ali", "ali")
	if fs.CurrentValue("ali") != "ali" {
		t.FailNow()
	}
}
