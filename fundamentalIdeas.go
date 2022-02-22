package ideas

import "time"

var CONCEPT_NOT_DETECTED = Concept{
	ID:                             "7788",
	EnglishHumanReadableExpression: "not detected",
	EnglishDescription:             "indicates that the observation or measurement did not find anything",
}

var CONCEPT_DETECTED = Concept{
	ID:                             "7789",
	EnglishHumanReadableExpression: "detected",
	EnglishDescription:             "indicates that the observation or measurement did detect what it was meant to detect",
}

var CONCEPT_DATA_NOT_AVAILABLE = Concept{
	ID:                             "7790",
	EnglishHumanReadableExpression: "data not available",
	EnglishDescription:             "indicates that the observation or measurement result will not be available or accissible and we need to do without them",
}

var CONCEPT_DATA_NEEDED = Concept{
	ID:                             "7791",
	EnglishHumanReadableExpression: "data needed",
	EnglishDescription:             "indicates that the observation or measurement result is missing or not yet ready and we may want to wait for it to become available",
}

var CONCEPT_APPLICABLE_HERE = Concept{
	ID:                             "7792",
	EnglishHumanReadableExpression: "applicable here",
	EnglishDescription:             "indicates that the observation or measurement is applicable here",
}

var CONCEPT_NOT_APPLICABLE_HERE = Concept{
	ID:                             "7793",
	EnglishHumanReadableExpression: "not applicable here",
	EnglishDescription:             "indicates that the observation or measurement is not applicable here",
}

func MeasurementIsValueNow(value interface{}) *Measurement {
	return &Measurement{ID: newStrID(), Timestamp: time.Now().UnixNano(), Value: value}
}
