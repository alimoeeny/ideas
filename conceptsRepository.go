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
	DEFAULT_CONCEPT_REPO.SetConcept(CONCEPT_HEMOLOBIN_CONCENTRATION)
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

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Hemoglobin-concentration", EnglishHumanReadableExpression: "Hemoglobin-concentration", Short1: "Hb", EnglishDescription: ""})
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

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-Smudge-Cell-count-over-WBC-count", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-Smudge-Cell-count-over-WBC-count", UNIT_RATIO.ID, UNIT_PERCENTAGE.ID)

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "SENS-History-of-Lymphoproliferative-Disorders ", EnglishHumanReadableExpression: "", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("SENS-History-of-Lymphoproliferative-Disorders ", UNIT_POSITIVE_NEGATIVE_MISSING.ID)

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

	DEFAULT_CONCEPT_REPO.SetConcept(Concept{ID: "FACT-Retic", EnglishHumanReadableExpression: "Reticulocyte", Short1: "", EnglishDescription: ""})
	DEFAULT_CONCEPT_REPO.SetAvailableUnitsForConceptWithID("FACT-Retic")

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

}
