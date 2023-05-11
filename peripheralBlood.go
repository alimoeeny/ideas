package ideas

// CONCEPTS
// TEMP:
var tempConcept1 = Concept{ID: "42010768", EnglishHumanReadableExpression: "Microcytic anemia with frequent (10%) spherocytes"}
var tempConcept2 = Concept{ID: "42010769", EnglishHumanReadableExpression: "occasional fragments (2%) and acantocytes."}
var tempConcept3 = Concept{ID: "42010770", EnglishHumanReadableExpression: "Granulocytes are left shifted with neutrophilia and toxic changes."}
var tempConcept4 = Concept{ID: "42010771", EnglishHumanReadableExpression: "No circulating blasts encountered."}
var tempConcept5 = Concept{ID: "42010772", EnglishHumanReadableExpression: "Marked thrombocytopenia}"}

var CONCEPT_PERIPHERAL_BLOOD_FILM_FINDINGS = ConceptSet{
	ID:                             "131314",
	EnglishHumanReadableExpression: "Blood Film Findings",
	EnglishDescription:             "Peripheral Blood Film Description",
	Concepts: []Concept{
		// TODO: can we have an exhastive list?
		tempConcept1, tempConcept2, tempConcept3, tempConcept4, tempConcept5,
	},
}
var CONCEPT_CBC = ConceptSet{
	ID:                             "131313",
	EnglishHumanReadableExpression: "Complete Blood Count",
	EnglishDescription:             "Complete Blood Cound",
	Concepts: []Concept{
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

// /----------------------------
var CONCEPT_RBC_COUNT = Concept{ID: "131314", EnglishHumanReadableExpression: "RBC Count", EnglishDescription: "Average Red Blood Cell Count per volume of whole blood"}

var CONCEPT_NORMAL_RBCCount = Concept{ID: "42000768", EnglishHumanReadableExpression: "Normal RBC count"}
var CONCEPT_HIGH_RBCCount = Concept{ID: "42000769", EnglishHumanReadableExpression: "High RBC count"}
var CONCEPT_LOW_RBCCount = Concept{ID: "42000770", EnglishHumanReadableExpression: "Low RBC count"}
var RBCCountXSet = MutuallyExclusiveConceptSet{ConceptSet{ID: "42000771", Concepts: []Concept{CONCEPT_NORMAL_RBCCount, CONCEPT_HIGH_RBCCount, CONCEPT_LOW_RBCCount}}}

var CONCEPT_rbcShapeTargetCell = Concept{ID: "42000772", EnglishHumanReadableExpression: "target cell"}

var CONCEPT_SICKLE_CELL_COUNT = Concept{ID: "42000793", EnglishHumanReadableExpression: "Sickle Cell Count", EnglishDescription: "Sickle Cell Count per volume of whole blood"}

// /----------------------------
var CONCEPT_WBC_COUNT = Concept{ID: "131315", EnglishHumanReadableExpression: "WBC Count", EnglishDescription: "White Blood Cell Count per volume of whole blood"}
var CONCEPT_NORMAL_WBC_COUNT = Concept{ID: "42000873", EnglishHumanReadableExpression: "Normal WBC count"}
var CONCEPT_HIGH_WBC_COUNT = Concept{ID: "42000874", EnglishHumanReadableExpression: "High WBC count"}
var CONCEPT_LOW_WBC_COUNT = Concept{ID: "42000875", EnglishHumanReadableExpression: "Low WBC count"}

// /----------------------------
var CONCEPT_HEMOLOBIN_CONCENTRATION = Concept{ID: "SENS-Hemoglobin-concentration", EnglishHumanReadableExpression: "Hemoglobin Concentration", EnglishDescription: "Concentration of Hemoglobin in whole blood", Short1: "Hb"}
var CONCEPT_NORMAL_HEMOLOBIN_CONCENTRATION = Concept{ID: "131317", EnglishHumanReadableExpression: "Hemoglobin Concentration", EnglishDescription: "Concentration of Hemoglobin in whole blood"}
var CONCEPT_HIGH_HEMOLOBIN_CONCENTRATION = Concept{ID: "131318", EnglishHumanReadableExpression: "Hemoglobin Concentration", EnglishDescription: "Concentration of Hemoglobin in whole blood"}
var CONCEPT_LOW_HEMOLOBIN_CONCENTRATION = Concept{ID: "131319", EnglishHumanReadableExpression: "Hemoglobin Concentration", EnglishDescription: "Concentration of Hemoglobin in whole blood"}

// /----------------------------
var CONCEPT_MVC = Concept{ID: "131320", EnglishHumanReadableExpression: "Mean corpuscular volume", Short1: "MCV", EnglishDescription: "Mean corpuscular volume"}
var CONCEPT_NORMAL_MVC = Concept{ID: "131321", EnglishHumanReadableExpression: "Normal mean corpuscular volume", Short1: "Normal MCV", EnglishDescription: ""}
var CONCEPT_HIGH_MVC = Concept{ID: "131322", EnglishHumanReadableExpression: "High mean corpuscular volume", Short1: "High MCV", EnglishDescription: ""}
var CONCEPT_LOW_MVC = Concept{ID: "131323", EnglishHumanReadableExpression: "Low mean corpuscular volume", Short1: "Low MCV", EnglishDescription: ""}

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

var CONCEPT_MALARIA_PARASITE_PRESENT = Concept{ID: "131420", EnglishHumanReadableExpression: "Malaria Parasite Detected", EnglishDescription: ""}
var CONCEPT_MALARIA_ASEXUAL_PARASITE_PRESENT = Concept{ID: "131421", EnglishHumanReadableExpression: "Asexual form of Malaria Parasite Detected", EnglishDescription: ""}
var CONCEPT_MALARIA_GAMETOCYTES_PRESENT = Concept{ID: "131422", EnglishHumanReadableExpression: "Malaria Gametocytes Detected", EnglishDescription: ""}
var CONCEPT_MALARIA_FALCIPARUM_PARASITE_PRESENT = Concept{ID: "131423", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_MALARIA_VIVAX_PARASITE_PRESENT = Concept{ID: "131424", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_MALARIA_OVALE_PARASITE_PRESENT = Concept{ID: "131425", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_MALARIA_MALARIAE_PARASITE_PRESENT = Concept{ID: "131426", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_MALARIA_KNOWLESI_PARASITE_PRESENT = Concept{ID: "131427", EnglishHumanReadableExpression: "", EnglishDescription: ""}
var CONCEPT_MALARIA_PARASITE_CLASSIFICATION = Concept{ID: "131428", EnglishHumanReadableExpression: "", EnglishDescription: ""}

// ---------------------------

var CONCEPT_BLAST_CELL_DETECTED = Concept{ID: "131480", EnglishHumanReadableExpression: "Blast Cell Detected", EnglishDescription: ""}

type PeripheralBloodFilmFindings struct {
	IdeaSet           `json:"idea_set,omitempty"`
	PatientIdentifier Identifier `json:"patient_identifier,omitempty"`
}

type CBC struct {
	IdeaSet           `json:"idea_set,omitempty"`
	PatientIdentifier Identifier `json:"patient_identifier,omitempty"`
}

var CBCIdeaSet = NewIdeaSet(
	"1645524588060025001",
	[]*Idea{NewIdea(
		"sqwvmbncwu",
		"RBC Count",
		map[string]*Measurement{
			CONCEPT_RBC_COUNT.ID:               nil,
			CONCEPT_WBC_COUNT.ID:               nil,
			CONCEPT_HEMOLOBIN_CONCENTRATION.ID: nil,
			CONCEPT_MVC.ID:                     nil,
			CONCEPT_MCH.ID:                     nil,
			CONCEPT_MCHC.ID:                    nil,
			CONCEPT_NEUTROPHIL_COUNT.ID:        nil,
			CONCEPT_LYMPHOCYTE_COUNT.ID:        nil,
			CONCEPT_MONOCYTE_COUNT.ID:          nil,
			CONCEPT_EOSINOPHIL_COUNT.ID:        nil,
			CONCEPT_EOSINOPHIL_COUNT.ID:        nil,
			CONCEPT_RETICULOCYTE_COUNT.ID:      nil,
		},
	)})

var alisCBC = CBC{
	PatientIdentifier: newIdentifier(),
	IdeaSet: NewIdeaSet(
		"1645524588060025000",
		[]*Idea{NewIdea(
			"sqwvmbncwu",
			"RBC Count",
			map[string]*Measurement{
				CONCEPT_RBC_COUNT.ID: {
					ID:        "fvvzwarsyx",
					Timestamp: 1645525355974859000,
					Value:     5,
					Unit:      UNIT_COUNT_PER_UL,
				},
				CONCEPT_WBC_COUNT.ID: {
					ID:        "zrblcflizl",
					Timestamp: 1645538338859452000,
					Value:     8569,
					Unit:      UNIT_COUNT_PER_UL,
				},
			},
		)}),
}
