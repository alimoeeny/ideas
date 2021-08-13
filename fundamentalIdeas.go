package ideas

var CONCEPT_NOT_DETECTED = &Concept{
	id:                             "7788",
	englishHumanReadableExpression: "not detected",
	englishDescription:             "indicates that the observation or measurement did not find anything",
}

var CONCEPT_DETECTED = &Concept{
	id:                             "7789",
	englishHumanReadableExpression: "detected",
	englishDescription:             "indicates that the observation or measurement did detect what it was meant to detect",
}

var CONCEPT_DATA_NOT_AVAILABLE = &Concept{
	id:                             "7790",
	englishHumanReadableExpression: "data not available",
	englishDescription:             "indicates that the observation or measurement result will not be available or accissible and we need to do without them",
}

var CONCEPT_DATA_NEEDED = &Concept{
	id:                             "7791",
	englishHumanReadableExpression: "data needed",
	englishDescription:             "indicates that the observation or measurement result is missing or not yet ready and we may want to wait for it to become available",
}

var CONCEPT_APPLICABLE_HERE = &Concept{
	id:                             "7792",
	englishHumanReadableExpression: "applicable here",
	englishDescription:             "indicates that the observation or measurement is applicable here",
}

var CONCEPT_NOT_APPLICABLE_HERE = &Concept{
	id:                             "7793",
	englishHumanReadableExpression: "not applicable here",
	englishDescription:             "indicates that the observation or measurement is not applicable here",
}
