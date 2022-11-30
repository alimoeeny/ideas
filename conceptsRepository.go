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
	//----------
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_RBC_COUNT)
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

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Spherocyte-over-RBC-count", EnglishHumanReadableExpression: "Spherocyte over RBC count", Short1: "Spherocyte", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Spherocyte-over-RBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Acanthocyte-over-RBC-count", EnglishHumanReadableExpression: "Acanthocyte-over-RBC-count", Short1: "Acanthocyte", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Spherocyte-over-RBC-count", UNIT_COUNT.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Fragment-over-RBC-count", EnglishHumanReadableExpression: "Fragment-over-RBC-count", Short1: "Fragment", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Fragment-over-RBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Platelet-count", EnglishHumanReadableExpression: "Platelet-count", Short1: "PLT", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Platelet-count", UNIT_COUNT.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Reticulocyte-over-RBC-count", EnglishHumanReadableExpression: "Reticulocyte-over-RBC-count", Short1: "Retic", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Reticulocyte-over-RBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_HEMOLOBIN_CONCENTRATION)
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Hemoglobin-concentration")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Hemoglobin-delta-7-days", EnglishHumanReadableExpression: "Hemoglobin-delta-7-days", Short1: "Hb ∆ 7 days", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Hemoglobin-delta-7-days")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Haptoglobin-concentration", EnglishHumanReadableExpression: "Haptoglobin-concentration", Short1: "Haptoglobin", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Haptoglobin-concentration")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Direct-Antiglobulin-Testing", EnglishHumanReadableExpression: "Direct-Antiglobulin-Testing", Short1: "DAT", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Direct-Antiglobulin-Testing")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Bilirubin-serum-concentration", EnglishHumanReadableExpression: "Bilirubin serum concentration", Short1: "Bilirubin", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Bilirubin-serum-concentration")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Mean-Corpuscular-Volume", EnglishHumanReadableExpression: "Mean-Corpuscular-Volume", Short1: "MCV", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Mean-Corpuscular-Volume")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Sickle-Cell-over-RBC-count", EnglishHumanReadableExpression: "Sickle-Cell-over-RBC-count", Short1: "Sickle cell", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Sickle-Cell-over-RBC-count")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Neutrophil-count-over-WBC-count", EnglishHumanReadableExpression: "Neutrophil count over WBC count", Short1: "Neutrophil %", EnglishDescription: "Nuetrophil / WBC"})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Neutrophil-count-over-WBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Neutrophil-count-absolute", EnglishHumanReadableExpression: "Neutrophil count", Short1: "Neutrophil count", EnglishDescription: "Nuetrophil Count"})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Neutrophil-count-absolute", UNIT_COUNT.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Hypogranular-Neutrophils-over-WBC-count", EnglishHumanReadableExpression: "Hypogranular-Neutrophils-over-WBC-count", Short1: "Hypogranular Neutrophils", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Hypogranular-Neutrophils-over-WBC-count")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Hypersegmented-Neurophil-over-WBC-count", EnglishHumanReadableExpression: "Hypersegmented-Neurophil-over-WBC-count", Short1: "Hypersegmented Neurophil", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Hypersegmented-Neurophil-over-WBC-count")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Hypogranular-Platelet-over-RBC-count", EnglishHumanReadableExpression: "Hypogranular-Platelet-over-RBC-count", Short1: "Hypogranular Platelet", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Hypogranular-Platelet-over-RBC-count")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Most-Recent-date-time-of-diagnosis-MDS", EnglishHumanReadableExpression: "Most-Recent-date-time-of-diagnosis-MDS", Short1: "History of MDS", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Most-Recent-date-time-of-diagnosis-MDS", UNIT_TIMESTAMP.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-NRBC-count-over-RBC-count", EnglishHumanReadableExpression: "NRBC-count-over-RBC-count", Short1: "NRBC", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-NRBC-count-over-RBC-count")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Rouleaux-Formation-count-over-RBC-count", EnglishHumanReadableExpression: "Rouleaux-Formation-count-over-RBC-count", Short1: "Rouleaux", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Rouleaux-Formation-count-over-RBC-count", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Metamyelocyte-over-WBC-count", EnglishHumanReadableExpression: "Metamyelocyte-over-WBC-count", Short1: "Metamyelocyte", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Metamyelocyte-over-WBC-count")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Promyelocyte-count-over-WBC-count", EnglishHumanReadableExpression: "Promyelocyte-count-over-WBC-count", Short1: "Promyelocyte", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Promyelocyte-count-over-WBC-count")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Myelocyte-count-over-WBC-count", EnglishHumanReadableExpression: "Myelocyte-over-WBC-count", Short1: "Myelocyte", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Myelocyte-count-over-WBC-count")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Teardrop-RBC-count-over-RBC-count", EnglishHumanReadableExpression: "Teardrop-RBC-over-RBC-count", Short1: "Teardrop RBC", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Teardrop-RBC-count-over-RBC-count")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Most-Recent-date-time-of-diagnosis-PMN", EnglishHumanReadableExpression: "Most-Recent-date-time-of-diagnosis-PMN", Short1: "History of PMN", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Most-Recent-date-time-of-diagnosis-PMN", UNIT_TIMESTAMP.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Monocyte-count-over-RBC-count", EnglishHumanReadableExpression: "Monocyte-count-over-RBC-count", Short1: "Monocyte", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Monocyte-count-over-RBC-count")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Most-Recent-date-time-of-diagnosis-High-Monocytes", EnglishHumanReadableExpression: "Most-Recent-date-time-of-diagnosis-High-Monocytes", Short1: "History of High Monocytes", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Most-Recent-date-time-of-diagnosis-High-Monocytes", UNIT_TIMESTAMP.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Blast-count-over-WBC-count", EnglishHumanReadableExpression: "Blast count over WBC count", Short1: "Blast / WBC", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Blast-count-over-WBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Hypochromic-RBC-over-RBC-count", EnglishHumanReadableExpression: "Hypochromic RBC over RBC count", Short1: "Hypochromic RBC", EnglishDescription: "Hypochromic RBC over RBC count"})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Hypochromic-RBC-over-RBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Teardrop-RBC-over-RBC-count", EnglishHumanReadableExpression: "Teardrop RBC over RBC count", Short1: "Teardrop-RBC-over-RBC-count", EnglishDescription: "Teardrop-RBC-over-RBC-count"})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Teardrop-RBC-over-RBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Blasts-with-vacuoles-count-over-WBC-count", EnglishHumanReadableExpression: "Blasts with vacuoles-count-over-WBC-count", Short1: "Blasts with vacuoles / WBC", EnglishDescription: "SENS-Blasts with vacuoles-count-over-WBC-count"})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Blasts-with-vacuoles-count-over-WBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Basophilic-Metamyelocytes-count-over-WBC-count", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Basophilic Metamyelocytes-count-over-WBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Myeloblast-count-over-WBC-count", EnglishHumanReadableExpression: "Myeloblast count over WBC count", Short1: "Myeloblast / WBC", EnglishDescription: "SENS-Myeloblast-count-over-WBC-count"})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Myeloblast-count-over-WBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Neutrophilic-Metamylocyte-count-over-WBC-count", EnglishHumanReadableExpression: "Neutrophilic Metamylocyte count over WBC count", Short1: "Neutrophilic Metamylocyte / WBC", EnglishDescription: "Neutrophilic Metamylocyte over WBC count"})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Neutrophilic-Metamylocyte-count-over-WBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Monoblast-count-over-WBC-count", EnglishHumanReadableExpression: "Monoblast-count-over-WBC-count", Short1: "Monoblast / WBC", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Monoblast-count-over-WBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Promonocytes-count-over-WBC-count", EnglishHumanReadableExpression: "Promonocytes-count-over-WBC-count", Short1: "Promonocytes / WBC", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Promonocytes-count-over-WBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Auer-Rods", EnglishHumanReadableExpression: "Auer-Rods", Short1: "Auer rods", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Auer-Rods")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Most-Recent-date-time-of-diagnosis-MDS-MPN-CMML", EnglishHumanReadableExpression: "Most-Recent-date-time-of-diagnosis-MDS-MPN-CMML", Short1: "History of MDS/MPN (CMML)", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Most-Recent-date-time-of-diagnosis-MDS-MPN-CMML", UNIT_TIMESTAMP.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Most-Recent-date-time-of-diagnosis-Leukemia", EnglishHumanReadableExpression: "Most-Recent-date-time-of-diagnosis-Leukemia", Short1: "History of Leukemia", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Most-Recent-date-time-of-diagnosis-Leukemia", UNIT_TIMESTAMP.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Serum-concentration-of-Creatinine", EnglishHumanReadableExpression: "Serum-concentration-of-Creatinine", Short1: "Serum Creatinine", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Serum-concentration-of-Creatinine")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Serum-concentration-of-Urea", EnglishHumanReadableExpression: "Serum-concentration-of-Urea", Short1: "Serum Urea", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Serum-concentration-of-Urea")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Serum-concentration-of-Potassium", EnglishHumanReadableExpression: "Serum-concentration-of-Potassium", Short1: "Serum K", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Serum-concentration-of-Potassium")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Serum-concentration-of-Lactate-DeHydrogenase", EnglishHumanReadableExpression: "Serum-concentration-of-Lactate-DeHydrogenase", Short1: "Serum LDH", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Serum-concentration-of-Lactate-DeHydrogenase")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Prothrombin-Time", EnglishHumanReadableExpression: "Prothrombin-Time", Short1: "PT", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Prothrombin-Time")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-activated-partial-thromboplastin-time", EnglishHumanReadableExpression: "activated-partial-thromboplastin-time", Short1: "aPTT", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-activated-partial-thromboplastin-time")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Toxic-Changes-in-Neutrophils", EnglishHumanReadableExpression: "Toxic-Changes-in-Neutrophils", Short1: "Toxic changes", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Toxic-Changes-in-Neutrophils", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Most-Recent-date-time-of-Chemotherapy", EnglishHumanReadableExpression: "Most-Recent-date-time-of-Chemotherapy", Short1: "History of recent chemotherapy", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Most-Recent-date-time-of-Chemotherapy", UNIT_TIMESTAMP.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Most-Recent-date-time-of-GCSF-infusion", EnglishHumanReadableExpression: "Most-Recent-date-time-of-GCSF-infusion", Short1: "History of recent GCSF", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Most-Recent-date-time-of-GCSF-infusion", UNIT_TIMESTAMP.ID, UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Target-Cell-count-over-RBC-count", EnglishHumanReadableExpression: "Target cell count over RBC count", Short1: "TRG / RBC", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Target-Cell-count-over-RBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Serum-Ferritin-level", EnglishHumanReadableExpression: "Serum Ferritin Level", Short1: "Serum Ferritin", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Serum-Ferritin-level", UNIT_MICROGRAM_PER_LITER.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Serum-Iron-level", EnglishHumanReadableExpression: "Serum Iron Level", Short1: "Serum Iron", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Serum-Iron-level", UNIT_MICROGRAM_PER_DECILITER.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Total-Iron-Binding-Capacity", EnglishHumanReadableExpression: "Total iron-binding capacity", Short1: "TIBC", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Total-Iron-Binding-Capacity", UNIT_MICROGRAM_PER_DECILITER.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Blood-Oxygen-Saturation", EnglishHumanReadableExpression: "Blood Oxygen Saturation", Short1: "O2 Sat", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Blood-Oxygen-Saturation", UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Red-Cell-Distribution-Width", EnglishHumanReadableExpression: "Red Cell Distribution Width", Short1: "RDW", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Red-Cell-Distribution-Width", UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Elevated-LFTs", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Elevated-LFTs", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-History-of-Splenectomy", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-History-of-Splenectomy", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Basophilic-Stippling-count-over-RBC-count", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Basophilic-Stippling-count-over-RBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Howell-Jolly-Bodies-count-over-RBC-count", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Howell-Jolly-Bodies-count-over-RBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Alanine-Transaminase-blood-concentration", EnglishHumanReadableExpression: "ALT", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Alanine-Transaminase-blood-concentration", UNIT_INTERNATIONAL_UNITS_PER_LITER.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Aspartate-Transaminase-blood-concentration", EnglishHumanReadableExpression: "AST", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Aspartate-Transaminase-blood-concentration", UNIT_INTERNATIONAL_UNITS_PER_LITER.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Alkaline-Phosphatase", EnglishHumanReadableExpression: "ALP", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-alkaline phosphatase", UNIT_INTERNATIONAL_UNITS_PER_LITER.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Gamma-Glutamyl-Transferase", EnglishHumanReadableExpression: "GGT", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-gamma-glutamyl transferase ", UNIT_INTERNATIONAL_UNITS_PER_LITER.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-international normalized ratio", EnglishHumanReadableExpression: "INR", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-international normalized ratio", UNIT_RATIO.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Albumin-Serum-Concentration", EnglishHumanReadableExpression: "Albumin", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Albumin-serum-concentration", UNIT_GRAMS_PER_LITER.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Pappenheimer-Body-count-over-RBC-count", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Pappenheimer-Body-count-over-RBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-History-of-Hemoglobinopathy", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-History-of-Hemoglobinopathy", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Patient-Age", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Patient-Age", UNIT_DAY.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Atypical-Lymphocytes-count-over-WBC-count", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Atypical-Lymphocytes-count-over-WBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Mononucleosis-Spot-Test", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Mononucleosis-Spot-Test", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Accelerating-Lymphocytosis", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Accelerating-Lymphocytosis", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Smudge-Cell-count-over-WBC-count", EnglishHumanReadableExpression: "Smudge Cells over WBC", Short1: "Smudge Cells / WBC", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Smudge-Cell-count-over-WBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-History-of-Lymphoproliferative-Disorders", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-History-of-Lymphoproliferative-Disorders ", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Circulating-Plasma-Cell-count-over-WBC-count", EnglishHumanReadableExpression: "Circulating Plasma Cells over WBC count", Short1: "Circulating Plasma Cells / WBC", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Circulating-Plasma-Cell-count-over-WBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	// --------------------------------------------------

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Chest-pain", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Severe-lower chest-pain", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-upper-abdominal-pain", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Sudden-breathlessness, onset over seconds", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Breathlessness on minimal exertion", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-orthopnoea", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-paroxysmal nocturnal dyspnoea (PND)", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Palpitations", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Acute breathlessness", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-wheeze ± cough", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Cough and pink frothy sputum", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Syncope", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Leg pain on walking—intermittent claudication", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Leg pain on standing—relieved by lying down", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Unilateral calf swelling", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Bilateral ankle swelling", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Thoughts on interpreting cardiovascular signs", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Peripheral cyanosis", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Central cyanosis", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Tachycardia (e.g. pulse rate >20bpm) 200 Bradycardia (<60bpm) 202", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Pulse irregular", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Pulse volume high", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Pulse volume low", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Blood pressure high—hypertension", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Blood pressure very low", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Postural fall in blood pressure", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-BP/pulse difference between arms", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-BP/pulse difference between arm and legs", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Prominent leg veins ± unilateral leg swelling", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Unilateral leg and ankle swelling", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Bilateral leg and ankle swelling", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Raised jugular venous pressure", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Abnormal apex impulse", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Extra heart sounds", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Diastolic murmur", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Mid-systolic murmur", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Pansystolic murmur", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Murmurs not entirely in systole or diastole", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})

	// --------------------------------------------------

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Blast-Over20p", EnglishHumanReadableExpression: "Blast / WBC >= 20%", Short1: "Blast >= 20%", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Blast-Over20p")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Blast-Below20p", EnglishHumanReadableExpression: "Blast / WBC < 20%", Short1: "Blast < 20%", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Blast-Below20p")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Monoblast-Over20p", EnglishHumanReadableExpression: "Monoblast / WBC >= 20%", Short1: "Moboblast >= 20%", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Monoblast-Over20p")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Monoblast-Below20p", EnglishHumanReadableExpression: "Moboblast / WBC < 20%", Short1: "Monoblast < 20%", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Monoblast-Below20p")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Promonocytes-Over20p", EnglishHumanReadableExpression: "Promonocytes / WBC >= 20%", Short1: "Promonocytes >= 20%", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Promonocytes-Over20p")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Promonocytes-Below20p", EnglishHumanReadableExpression: "Promonocytes / WBC < 20%", Short1: "Promonocytes < 20%", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Promonocytes-Below20p")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Spherocytosis", EnglishHumanReadableExpression: "Spherocytosis", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Spherocytosis")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-RBCFragments", EnglishHumanReadableExpression: "RBC Fragments", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-RBCFragments")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Low-Platelet-count-Adult-150000-µl", EnglishHumanReadableExpression: "Thrombocytopenia", Short1: "Thrombocytopenia", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Low-Platelet-count-Adult-150000-µl")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Prolonged-Prothrombin-Time", EnglishHumanReadableExpression: "Prolonged Prothrombin Time", Short1: "Prolonged PT", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Prolonged-Prothrombin-Time")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Prolonged-activated-partial-thromboplastin-time", EnglishHumanReadableExpression: "Prolonged activated Partial Thromboplastin Time", Short1: "Prolonged aPTT", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Prolonged-activated-partial-thromboplastin-time")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Retic", EnglishHumanReadableExpression: "Reticulocyte", Short1: "Retic", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Retic")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Rouleaux-Formation", EnglishHumanReadableExpression: "Rouleaux Formation", Short1: "Rouleaux", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Rouleaux-Formation", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Circulating-Plasma-Cells-Over20p", EnglishHumanReadableExpression: "Circulating Plasma Cells over 20 percent", Short1: "Circulating Plasma Cells Over 20 percent", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Circulating-Plasma-Cells-Over20p")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Positive-History-of-Acute-Leukemia", EnglishHumanReadableExpression: "Positive History of Acute Leukemia", Short1: "History of Acute Leukemia +", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Positive-History-of-Acute-Leukemia")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Promyelocyte-Detected", EnglishHumanReadableExpression: "Promyelocyte Detected", Short1: "Promyelocyte Detected", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Promyelocyte-Detected")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Positive-History-of-recent-chemotherapy", EnglishHumanReadableExpression: "Positive history of recent chemotherapy", Short1: "Recent Chemotherapy  +", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Positive-History-of-recent-chemotherapy")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Positive-History-of-recent-GCSF", EnglishHumanReadableExpression: "Positive history of recent GCSF", Short1: "Recent GCSF  +", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Positive-History-of-recent-GCSF")

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Toxic-Changes-in-Neutrophils-Detected", EnglishHumanReadableExpression: "Toxic Changes in Neutrophils Detected", Short1: "Toxic changes +", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Toxic-Changes-in-Neutrophils-Detected", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-High-Serum-concentration-of-Lactate-DeHydrogenase", EnglishHumanReadableExpression: "High Serum concentration of Lactate DeHydrogenase", Short1: "High Serum LDH", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-High-Serum-concentration-of-Lactate-DeHydrogenase", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-High-Serum-concentration-of-Creatinine", EnglishHumanReadableExpression: "High Serum concentration of Creatinine", Short1: "High Serum Creatinine", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-High-Serum-concentration-of-Creatinine", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-High-Serum-concentration-of-Urea", EnglishHumanReadableExpression: "High Serum concentration of Urea", Short1: "High Serum Urea", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-High-Serum-concentration-of-Urea", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-High-Serum-concentration-of-Potassium", EnglishHumanReadableExpression: "High Serum concentration of Potassium", Short1: "High Serum K", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-High-Serum-concentration-of-Potassium", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Hypogranular-Neutrophils-detected", EnglishHumanReadableExpression: "Hypogranular Neutrophils detected", Short1: "Hypogranular Neutrophils +", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Hypogranular-Neutrophils-detected", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Hypersegmented-Neurophil-detected", EnglishHumanReadableExpression: "Hypersegmented Neurophil detected", Short1: "Hypersegmented Neurophils +", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Hypersegmented-Neurophil-detected", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Hypogranular-Platelet-detected", EnglishHumanReadableExpression: "Hypogranular Platelet detected", Short1: "Hypogranular Platelets +", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Hypogranular-Platelet-detected", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Positive-History-of-MDS", EnglishHumanReadableExpression: "Positive History of MDS", Short1: "History of MDS +", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Positive-History-of-MDS", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Positive-History-of-MDS-MPN-CMML", EnglishHumanReadableExpression: "Positive History of MDS-MPN-CMML", Short1: "History of MDS-MPN-CMML +", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Positive-History-of-MDS-MPN-CMML", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-NRBC-count-elevated", EnglishHumanReadableExpression: "Elevated NRBC count", Short1: "Elevated NRBC", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-NRBC-count-elevated", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Elevated-Metamyelocyte-over-WBC-count", EnglishHumanReadableExpression: "Elevated Metamyelocyte count", Short1: "Elevated Metamyelocyte", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Elevated-Metamyelocyte-over-WBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Elevated-Promyelocyte-count-over-WBC-count", EnglishHumanReadableExpression: "Elevated Promyelocyte count", Short1: "Elevated Promyelocyte", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Elevated-Promyelocyte-count-over-WBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Elevated-Myelocyte-count-over-WBC-count", EnglishHumanReadableExpression: "Elevated Myelocyte count", Short1: "Elevated Myelocyte", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Elevated-Myelocyte-count-over-WBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Elevated-Teardrop-RBC-count-over-RBC-count", EnglishHumanReadableExpression: "Elevated Teardrop count", Short1: "Elevated Teardrop RBC", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Elevated-Teardrop-RBC-count-over-RBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-MAHA", EnglishHumanReadableExpression: "MAHA", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("DDX-MAHA", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-IHA", EnglishHumanReadableExpression: "'IHA", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("DDX-IHA", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-HS", EnglishHumanReadableExpression: "HS", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("DDX-HS", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-AML", EnglishHumanReadableExpression: "Acute Myeloid Leukemia", Short1: "AML", EnglishDescription: "Acute Myeloid Leukemia"})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-ALL", EnglishHumanReadableExpression: "Acute Lymphocytic Leukemia", Short1: "ALL", EnglishDescription: "Acute Lymphocytic Leukemia"})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-CML", EnglishHumanReadableExpression: "Chronic Myeloid Leukemia", Short1: "CML", EnglishDescription: "Chronic Myeloid Leukemia"})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-JMML", EnglishHumanReadableExpression: "Juvenile Myelomonocytic Leukemia", Short1: "JMML", EnglishDescription: "Juvenile Myelomonocytic Leukemia"})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-CLL", EnglishHumanReadableExpression: "Chronic Lymphocytic Leukemia (CLL)", Short1: "CLL", EnglishDescription: "Chronic Lymphocytic Leukemia (CLL)"})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-CBL", EnglishHumanReadableExpression: "Chronic Basophilic Leukemia (CBL)", Short1: "CBL", EnglishDescription: "Chronic Basophilic Leukemia (CBL)"})

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Hypochromic-Anemia", EnglishHumanReadableExpression: "Hypochromic Anemia", Short1: "Hypochromic Anemia", EnglishDescription: "Hypochromic Anemia"})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Thrombocytopenia", EnglishHumanReadableExpression: "Thrombocytopenia", Short1: "Thrombocytopenia", EnglishDescription: "Thrombocytopenia"})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Neutropenia", EnglishHumanReadableExpression: "Neutropenia", Short1: "Neutropenia", EnglishDescription: "Neutropenia"})

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Plasma-Cell-Leukemia", EnglishHumanReadableExpression: "Plasma Cell Leukemia", Short1: "Plasma Cell Leukemia", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("DDX-Plasma-Cell-Leukemia", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Multiple-Myeloma", EnglishHumanReadableExpression: "Multiple Myeloma", Short1: "Multiple Myeloma", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("DDX-Multiple-Myeloma", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Acute-Leukemia", EnglishHumanReadableExpression: "Acute Leukemia", Short1: "Acute Leukemia", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("DDX-Acute-Leukemia", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Acute-Promyelocytic-Leukemia", EnglishHumanReadableExpression: "Acute Promyelocytic Leukemia", Short1: "Acute Promyelocytic Leukemia", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("DDX-Acute-Promyelocytic-Leukemia", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Myeloproliferative-Neoplasms", EnglishHumanReadableExpression: "Myeloproliferative neoplasms", Short1: "MPN", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("DDX-Myeloproliferative-Neoplasms", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Myelodysplastic-syndromes", EnglishHumanReadableExpression: "Myelodysplastic syndromes", Short1: "MDS", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("DDX-Myelodysplastic-syndromes", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Myelodysplastic-syndrome-myeloproliferative-neoplasm-overlap-syndromes", EnglishHumanReadableExpression: "MDS/MPN (CMML) Overlap syndroms", Short1: "MDS/MPN (CMML)", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("DDX-Myelodysplastic-syndrome-myeloproliferative-neoplasm-overlap-syndromes", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-History-of-recent-chemotherapy", EnglishHumanReadableExpression: "History of recent chemotherapy", Short1: "Recent Chemotherapy", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("DDX-History-of-recent-chemotherapy", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-History-of-recent-GCSF", EnglishHumanReadableExpression: "History of recent GCSF", Short1: "Recent GCSF", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("DDX-History-of-recent-GCSF", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Iron-Deficiency-Anemia", EnglishHumanReadableExpression: "Iron Deficiency Anemia", Short1: "Iron Deficiency Anemia", EnglishDescription: ""})

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Monoclonal-B-Lymphocytosis", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Sideroblastic-Anemia", EnglishHumanReadableExpression: "Sideroblastic anemia", Short1: "", EnglishDescription: ""})

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Angina", EnglishHumanReadableExpression: "Angina", Short1: "Angina", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-STEMI", EnglishHumanReadableExpression: "ST Elevated Myocardial Infarction", Short1: "STEMI", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-NSTEMI", EnglishHumanReadableExpression: "Non-ST Elevation Myocardial Infaction", Short1: "NSTEMI", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Esophagitis", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Esophagial-Spasm", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Pulmunary-Embolus", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Pulmunary-Infarction", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Atrial-Fibrilation", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Cardiac-Arythmia"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Pneumothorax", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Tension-Pneumothorax", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Dissecting-Thoracic-Aortic-Aneurism", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Chest-Wall-Pain", EnglishHumanReadableExpression: "Chest Wall Pain", Short1: "", EnglishDescription: "", AKA: []string{"DDX-Tietzes-Syndrom"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Tietzes-Syndrom", EnglishHumanReadableExpression: "Tietze's Syndrom", Short1: "", EnglishDescription: "", AKA: []string{"DDX-Chest-Wall-Pain"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Gastro-Esophageal-Reflux", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Gastritis", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Billiary-Colic", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Pancreatitis", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Inferior-Myocardial-Infarction", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Anaphylaxis", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Inhalation-of-Foreign-Body", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Pulmunary-Edema", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-COPD", EnglishHumanReadableExpression: "Chronic Obstructive Pulmonary Disease", Short1: "COPD", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Asthma", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Cardiac-Arythmia", EnglishHumanReadableExpression: "Cardiac Arythmia", Short1: "Arythmia", EnglishDescription: "", AKA: []string{}, Includes: []string{"DDX-Atrial-Fibrilation", "DDX-PSVT"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Eisodic-Heart-Block", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Cardiac-Arythmia"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Sinus-Tachycardia", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Cardiac-Arythmia"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-PSVT", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Cardiac-Arythmia"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-SVT", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Cardiac-Arythmia"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-VTAC", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Cardiac-Arythmia"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Ventricular-Fibrilation", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Cardiac-Arythmia"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Anxiety", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Thyrotoxicosis", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Hyperventilation", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Hypovolamia", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Infection", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Menopause", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Pheochromocytoma", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-PVC", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Cardiac-Arythmia"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-POTS-Syndrom", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Paraganglioma", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Exacerbation-of-Asthma", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Exacerbation-of-COPD", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Acute-Bronchitis", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}, Includes: []string{"DDX-Acute-Bacterial-Bronchitis", "DDX-Acute-Viral-Bronchitis"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Acute-Viral-Bronchitis", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Acute-Bronchitis"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Acute-Bacterial-Bronchitis", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Acute-Bronchitis"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Acute-Left-Ventricular-Failure", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Heart-Failure"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Heart-Failure", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}, Includes: []string{"DDX-Acute-Left-Ventricular-Failure"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Valvular-Disease", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}, Includes: []string{"DDX-Mitral-Stenosis"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Mitral-Stenosis", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}, IsIncludedIn: []string{"DDX-Valvular-Disease"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Vasovagal-Attack", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Postural-Hypotension", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Aortic-Stenosis", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Hyperthrophic-Cardiomyopathy", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Micturation-Syncope", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Cough-Syncope", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Carotid-Sinus-Syncope", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Hypoglycemia", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Epilepsy", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-CVA", EnglishHumanReadableExpression: "Cerebrovascular Accident", Short1: "CVA", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Arterial-disease-in-legs", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Spinal-Claudication", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Peripheral-Venous-Disease", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Spinal-Disk-Protrusion", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Deep-Venous-Thrombosis", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Lower-Exterimity-Varicoses", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Deep-Venous-Thrombosis", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Ruptured-Bakers-Cyst", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Cellulitis", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Abnormal-Lymphatic-Drainage", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Lymphoma", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Right-Ventricular-Failure", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Pulmonary-Hypertension", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Congestive-Heart-Failure", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Poor-Venous-return-due-to-abdominal-pelvic-mass", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Low-albumin-state", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Liver-Failure", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Nephrotic-Syndrom", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Malnutrition", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Malabsorption", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-IVC-Obstruction-due-to-clot", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Bilateral-Venous-thrombosis-of-lower-extremities", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Septicemia", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Valular-stenosis", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Valvular-incompetence", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Hemorhage", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Internal-Hemorhage", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-GI-Bleeding", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Raynauds-Phenomenon", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Peripheral-Arterial-Obstruction", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Diabetic-small-vessle-disease", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Low-Cardiac-output", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-MI", EnglishHumanReadableExpression: "Myocardial Infarction", Short1: "MI", EnglishDescription: "", AKA: []string{}, Includes: []string{"DDX-STEMI", "DDX-NSTEMI", "DDX-Inferior-Myocardial-Infarction"}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Right-To-Left-Cardiac-Shunt", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Tetralogy-of-Fallot", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Tricuspid-Atresia", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Pulmonary-AV-Fistula", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Right-To-Left-Pulmonary-shunt", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Hemoglobin-Abnormalities", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Hemoglobin-M-disease", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-NADH-Diaphorase", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Methemoglobinemia", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})
	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "DDX-Sulfhemoglobinemia", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: "", AKA: []string{}})

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
