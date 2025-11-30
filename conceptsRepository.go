package ideas

var DEFAULT_CONCEPT_REPO = ConceptsRepository{
	ID:                 "vczncfpjwr",
	conceptsDict:       map[string]*Concept{},
	unitsRepo:          &DEFAULT_UNIT_REPO,
	conceptsToUnitsMap: map[string][]string{},
}

func init() {
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_NOT_DETECTED)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_DETECTED)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_DATA_NOT_AVAILABLE)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_DATA_NEEDED)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_APPLICABLE_HERE)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_NOT_APPLICABLE_HERE)

	// Diagnostic Flow
	DEFAULT_CONCEPT_REPO.SetConcept(DIAGNOSTIC_FLOW_INPUT)
	DEFAULT_CONCEPT_REPO.SetConcept(DIAGNOSTIC_FLOW_OUTPUT)
	//----------
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_WBC_COUNT)

	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_MVC)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_MCH)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_MCHC)

	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_PLATELET_COUNT)

	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_NEUTROPHIL_COUNT)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_LYMPHOCYTE_COUNT)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_MONOCYTE_COUNT)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_EOSINOPHIL_COUNT)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_BASOPHIL_COUNT)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_RETICULOCYTE_COUNT)

	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_NORMAL_RBC_FINDINGS)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_NORMAL_WBC_FINDINGS)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_NORMAL_PLATELET_FINDINGS)

	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_NEUTROPHILIA_ABSOLUTE)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_NEUTROPENIA_ABSOLUTE)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_LYMPHOCYTOSIS_ABSOLUTE)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_LYMPHOPENIA_ABSOLUTE)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_MONOCYTOSIS_ABSOLUTE)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_MONOCYTOPENIA_ABSOLUTE)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_EOSINOPHILIA_ABSOLUTE)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_EOSINOPHILOPENIA_ABSOLUTE)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_BASOPHILIA_ABSOLUTE)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_BASOPHILOPENIA_ABSOLUTE)

	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_NEUTROPHILIA_RELATIVE)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_NEUTROPENIA_RELATIVE)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_LYMPHOCYTOSIS_RELATIVE)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_LYMPHOPENIA_RELATIVE)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_MONOCYTOSIS_RELATIVE)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_MONOCYTOPENIA_RELATIVE)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_EOSINOPHILIA_RELATIVE)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_EOSINOPHILOPENIA_RELATIVE)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_BASOPHILIA_RELATIVE)
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_BASOPHILOPENIA_RELATIVE)

	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_RETICULOCYTOSIS)

	//------------
	// DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_CBC)
	// DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_PERIPHERAL_BLOOD_FILM_FINDINGS)
	//------------
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_WEATHER_IS_RAINING)
	// TEMP
	DEFAULT_CONCEPT_REPO.SetConcept(tempConcept1)
	DEFAULT_CONCEPT_REPO.SetConcept(tempConcept2)
	DEFAULT_CONCEPT_REPO.SetConcept(tempConcept3)
	DEFAULT_CONCEPT_REPO.SetConcept(tempConcept4)
	DEFAULT_CONCEPT_REPO.SetConcept(tempConcept5)

	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_RBC_COUNT)
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID(CONCEPT_RBC_COUNT.ID, UNIT_COUNT.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_SICKLE_CELL_COUNT)
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID(CONCEPT_SICKLE_CELL_COUNT.ID, UNIT_COUNT.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_MALARIA_PARASITE_PRESENT)
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID(CONCEPT_MALARIA_PARASITE_PRESENT.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_MALARIA_ASEXUAL_PARASITE_PRESENT)
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID(CONCEPT_MALARIA_ASEXUAL_PARASITE_PRESENT.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_MALARIA_GAMETOCYTES_PRESENT)
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID(CONCEPT_MALARIA_GAMETOCYTES_PRESENT.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_MALARIA_FALCIPARUM_PARASITE_PRESENT)
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID(CONCEPT_MALARIA_FALCIPARUM_PARASITE_PRESENT.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_MALARIA_VIVAX_PARASITE_PRESENT)
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID(CONCEPT_MALARIA_VIVAX_PARASITE_PRESENT.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_MALARIA_OVALE_PARASITE_PRESENT)
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID(CONCEPT_MALARIA_OVALE_PARASITE_PRESENT.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_MALARIA_MALARIAE_PARASITE_PRESENT)
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID(CONCEPT_MALARIA_MALARIAE_PARASITE_PRESENT.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_MALARIA_KNOWLESI_PARASITE_PRESENT)
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID(CONCEPT_MALARIA_KNOWLESI_PARASITE_PRESENT.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_MALARIA_PARASITE_CLASSIFICATION)
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID(CONCEPT_MALARIA_PARASITE_CLASSIFICATION.ID, UNIT_CATEGORICAL.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_BLAST_CELL_DETECTED)
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID(CONCEPT_BLAST_CELL_DETECTED.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19359", EnglishHumanReadableExpression: "Spherocyte over RBC count", Short1: "Spherocyte-over-RBC-count, Spherocyte", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19359", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19360", EnglishHumanReadableExpression: "Acanthocyte-over-RBC-count", Short1: "Acanthocyte-over-RBC-count, Acanthocyte", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19359", UNIT_COUNT.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19361", EnglishHumanReadableExpression: "Fragment-over-RBC-count", Short1: "Fragment-over-RBC-count, Fragment", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19361", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19363", EnglishHumanReadableExpression: "Reticulocyte-over-RBC-count", Short1: "Reticulocyte-over-RBC-count, Retic", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19363", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_HEMOLOBIN_CONCENTRATION)
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("131320")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19364", EnglishHumanReadableExpression: "Hemoglobin-delta-7-days", Short1: "SENS-Hemoglobin-delta-7-days, Hb ∆ 7 days", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19364")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19365", EnglishHumanReadableExpression: "Haptoglobin-concentration", Short1: "SENS-Haptoglobin-concentration, Haptoglobin", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19365")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19366", EnglishHumanReadableExpression: "Direct-Antiglobulin-Testing", Short1: "SENS-Direct-Antiglobulin-Testing, DAT", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19366")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19367", EnglishHumanReadableExpression: "Bilirubin serum concentration", Short1: "Bilirubin-serum-concentration, Bilirubin", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19367")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19369", EnglishHumanReadableExpression: "Sickle-Cell-over-RBC-count", Short1: "Sickle-Cell-over-RBC-count, Sickle cell", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19369")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19372", EnglishHumanReadableExpression: "Hypogranular-Neutrophils-over-WBC-count", Short1: "Hypogranular-Neutrophils-over-WBC-count, Hypogranular Neutrophils", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19372")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19373", EnglishHumanReadableExpression: "Hypersegmented-Neurophil-over-WBC-count", Short1: "Hypersegmented-Neurophil-over-WBC-count, Hypersegmented Neurophil", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19373")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19374", EnglishHumanReadableExpression: "Hypogranular-Platelet-over-RBC-count", Short1: "Hypogranular-Platelet-over-RBC-count, Hypogranular Platelet", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19374")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19375", EnglishHumanReadableExpression: "Most-Recent-date-time-of-diagnosis-MDS", Short1: "Most-Recent-date-time-of-diagnosis-MDS, History of MDS", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19375", UNIT_TIMESTAMP.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19376", EnglishHumanReadableExpression: "NRBC-count-over-RBC-count", Short1: "NRBC-count-over-RBC-count, NRBC", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19376")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19377", EnglishHumanReadableExpression: "Rouleaux-Formation-count-over-RBC-count", Short1: "Rouleaux-Formation-count-over-RBC-count, Rouleaux", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19377", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19378", EnglishHumanReadableExpression: "Metamyelocyte-over-WBC-count", Short1: "SENS-Metamyelocyte-over-WBC-count, Metamyelocyte", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19378")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19379", EnglishHumanReadableExpression: "Promyelocyte-count-over-WBC-count", Short1: "SENS-Promyelocyte-count-over-WBC-count, Promyelocyte", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19379")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19380", EnglishHumanReadableExpression: "Myelocyte-over-WBC-count", Short1: "SENS-Myelocyte-count-over-WBC-count, Myelocyte", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19380")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19381", EnglishHumanReadableExpression: "Teardrop-RBC-over-RBC-count", Short1: "SENS-Teardrop-RBC-count-over-RBC-count, Teardrop RBC", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19381")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19382", EnglishHumanReadableExpression: "Most-Recent-date-time-of-diagnosis-PMN", Short1: "SENS-Most-Recent-date-time-of-diagnosis-PMN, History of PMN", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19382", UNIT_TIMESTAMP.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19383", EnglishHumanReadableExpression: "Monocyte-count-over-RBC-count", Short1: "SENS-Monocyte-count-over-RBC-count, Monocyte", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19383")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19384", EnglishHumanReadableExpression: "Most-Recent-date-time-of-diagnosis-High-Monocytes", Short1: "SENS-Most-Recent-date-time-of-diagnosis-High-Monocytes, History of High Monocytes", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19384", UNIT_TIMESTAMP.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19385", EnglishHumanReadableExpression: "Blast count over WBC count", Short1: "SENS-Blast-count-over-WBC-count, Blast / WBC", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19385", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19386", EnglishHumanReadableExpression: "Hypochromic RBC over RBC count", Short1: "SENS-Hypochromic-RBC-over-RBC-count, Hypochromic RBC", EnglishDescription: "Hypochromic RBC over RBC count"})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19386", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19387", EnglishHumanReadableExpression: "Teardrop RBC over RBC count", Short1: "SENS-Teardrop-RBC-over-RBC-count, Teardrop-RBC-over-RBC-count", EnglishDescription: "Teardrop-RBC-over-RBC-count"})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19387", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19388", EnglishHumanReadableExpression: "Blasts with vacuoles-count-over-WBC-count", Short1: "SENS-Blasts-with-vacuoles-count-over-WBC-count, Blasts with vacuoles / WBC", EnglishDescription: "SENS-Blasts with vacuoles-count-over-WBC-count"})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19388", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19389", EnglishHumanReadableExpression: "", Short1: "SENS-Basophilic-Metamyelocytes-count-over-WBC-count", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Basophilic Metamyelocytes-count-over-WBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19390", EnglishHumanReadableExpression: "Myeloblast count over WBC count", Short1: "SENS-Myeloblast-count-over-WBC-count, Myeloblast / WBC", EnglishDescription: "SENS-Myeloblast-count-over-WBC-count"})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19390", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19391", EnglishHumanReadableExpression: "Neutrophilic Metamylocyte count over WBC count", Short1: "SENS-Neutrophilic-Metamylocyte-count-over-WBC-count, Neutrophilic Metamylocyte / WBC", EnglishDescription: "Neutrophilic Metamylocyte over WBC count"})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19391", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19392", EnglishHumanReadableExpression: "Monoblast-count-over-WBC-count", Short1: "SENS-Monoblast-count-over-WBC-count, Monoblast / WBC", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19392", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19393", EnglishHumanReadableExpression: "Promonocytes-count-over-WBC-count", Short1: "SENS-Promonocytes-count-over-WBC-count, Promonocytes / WBC", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19393", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19394", EnglishHumanReadableExpression: "Auer-Rods", Short1: "SENS-Auer-Rods, Auer rods", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19394")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19395", EnglishHumanReadableExpression: "Most-Recent-date-time-of-diagnosis-MDS-MPN-CMML", Short1: "SENS-Most-Recent-date-time-of-diagnosis-MDS-MPN-CMML, History of MDS/MPN (CMML)", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19395", UNIT_TIMESTAMP.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19396", EnglishHumanReadableExpression: "Most-Recent-date-time-of-diagnosis-Leukemia", Short1: "SENS-Most-Recent-date-time-of-diagnosis-Leukemia, History of Leukemia", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19396", UNIT_TIMESTAMP.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19397", EnglishHumanReadableExpression: "Serum-concentration-of-Creatinine", Short1: "SENS-Serum-concentration-of-Creatinine, Serum Creatinine", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19397")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19398", EnglishHumanReadableExpression: "Serum-concentration-of-Urea", Short1: "SENS-Serum-concentration-of-Urea, Serum Urea", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19398")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19399", EnglishHumanReadableExpression: "Serum-concentration-of-Potassium", Short1: "SENS-Serum-concentration-of-Potassium, Serum K", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19399")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19400", EnglishHumanReadableExpression: "Serum-concentration-of-Lactate-DeHydrogenase", Short1: "SENS-Serum-concentration-of-Lactate-DeHydrogenase, Serum LDH", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19400")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19401", EnglishHumanReadableExpression: "Prothrombin-Time", Short1: "SENS-Prothrombin-Time, PT", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19401")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19402", EnglishHumanReadableExpression: "activated-partial-thromboplastin-time", Short1: "SENS-activated-partial-thromboplastin-time, aPTT", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19402")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19403", EnglishHumanReadableExpression: "Toxic-Changes-in-Neutrophils", Short1: "SENS-Toxic-Changes-in-Neutrophils, Toxic changes", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19403", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19404", EnglishHumanReadableExpression: "Most-Recent-date-time-of-Chemotherapy", Short1: "SENS-Most-Recent-date-time-of-Chemotherapy, History of recent chemotherapy", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19404", UNIT_TIMESTAMP.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19405", EnglishHumanReadableExpression: "Most-Recent-date-time-of-GCSF-infusion", Short1: "SENS-Most-Recent-date-time-of-GCSF-infusion, History of recent GCSF", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19405", UNIT_TIMESTAMP.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19406", EnglishHumanReadableExpression: "Target cell count over RBC count", Short1: "SENS-Target-Cell-count-over-RBC-count, TRG / RBC", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19406", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19407", EnglishHumanReadableExpression: "Serum Ferritin Level", Short1: "SENS-Serum-Ferritin-level, Serum Ferritin", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19407", UNIT_MICROGRAM_PER_LITER.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19408", EnglishHumanReadableExpression: "Serum Iron Level", Short1: "SENS-Serum-Iron-level, Serum Iron", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19408", UNIT_MICROGRAM_PER_DECILITER.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19409", EnglishHumanReadableExpression: "Total iron-binding capacity", Short1: "SENS-Total-Iron-Binding-Capacity, TIBC", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19409", UNIT_MICROGRAM_PER_DECILITER.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19410", EnglishHumanReadableExpression: "Blood Oxygen Saturation", Short1: "SENS-Blood-Oxygen-Saturation, O2 Sat", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19410", UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19411", EnglishHumanReadableExpression: "Red Cell Distribution Width", Short1: "SENS-Red-Cell-Distribution-Width, RDW", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19411", UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19412", EnglishHumanReadableExpression: "", Short1: "SENS-Elevated-LFTs", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19412", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19413", EnglishHumanReadableExpression: "", Short1: "SENS-History-of-Splenectomy", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19413", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19414", EnglishHumanReadableExpression: "", Short1: "SENS-Basophilic-Stippling-count-over-RBC-count", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19414", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19415", EnglishHumanReadableExpression: "", Short1: "SENS-Howell-Jolly-Bodies-count-over-RBC-count", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19415", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19416", EnglishHumanReadableExpression: "ALT", Short1: "SENS-Alanine-Transaminase-blood-concentration", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19416", UNIT_INTERNATIONAL_UNITS_PER_LITER.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19417", EnglishHumanReadableExpression: "AST", Short1: "SENS-Aspartate-Transaminase-blood-concentration", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19417", UNIT_INTERNATIONAL_UNITS_PER_LITER.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19418", EnglishHumanReadableExpression: "ALP", Short1: "SENS-Alkaline-Phosphatase", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-alkaline phosphatase", UNIT_INTERNATIONAL_UNITS_PER_LITER.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19419", EnglishHumanReadableExpression: "GGT", Short1: "SENS-Gamma-Glutamyl-Transferase", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-gamma-glutamyl transferase ", UNIT_INTERNATIONAL_UNITS_PER_LITER.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19420", EnglishHumanReadableExpression: "INR", Short1: "SENS-international normalized ratio", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19420", UNIT_RATIO.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19421", EnglishHumanReadableExpression: "Albumin", Short1: "SENS-Albumin-Serum-Concentration, Albumin", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Albumin-serum-concentration", UNIT_GRAMS_PER_LITER.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19422", EnglishHumanReadableExpression: "", Short1: "SENS-Pappenheimer-Body-count-over-RBC-count", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19422", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19423", EnglishHumanReadableExpression: "", Short1: "SENS-History-of-Hemoglobinopathy", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19423", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19424", EnglishHumanReadableExpression: "", Short1: "SENS-Patient-Age", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19424", UNIT_DAY.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19425", EnglishHumanReadableExpression: "", Short1: "SENS-Atypical-Lymphocytes-count-over-WBC-count", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19425", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19426", EnglishHumanReadableExpression: "", Short1: "SENS-Mononucleosis-Spot-Test", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19426", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19427", EnglishHumanReadableExpression: "", Short1: "SENS-Accelerating-Lymphocytosis", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19427", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19428", EnglishHumanReadableExpression: "Smudge Cells over WBC", Short1: "SENS-Smudge-Cell-count-over-WBC-count, Smudge Cells / WBC", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19428", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19429", EnglishHumanReadableExpression: "", Short1: "SENS-History-of-Lymphoproliferative-Disorders", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-History-of-Lymphoproliferative-Disorders ", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19430", EnglishHumanReadableExpression: "Circulating Plasma Cells over WBC count", Short1: "SENS-Circulating-Plasma-Cell-count-over-WBC-count, Circulating Plasma Cells / WBC", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19430", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	// --------------------------------------------------

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19431", EnglishHumanReadableExpression: "", Short1: "SENS-Chest-pain", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19432", EnglishHumanReadableExpression: "", Short1: "SENS-Severe-lower chest-pain", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19433", EnglishHumanReadableExpression: "", Short1: "SENS-upper-abdominal-pain", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19434", EnglishHumanReadableExpression: "", Short1: "SENS-Sudden-breathlessness, onset over seconds", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19435", EnglishHumanReadableExpression: "", Short1: "SENS-Breathlessness on minimal exertion", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19436", EnglishHumanReadableExpression: "", Short1: "SENS-orthopnoea", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19437", EnglishHumanReadableExpression: "", Short1: "SENS-paroxysmal nocturnal dyspnoea (PND)", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19438", EnglishHumanReadableExpression: "", Short1: "SENS-Palpitations", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19439", EnglishHumanReadableExpression: "", Short1: "SENS-Acute breathlessness", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19440", EnglishHumanReadableExpression: "", Short1: "SENS-wheeze \u00b1 cough", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19441", EnglishHumanReadableExpression: "", Short1: "SENS-Cough and pink frothy sputum", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19442", EnglishHumanReadableExpression: "", Short1: "SENS-Syncope", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19443", EnglishHumanReadableExpression: "", Short1: "SENS-Leg pain on walking\u2014intermittent claudication", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19444", EnglishHumanReadableExpression: "", Short1: "SENS-Leg pain on standing\u2014relieved by lying down", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19445", EnglishHumanReadableExpression: "", Short1: "SENS-Unilateral calf swelling", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19446", EnglishHumanReadableExpression: "", Short1: "SENS-Bilateral ankle swelling", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19447", EnglishHumanReadableExpression: "", Short1: "SENS-Thoughts on interpreting cardiovascular signs", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19448", EnglishHumanReadableExpression: "", Short1: "SENS-Peripheral cyanosis", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19449", EnglishHumanReadableExpression: "", Short1: "SENS-Central cyanosis", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19450", EnglishHumanReadableExpression: "", Short1: "SENS-Tachycardia (e.g. pulse rate >\u000020bpm) 200 Bradycardia (<60bpm) 202", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19451", EnglishHumanReadableExpression: "", Short1: "SENS-Pulse irregular", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19456", EnglishHumanReadableExpression: "", Short1: "SENS-Postural fall in blood pressure", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19457", EnglishHumanReadableExpression: "", Short1: "SENS-BP/pulse difference between arms", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19458", EnglishHumanReadableExpression: "", Short1: "SENS-BP/pulse difference between arm and legs", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19454", EnglishHumanReadableExpression: "", Short1: "SENS-Blood pressure high\u2014hypertension", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19460", EnglishHumanReadableExpression: "", Short1: "SENS-Unilateral leg and ankle swelling", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19461", EnglishHumanReadableExpression: "", Short1: "SENS-Bilateral leg and ankle swelling", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19462", EnglishHumanReadableExpression: "", Short1: "SENS-Raised jugular venous pressure", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19463", EnglishHumanReadableExpression: "", Short1: "SENS-Abnormal apex impulse", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19464", EnglishHumanReadableExpression: "", Short1: "SENS-Extra heart sounds", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19465", EnglishHumanReadableExpression: "", Short1: "SENS-Diastolic murmur", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19466", EnglishHumanReadableExpression: "", Short1: "SENS-Mid-systolic murmur", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19467", EnglishHumanReadableExpression: "", Short1: "SENS-Pansystolic murmur", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19468", EnglishHumanReadableExpression: "", Short1: "SENS-Murmurs not entirely in systole or diastole", EnglishDescription: ""})

	// --------------------------------------------------

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19469", EnglishHumanReadableExpression: "Blast / WBC >= 20%", Short1: "FACT-Blast-Over20p, Blast >= 20%", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19469")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19470", EnglishHumanReadableExpression: "Blast / WBC < 20%", Short1: "FACT-Blast-Below20p, Blast < 20%", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19470")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19471", EnglishHumanReadableExpression: "Monoblast / WBC >= 20%", Short1: "FACT-Monoblast-Over20p, Moboblast >= 20%", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19471")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19472", EnglishHumanReadableExpression: "Moboblast / WBC < 20%", Short1: "FACT-Monoblast-Below20p, Monoblast < 20%", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19472")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19473", EnglishHumanReadableExpression: "Promonocytes / WBC >= 20%", Short1: "FACT-Promonocytes-Over20p, Promonocytes >= 20%", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19473")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19474", EnglishHumanReadableExpression: "Promonocytes / WBC < 20%", Short1: "FACT-Promonocytes-Below20p, Promonocytes < 20%", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19474")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19475", EnglishHumanReadableExpression: "Spherocytosis", Short1: "FACT-Spherocytosis", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19475")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19476", EnglishHumanReadableExpression: "RBC Fragments", Short1: "FACT-RBCFragments", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19476")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19477", EnglishHumanReadableExpression: "Thrombocytopenia", Short1: "FACT-Low-Platelet-count-Adult-150000-µl, Thrombocytopenia", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19477")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19478", EnglishHumanReadableExpression: "Prolonged Prothrombin Time", Short1: "FACT-Prolonged-Prothrombin-Time, Prolonged PT", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19478")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19479", EnglishHumanReadableExpression: "Prolonged activated Partial Thromboplastin Time", Short1: "FACT-Prolonged-activated-partial-thromboplastin-time, Prolonged aPTT", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19479")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19480", EnglishHumanReadableExpression: "Reticulocyte", Short1: "FACT-Retic, Retic", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19480")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19481", EnglishHumanReadableExpression: "Rouleaux Formation", Short1: "FACT-Rouleaux-Formation, Rouleaux", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19481", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19482", EnglishHumanReadableExpression: "Circulating Plasma Cells over 20 percent", Short1: "FACT-Circulating-Plasma-Cells-Over20p, Circulating Plasma Cells Over 20 percent", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19482")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19483", EnglishHumanReadableExpression: "Positive History of Acute Leukemia", Short1: "Positive-History-of-Acute-Leukemia, History of Acute Leukemia +", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19483")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19484", EnglishHumanReadableExpression: "Promyelocyte Detected", Short1: "FACT-Promyelocyte-Detected, Promyelocyte Detected", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19484")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19485", EnglishHumanReadableExpression: "Positive history of recent chemotherapy", Short1: "Positive-History-of-recent-chemotherapy, Recent Chemotherapy  +", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19485")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19486", EnglishHumanReadableExpression: "Positive history of recent GCSF", Short1: "Positive-History-of-recent-GCSF, Recent GCSF  +", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19486")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19487", EnglishHumanReadableExpression: "Toxic Changes in Neutrophils Detected", Short1: "Toxic-Changes-in-Neutrophils-Detected, Toxic changes +", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19487", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19488", EnglishHumanReadableExpression: "High Serum concentration of Lactate DeHydrogenase", Short1: "High-Serum-concentration-of-Lactate-DeHydrogenase, High Serum LDH", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19488", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19489", EnglishHumanReadableExpression: "High Serum concentration of Creatinine", Short1: "High-Serum-concentration-of-Creatinine, High Serum Creatinine", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19489", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19490", EnglishHumanReadableExpression: "High Serum concentration of Urea", Short1: "High-Serum-concentration-of-Urea, High Serum Urea", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19490", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19491", EnglishHumanReadableExpression: "High Serum concentration of Potassium", Short1: "High-Serum-concentration-of-Potassium, High Serum K", EnglishDescription: "", Categories: []Concept{DIAGNOSTIC_FLOW_INPUT, DIAGNOSTIC_FLOW_OUTPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19491", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19492", EnglishHumanReadableExpression: "Hypogranular Neutrophils detected", Short1: "FACT-Hypogranular-Neutrophils-detected, Hypogranular Neutrophils +", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19492", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19493", EnglishHumanReadableExpression: "Hypersegmented Neurophil detected", Short1: "FACT-Hypersegmented-Neurophil-detected, Hypersegmented Neurophils +", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19493", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19494", EnglishHumanReadableExpression: "Hypogranular Platelet detected", Short1: "FACT-Hypogranular-Platelet-detected, Hypogranular Platelets +", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19494", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19495", EnglishHumanReadableExpression: "Positive History of MDS", Short1: "FACT-Positive-History-of-MDS, History of MDS +", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19495", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19496", EnglishHumanReadableExpression: "Positive History of MDS-MPN-CMML", Short1: "FACT-Positive-History-of-MDS-MPN-CMML, History of MDS-MPN-CMML +", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19496", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19497", EnglishHumanReadableExpression: "Elevated NRBC count", Short1: "FACT-NRBC-count-elevated, Elevated NRBC", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19497", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19498", EnglishHumanReadableExpression: "Elevated Metamyelocyte count", Short1: "FACT-Elevated-Metamyelocyte-over-WBC-count, Elevated Metamyelocyte", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19498", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19499", EnglishHumanReadableExpression: "Elevated Promyelocyte count", Short1: "FACT-Elevated-Promyelocyte-count-over-WBC-count, Elevated Promyelocyte", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19499", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19500", EnglishHumanReadableExpression: "Elevated Myelocyte count", Short1: "FACT-Elevated-Myelocyte-count-over-WBC-count, Elevated Myelocyte", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19500", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19501", EnglishHumanReadableExpression: "Elevated Teardrop count", Short1: "FACT-Elevated-Teardrop-RBC-count-over-RBC-count, Elevated Teardrop RBC", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19501", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19234", EnglishHumanReadableExpression: "MAHA", Short1: "DDX-MAHA", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19234", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19235", EnglishHumanReadableExpression: "'IHA", Short1: "DDX-IHA", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19235", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19452", EnglishHumanReadableExpression: "", Short1: "SENS-Pulse volume high", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19236", EnglishHumanReadableExpression: "HS", Short1: "DDX-HS", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19236", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19237", EnglishHumanReadableExpression: "Acute Myeloid Leukemia", Short1: "AML, AML", EnglishDescription: "Acute Myeloid Leukemia", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19238", EnglishHumanReadableExpression: "Acute Lymphocytic Leukemia", Short1: "ALL, ALL", EnglishDescription: "Acute Lymphocytic Leukemia", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19239", EnglishHumanReadableExpression: "Chronic Myeloid Leukemia", Short1: "CML, CML", EnglishDescription: "Chronic Myeloid Leukemia", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19240", EnglishHumanReadableExpression: "Juvenile Myelomonocytic Leukemia", Short1: "JMML, JMML", EnglishDescription: "Juvenile Myelomonocytic Leukemia", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19241", EnglishHumanReadableExpression: "Chronic Lymphocytic Leukemia (CLL)", Short1: "CLL, CLL", EnglishDescription: "Chronic Lymphocytic Leukemia (CLL)", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19242", EnglishHumanReadableExpression: "Chronic Basophilic Leukemia (CBL)", Short1: "CBL, CBL", EnglishDescription: "Chronic Basophilic Leukemia (CBL)", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19243", EnglishHumanReadableExpression: "Hypochromic Anemia", Short1: "Hypochromic-Anemia, Hypochromic Anemia", EnglishDescription: "Hypochromic Anemia", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19244", EnglishHumanReadableExpression: "Thrombocytopenia", Short1: "Thrombocytopenia, Thrombocytopenia", EnglishDescription: "Thrombocytopenia", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19245", EnglishHumanReadableExpression: "Neutropenia", Short1: "Neutropenia, Neutropenia", EnglishDescription: "Neutropenia", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19246", EnglishHumanReadableExpression: "Plasma Cell Leukemia", Short1: "Plasma-Cell-Leukemia, Plasma Cell Leukemia", EnglishDescription: "Plasma Cell Leukemia", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19246", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19247", EnglishHumanReadableExpression: "Multiple Myeloma", Short1: "Multiple-Myeloma, Multiple Myeloma", EnglishDescription: "Multiple Myeloma", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19247", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19248", EnglishHumanReadableExpression: "Acute Leukemia", Short1: "Acute-Leukemia, Acute Leukemia", EnglishDescription: "Acute Leukemia", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19248", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19249", EnglishHumanReadableExpression: "Acute Promyelocytic Leukemia", Short1: "Acute-Promyelocytic-Leukemia, Acute Promyelocytic Leukemia", EnglishDescription: "Acute Promyelocytic Leukemia", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19249", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19250", EnglishHumanReadableExpression: "Myeloproliferative neoplasms", Short1: "Myeloproliferative-Neoplasms, MPN", EnglishDescription: "Myeloproliferative neoplasms", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19250", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19251", EnglishHumanReadableExpression: "Myelodysplastic syndromes", Short1: "Myelodysplastic-syndromes, MDS", EnglishDescription: "Myelodysplastic syndromes", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19251", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19252", EnglishHumanReadableExpression: "MDS/MPN (CMML) Overlap syndroms", Short1: "Myelodysplastic-syndrome-myeloproliferative-neoplasm-overlap-syndromes, MDS/MPN (CMML)", EnglishDescription: "MDS/MPN (CMML) Overlap syndroms", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19252", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19253", EnglishHumanReadableExpression: "History of recent chemotherapy", Short1: "DDX-History-of-recent-chemotherapy, Recent Chemotherapy", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19253", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19254", EnglishHumanReadableExpression: "History of recent GCSF", Short1: "History-of-recent-GCSF, Recent GCSF", EnglishDescription: "History of recent GCSF", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("19254", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19255", EnglishHumanReadableExpression: "Iron Deficiency Anemia", Short1: "Iron-Deficiency-Anemia, Iron Deficiency Anemia", EnglishDescription: "Iron Deficiency Anemia", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19256", EnglishHumanReadableExpression: "Monoclonal B-Lymphocytosis", Short1: "Monoclonal-B-Lymphocytosis", EnglishDescription: "Monoclonal B-Lymphocytosis", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19257", EnglishHumanReadableExpression: "Sideroblastic anemia", Short1: "Sideroblastic-Anemia", EnglishDescription: "Sideroblastic anemia", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19258", EnglishHumanReadableExpression: "Angina", Short1: "Angina", EnglishDescription: "Angina", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19259", EnglishHumanReadableExpression: "ST Elevated Myocardial Infarction", Short1: "STEMI, STEMI", EnglishDescription: "ST Elevated Myocardial Infarction", Categories: []Concept{DIAGNOSTIC_FLOW_OUTPUT, DIAGNOSTIC_FLOW_INPUT}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19260", EnglishHumanReadableExpression: "Non-ST Elevation Myocardial Infaction", Short1: "DDX-NSTEMI, NSTEMI", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19261", EnglishHumanReadableExpression: "", Short1: "DDX-Esophagitis", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19262", EnglishHumanReadableExpression: "", Short1: "DDX-Esophagial-Spasm", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19263", EnglishHumanReadableExpression: "", Short1: "DDX-Pulmunary-Embolus", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19264", EnglishHumanReadableExpression: "", Short1: "DDX-Pulmunary-Infarction", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19265", EnglishHumanReadableExpression: "", Short1: "DDX-Atrial-Fibrilation", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Cardiac-Arythmia"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19266", EnglishHumanReadableExpression: "", Short1: "DDX-Pneumothorax", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19267", EnglishHumanReadableExpression: "", Short1: "DDX-Tension-Pneumothorax", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19268", EnglishHumanReadableExpression: "", Short1: "DDX-Dissecting-Thoracic-Aortic-Aneurism", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19269", EnglishHumanReadableExpression: "Chest Wall Pain", Short1: "DDX-Chest-Wall-Pain", EnglishDescription: "", AKA: []string{"DDX-Tietzes-Syndrom"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19270", EnglishHumanReadableExpression: "Tietze's Syndrom", Short1: "DDX-Tietzes-Syndrom", EnglishDescription: "", AKA: []string{"DDX-Chest-Wall-Pain"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19271", EnglishHumanReadableExpression: "", Short1: "DDX-Gastro-Esophageal-Reflux", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19272", EnglishHumanReadableExpression: "", Short1: "DDX-Gastritis", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19273", EnglishHumanReadableExpression: "", Short1: "DDX-Billiary-Colic", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19274", EnglishHumanReadableExpression: "", Short1: "DDX-Pancreatitis", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19275", EnglishHumanReadableExpression: "", Short1: "DDX-Inferior-Myocardial-Infarction", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19276", EnglishHumanReadableExpression: "", Short1: "DDX-Anaphylaxis", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19277", EnglishHumanReadableExpression: "", Short1: "DDX-Inhalation-of-Foreign-Body", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19278", EnglishHumanReadableExpression: "", Short1: "DDX-Pulmunary-Edema", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19279", EnglishHumanReadableExpression: "Chronic Obstructive Pulmonary Disease", Short1: "DDX-COPD, COPD", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19280", EnglishHumanReadableExpression: "", Short1: "DDX-Asthma", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19281", EnglishHumanReadableExpression: "Cardiac Arythmia", Short1: "DDX-Cardiac-Arythmia, Arythmia", EnglishDescription: "", AKA: []string{}, Includes: []string{"DDX-Atrial-Fibrilation", "DDX-PSVT"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19282", EnglishHumanReadableExpression: "", Short1: "DDX-Eisodic-Heart-Block", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Cardiac-Arythmia"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19283", EnglishHumanReadableExpression: "", Short1: "DDX-Sinus-Tachycardia", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Cardiac-Arythmia"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19284", EnglishHumanReadableExpression: "", Short1: "DDX-PSVT", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Cardiac-Arythmia"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19285", EnglishHumanReadableExpression: "", Short1: "DDX-SVT", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Cardiac-Arythmia"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19286", EnglishHumanReadableExpression: "", Short1: "DDX-VTAC", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Cardiac-Arythmia"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19287", EnglishHumanReadableExpression: "", Short1: "DDX-Ventricular-Fibrilation", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Cardiac-Arythmia"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19288", EnglishHumanReadableExpression: "", Short1: "DDX-Anxiety", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19289", EnglishHumanReadableExpression: "", Short1: "DDX-Thyrotoxicosis", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19290", EnglishHumanReadableExpression: "", Short1: "DDX-Hyperventilation", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19291", EnglishHumanReadableExpression: "", Short1: "DDX-Hypovolamia", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19292", EnglishHumanReadableExpression: "", Short1: "DDX-Infection", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19293", EnglishHumanReadableExpression: "", Short1: "DDX-Menopause", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19294", EnglishHumanReadableExpression: "", Short1: "DDX-Pheochromocytoma", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19295", EnglishHumanReadableExpression: "", Short1: "DDX-PVC", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Cardiac-Arythmia"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19296", EnglishHumanReadableExpression: "", Short1: "DDX-POTS-Syndrom", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19297", EnglishHumanReadableExpression: "", Short1: "DDX-Paraganglioma", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19298", EnglishHumanReadableExpression: "", Short1: "DDX-Exacerbation-of-Asthma", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19299", EnglishHumanReadableExpression: "", Short1: "DDX-Exacerbation-of-COPD", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19300", EnglishHumanReadableExpression: "", Short1: "DDX-Acute-Bronchitis", EnglishDescription: "", AKA: []string{}, Includes: []string{"DDX-Acute-Bacterial-Bronchitis", "DDX-Acute-Viral-Bronchitis"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19301", EnglishHumanReadableExpression: "", Short1: "DDX-Acute-Viral-Bronchitis", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Acute-Bronchitis"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19302", EnglishHumanReadableExpression: "", Short1: "DDX-Acute-Bacterial-Bronchitis", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Acute-Bronchitis"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19303", EnglishHumanReadableExpression: "", Short1: "DDX-Acute-Left-Ventricular-Failure", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Heart-Failure"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19304", EnglishHumanReadableExpression: "", Short1: "DDX-Heart-Failure", EnglishDescription: "", AKA: []string{}, Includes: []string{"DDX-Acute-Left-Ventricular-Failure"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19305", EnglishHumanReadableExpression: "", Short1: "DDX-Valvular-Disease", EnglishDescription: "", AKA: []string{}, Includes: []string{"DDX-Mitral-Stenosis"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19306", EnglishHumanReadableExpression: "", Short1: "DDX-Mitral-Stenosis", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Valvular-Disease"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19307", EnglishHumanReadableExpression: "", Short1: "DDX-Vasovagal-Attack", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19308", EnglishHumanReadableExpression: "", Short1: "DDX-Postural-Hypotension", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19309", EnglishHumanReadableExpression: "", Short1: "DDX-Aortic-Stenosis", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19310", EnglishHumanReadableExpression: "", Short1: "DDX-Hyperthrophic-Cardiomyopathy", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19311", EnglishHumanReadableExpression: "", Short1: "DDX-Micturation-Syncope", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19312", EnglishHumanReadableExpression: "", Short1: "DDX-Cough-Syncope", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19313", EnglishHumanReadableExpression: "", Short1: "DDX-Carotid-Sinus-Syncope", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19314", EnglishHumanReadableExpression: "", Short1: "DDX-Hypoglycemia", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19315", EnglishHumanReadableExpression: "", Short1: "DDX-Epilepsy", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19316", EnglishHumanReadableExpression: "Cerebrovascular Accident", Short1: "DDX-CVA, CVA", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19317", EnglishHumanReadableExpression: "", Short1: "DDX-Arterial-disease-in-legs", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19318", EnglishHumanReadableExpression: "", Short1: "DDX-Spinal-Claudication", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19319", EnglishHumanReadableExpression: "", Short1: "DDX-Peripheral-Venous-Disease", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19320", EnglishHumanReadableExpression: "", Short1: "DDX-Spinal-Disk-Protrusion", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19321", EnglishHumanReadableExpression: "", Short1: "DDX-Deep-Venous-Thrombosis", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19322", EnglishHumanReadableExpression: "", Short1: "DDX-Lower-Exterimity-Varicoses", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19323", EnglishHumanReadableExpression: "", Short1: "DDX-Ruptured-Bakers-Cyst", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19324", EnglishHumanReadableExpression: "", Short1: "DDX-Cellulitis", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19325", EnglishHumanReadableExpression: "", Short1: "DDX-Abnormal-Lymphatic-Drainage", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19326", EnglishHumanReadableExpression: "", Short1: "DDX-Lymphoma", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19327", EnglishHumanReadableExpression: "", Short1: "DDX-Right-Ventricular-Failure", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19328", EnglishHumanReadableExpression: "", Short1: "DDX-Pulmonary-Hypertension", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19329", EnglishHumanReadableExpression: "", Short1: "DDX-Congestive-Heart-Failure", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19330", EnglishHumanReadableExpression: "", Short1: "DDX-Poor-Venous-return-due-to-abdominal-pelvic-mass", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19331", EnglishHumanReadableExpression: "", Short1: "DDX-Low-albumin-state", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19332", EnglishHumanReadableExpression: "", Short1: "DDX-Liver-Failure", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19333", EnglishHumanReadableExpression: "", Short1: "DDX-Nephrotic-Syndrom", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19334", EnglishHumanReadableExpression: "", Short1: "DDX-Malnutrition", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19335", EnglishHumanReadableExpression: "", Short1: "DDX-Malabsorption", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19336", EnglishHumanReadableExpression: "", Short1: "DDX-IVC-Obstruction-due-to-clot", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19337", EnglishHumanReadableExpression: "", Short1: "DDX-Bilateral-Venous-thrombosis-of-lower-extremities", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19338", EnglishHumanReadableExpression: "", Short1: "DDX-Septicemia", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19339", EnglishHumanReadableExpression: "", Short1: "DDX-Valular-stenosis", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19340", EnglishHumanReadableExpression: "", Short1: "DDX-Valvular-incompetence", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19341", EnglishHumanReadableExpression: "", Short1: "DDX-Hemorhage", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19342", EnglishHumanReadableExpression: "", Short1: "DDX-Internal-Hemorhage", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19343", EnglishHumanReadableExpression: "", Short1: "DDX-GI-Bleeding", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19344", EnglishHumanReadableExpression: "", Short1: "DDX-Raynauds-Phenomenon", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19345", EnglishHumanReadableExpression: "", Short1: "DDX-Peripheral-Arterial-Obstruction", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19346", EnglishHumanReadableExpression: "", Short1: "DDX-Diabetic-small-vessle-disease", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19347", EnglishHumanReadableExpression: "", Short1: "DDX-Low-Cardiac-output", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19348", EnglishHumanReadableExpression: "Myocardial Infarction", Short1: "DDX-MI, MI", EnglishDescription: "", AKA: []string{}, Includes: []string{"DDX-STEMI", "DDX-NSTEMI", "DDX-Inferior-Myocardial-Infarction"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19349", EnglishHumanReadableExpression: "", Short1: "DDX-Right-To-Left-Cardiac-Shunt", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19350", EnglishHumanReadableExpression: "", Short1: "DDX-Tetralogy-of-Fallot", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19351", EnglishHumanReadableExpression: "", Short1: "DDX-Tricuspid-Atresia", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19352", EnglishHumanReadableExpression: "", Short1: "DDX-Pulmonary-AV-Fistula", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19353", EnglishHumanReadableExpression: "", Short1: "DDX-Right-To-Left-Pulmonary-shunt", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19354", EnglishHumanReadableExpression: "", Short1: "DDX-Hemoglobin-Abnormalities", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19355", EnglishHumanReadableExpression: "", Short1: "DDX-Hemoglobin-M-disease", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19356", EnglishHumanReadableExpression: "", Short1: "DDX-NADH-Diaphorase", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19357", EnglishHumanReadableExpression: "", Short1: "DDX-Methemoglobinemia", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "19358", EnglishHumanReadableExpression: "", Short1: "DDX-Sulfhemoglobinemia", EnglishDescription: "", AKA: []string{}})

	// -------------------------------------------------     EVAL-  labs and paraclinic and radiology and such
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "EVAL-", EnglishHumanReadableExpression: "abdominal ultrasound", Short1: "Abdominal sonogram", EnglishDescription: "", AKA: []string{}})

	// acetylcholine receptor antibody panel aka (AChR Ab, Anti–AChR antibody)
	// acid phosphatase aka (Prostatic acid phosphatase [PAP],Tartrate-resistant acid phosphatase [TRAP])
	// activated clotting time (ACT, Activated coagulation time)
	// adrenocorticotropic hormone (ACTH, Corticotropin)
	// ACTH stimulation test with cosyntropin aka adrenocorticotropic hormone stimulation test with cosyntropin (ACTH stimulation test, Cortisol stimulation test)
	// ACTH stimulation test with metyrapone aka adrenocorticotropic hormone stimulation test with metyrapone (ACTH stimulation test with metyrapone, Metyrapone test)
	// adrenal steroid precursors aka (Androstenediones [AD], Dehydroepiandrosterone [DHEA], Dehydroepiandrosterone sulfate [DHEA S], 11-Deoxycortisol, 17-Hydroxyprogesterone, 17-Hydroxypregnenolone, Pregnenolone)
	// age-related macular degeneration risk analysis aka (ARMD risk analysis, Y402H, and A69S)
	// alanine aminotransferase aka (ALT, formerly Serum glutamic-pyruvic transaminase [SGPT])
	// aldolase
	// aldosterone
	// alkaline phosphatase (ALP)
	// allergy blood testing aka (IgE antibody test, Radioallergosorbent test [RAST])
	// allergy skin testing
	// alpha1-antitrypsin aka (A1AT, AAT, Alpha1-antitrypsin phenotyping)
	// aluminum
	// amino acid profiles aka (Amino acid screen)
	// 	ammonia level
	// 	amniocentesis aka (Amniotic fluid analysis)
	// amylase
	// amyloid beta protein precursor, soluble (sBPP)
	// angiotensin
	// angiotensin-converting enzyme angiotensin-converting enzyme (ACE, Serum angiotensin-
	// converting enzyme [SACE])
	// anion gap
	// anion gap (AG, R factor)
	// 64 anticardiolipin antibodies
	// anticardiolipin antibodies (aCL antibodies, ACA,
	// Antiphospholipid antibodies, Lupus anticoagulant)
	// 66 anticentromere antibody test
	// anticentromere antibody test (Centromere antibody)
	// antichromatin antibody test 67 antichromatin antibody test (Antinucleosome antibody A
	// [anti-NCS], Antihistone antibody test [anti-HST, AHA])
	// anti–cyclic citrullinated peptide antibody
	// 69
	// 	anti–cyclic citrullinated peptide antibody (Cyclic citrullinated peptide antibody, CCP IgG anti-CCP)
	// 70 antidiuretic hormone
	// antidiuretic hormone (ADH, Vasopressin)
	// anti-DNA antibody test 73
	// 	anti-DNA antibody test (Anti–deoxyribonucleic acid antibodies, Antibody to double-stranded DNA, Anti–double-stranded DNA, Anti–ds-DNA, DNA antibody, Native double-stranded DNA)
	// anti–extractable nuclear antigens 75
	// 	anti–extractable nuclear antigens (Anti-ENAs, Antibodies to extractable nuclear antigens, Antihistidyl transfer synthase [anti–Jo-1], Antiribonucleoprotein [anti-RNP], Anti-Smith [anti-SM])
	// anti–glomerular basement membrane antibodies 77
	// 	anti–glomerular basement membrane antibodies (Anti-GBM antibody, AGBM, Glomerular basement antibody, Goodpasture antibody)
	// A
	// anti-glycan antibodies
	// 79
	// 	anti-glycan antibodies (Crohn disease prognostic panel, Multiple sclerosis antibody panel)
	// 80 anti-liver/kidney microsomal type 1 antibodies anti-liver/kidney microsomal type 1 antibodies
	// (Anti-LKM-1 antibodies)
	// 82 antimitochondrial antibody antimitochondrial antibody (AMA)
	// antimyocardial antibody 83 antimyocardial antibody (AMA)
	// 84 antineutrophil cytoplasmic antibody antineutrophil cytoplasmic antibody (ANCA)
	// 86 antinuclear antibody antinuclear antibody (ANA)
	// 90 antiparietal cell antibody antiparietal cell antibody (APCA)
	// antiscleroderma antibody
	// 91
	// 	antiscleroderma antibody (Scl-70 antibody, Scleroderma antibody, RNA polymerase III antibody)
	// anti–smooth muscle antibody 93 anti–smooth muscle antibody (ASMA)
	// 94 antispermatozoal antibody
	//  antispermatozoal antibody (Sperm agglutination and inhibition, Sperm antibodies, Antisperm antibodies, Infertility screen)
	// 96 anti–SS-A (Ro), anti–SS-B (La), and anti–SS-C antibodies anti–SS-A (Ro), anti–SS-B (La), and anti–SS-C
	//  antibodies (Anti-Ro, Anti-La, Sjögren antibodies)
	// 98 antithrombin activity and antigen assay
	// antithrombin activity and antigen assay (Antithrombin III [AT-III] activity/assay, Functional antithrombin III assay, Heparin cofactor, Immunologic antithrombin III, Serine
	// protease inhibitor)
	// antithyroglobulin antibody 101 antithyroglobulin antibody (Thyroid autoantibody, Thyroid A
	// antithyroglobulin antibody, Thyroglobulin antibody)
	// 102 antithyroid peroxidase antibody
	// antithyroid peroxidase antibody (Anti-TPO, TPO-Ab,
	// Antithyroid microsomal antibody, Thyroid autoantibody)
	// 104 apolipoproteins apolipoproteins
	// 108 Apt test
	// Apt test (Downey test, Qualitative fetal hemoglobin stool test,
	// Stool for swallowed blood)
	// arterial blood gases 109 arterial blood gases (ABGs, blood gases)
	// 	arteriography (Angiography)
	// arteriography 117
	// A
	// arthrocentesis with synovial fluid analysis 121 arthrocentesis with synovial fluid analysis
	// 	arthroscopy
	// arthroscopy 125
	// A
	// aspartate aminotransferase 129 aspartate aminotransferase (AST; Formerly called serum A
	// glutamic-oxaloacetic transaminase [SGOT])
	// 132 barium enema
	// barium enema (BE, Lower GI series)
	// 136 barium swallow barium swallow
	// Bence Jones protein 139 Bence Jones protein (Free kappa and lambda light chains)
	// 11 beta-prostaglandin F(2) alpha 141 11 beta-prostaglandin F(2) alpha
	// 142 bilirubin bilirubin
	// 146 bioterrorism infectious agents testing bioterrorism infectious agents testing
	// 152 bladder cancer markers
	// bladder cancer markers (Bladder tumor antigen [BTA],
	// Nuclear matrix protein 22 [NMP22])
	// 154 blood culture and sensitivity blood culture and sensitivity
	// 156 blood smear
	// blood smear (Peripheral blood smear, Red blood cell [RBC]
	// morphology, RBC smear)
	//  blood typing (Blood group microarray testing)
	// 162 bone densitometry
	// bone densitometry (Bone mineral content [BMC], Bone mineral
	// density [BMD], DEXA scan)
	// 166 bone marrow biopsy
	// bone marrow biopsy (Bone marrow examination, Bone
	// marrow aspiration)
	//  bone scan
	// 174 bone turnover markers
	// bone turnover markers (BTMs, N-telopeptide [NTx], Bone collagen equivalents (BCEs), Osteocalcin [bone G1 a protein, BGP, osteocalc], Pyridinium [PYD] crosslinks, Bone-specific alkaline phosphatase [BSAP], Amino-terminal propeptide of type I procollagen [P1NP], C-telopeptide [CTx])
	// 178 bone x-ray bone x-ray
	// 180 brain scan
	// brain scan (Cisternogram, Cerebral blood flow, DaT scan)
	// 182 breast cancer genomics
	// breast cancer genomics (Oncotype DX, Genotyping,
	// MammaPrint)
	// 184 breast cancer tumor analysis
	//  breast cancer tumor analysis (DNA ploidy status, S-phase fraction, Cathepsin D, HER-2 [c erbB2, neu] protein, p53 protein, Ki67 protein)
	//  breast ductal lavage
	// breast ductal lavage 187
	// B
	// breast ultrasonography 189 breast ultrasonography (Ultrasound mammography,
	// Breast sonogram)
	//  bronchoscopy
	// bronchoscopy 191
	// B
	// 196 CA 15-3 and CA 27.29 tumor markers
	// CA 15-3 and CA 27.29 tumor markers (Cancer antigen
	// 15-3, Cancer antigen 27.29)
	// CA 19-9 tumor marker 197 CA 19-9 tumor marker (Cancer antigen 19-9)
	// CA-125 tumor marker 199 CA-125 tumor marker (Cancer antigen-125)
	// calcitonin 201 calcitonin (Human calcitonin [HCT], Thyrocalcitonin)
	// 206 caloric study
	// caloric study (Oculovestibular reflex study)
	// 208 carbon dioxide content
	// carbon dioxide content (CO2 content, CO2 combining
	// power)
	// 210 carboxyhemoglobin carboxyhemoglobin (COHb, Carbon monoxide)
	// 212 carcinoembryonic antigen carcinoembryonic antigen (CEA)
	// 214 cardiac catheterization
	// cardiac catheterization (Coronary angiography,
	// Ventriculography)
	// cardiac nuclear scan (Myocardial perfusion scan, Myocardial perfusion imaging, Myocardial scan, Cardiac scan, Heart scan, Thallium scan, MUGA scan, Isonitrile scan, Sestamibi cardiac scan, Cardiac flow studies, and Nuclear stress test)
	// cardiac stress testing 225 cardiac stress testing (Exercise stress testing; Nuclear stress
	// testing; Echo stress testing)
	// carotid artery duplex scanning 229 carotid artery duplex scanning (Carotid ultrasound)
	// cell culture drug resistance testing 231 cell culture drug resistance testing (CCDRT,
	// Chemosensitivity assay, Drug response assay)
	// 232 cell-free maternal DNA testing
	// cell-free maternal DNA testing (Noninvasive prenatal
	// testing [NIPT, cell-free DNA in maternal blood])
	// 234 cell surface immunophenotyping
	//  cell surface immunophenotyping (Flow cytometry cell surface immunophenotyping, Lymphocyte immunophenotyping, AIDS T-lymphocyte cell markers, CD4 marker, CD4/CD8 ratio, CD4 percentage)
	// cervical biopsy 237 cervical biopsy (LEEP procedure, Cone biopsy)
	// 240 chest x-ray chest x-ray (CXR)
	// Chlamydia
	// 246 chloride, blood chloride, blood (Cl)
	// 248 cholesterol cholesterol
	// cholinesterase (CHS, Pseudocholinesterase [PChE], Cholinesterase RBC, Red blood cell cholinesterase, Acetylcholinesterase)
	// 254 chorionic villus sampling
	// chorionic villus sampling (CVS, Chorionic villus biopsy [CVB])
	// chromosome karyotype 257 chromosome karyotype (Blood chromosome analysis,
	// Chromosome studies, Cytogenetics, Karyotype)
	// Clostridium difficile 259 Clostridium difficile (C. diff, Clostridial toxin assay)
	// coagulating factors concentration 261 coagulating factors concentration (Coagulating factors,
	// Blood-clotting factors)
	// cold agglutinins
	// colon cancer tumor analysis (Microsatellite instability [MSI] testing, DNA mismatch repair [MMR] genetic testing, BRAF mutation analysis, Oncotype DX colon cancer assay)
	// colonoscopy
	// 274 colposcopy colposcopy
	// complement assay 277 complement assay (Total, C3 and C4 complement)
	// computed tomography of the abdomen and pelvis 281
	//  computed tomography of the abdomen and pelvis (CT scan of the abdomen and pelvis, Helical/spiral CT scan of the abdomen and pelvis)
	// computed tomography of the brain 287 computed tomography of the brain (CT scan
	// of the brain, Helical/spiral CT scan of the brain)
	// 290 computed tomography of the chest
	// computed tomography of the chest (Chest CT scan,
	// Helical/spiral CT scan of the chest)
	// computed tomography of the heart 293 computed tomography of the heart (Coronary CT
	// angiography, Coronary calcium score)
	// Coombs test, direct 297 Coombs test, direct (Direct antiglobulin test [DAT])
	// Coombs test, indirect 299 Coombs test, indirect (Blood antibody screening,
	// Indirect antiglobulin test [IAT])
	// cortisol, blood, urine, saliva 301 cortisol, blood, urine, saliva (Hydrocortisone, Serum
	// cortisol, Salivary cortisol)
	// 304 C-peptide
	// C-peptide (Connecting peptide insulin, Insulin C-peptide,
	// Proinsulin C-peptide)
	// 306 C-reactive protein test
	// C-reactive protein test (CRP, High-sensitivity C-reactive
	// protein [hs-CRP])
	// 308 creatine kinase
	// creatine kinase (CK, Creatine phosphokinase [CPK])
	// 312 creatinine, blood
	// creatinine, blood (Serum creatinine)
	// creatinine clearance 315 creatinine clearance (CC, Estimated glomerular filtration
	// rate [eGFR])
	// cryoglobulin
	// cystography 321 cystography (Cystourethrography, Voiding cystography, Voiding
	// cystourethrography [VCUG])
	// cystometry 323
	//  cystometry (Cystometrogram [CMG])
	// 326 cystoscopy cystoscopy (Endourology)
	// cytokines
	// cytomegalovirus (CMV)
	// D-dimer 335 D-dimer (Fragment D-dimer, Fibrin degradation product [FDP],
	// Fibrin split products)
	// delta-aminolevulinic acid 337 delta-aminolevulinic acid (Aminolevulinic acid [ALA], d-ALA)
	// dexamethasone suppression test 339 dexamethasone suppression test (DST, Prolonged/rapid
	// DST, Cortisol suppression test, ACTH suppression test)
	// 342 diabetes mellitus autoantibody panel
	//  diabetes mellitus autoantibody panel (Insulin autoantibody [IAA], Islet cell antibody [ICA], Glutamic acid decarboxylase antibody [GAD Ab])
	// 2,3-diphosphoglycerate 343 2,3-diphosphoglycerate (2,3-DPG in erythrocytes)
	// disseminated intravascular coagulation screening 345 disseminated intravascular coagulation screening
	// (DIC screening)
	// drug monitoring 347 drug monitoring (Therapeutic drug monitoring [TDM])
	// 352 drug sensitivity genotype testing
	// drug sensitivity genotype testing (AccuType)
	// ductoscopy (Mammary ductoscopy)
	// 356 echocardiography
	// echocardiography (Cardiac echo, Heart sonogram, Transthoracic
	// echocardiography [TTE])
	// electrocardiography 359 electrocardiography (Electrocardiogram [ECG, EKG])
	// electroencephalography 365 electroencephalography (Electroencephalogram [EEG])
	// 368 electromyography electromyography (EMG)
	// electromyography of the pelvic floor sphincter 371
	//  electromyography of the pelvic floor sphincter (Pelvic floor sphincter electromyography, Pelvic floor sphincter EMG, Rectal EMG procedure)
	// electroneurography 373 electroneurography (ENG, Nerve conduction studies)
	// 376 electronystagmography electronystagmography (Electrooculography)
	// 378 electrophysiologic study electrophysiologic study (EPS, Cardiac mapping)
	// 382 endometrial biopsy endometrial biopsy
	// 384 endoscopic retrograde cholangiopancreatography endoscopic retrograde cholangiopancreatography
	// (ERCP)
	// 388 Epstein-Barr virus testing
	// Epstein-Barr virus testing (EBV antibody titer)
	// erythrocyte fragility 391 erythrocyte fragility (Osmotic fragility [OF], Red blood cell
	// fragility)
	// erythrocyte sedimentation rate 393 erythrocyte sedimentation rate (ESR, Sed rate test)
	// erythropoietin (EPO)
	// esophageal function studies 397 esophageal function studies (Esophageal manometry,
	// Esophageal motility studies)
	// esophagogastroduodenoscopy 401 esophagogastroduodenoscopy (EGD, Upper
	// gastrointestinal [UGI] endoscopy, Gastroscopy)
	// estrogen fractions 405 estrogen fractions (Estriol excretion, Estradiol, Estrone)
	// 408 estrogen receptor assay
	// estrogen receptor assay (ER assay, ERA, Estradiol
	// receptor)
	// 410 ethanol
	// ethanol (Ethyl alcohol, Blood alcohol, Blood EtOH)
	// 412 evoked potential studies
	// evoked potential studies (EP studies, Evoked brain
	// potentials, Evoked responses)
	// 416 factor V Leiden
	// factor V Leiden (FVL, Mutation analysis)
	// 418 febrile antibodies
	// febrile antibodies (Febrile agglutinins)
	// fecal fat (Fat absorption, Quantitative stool fat determination)
	// ferritin
	// fetal biophysical profile 423 fetal biophysical profile (Biophysical profile [BPP])
	// 426 fetal contraction stress test
	// fetal contraction stress test (Contraction stress test [CST],
	// Oxytocin challenge test [OCT])
	// Fibronectin may help with implantation of the fertilized egg into the uterine lining. Normally, fibronectin cannot be identi- fied in vaginal secretions after 22 weeks of pregnancy. However, concentrations are very high in the amniotic fluid. If fibronectin is identified in vaginal secretions after 24 weeks, the patient is at high risk for preterm delivery. The use of fFN is limited to symptomatic women with contractions whose membranes are intact and who have cervical dilatation of less than 3cm.
	// Procedure and patient care
	// Before
	// Explain the procedure to the patient.
	// Tell the patient that no fasting is required.
	// • Determine whether the patient has had a cervical exam
	// or intercourse within the past 24 hours. Results may be inaccurate.
	// During
	// • Note the following procedural steps:
	// 1. The patient is placed in the lithotomy position.
	// 2. A vaginal speculum is inserted to expose the cervix.
	// 3. Vaginal secretions are collected from the posterior vagina
	// and paracervical area using a swab from a kit.
	// 4. The container with appropriate medium is labeled with the
	// patient’s name, age, and estimated date of confinement. Tell the patient that no discomfort, except for insertion of the speculum, is associated with this procedure.
	// • Note that this procedure is performed by a physician or other licensed health care provider in several minutes.
	// After
	// Tell the patient results will be available the same or next day. Educate the patient about the signs of preterm labor.
	// Abnormal findings
	// High risk for preterm premature delivery
	// F
	// fetal fibronectin 429
	//  fetal fibronectin (fFN)
	// 430 fetal hemoglobin
	// fetal hemoglobin (Kleihauer-Betke)
	// 432 fetal nonstress test
	// fetal nonstress test (Nonstress test, NST, Fetal activity
	// determination)
	// 434 fetal scalp blood pH fetal scalp blood pH
	// 436 fetoscopy fetoscopy
	// fibrinogen (Factor I, Quantitative fibrinogen)
	// fluorescein angiography 441 fluorescein angiography (FA, Ocular photography)
	// folic acid (Folate)
	// β-D-glucan: Negative
	// Indeterminate Positive
	// Less than 60pg/mL 60-79 pg/mL
	// ≥80 pg/mL
	// F
	// fungal testing
	// 445
	//  fungal testing (Antifungal antibodies; Beta-D-glucan (1,3)-ß-D-glucan, Fungitell, Fungal culture, Fungal antigen assay, Fungal PCR testing)
	// galectin-3 (GAL-3)
	// 448 gallbladder nuclear scanning
	//  gallbladder nuclear scanning (Hepatobiliary scintigraphy, Cholescintigraphy, DISIDA scanning, HIDA scanning)
	// 450 gallium scan gallium scan
	// 452 gamma-glutamyl transpeptidase gamma-glutamyl transpeptidase (GGTP, γ-GTP,
	// Gamma-glutamyl transferase [GGT])
	// 454 gastric emptying scan gastric emptying scan
	// 456 gastrin gastrin
	// 458 gastroesophageal reflux scan gastroesophageal reflux scan (GE reflux scan,
	// Aspiration scan)
	// 460 gastrointestinal bleeding scan gastrointestinal bleeding scan
	// (Abdominal scintigraphy, GI scintigraphy)
	// 462 genetic testing
	// genetic testing (Breast cancer [BRCA] and ovarian cancer,
	// Colon cancer, Cardiovascular disease, Tay-Sachs disease, Cystic fibrosis, Melanoma, Hemochromatosis, Thyroid cancer, Paternity [parentage analysis], Forensic genetic testing)
	// 470 gliadin, endomysial, and tissue transglutaminase antibodies gliadin, endomysial, and tissue transglutaminase
	//  antibodies
	// 472 glucagon glucagon
	// 474 glucose, blood
	// glucose, blood (Blood sugar, Fasting blood sugar [FBS])
	// glucose, postprandial 477 glucose, postprandial (2-Hour postprandial glucose
	// [2-Hour PPG], 1-Hour glucose screen)
	// glucose tolerance test 479 glucose tolerance test (GTT, Oral glucose tolerance test
	// [OGTT])
	// glycosylated hemoglobin (GHb, GHB, Glycohemoglobin, Glycolated hemoglobin, Hemoglobin [HbA1c], Diabetic
	// control index)
	// 486 growth hormone
	// growth hormone (GH, Human growth hormone [HGH],
	// Somatotropin hormone [SH])
	// growth hormone stimulation test 489 growth hormone stimulation test (GH provocation test,
	// Insulin tolerance test [ITT], Arginine test)
	// haptoglobin
	// Heinz body preparation
	// 494 Helicobacter pylori testing
	//  Helicobacter pylori testing (Anti-Helicobacter pylori antibody, Campylobacter-like organism [CLO] test, H. pylori stool antigen)
	// hematocrit 497 hematocrit (Hct, Packed red blood cell volume, Packed cell
	// volume [PCV])
	// 500 hemoglobin hemoglobin (Hb, Hgb)
	// 502 hemoglobin electrophoresis
	// hemoglobin electrophoresis (Hgb electrophoresis)
	// hepatitis virus studies
	// 510 herpes simplex
	// herpes simplex (Herpesvirus types 1 and 2, Herpes simplex virus
	// types 1 and 2 [HSV 1, HSV 2], Herpes genitalis)
	// hexosaminidase 513 hexosaminidase (Hexosaminidase A, Hex A, Total
	// hexosaminidase, Hexosaminidase A and B)
	// HIV drug resistance testing 515 HIV drug resistance testing (HIV genotype)
	// HIV RNA quantification 517 HIV RNA quantification (HIV viral load)
	// HIV serologic and virologic testing 521
	// HIV serologic and virologic testing (AIDS serology, Acquired immunodeficiency serology, AIDS screen, Human immunodeficiency virus [HIV] antibody test, Western blot test, p24 direct antigen, HIV-RNA viral test)
	// Holter monitoring is a continuous recording of the electrical activity of the heart. This can be performed for periods of up to 72 hours. With this technique, an electrocardiogram (ECG) is recorded continuously on magnetic tape during unrestricted activity, rest, and sleep (Figure 24). The Holter monitor is equipped with a clock that permits accurate time monitoring on the ECG tape. The patient is asked to carry a diary and to record daily activities as well as any cardiac symptoms that may develop during the period of monitoring.
	// Most units are equipped with an event marker. This is a but- ton the patient can push when such symptoms as chest pain, syn- cope, or palpitations are experienced. This type of monitor is referred to as an event recorder.
	// H
	// 	 Wires
	// Holter monitoring 525
	//  Holter monitoring
	// 528 homocysteine homocysteine (Hcy)
	// 530 human chorionic gonadotropin
	// human chorionic gonadotropin (hCG, Pregnancy tests,
	// hCG beta subunit)
	// human lymphocyte antigen B27 533
	//  human lymphocyte antigen B27 (HLA-B27 antigen, Human leukocyte A antigen, White blood cell antigens, Histocompatibility leukocyte A antigen)
	// 534 human papillomavirus
	// human papillomavirus (HPV test, HPV DNA testing)
	// 538 human placental lactogen
	// human placental lactogen (hPL, Human chorionic
	// somatomammotropin [HCS])
	// 540 human T-cell lymphotrophic virus
	// human T-cell lymphotrophic virus (HTLV) I/II antibody
	// 17-hydroxycorticosteroids 541 17-hydroxycorticosteroids (17-OCHS)
	// 5-hydroxyindoleacetic acid 543 5-hydroxyindoleacetic acid (5-HIAA)
	// 21-hydroxylase antibodies
	// 546 hysterosalpingography hysterosalpingography (Uterotubography,
	// Uterosalpingography, Hysterogram)
	// 548 hysteroscopy hysteroscopy
	// immunoglobulin quantification 551 immunoglobulin quantification
	// 554 insulin assay insulin assay
	// 556 insulin-like growth factor
	// insulin-like growth factor (IGF-1, Somatomedin C, Insulin-
	//  like growth factor binding proteins [IGF BP])
	// 558 intravascular ultrasound intravascular ultrasound (IVUS)
	// 560 intrinsic factor antibody intrinsic factor antibody (IF ab)
	// iron level and total iron-binding capacity 561 iron level and total iron-binding capacity (Fe and TIBC,
	// Transferrin saturation, Transferrin)
	// ischemia-modified albumin 565 ischemia-modified albumin (IMA)
	// 566 laboratory genetics laboratory genetics
	// Under conditions of normal oxygen availability to tissues, glucose is metabolized to CO2 and H2O for energy. When oxy- gen to the tissues is diminished, anaerobic metabolism of glucose occurs, and lactate (lactic acid) is formed instead of CO2 and H2O. To compound the problem of lactic acid buildup, when the liver is hypoxic, it fails to clear the lactic acid. Lactic acid levels accumulate, causing lactic acidosis (LA). Therefore, blood lactate is a fairly sensitive and reliable indicator of tissue hypoxia. The hypoxia may be caused by local tissue hypoxia (e.g., mesenteric ischemia, extremity ischemia) or generalized tissue hypoxia (e.g., that which exists in shock). Lactic acid blood levels are used to document the presence of tissue hypoxia, determine the degree of hypoxia, and monitor the effect of therapy.
	// Type I LA is caused by diseases or factors that increase lactate but are not hypoxic related (e.g., glycogen storage diseases, liver diseases, or drugs). LA caused by hypoxia is classified as type II. Shock, convulsions, and extremity ischemia are the most com- mon causes of type II LA. Type III LA is idiopathic and is most commonly seen in nonketotic patients with diabetes. The patho- physiology of lactic acid accumulation in type III is not known.
	// Interfering factors
	// • The prolonged use of a tourniquet or clenching of hands increases lactate levels.
	// Drugs that increase levels include aspirin, cyanide, ethanol (chronic use), nalidixic acid, and phenformin.
	// Procedure and patient care
	// • See inside front cover for Routine Blood Testing.
	// • Fasting: no.
	// • Blood tube commonly used: red.
	// Instruct the patient to avoid making a fist before and while
	// blood is being withdrawn.
	// • Avoid the use of a tourniquet if possible.
	// • Collect a venous blood sample or arterial blood sample in a
	// L
	// lactic acid 569
	//  lactic acid (Lactate)
	// lactic dehydrogenase 571 lactic dehydrogenase (LDH, Lactate dehydrogenase)
	// lactoferrin
	// lactose tolerance test 575 lactose tolerance test (Hydrogen breath test)
	// 578 laparoscopy laparoscopy
	// lead
	// 582 legionnaires disease antibody test legionnaires disease antibody test
	// leucine aminopeptidase (LAP)
	// 584 lipase lipase
	// 586 lipoprotein-associated phospholipase A2 lipoprotein-associated phospholipase A2 (Lp-PLA2
	// PLAC test)
	// lipoproteins 587
	// lipoproteins (High-density lipoproteins [HDLs, HDL-C], Low-density lipoproteins [LDLs, LDL-C], Very-low-density lipoproteins [VLDLs], Lipoprotein electrophoresis, Lipoprotein phenotyping, Lipid fractionation, Non-HDL cholesterol, Lipid profile)
	// 14- to 18-gauge needle Liver
	// Diaphragm
	// liver biopsy 591
	//  liver biopsy
	// 594 liver/spleen scanning
	// liver/spleen scanning (Liver scanning)
	// 596 lumbar puncture and cerebrospinal fluid examination
	//  lumbar puncture and cerebrospinal fluid examination (LP and CSF examination, Spinal tap, Cerebrospinal fluid analysis)
	// 604 lung biopsy lung biopsy
	// lung scan 607 lung scan (Ventilation/perfusion scanning [VPS], Pulmonary
	// scintiphotography, V/Q scan)
	// 610 luteinizing hormone
	// luteinizing hormone (LH assay, Lutropin) and
	// follicle-stimulating hormone (FSH) assay
	// Lyme disease test
	// magnesium
	// magnetic resonance imaging 617 magnetic resonance imaging (MRI, Nuclear magnetic
	// resonance imaging [NMRI])
	// 624 mammography mammography (Mammogram)
	// 628 maternal screen testing
	// maternal screen testing (Maternal triple screen, Maternal
	// quadruple screen)
	// measles rubeola antibody
	// 632 Meckel diverticulum nuclear scan Meckel diverticulum nuclear scan
	// 634 mediastinoscopy mediastinoscopy
	// 636 metanephrine, plasma free
	// metanephrine, plasma free (Fractionated metanephrines)
	// 638 methemoglobin methemoglobin (Hemoglobin M)
	// 640 methylated septin 9 DNA
	// methylated septin 9 DNA (mSEPT9, ColoVantage)
	// microalbumin (MA)
	// microglobulin 643 microglobulin (Beta2-microglobulin [β2m], Alpha-1-microglobulin,
	// Retinol binding protein)
	// mononucleosis rapid test 645 mononucleosis rapid test (Mononuclear heterophil test,
	// Heterophil antibody test, Monospot test)
	// 646 Mycoplasma pneumoniae antibodies, IgG and IgM Mycoplasma pneumoniae antibodies, IgG and IgM
	// myelography (Myelogram)
	// 650 myoglobin myoglobin
	// natriuretic peptides 651
	// natriuretic peptides (Atrial natriuretic peptide [ANP], Brain natriuretic peptide [BNP], N-terminal fragment of pro–brain [B-type] natriuretic peptide [NT–pro-BNP], C-type natriuretic peptide [CNP], Ventricular natriuretic peptide, CHF peptides)
	// neuron-specific enolase (NSE)
	// 654 neutrophil antibody screen
	// neutrophil antibody screen (Granulocyte antibodies, Polymorphonucleocyte antibodies [PMN ab], Antigranulocyte antibodies, Antineutrophil antibodies, Neutrophil antibodies, Leukoagglutinin)
	// neutrophil gelatinase–associated lipocalin 655 neutrophil gelatinase–associated lipocalin (NGAL,
	// Lipocalin-2)
	// newborn metabolic screening 657 newborn metabolic screening
	// nicotine and metabolites 661 nicotine and metabolites (Nicotine, Continine, 3-Hydroxy
	//  cotinine, Nornicotine, Anabasine)
	// 664 5′-nucleotidase 5′-nucleotidase
	// obstruction series
	// octreotide scan 667 octreotide scan (Carcinoid nuclear scan, Neuroendocrine
	// nuclear scan)
	// osmolality, blood 669
	//  osmolality, blood (Serum osmolality)
	// osmolality, urine 671
	//  osmolality, urine (Urine osmolality)
	// oximetry (Pulse oximetry, Oxygen saturation)
	// pancreatic enzymes 675 pancreatic enzymes (Pancreatic secretory test, Amylase, Lipase,
	// Trypsin, Chymotrypsin)
	// 678 pancreatobiliary FISH testing pancreatobiliary FISH testing
	// 680 Papanicolaou smear
	// Papanicolaou smear (Pap smear, Pap test, Cytologic test for
	// cancer, Liquid-based cervical cytology [LBCC], ThinPrep)
	// 684 paracentesis
	// paracentesis (Peritoneal fluid analysis, Ascitic fluid cytology,
	// Peritoneal tap)
	// parathyroid hormone 689 parathyroid hormone (PTH, Parathormone)
	// parathyroid scan 691 parathyroid scan (Parathyroid scintigraphy)
	// partial thromboplastin time, activated 693 partial thromboplastin time, activated (APTT, Partial
	// thromboplastin time [PTT])
	// 696 parvovirus B19 antibody parvovirus B19 antibody
	// pelvic ultrasonography 697 pelvic ultrasonography (Pelvic ultrasonography in pregnancy,
	// Obstetric ultrasonography, Vaginal ultrasound)
	// 700 pepsinogen pepsinogen
	// pericardiocentesis
	// 704 pheochromocytoma suppression and provocative testing pheochromocytoma suppression and provocative
	//  testing (Clonidine suppression test [CST], Glucagon stimulation test)
	// 706 phosphate
	// phosphate (PO4, Phosphorus [P])
	// 708 PI-linked antigen
	// PI-linked antigen (Phosphatidylinositol antigen)
	// placental growth factor (PGF)
	// 710 plasminogen plasminogen (Fibrinolysin)
	// 712 plasminogen activator inhibitor 1 plasminogen activator inhibitor 1 (PAI-1)
	// 714 platelet aggregation test platelet aggregation test
	// 716 platelet antibody detection
	// platelet antibody detection (Antiplatelet antibody
	// detection)
	// 718 platelet count
	// platelet count (Thrombocyte count)
	// platelet function assay 721 platelet function assay (Platelet closure time, PCT, Aspirin
	// resistance tests, Bleeding time [BT], 11-Dehydro-thromboxane B2)
	// 724 platelet volume, mean
	// platelet volume, mean (Mean platelet volume [MPV])
	// plethysmography, arterial 725 plethysmography, arterial (Ankle-brachial index [ABI])
	// pleural biopsy
	// porphyrins and porphobilinogens 729 porphyrins and porphobilinogens
	// positron emission tomography 731 positron emission tomography (PET scan)
	// 736 potassium, blood potassium, blood (K)
	// potassium, urine (K)
	// 740 prealbumin
	// prealbumin (PAB, Thyroxine-binding prealbumin [TBPA],
	// Thyretin, Transthyretin)
	// 742 pregnancy-associated plasma protein-A (PAPP-A) pregnancy-associated plasma protein-A (PAPP-A)
	// 744 pregnanediol pregnanediol
	// 746 progesterone assay progesterone assay
	// 748 progesterone receptor assay
	// progesterone receptor assay (PR assay, PRA, PgR)
	// 750 prolactin levels prolactin levels (PRLs)
	// 752 ProstaScint scan
	// ProstaScint scan (Radioimmunoscintigraphy [RIS])
	// 754 prostate/rectal sonogram
	// prostate/rectal sonogram (Ultrasound prostate)
	// 756 prostate-specific antigen prostate-specific antigen (PSA)
	// 760 protein
	//  protein (Protein electrophoresis, Immunofixation electrophoresis [IFE], Serum protein electrophoresis [SPEP], Albumin, Globulin, Total protein)
	// protein C, protein S
	// prothrombin time 767 prothrombin time (PT, Pro-time, International normalized
	// ratio [INR])
	// pulmonary angiography 771 pulmonary angiography (Pulmonary arteriography)
	// pulmonary function tests 773 pulmonary function tests (PFTs)
	// 778 pyelography
	//  pyelography (Intravenous pyelography [IVP], Excretory urography [EUG], Intravenous urography [IUG, IVU], Retrograde pyelography, Antegrade pyelography)
	// 784 rabies-neutralizing antibody test rabies-neutralizing antibody test
	// red blood cell count 785 red blood cell count (RBC count, Erythrocyte count)
	// 788 red blood cell indices
	// red blood cell indices (RBC indices, Blood indices)
	// 792 renal biopsy
	// renal biopsy (Kidney biopsy)
	// 796 renal scanning
	//  renal scanning (Kidney scan, Radiorenography, Radionuclide renal imaging, Nuclear imaging of the kidney, DSMA renal scan, DTPA renal scan, Captopril renal scan)
	// 800 renin assay, plasma
	// renin assay, plasma (Plasma renin activity [PRA], Plasma renin
	// concentration [PRC])
	// reticulocyte count 805
	//  reticulocyte count (Retic count)
	// rheumatoid factor (RF)
	// ribosome P antibodies 809 ribosome P antibodies (Ribosomal P Ab, Anti–ribosome
	// P antibodies)
	// 810 rubella antibody test
	// rubella antibody test (German measles test)
	// 812 salivary gland nuclear imaging
	// salivary gland nuclear imaging (Parotid gland nuclear
	// imaging)
	// 814 SARS viral testing SARS viral testing
	// 816 scrotal ultrasound
	// scrotal ultrasound (Ultrasound of testes)
	// 818 semen analysis
	// semen analysis (Sperm count, Sperm examination)
	// sentinel lymph node biopsy 821 sentinel lymph node biopsy (SLNB, Lymphoscintigraphy)
	// serotonin (5-hydroxytryptamine, 5-HT) and chromogranin A 823 serotonin (5-hydroxytryptamine, 5-HT) and
	// sexual assault testing
	// 828 sexually transmitted disease testing (STD testing) sexually transmitted disease testing (STD testing)
	// 834 sialography sialography
	// 836 sickle cell screen
	// sickle cell screen (Sickledex, Hemoglobin [Hgb] S test)
	// 838 sigmoidoscopy
	// sigmoidoscopy (Proctoscopy, Anoscopy)
	// 840 Sims-Huhner test
	// Sims-Huhner test (Postcoital test, Postcoital cervical mucus
	// test, Cervical mucus sperm penetration test)
	// 842 skin biopsy
	//  skin biopsy (Cutaneous immunofluorescence biopsy, Skin biopsy antibodies, Skin immunohistopathology, Direct immunofluorescence antibody test)
	// 844 skull x-ray skull x-ray
	// sleep studies (Polysomnography [PSG])
	// small bowel follow-through 849 small bowel follow-through (SBF, Small bowel enema)
	// 852 sodium (Na), blood sodium (Na), blood
	// 854 sodium (Na), urine sodium (Na), urine
	// 856 spinal x-ray
	// spinal x-ray (Cervical, thoracic, lumbar, sacral, and
	// coccygeal x-ray studies)
	// 858 sputum culture and sensitivity
	// sputum culture and sensitivity (C&S, Culture and Gram
	// stain)
	// 860 sputum cytology sputum cytology
	// squamous cell carcinoma antigen 861 squamous cell carcinoma antigen (SCC antigen)
	// 862 stool culture
	// stool culture (Stool for culture and sensitivity [C&S], Stool
	// for ova and parasites [O&P])
	// 864 stool for occult blood
	// stool for occult blood (Stool for OB, Fecal occult blood test
	// [FOBT], Fecal immunotest [FIT], DNA stool sample)
	// Streptococcus serologic testing 867
	//  Streptococcus serologic testing (Antistreptolysin O titer [ASO], Antideoxyribonuclease-B titer, [Anti-DNase-B, ADNase-B, ADB], Streptococcus group B antigen detection, Streptozyme)
	// substance abuse testing 869 substance abuse testing (Urine drug testing, Drug
	// screening, Toxicology screening)
	// swallowing examination 873 swallowing examination (Videofluoroscopy swallowing
	// examination)
	// sweat electrolytes test 875 sweat electrolytes test (lontophoretic sweat test)
	// syphilis detection test (Serologic test for syphilis [STS], Venereal Disease Research Laboratory [VDRL], Rapid plasma reagin [RPR], Fluorescent treponemal antibody test [FTA])
	// testosterone (Dihydrotestosterone [DHT])
	// thoracentesis and pleural fluid analysis 883 thoracentesis and pleural fluid analysis (Pleural tap)
	// thoracoscopy
	// throat and nose cultures
	// thromboelastography 893 thromboelastography (Thromboelastometry)
	// 896 thrombosis indicators
	//  thrombosis indicators
	// fibrin monomers (Fibrin degradation products [FDPs], Fibrin split products [FSPs])
	// fibrinopeptide A (FPA)
	// prothrombin fragment (F1 + 2)
	// 898 thyroglobulin
	// thyroglobulin (Tg, Thyrogen-stimulated thyroglobulin)
	// thyroid scanning (Thyroid scintiscan)
	// 904 thyroid-stimulating hormone thyroid-stimulating hormone (TSH, Thyrotropin,
	// TRH stimulation test)
	// 906 thyroid-stimulating hormone stimulation test thyroid-stimulating hormone stimulation test (TSH
	// stimulation test)
	// thyroid-stimulating immunoglobulins 907
	//  thyroid-stimulating immunoglobulins (TSIs,
	// Long-acting thyroid stimulator [LATS], Thyroid-binding inhibitory immunoglobulin [TBII], Thyrotropin receptor antibody, Thyroid-stimulating hormone receptor [TSHR] antibodies)
	// thyroid ultrasound 909 thyroid ultrasound (Thyroid echogram, Thyroid sonogram)
	// 910 thyrotropin-releasing hormone stimulation test
	//  thyrotropin-releasing hormone stimulation
	// test (TRH stimulation test, Thyrotropin-releasing factor stimulation test [TRF stimulation test])
	// 912 thyroxine, total and free
	// thyroxine, total and free (T4, Thyroxine screen, FT4)
	// thyroxine-binding globulin 915 thyroxine-binding globulin (TBG, Thyroid-binding globulin)
	// 918 total blood volume
	// total blood volume (TBV, Red blood cell [RBC] volume)
	// 920 toxoplasmosis antibody titer toxoplasmosis antibody titer
	// transesophageal echocardiography 921 transesophageal echocardiography (TEE)
	// transferrin receptor (TfR) assay 925 transferrin receptor (TfR) assay
	// triglycerides (TGs)
	// triiodothyronine (T3 radioimmunoassay)
	// troponins (Cardiac-specific troponin T [cTnT], Cardiac-specific troponin I [cTnI])
	// 934 tuberculin skin test
	// tuberculin skin test (TST, Tuberculin test, Mantoux test,
	// Purified protein derivative [PPD] test)
	// 936 tuberculosis culture
	// tuberculosis culture (TB culture, Acid-fast bacillus [AFB] smear)
	// 938 tuberculosis testing
	//  tuberculosis testing (TB testing, Interferon gamma release assay [IGRA], QuantiFERON-TB Gold [QFT, QFT-G, TB Gold test,
	// TB blood test], Nucleic acid amplification for TB [NAAT], TB antibody)
	// upper gastrointestinal x-ray study 941 upper gastrointestinal x-ray study (Upper GI series, UGI)
	// 944 urea breath test
	// urea breath test (UBT, H. pylori breath test)
	// 946 urea nitrogen blood test
	// urea nitrogen blood test (Blood urea nitrogen [BUN])
	// uric acid, blood and urine
	// 952 urinalysis urinalysis (UA)
	// 966 urinary stone analysis
	// urinary stone analysis (Renal calculus analysis)
	// 968 urine culture and sensitivity
	// urine culture and sensitivity (C&S)
	// 970 urine flow studies
	// urine flow studies (Uroflowmetry, Urodynamic studies)
	// 972 uroporphyrinogen-1-synthase uroporphyrinogen-1-synthase
	// vanillylmandelic acid, homovanillic acid, and catecholamines 973
	//  vanillylmandelic acid (VMA), homovanillic acid (HVA), and catecholamines (Epinephrine, Norepinephrine, Metanephrine, Normetanephrine, Dopamine)
	// vascular ultrasound studies 977 vascular ultrasound studies (Venous Doppler ultrasound,
	// Venous duplex scan)
	// 980 venography of lower extremities
	// venography of lower extremities (Phlebography,
	// Venogram)
	// 982 viral cultures viral cultures
	// 984 virus testing virus testing
	// 986 vitamin B12
	// vitaminB12 (Cyanocobalamin)andmethylmalonicacid(MMA)
	// 988 vitamin D
	// vitamin D (25-hydroxy vitamin D2 and D3; 1,25-dihydroxyvitamin D
	// [1,25(OH)2D])
	// white blood cell count and differential count 991 white blood cell count and differential count
	// (WBC and differential, Leukocyte count)
	// 996 white blood cell scan
	// white blood cell scan (WBC scan, Inflammatory scan)
	// 998 wound culture and sensitivity wound culture and sensitivity (C&S)
	// 1000 d-xylose absorption test
	// d-xylose absorption test (Xylose tolerance test)
	// 1002 zinc protoporphyrin zinc protoporphyrin (zpp)
	// R RAST RBC
	// RDW RF RIA RPR
	// S S&A SACE
	// SBF SGOT SGPT SLNB SPA SPECT STS
	// T T3
	// T4 T&C T&S TBG TBPA TEE Tg TGs TIBC TRF TRH TSH TSI TTE
	// U UA
	// UGI series
	// UPP
	// US
	// V VCUG
	// VDRL VLDL VMA VPS
	// W WBC
	// Radioallergosorbent test
	// Red blood cell
	// Red blood cell distribution width
	// Rheumatoid factor
	// Radioimmunoassay
	// Rapid plasma reagin test
	// Sugar and acetone
	// Serum angiotensin-converting enzyme
	// Small bowel follow-through
	// Serum glutamic-oxaloacetic transaminase Serum glutamic-pyruvic transaminase
	// Sentinel lymph node biopsy
	// Sperm penetration assay
	// Single-photon emission computed tomography Serologic test for syphilis
	// Triiodothyronine
	// Thyroxine
	// Type and crossmatch
	// Type and screen
	// Thyroxine-binding globulin Thyroxine-binding prealbumin Transesophageal echocardiography Thyroglobulin
	// Triglycerides
	// Total iron-binding capacity Thyrotropin-releasing factor Thyrotropin-releasing hormone Thyroid-stimulating hormone Thyroid-stimulating immunoglobulins Transthoracic echocardiography
	// Urinalysis
	// Upper gastrointestinal series
	// Urethral pressure profile
	// Ultrasound
	// Voiding cystourethrography
	// Venereal Disease Research Laboratory Very-low-density lipoprotein
	// Vanillylmandelic acid
	// Ventilation/perfusion scanning
	// White blood cell
}
