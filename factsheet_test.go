package ideas

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

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

func Test_FactsheetJSON(t *testing.T) {
	fs := NewDictionaryFactsheet(newStrID())
	jb, err := json.Marshal(fs)
	if err != nil {
		t.Errorf("Error marshaling empty factsheet: %v", err)
	}
	if !strings.Contains(string(jb), "\"_ID\":\"") {
		t.Errorf("JSON should contain _ID, but does not: %s", string(jb))
	}
	if !strings.Contains(string(jb), "\"_VERSION\":0") {
		t.Errorf("JSON should contain _VERSION, but does not: %s", string(jb))
	}
	if !strings.Contains(string(jb), "\"_TIMESTAMP\":") {
		t.Errorf("JSON should contain _TIMESTAMP, but does not: %s", string(jb))
	}
	fs.SetValue("ali", "ali")
	jb, err = json.Marshal(fs)
	if err != nil {
		t.Errorf("Error marshaling factsheet: %v", err)
	}
	if !strings.Contains(string(jb), "\"ali\":\"") {
		t.Errorf("JSON should contain ali, but does not: %s", string(jb))
	}
	fmt.Println(string(jb))
}
