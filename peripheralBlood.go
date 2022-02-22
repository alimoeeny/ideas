package ideas

// CONCEPTS

var CONCEPT_CBC = ConceptSet{
	ID:                             "131313",
	EnglishHumanReadableExpression: "Complete Blood Count",
	EnglishDescription:             "Complete Blood Cound",
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
var CONCEPT_RBC_COUNT = Concept{ID: "131314", EnglishHumanReadableExpression: "RBC Count", EnglishDescription: "Average Red Blood Cell Count per volume of whole blood"}

var CONCEPT_NORMAL_RBCCount = Concept{ID: "42000768", EnglishHumanReadableExpression: "Normal RBC count"}
var CONCEPT_HIGH_RBCCount = Concept{ID: "42000769", EnglishHumanReadableExpression: "High RBC count"}
var CONCEPT_LOW_RBCCount = Concept{ID: "42000770", EnglishHumanReadableExpression: "Low RBC count"}
var RBCCountXSet = MutuallyExclusiveConceptSet{ConceptSet{ID: "42000771", concepts: []Concept{CONCEPT_NORMAL_RBCCount, CONCEPT_HIGH_RBCCount, CONCEPT_LOW_RBCCount}}}

var CONCEPT_rbcShapeTargetCell = Concept{ID: "42000772", EnglishHumanReadableExpression: "target cell"}

///----------------------------
var CONCEPT_WBC_COUNT = Concept{ID: "131315", EnglishHumanReadableExpression: "WBC Count", EnglishDescription: "White Blood Cell Count per volume of whole blood"}
var CONCEPT_NORMAL_WBC_COUNT = Concept{ID: "42000873", EnglishHumanReadableExpression: "Normal WBC count"}
var CONCEPT_HIGH_WBC_COUNT = Concept{ID: "42000874", EnglishHumanReadableExpression: "High WBC count"}
var CONCEPT_LOW_WBC_COUNT = Concept{ID: "42000875", EnglishHumanReadableExpression: "Low WBC count"}

///----------------------------
var CONCEPT_HEMOLOBIN_CONCENTRATION = Concept{ID: "131316", EnglishHumanReadableExpression: "Hemoglobin Concentration", EnglishDescription: "Concentration of Hemoglobin in whole blood"}
var CONCEPT_NORMAL_HEMOLOBIN_CONCENTRATION = Concept{ID: "131317", EnglishHumanReadableExpression: "Hemoglobin Concentration", EnglishDescription: "Concentration of Hemoglobin in whole blood"}
var CONCEPT_HIGH_HEMOLOBIN_CONCENTRATION = Concept{ID: "131318", EnglishHumanReadableExpression: "Hemoglobin Concentration", EnglishDescription: "Concentration of Hemoglobin in whole blood"}
var CONCEPT_LOW_HEMOLOBIN_CONCENTRATION = Concept{ID: "131319", EnglishHumanReadableExpression: "Hemoglobin Concentration", EnglishDescription: "Concentration of Hemoglobin in whole blood"}

///----------------------------
var CONCEPT_MVC = Concept{ID: "131320", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_NORMAL_MVC = Concept{ID: "131321", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_HIGH_MVC = Concept{ID: "131322", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_LOW_MVC = Concept{ID: "131323", EnglishHumanReadableExpression: "", EnglishDescription: ""}

// ---------------------------
var CONCEPT_MCH = Concept{ID: "131330", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_NORMAL_MCH = Concept{ID: "131331", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEP_MCH = Concept{ID: "131332", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_LOW_MCH = Concept{ID: "131333", EnglishHumanReadableExpression: "", EnglishDescription: ""}

// ---------------------------
var CONCEPT_MCHC = Concept{ID: "131340", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_NORMAL_MCHC = Concept{ID: "131341", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEP_MCHC = Concept{ID: "131342", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_LOW_MCHC = Concept{ID: "131343", EnglishHumanReadableExpression: "", EnglishDescription: ""}

// ---------------------------
var CONCEPT_PLATELET_COUNT = Concept{ID: "131350", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_NORMAL_PLATELET_COUNT = Concept{ID: "131351", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEP_PLATELET_COUNT = Concept{ID: "131352", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_LOW_PLATELET_COUNT = Concept{ID: "131353", EnglishHumanReadableExpression: "", EnglishDescription: ""}

// ---------------------------
var CONCEPT_NEUTROPHIL_COUNT = Concept{ID: "131360", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_NORMAL_NEUTROPHIL_COUNT = Concept{ID: "131361", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEP_NEUTROPHIL_COUNT = Concept{ID: "131362", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_LOW_NEUTROPHIL_COUNT = Concept{ID: "131363", EnglishHumanReadableExpression: "", EnglishDescription: ""}

// ---------------------------
var CONCEPT_LYMPHOCYTE_COUNT = Concept{ID: "131370", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_NORMAL_LYMPHOCYTE_COUNT = Concept{ID: "131371", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEP_LYMPHOCYTE_COUNT = Concept{ID: "131372", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_LOW_LYMPHOCYTE_COUNT = Concept{ID: "131373", EnglishHumanReadableExpression: "", EnglishDescription: ""}

// ---------------------------
var CONCEPT_MONOCYTE_COUNT = Concept{ID: "131380", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_NORMAL_MONOCYTE_COUNT = Concept{ID: "131381", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEP_MONOCYTE_COUNT = Concept{ID: "131382", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_LOW_MONOCYTE_COUNT = Concept{ID: "131383", EnglishHumanReadableExpression: "", EnglishDescription: ""}

// ---------------------------
var CONCEPT_EOSINOPHIL_COUNT = Concept{ID: "131390", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_NORMAL_EOSINOPHIL_COUNT = Concept{ID: "131391", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEP_EOSINOPHIL_COUNT = Concept{ID: "131392", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_LOW_EOSINOPHIL_COUNT = Concept{ID: "131393", EnglishHumanReadableExpression: "", EnglishDescription: ""}

// ---------------------------
var CONCEPT_BASOPHIL_COUNT = Concept{ID: "131400", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_NORMAL_BASOPHIL_COUNT = Concept{ID: "131401", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEP_BASOPHIL_COUNT = Concept{ID: "131402", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_LOW_BASOPHIL_COUNT = Concept{ID: "131403", EnglishHumanReadableExpression: "", EnglishDescription: ""}

// ---------------------------
var CONCEPT_RETICULOCYTE_COUNT = Concept{ID: "131410", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_NORMAL_RETICULOCYTE_COUNT = Concept{ID: "131411", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEP_RETICULOCYTE_COUNT = Concept{ID: "131412", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_LOW_RETICULOCYTE_COUNT = Concept{ID: "131413", EnglishHumanReadableExpression: "", EnglishDescription: ""}

// ---------------------------

type CBC struct {
	IdeaSet           `json:"idea_set,omitempty"`
	PatientIdentifier Identifier `json:"patient_identifier,omitempty"`
}

var alisCBC = CBC{
	PatientIdentifier: newIdentifier(),
	IdeaSet: NewIdeaSet(
		"1645524588060025000",
		[]*Idea{NewIdea(
			"sqwvmbncwu",
			"RBC Count",
			map[Concept]*Measurement{
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
		)}),
}
