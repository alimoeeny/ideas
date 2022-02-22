package ideas

import (
	"encoding/json"
	"testing"
	"time"
)

func Test_concept(t *testing.T) {
	nid := newStrID()
	c := Concept{
		ID:                             nid,
		EnglishHumanReadableExpression: "RBC Count",
		EnglishDescription:             "Average number of RBCs per µl",
	}
	if c.ID != nid {
		t.FailNow()
	}
}

func Test_idea(t *testing.T) {
	crbcc := Concept{
		ID:                             newStrID(),
		EnglishHumanReadableExpression: "RBC Count",
		EnglishDescription:             "Average number of RBCs per µl",
	}
	p := Concept{
		ID:                             newStrID(),
		EnglishHumanReadableExpression: "Related to Patient with ID",
		EnglishDescription:             "to express association of other concepts in the same idea to a particular patient with id",
	}

	nid := newStrID()
	i := Idea{
		ID:                             nid,
		EnglishHumanReadableExpression: "Patient with id 123 had RBC count of 5.3 µl at time t",
		Facts: map[string]*Measurement{
			p.ID: {
				ID:        newStrID(),
				Timestamp: 0,
				Value:     "PATIENT-123",
				Unit:      UNIT_IDENTIFIER,
			},
			crbcc.ID: {
				ID:        newStrID(),
				Timestamp: time.Now().UnixNano(),
				Value:     5900000,
				Unit:      UNIT_COUNT_PER_UL,
			},
		},
	}
	if i.ID != nid {
		t.FailNow()
	}
}

func Test_FactCheck(t *testing.T) {
	i := NewIdea("", "something", map[string]*Measurement{})
	if i.FactCheck(CONCEPT_APPLICABLE_HERE.ID) != nil {
		t.Errorf("Expected nil, got %v\n", i.FactCheck(CONCEPT_APPLICABLE_HERE.ID))
	}
}

func Test_Idea_JSON(t *testing.T) {
	i := NewIdea("xyz", "abc", map[string]*Measurement{CONCEPT_DETECTED.ID: {Value: true}})
	j, err := json.Marshal(i)
	if err != nil {
		t.Errorf("Error marshaling idea into json: %v\n", err)
		t.FailNow()
	}
	if len(j) < 10 {
		t.Errorf("this is not a correct json serialization of the Idea %s\n", string(j))
		t.FailNow()
	}

	var iBack Idea
	err = json.Unmarshal(j, &iBack)
	if err != nil {
		t.Errorf("Error unmarshaling idea %v \n", err)
		t.FailNow()
	}

	if i.ID != iBack.ID {
		t.Errorf("expected id to be %s but got %s\n", i.ID, iBack.ID)
	}
	if i.EnglishHumanReadableExpression != iBack.EnglishHumanReadableExpression {
		t.Errorf("expected EnglishHumanReadableExpression to be %s but got %s\n", i.EnglishHumanReadableExpression, iBack.EnglishHumanReadableExpression)
	}
	for idx, f := range i.Facts {
		if f.Value != iBack.Facts[idx].Value || f.ID != iBack.Facts[idx].ID {
			t.Errorf("expected facts to be %s but got %s\n", f, iBack.Facts[idx])
		}
	}
	for idx, f := range iBack.Facts {
		if f.Value != i.Facts[idx].Value || f.ID != iBack.Facts[idx].ID {
			t.Errorf("REVERSE expected facts to be %s but got %s\n", f, i.Facts[idx])
		}
	}
}

func Test_NewConcept(t *testing.T) {
	c := DEFAULT_CONCEPT_REPO.NewConcept("a", "b", "c")
	if DEFAULT_CONCEPT_REPO[c.ID].ID != c.ID || DEFAULT_CONCEPT_REPO[c.ID].EnglishHumanReadableExpression != c.EnglishHumanReadableExpression {
		t.Fatalf("expected c to be in the default repo, but it is not!!!\n %v != %v", DEFAULT_CONCEPT_REPO[c.ID], &c)
		t.FailNow()
	}
}
