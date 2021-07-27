package ideas

import (
	"testing"
	"time"
)

func Test_concept(t *testing.T) {
	nid := newID()
	c := Concept{
		id:                             nid,
		englishHumanReadableExpression: "RBC Count",
		englishDescription:             "Average number of RBCs per µl",
	}
	if c.id != nid {
		t.Fail()
	}
}

func Test_idea(t *testing.T) {
	crbcc := Concept{
		id:                             newID(),
		englishHumanReadableExpression: "RBC Count",
		englishDescription:             "Average number of RBCs per µl",
	}
	p := Concept{
		id:                             newID(),
		englishHumanReadableExpression: "Related to Patient with ID",
		englishDescription:             "to express association of other concepts in the same idea to a particular patient with id",
	}

	nid := newID()
	i := Idea{
		id:                             nid,
		englishHumanReadableExpression: "Patient with id 123 had RBC count of 5.3 µl at time t",
		facts: map[*Concept]Measurment{
			&p: {
				id:        newID(),
				timestamp: 0,
				value:     "PATIENT-123",
				unit:      UNIT_IDENTIFIER,
			},
			&crbcc: {
				id:        newID(),
				timestamp: time.Now().UnixNano(),
				value:     5900000,
				unit:      UNIT_COUNT_PER_UL,
			},
		},
	}
	if i.id != nid {
		t.Fail()
	}
}
