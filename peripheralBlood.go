package ideas

// CONCEPTS
// TEMP:
var tempConcept1 = Concept{ID: "42010768", EnglishHumanReadableExpression: "Microcytic anemia with frequent (10%) spherocytes"}
var tempConcept2 = Concept{ID: "42010769", EnglishHumanReadableExpression: "occasional fragments (2%) and acantocytes."}
var tempConcept3 = Concept{ID: "42010770", EnglishHumanReadableExpression: "Granulocytes are left shifted with neutrophilia and toxic changes."}
var tempConcept4 = Concept{ID: "42010771", EnglishHumanReadableExpression: "No circulating blasts encountered."}
var tempConcept5 = Concept{ID: "42010772", EnglishHumanReadableExpression: "Marked thrombocytopenia}"}

var CONCEPT_PERIPHERAL_BLOOD_FILM_FINDINGS = ConceptSet{
	ID:                             "131312",
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
var CONCEPT_RBC_COUNT = Concept{ID: "131314", EnglishHumanReadableExpression: "RBC Count", EnglishDescription: "Average Red Blood Cell Count per volume of whole blood", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}

var CONCEPT_NORMAL_RBCCount = Concept{ID: "42000768", EnglishHumanReadableExpression: "Normal RBC count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_HIGH_RBCCount = Concept{ID: "42000769", EnglishHumanReadableExpression: "High RBC count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_LOW_RBCCount = Concept{ID: "42000770", EnglishHumanReadableExpression: "Low RBC count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var RBCCountXSet = MutuallyExclusiveConceptSet{ConceptSet{ID: "42000771", Concepts: []Concept{CONCEPT_NORMAL_RBCCount, CONCEPT_HIGH_RBCCount, CONCEPT_LOW_RBCCount}}}

var CONCEPT_rbcShapeTargetCell = Concept{ID: "42000772", EnglishHumanReadableExpression: "target cell", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}

var CONCEPT_SICKLE_CELL_COUNT = Concept{ID: "42000793", EnglishHumanReadableExpression: "Sickle Cell Count", EnglishDescription: "Sickle Cell Count per volume of whole blood", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}

// /----------------------------
var CONCEPT_WBC_COUNT = Concept{ID: "131315", EnglishHumanReadableExpression: "WBC Count", EnglishDescription: "White Blood Cell Count per volume of whole blood", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}
var CONCEPT_NORMAL_WBC_COUNT = Concept{ID: "42000873", EnglishHumanReadableExpression: "Normal WBC count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_HIGH_WBC_COUNT = Concept{ID: "42000874", EnglishHumanReadableExpression: "High WBC count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_LOW_WBC_COUNT = Concept{ID: "42000875", EnglishHumanReadableExpression: "Low WBC count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}

// /----------------------------
var CONCEPT_HEMOLOBIN_CONCENTRATION = Concept{ID: "131316", EnglishHumanReadableExpression: "Hemoglobin Concentration", EnglishDescription: "Concentration of Hemoglobin in whole blood", Short1: "Hb", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}
var CONCEPT_NORMAL_HEMOLOBIN_CONCENTRATION = Concept{ID: "131317", EnglishHumanReadableExpression: "Normal Hemoglobin Concentration", EnglishDescription: "Concentration of Hemoglobin in whole blood", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_HIGH_HEMOLOBIN_CONCENTRATION = Concept{ID: "131318", EnglishHumanReadableExpression: "High Hemoglobin Concentration", EnglishDescription: "Concentration of Hemoglobin in whole blood", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_LOW_HEMOLOBIN_CONCENTRATION = Concept{ID: "131319", EnglishHumanReadableExpression: "Low Hemoglobin Concentration", EnglishDescription: "Concentration of Hemoglobin in whole blood", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}

// /----------------------------
var CONCEPT_MVC = Concept{ID: "131320", EnglishHumanReadableExpression: "Mean corpuscular volume", Short1: "MCV", EnglishDescription: "Mean corpuscular volume", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}
var CONCEPT_NORMAL_MVC = Concept{ID: "131321", EnglishHumanReadableExpression: "Normal mean corpuscular volume", Short1: "Normal MCV", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_HIGH_MVC = Concept{ID: "131322", EnglishHumanReadableExpression: "High mean corpuscular volume", Short1: "High MCV", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_LOW_MVC = Concept{ID: "131323", EnglishHumanReadableExpression: "Low mean corpuscular volume", Short1: "Low MCV", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}

// ---------------------------
var CONCEPT_MCH = Concept{ID: "131330", EnglishHumanReadableExpression: "Mean corpuscular hemoglobin", Short1: "MCH", EnglishDescription: "Mean corpuscular hemoglobin", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}
var CONCEPT_NORMAL_MCH = Concept{ID: "131331", EnglishHumanReadableExpression: "Normal mean corpuscular hemoglobin", Short1: "Normal MCH", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_HIGH_MCH = Concept{ID: "131332", EnglishHumanReadableExpression: "High mean corpuscular hemoglobin", Short1: "High MCH", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_LOW_MCH = Concept{ID: "131333", EnglishHumanReadableExpression: "Low mean corpuscular hemoglobin", Short1: "Low MCH", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}

// ---------------------------
var CONCEPT_MCHC = Concept{ID: "131340", EnglishHumanReadableExpression: "Mean corpuscular hemoglobin concentration", Short1: "MCHC", EnglishDescription: "Mean corpuscular hemoglobin concentration", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}
var CONCEPT_NORMAL_MCHC = Concept{ID: "131341", EnglishHumanReadableExpression: "Normal mean corpuscular hemoglobin concentration", Short1: "Normal MCHC", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_HIGH_MCHC = Concept{ID: "131342", EnglishHumanReadableExpression: "High mean corpuscular hemoglobin concentration", Short1: "High MCHC", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_LOW_MCHC = Concept{ID: "131343", EnglishHumanReadableExpression: "Low mean corpuscular hemoglobin concentration", Short1: "Low MCHC", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}

// ---------------------------
var CONCEPT_PLATELET_COUNT = Concept{ID: "131350", EnglishHumanReadableExpression: "Platelet Count", Short1: "PLT", EnglishDescription: "Platelet Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}
var CONCEPT_NORMAL_PLATELET_COUNT = Concept{ID: "131351", EnglishHumanReadableExpression: "Normal platelet count", Short1: "Normal PLT", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_HIGH_PLATELET_COUNT = Concept{ID: "131352", EnglishHumanReadableExpression: "High platelet count", Short1: "High PLT", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_LOW_PLATELET_COUNT = Concept{ID: "131353", EnglishHumanReadableExpression: "Low platelet count", Short1: "Low PLT", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}

// ---------------------------
var CONCEPT_MEAN_PLATELET_VOLUME = Concept{ID: "13264", EnglishHumanReadableExpression: "Mean Platelet Volume", Short1: "MPV", EnglishDescription: "Mean Platelet Volume", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}
var CONCEPT_NORMAL_MEAN_PLATELET_VOLUME = Concept{ID: "13265", EnglishHumanReadableExpression: "Normal mean platelet volume", Short1: "Normal MPV", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_HIGH_MEAN_PLATELET_VOLUME = Concept{ID: "13266", EnglishHumanReadableExpression: "High mean platelet volume", Short1: "High MPV", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_LOW_MEAN_PLATELET_VOLUME = Concept{ID: "13267", EnglishHumanReadableExpression: "Low mean platelet volume", Short1: "Low MPV", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}

// ---------------------------
var CONCEPT_NEUTROPHIL_COUNT = Concept{ID: "131360", EnglishHumanReadableExpression: "Neutrophil Count", Short1: "NEU", EnglishDescription: "Neutrophil Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}
var CONCEPT_NORMAL_NEUTROPHIL_COUNT = Concept{ID: "131361", EnglishHumanReadableExpression: "Normal neutrophil count", Short1: "Normal NEU", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_HIGH_NEUTROPHIL_COUNT = Concept{ID: "131362", EnglishHumanReadableExpression: "High neutrophil count", Short1: "High NEU", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_LOW_NEUTROPHIL_COUNT = Concept{ID: "131363", EnglishHumanReadableExpression: "Low neutrophil count", Short1: "Low NEU", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_NEUTROPHIL_PERCENTAGE = Concept{ID: "13268", EnglishHumanReadableExpression: "Neutrophil Percentage", Short1: "NEU%", EnglishDescription: "Neutrophil Percentage", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}

// ---------------------------
var CONCEPT_LYMPHOCYTE_COUNT = Concept{ID: "131370", EnglishHumanReadableExpression: "Lymphocyte Count", Short1: "LYM", EnglishDescription: "Lymphocyte Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}
var CONCEPT_NORMAL_LYMPHOCYTE_COUNT = Concept{ID: "131371", EnglishHumanReadableExpression: "Normal lymphocyte count", Short1: "Normal LYM", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_HIGH_LYMPHOCYTE_COUNT = Concept{ID: "131372", EnglishHumanReadableExpression: "High lymphocyte count", Short1: "High LYM", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_LOW_LYMPHOCYTE_COUNT = Concept{ID: "131373", EnglishHumanReadableExpression: "Low lymphocyte count", Short1: "Low LYM", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_LYMPHOCYTE_PERCENTAGE = Concept{ID: "13269", EnglishHumanReadableExpression: "Lymphocyte Percentage", Short1: "LYM%", EnglishDescription: "Lymphocyte Percentage", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}

// ---------------------------
var CONCEPT_MONOCYTE_COUNT = Concept{ID: "131380", EnglishHumanReadableExpression: "Monocyte Count", Short1: "MON", EnglishDescription: "Monocyte Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}
var CONCEPT_NORMAL_MONOCYTE_COUNT = Concept{ID: "131381", EnglishHumanReadableExpression: "Normal monocyte count", Short1: "Normal MON", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_HIGH_MONOCYTE_COUNT = Concept{ID: "131382", EnglishHumanReadableExpression: "High monocyte count", Short1: "High MON", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_LOW_MONOCYTE_COUNT = Concept{ID: "131383", EnglishHumanReadableExpression: "Low monocyte count", Short1: "Low MON", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_MONOCYTE_PERCENTAGE = Concept{ID: "13270", EnglishHumanReadableExpression: "Monocyte Percentage", Short1: "MON%", EnglishDescription: "Monocyte Percentage", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}

// ---------------------------
var CONCEPT_EOSINOPHIL_COUNT = Concept{ID: "131390", EnglishHumanReadableExpression: "Eosinophil Count", Short1: "EOS", EnglishDescription: "Eosinophil Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}
var CONCEPT_NORMAL_EOSINOPHIL_COUNT = Concept{ID: "131391", EnglishHumanReadableExpression: "Normal eosinophil count", Short1: "Normal EOS", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_HIGH_EOSINOPHIL_COUNT = Concept{ID: "131392", EnglishHumanReadableExpression: "High eosinophil count", Short1: "High EOS", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_LOW_EOSINOPHIL_COUNT = Concept{ID: "131393", EnglishHumanReadableExpression: "Low eosinophil count", Short1: "Low EOS", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_EOSINOPHIL_PERCENTAGE = Concept{ID: "13271", EnglishHumanReadableExpression: "Eosinophil Percentage", Short1: "EOS%", EnglishDescription: "Eosinophil Percentage", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}

// ---------------------------
var CONCEPT_BASOPHIL_COUNT = Concept{ID: "131400", EnglishHumanReadableExpression: "Basophil Count", Short1: "BAS", EnglishDescription: "Basophil Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}
var CONCEPT_NORMAL_BASOPHIL_COUNT = Concept{ID: "131401", EnglishHumanReadableExpression: "Normal basophil count", Short1: "Normal BAS", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_HIGH_BASOPHIL_COUNT = Concept{ID: "131402", EnglishHumanReadableExpression: "High basophil count", Short1: "High BAS", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_LOW_BASOPHIL_COUNT = Concept{ID: "131403", EnglishHumanReadableExpression: "Low basophil count", Short1: "Low BAS", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_BASOPHIL_PERCENTAGE = Concept{ID: "13272", EnglishHumanReadableExpression: "Basophil Percentage", Short1: "BAS%", EnglishDescription: "Basophil Percentage", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}

// ---------------------------
var CONCEPT_RETICULOCYTE_COUNT = Concept{ID: "131410", EnglishHumanReadableExpression: "Reticulocyte Count", Short1: "RET", EnglishDescription: "Reticulocyte Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}
var CONCEPT_NORMAL_RETICULOCYTE_COUNT = Concept{ID: "131411", EnglishHumanReadableExpression: "Normal reticulocyte count", Short1: "Normal RET", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_HIGH_RETICULOCYTE_COUNT = Concept{ID: "131412", EnglishHumanReadableExpression: "High reticulocyte count", Short1: "High RET", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_LOW_RETICULOCYTE_COUNT = Concept{ID: "131413", EnglishHumanReadableExpression: "Low reticulocyte count", Short1: "Low RET", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_RETICULOCYTE_PERCENTAGE = Concept{ID: "13273", EnglishHumanReadableExpression: "Reticulocyte Percentage", Short1: "RET%", EnglishDescription: "Reticulocyte Percentage", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}}

// ---------------------------

var CONCEPT_MALARIA_PARASITE_PRESENT = Concept{ID: "131420", EnglishHumanReadableExpression: "Malaria Parasite Detected", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_MALARIA_ASEXUAL_PARASITE_PRESENT = Concept{ID: "131421", EnglishHumanReadableExpression: "Asexual form of Malaria Parasite Detected", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_MALARIA_GAMETOCYTES_PRESENT = Concept{ID: "131422", EnglishHumanReadableExpression: "Malaria Gametocytes Detected", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_MALARIA_FALCIPARUM_PARASITE_PRESENT = Concept{ID: "131423", EnglishHumanReadableExpression: "", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_MALARIA_VIVAX_PARASITE_PRESENT = Concept{ID: "131424", EnglishHumanReadableExpression: "", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_MALARIA_OVALE_PARASITE_PRESENT = Concept{ID: "131425", EnglishHumanReadableExpression: "", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_MALARIA_MALARIAE_PARASITE_PRESENT = Concept{ID: "131426", EnglishHumanReadableExpression: "", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_MALARIA_KNOWLESI_PARASITE_PRESENT = Concept{ID: "131427", EnglishHumanReadableExpression: "", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_MALARIA_PARASITE_CLASSIFICATION = Concept{ID: "131428", EnglishHumanReadableExpression: "", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}

// ---------------------------
// ---------------------------

var CONCEPT_NORMAL_RBC_FINDINGS = Concept{ID: "13211", EnglishHumanReadableExpression: "RBCs Are Normal", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_NORMAL_WBC_FINDINGS = Concept{ID: "13212", EnglishHumanReadableExpression: "WBCs Are Normal", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_NORMAL_PLATELET_FINDINGS = Concept{ID: "13213", EnglishHumanReadableExpression: "Platelets Are Normal", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}

var CONCEPT_NEUTROPHILIA_ABSOLUTE = Concept{ID: "13222", EnglishHumanReadableExpression: "Absolute Neutrophilia", EnglishDescription: "Increased Neutrophiol Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_NEUTROPENIA_ABSOLUTE = Concept{ID: "13223", EnglishHumanReadableExpression: "Absolute Neutropenia", EnglishDescription: "Decreased Neutrophiol Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_LYMPHOCYTOSIS_ABSOLUTE = Concept{ID: "13224", EnglishHumanReadableExpression: "Absolute Lymphocytosis", EnglishDescription: "Increased Lymphocyte Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_LYMPHOPENIA_ABSOLUTE = Concept{ID: "13225", EnglishHumanReadableExpression: "Absolute Lymphopenia", EnglishDescription: "Decreased Lymphocyte Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_MONOCYTOSIS_ABSOLUTE = Concept{ID: "13226", EnglishHumanReadableExpression: "Absolute Monocytosis", EnglishDescription: "Increased Monocyte Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_MONOCYTOPENIA_ABSOLUTE = Concept{ID: "13227", EnglishHumanReadableExpression: "Absolute Monocytopenia", EnglishDescription: "Decreased Monocyte Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_EOSINOPHILIA_ABSOLUTE = Concept{ID: "13228", EnglishHumanReadableExpression: "Absolute Eosinophilia", EnglishDescription: "Increased Eosinophil Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_EOSINOPHILOPENIA_ABSOLUTE = Concept{ID: "13229", EnglishHumanReadableExpression: "Absolute Eosinopenia", EnglishDescription: "Decreased Eosinophil Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_BASOPHILIA_ABSOLUTE = Concept{ID: "13230", EnglishHumanReadableExpression: "Absolute Basophilia", EnglishDescription: "Increased Basophil Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_BASOPHILOPENIA_ABSOLUTE = Concept{ID: "13231", EnglishHumanReadableExpression: "Absolute Basophilopenia", EnglishDescription: "Decreased Basophil Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}

var CONCEPT_NEUTROPHILIA_RELATIVE = Concept{ID: "13252", EnglishHumanReadableExpression: "Relative Neutrophilia", EnglishDescription: "Increased Neutrophiol Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_NEUTROPENIA_RELATIVE = Concept{ID: "13253", EnglishHumanReadableExpression: "Relative Neutropenia", EnglishDescription: "Decreased Neutrophiol Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_LYMPHOCYTOSIS_RELATIVE = Concept{ID: "13254", EnglishHumanReadableExpression: "Relative Lymphocytosis", EnglishDescription: "Increased Lymphocyte Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_LYMPHOPENIA_RELATIVE = Concept{ID: "13255", EnglishHumanReadableExpression: "Relative Lymphopenia", EnglishDescription: "Decreased Lymphocyte Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_MONOCYTOSIS_RELATIVE = Concept{ID: "13256", EnglishHumanReadableExpression: "Relative Monocytosis", EnglishDescription: "Increased Monocyte Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_MONOCYTOPENIA_RELATIVE = Concept{ID: "13257", EnglishHumanReadableExpression: "Relative Monocytopenia", EnglishDescription: "Decreased Monocyte Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_EOSINOPHILIA_RELATIVE = Concept{ID: "13258", EnglishHumanReadableExpression: "Relative Eosinophilia", EnglishDescription: "Increased Eosinophil Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_EOSINOPHILOPENIA_RELATIVE = Concept{ID: "13259", EnglishHumanReadableExpression: "Relative Eosinopenia", EnglishDescription: "Decreased Eosinophil Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_BASOPHILIA_RELATIVE = Concept{ID: "13260", EnglishHumanReadableExpression: "Relative Basophilia", EnglishDescription: "Increased Basophil Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}
var CONCEPT_BASOPHILOPENIA_RELATIVE = Concept{ID: "13261", EnglishHumanReadableExpression: "Relative Basophilopenia", EnglishDescription: "Decreased Basophil Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}

var CONCEPT_RETICULOCYTOSIS = Concept{ID: "13262", EnglishHumanReadableExpression: "Reticulocytosis", EnglishDescription: "Increased Reticulocyte Count", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}

var CONCEPT_BLAST_CELL_DETECTED = Concept{ID: "13263", EnglishHumanReadableExpression: "Blast Cell Detected", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}}

// ----------------------------------------
// ---------------------------------------- Comments
var CONCEPT_CBC_COMMENT_REACTICE_CHANGES = Concept{
	ID:                             "135691",
	EnglishHumanReadableExpression: "The Neutrophilia with reactive changes is suggestive of systemic illness, or infection. Clinical correlation is recommended.",
	EnglishDescription:             "The Neutrophilia with reactive changes is suggestive of systemic illness, or infection. Clinical correlation is recommended.",
	Short1:                         "Reactive Changes",
	Categories:                     []Concept{DIAGNOSTIC_FLOW_OUTPUT},
}

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
