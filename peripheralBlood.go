package ideas

// CONCEPTS

var CONCEPT_CBC = ConceptSet{
	ID:                             "131313",
	englishHumanReadableExpression: "Complete Blood Count",
	englishDescription:             "Complete Blood Cound",
	concepts: []Concept{
		CONCEPT_RBC_COUNT,
		CONCEPT_WBC_COUNT,
		CONCEPT_HEMOLOBIN_CONCENTRATION,
		CONCEPT_MVC,
		CONCEPT_MCH,
		CONCEPT_MCHC,
		CONCEPT_PLATELET_COUNT,
		CONCEPT_NEUTROPHIL_COUNT,
		CONCEPT_LYMPHOCYTE_COUNT,
		CONCEPT_MONOCYTE_COUNT,
		CONCEPT_EOSINOPHIL_COUNT,
		CONCEPT_BASOPHIL_COUNT,
		CONCEPT_RETICULOCYTE_COUNT,
	},
}

///----------------------------
var CONCEPT_RBC_COUNT = Concept{id: "131314", englishHumanReadableExpression: "RBC Count", englishDescription: "Average Red Blood Cell Count per volume of whole blood"}

var CONCEPT_NORMAL_RBCCount = Concept{id: "42000768", englishHumanReadableExpression: "Normal RBC count"}
var CONCEPT_HIGH_RBCCount = Concept{id: "42000769", englishHumanReadableExpression: "High RBC count"}
var CONCEPT_LOW_RBCCount = Concept{id: "42000770", englishHumanReadableExpression: "Low RBC count"}
var RBCCountXSet = MutuallyExclusiveConceptSet{ConceptSet{ID: "42000771", concepts: []Concept{CONCEPT_NORMAL_RBCCount, CONCEPT_HIGH_RBCCount, CONCEPT_LOW_RBCCount}}}

var CONCEPT_rbcShapeTargetCell = Concept{id: "42000772", englishHumanReadableExpression: "target cell"}

///----------------------------
var CONCEPT_WBC_COUNT = Concept{id: "131315", englishHumanReadableExpression: "WBC Count", englishDescription: "White Blood Cell Count per volume of whole blood"}
var CONCEPT_NORMAL_WBC_COUNT = Concept{id: "42000873", englishHumanReadableExpression: "Normal WBC count"}
var CONCEPT_HIGH_WBC_COUNT = Concept{id: "42000874", englishHumanReadableExpression: "High WBC count"}
var CONCEPT_LOW_WBC_COUNT = Concept{id: "42000875", englishHumanReadableExpression: "Low WBC count"}

///----------------------------
var CONCEPT_HEMOLOBIN_CONCENTRATION = Concept{id: "131316", englishHumanReadableExpression: "Hemoglobin Concentration", englishDescription: "Concentration of Hemoglobin in whole blood"}
var CONCEPT_NORMAL_HEMOLOBIN_CONCENTRATION = Concept{id: "131317", englishHumanReadableExpression: "Hemoglobin Concentration", englishDescription: "Concentration of Hemoglobin in whole blood"}
var CONCEPT_HIGH_HEMOLOBIN_CONCENTRATION = Concept{id: "131318", englishHumanReadableExpression: "Hemoglobin Concentration", englishDescription: "Concentration of Hemoglobin in whole blood"}
var CONCEPT_LOW_HEMOLOBIN_CONCENTRATION = Concept{id: "131319", englishHumanReadableExpression: "Hemoglobin Concentration", englishDescription: "Concentration of Hemoglobin in whole blood"}

///----------------------------
var CONCEPT_MVC = Concept{id: "131320", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEPT_NORMAL_MVC = Concept{id: "131321", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEPT_HIGH_MVC = Concept{id: "131322", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEPT_LOW_MVC = Concept{id: "131323", englishHumanReadableExpression: "", englishDescription: ""}

// ---------------------------
var CONCEPT_MCH = Concept{id: "131330", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEPT_NORMAL_MCH = Concept{id: "131331", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEP_MCH = Concept{id: "131332", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEPT_LOW_MCH = Concept{id: "131333", englishHumanReadableExpression: "", englishDescription: ""}

// ---------------------------
var CONCEPT_MCHC = Concept{id: "131340", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEPT_NORMAL_MCHC = Concept{id: "131341", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEP_MCHC = Concept{id: "131342", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEPT_LOW_MCHC = Concept{id: "131343", englishHumanReadableExpression: "", englishDescription: ""}

// ---------------------------
var CONCEPT_PLATELET_COUNT = Concept{id: "131350", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEPT_NORMAL_PLATELET_COUNT = Concept{id: "131351", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEP_PLATELET_COUNT = Concept{id: "131352", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEPT_LOW_PLATELET_COUNT = Concept{id: "131353", englishHumanReadableExpression: "", englishDescription: ""}

// ---------------------------
var CONCEPT_NEUTROPHIL_COUNT = Concept{id: "131360", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEPT_NORMAL_NEUTROPHIL_COUNT = Concept{id: "131361", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEP_NEUTROPHIL_COUNT = Concept{id: "131362", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEPT_LOW_NEUTROPHIL_COUNT = Concept{id: "131363", englishHumanReadableExpression: "", englishDescription: ""}

// ---------------------------
var CONCEPT_LYMPHOCYTE_COUNT = Concept{id: "131370", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEPT_NORMAL_LYMPHOCYTE_COUNT = Concept{id: "131371", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEP_LYMPHOCYTE_COUNT = Concept{id: "131372", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEPT_LOW_LYMPHOCYTE_COUNT = Concept{id: "131373", englishHumanReadableExpression: "", englishDescription: ""}

// ---------------------------
var CONCEPT_MONOCYTE_COUNT = Concept{id: "131380", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEPT_NORMAL_MONOCYTE_COUNT = Concept{id: "131381", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEP_MONOCYTE_COUNT = Concept{id: "131382", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEPT_LOW_MONOCYTE_COUNT = Concept{id: "131383", englishHumanReadableExpression: "", englishDescription: ""}

// ---------------------------
var CONCEPT_EOSINOPHIL_COUNT = Concept{id: "131390", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEPT_NORMAL_EOSINOPHIL_COUNT = Concept{id: "131391", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEP_EOSINOPHIL_COUNT = Concept{id: "131392", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEPT_LOW_EOSINOPHIL_COUNT = Concept{id: "131393", englishHumanReadableExpression: "", englishDescription: ""}

// ---------------------------
var CONCEPT_BASOPHIL_COUNT = Concept{id: "131400", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEPT_NORMAL_BASOPHIL_COUNT = Concept{id: "131401", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEP_BASOPHIL_COUNT = Concept{id: "131402", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEPT_LOW_BASOPHIL_COUNT = Concept{id: "131403", englishHumanReadableExpression: "", englishDescription: ""}

// ---------------------------
var CONCEPT_RETICULOCYTE_COUNT = Concept{id: "131410", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEPT_NORMAL_RETICULOCYTE_COUNT = Concept{id: "131411", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEP_RETICULOCYTE_COUNT = Concept{id: "131412", englishHumanReadableExpression: "", englishDescription: ""}
var CONCEPT_LOW_RETICULOCYTE_COUNT = Concept{id: "131413", englishHumanReadableExpression: "", englishDescription: ""}

// ---------------------------

type CBC struct {
	IdeaSet
	PatientIdentifier Identifier
}

var alisCBC = CBC{
	PatientIdentifier: newIdentifier(),
	IdeaSet: IdeaSet{
		id: 1645524588060025000,
		ideas: []*Idea{{
			id:                             "sqwvmbncwu",
			englishHumanReadableExpression: "RBC Count",
			facts: map[Concept]*Measurement{
				CONCEPT_RBC_COUNT: {
					ID:        "fvvzwarsyx",
					Timestamp: 1645525355974859000,
					Value:     5,
					Unit:      UNIT_COUNT_PER_UL,
				},
				CONCEPT_WBC_COUNT: {
					ID:        "zrblcflizl",
					Timestamp: 1645538338859452000,
					Value:     8569,
					Unit:      UNIT_COUNT_PER_UL,
				},
			},
		}},
	},
}
