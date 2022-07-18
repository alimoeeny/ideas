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

func Test_Factsheet_Merge(t *testing.T) {
	// emtpy
	{
		fs := MergeFactsheetsAndOverwriteLeft(NewDictionaryFactsheet(""))
		if fs.ID() != "" {
			t.Errorf("Expect '' but got %s\n", fs.ID())
			t.FailNow()
		}

		fs = MergeFactsheetsAndOverwriteLeft(NewDictionaryFactsheet("12"))
		if fs.ID() != "12" {
			t.Errorf("Expect '12' but got %s\n", fs.ID())
			t.FailNow()
		}

		fs = MergeFactsheetsAndOverwriteLeft(NewDictionaryFactsheet("12"), NewDictionaryFactsheet("13"))
		if fs.ID() != "12" {
			t.Errorf("Expect '12' but got %s\n", fs.ID())
			t.FailNow()
		}
	}
	// basics
	{
		A := NewDictionaryFactsheet("12")
		A.SetValue("a", 1)
		A.SetValue("b", 2)
		B := NewDictionaryFactsheet("13")
		B.SetValue("A", 11)
		fs := MergeFactsheetsAndOverwriteLeft(A, B)
		if fs.ID() != "12" {
			t.Errorf("Expect '12' but got %s\n", fs.ID())
			t.FailNow()
		}
		if fs.CurrentValue("a") != 1 {
			t.Errorf("Expect '1' for 'a' but got %s\n", fs.CurrentValue("a"))
			t.FailNow()
		}
		if fs.CurrentValue("b") != 2 {
			t.Errorf("Expect '2' for 'b' but got %s\n", fs.CurrentValue("b"))
			t.FailNow()
		}
		if fs.CurrentValue("A") != 11 {
			t.Errorf("Expect '11' for 'A' but got %s\n", fs.CurrentValue("A"))
			t.FailNow()
		}
	}
	// left overwrite
	{
		A := NewDictionaryFactsheet("12")
		A.SetValue("a", 1)
		A.SetValue("b", 2)
		B := NewDictionaryFactsheet("13")
		B.SetValue("A", 11)
		B.SetValue("b", 12)
		fs := MergeFactsheetsAndOverwriteLeft(A, B)
		if fs.ID() != "12" {
			t.Errorf("Expect '12' but got %s\n", fs.ID())
			t.FailNow()
		}
		if fs.CurrentValue("a") != 1 {
			t.Errorf("Expect '1' for 'a' but got %s\n", fs.CurrentValue("a"))
			t.FailNow()
		}
		if fs.CurrentValue("b") != 12 {
			t.Errorf("Expect '12' for 'b' but got %s\n", fs.CurrentValue("b"))
			t.FailNow()
		}
		if fs.CurrentValue("A") != 11 {
			t.Errorf("Expect '11' for 'A' but got %s\n", fs.CurrentValue("A"))
			t.FailNow()
		}
	}
}
