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
		Facts: map[Concept]*Measurement{
			p: {
				ID:        newStrID(),
				Timestamp: 0,
				Value:     "PATIENT-123",
				Unit:      UNIT_IDENTIFIER,
			},
			crbcc: {
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
	i := NewIdea("", "something", map[Concept]*Measurement{})
	if i.FactCheck(CONCEPT_APPLICABLE_HERE) != nil {
		t.Errorf("Expected nil, got %v", i.FactCheck(CONCEPT_APPLICABLE_HERE))
	}
}

func Test_Idea_JSON(t *testing.T) {
	i := NewIdea("xyz", "abc", map[Concept]*Measurement{CONCEPT_DETECTED: &Measurement{Value: true}})
	j, err := json.Marshal(i)
	if err != nil {
		t.Errorf("Error marshaling idea into json: %v", err)
		t.FailNow()
	}
	if len(j) < 10 {
		t.Errorf("this is not a correct json serialization of the Idea %s", string(j))
		t.FailNow()
	}
}
