package ideas

import (
	"testing"
	"time"
)

func Test_concept(t *testing.T) {
	nid := newStrID()
	c := Concept{
		id:                             nid,
		englishHumanReadableExpression: "RBC Count",
		englishDescription:             "Average number of RBCs per µl",
	}
	if c.id != nid {
		t.FailNow()
	}
}

func Test_idea(t *testing.T) {
	crbcc := Concept{
		id:                             newStrID(),
		englishHumanReadableExpression: "RBC Count",
		englishDescription:             "Average number of RBCs per µl",
	}
	p := Concept{
		id:                             newStrID(),
		englishHumanReadableExpression: "Related to Patient with ID",
		englishDescription:             "to express association of other concepts in the same idea to a particular patient with id",
	}

	nid := newStrID()
	i := Idea{
		id:                             nid,
		englishHumanReadableExpression: "Patient with id 123 had RBC count of 5.3 µl at time t",
		facts: map[Concept]*Measurement{
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
	if i.id != nid {
		t.FailNow()
	}
}

func Test_FactCheck(t *testing.T) {
	i := NewIdea("something", map[Concept]*Measurement{})
	if i.FactCheck(CONCEPT_APPLICABLE_HERE) != nil {
		t.Errorf("Expected nil, got %v", i.FactCheck(CONCEPT_APPLICABLE_HERE))
	}
}
