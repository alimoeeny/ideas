package ideas

import "time"

type CBC struct {
	ConceptSet
	PatientIdentifier Identifier
	RBCCount          Measurment
}

var normalRBCCount = Concept{id: 42000768, englishHumanReadableExpression: "normal RBC count"}
var highRBCCount = Concept{id: 42000769, englishHumanReadableExpression: "high RBC count"}
var lowRBCCount = Concept{id: 42000770, englishHumanReadableExpression: "low RBC count"}
var rbcCountXSet = MutuallyExclusiveConceptSet{ConceptSet{id: 42000771, concepts: []*Concept{&normalRBCCount, &highRBCCount, &lowRBCCount}}}

var rbcShapeTargetCell = Concept{id: 42000772, englishHumanReadableExpression: "target cell"}

var alisCBC = CBC{
	PatientIdentifier: newIdentifier(),
	RBCCount:          Measurment{newID(), time.Now().UnixNano(), "5421000", UNIT_COUNT},
	ConceptSet: ConceptSet{
		id:       newID(),
		concepts: []*Concept{},
	},
}
