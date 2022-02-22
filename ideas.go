package ideas

func NewConcept(id string, expression string, description string) Concept {
	return Concept{id, expression, description}
}

// Concept represents a bit of knowledge that can be queried form the user or from the machines
// an example concept is RBC count, a workflow can ask for RBC count from an algorithm or from the user
type Concept struct {
	ID                             string `json:"id,omitempty"`
	EnglishHumanReadableExpression string `json:"english_human_readable_expression,omitempty"`
	EnglishDescription             string `json:"english_description,omitempty"`
}

// func (c Concept) MarshalText() (text []byte, err error) {
// 	key := c.ID
// 	if key == "" {
// 		key += c.EnglishHumanReadableExpression + c.EnglishDescription
// 	}
// 	return []byte(key), nil
// }

func (c Concept) String() string {
	return c.EnglishHumanReadableExpression
}

// ConceptSet represents a group of related Concepts
type ConceptSet struct {
	ID                             string    `json:"id,omitempty"`
	EnglishHumanReadableExpression string    `json:"english_human_readable_expression,omitempty"`
	EnglishDescription             string    `json:"english_description,omitempty"`
	Concepts                       []Concept `json:"concepts,omitempty"`
}

// MutuallyExclusiveConceptSet represents a group of Concepts where only one of them can be true at a time
// like macrocytosis, microcytosis and normocytosis
type MutuallyExclusiveConceptSet struct {
	ConceptSet
}

func NewIdea(id string, description string, facts map[string]*Measurement) *Idea {
	if id == "" {
		id = newStrID()
	}
	return &Idea{id, description, facts}
}

// Idea represents an instant of one or more Concepts being realized
// for example when we are talking about someones RBC count at a specific time,
// basically we are associating the concept of RBC count with the measurement of their rbcs at a particular time
// like the Idea is this patient has an RBC count of 4.8M /µl at this timestamp
// Facts are maps of concept ids to measurments, the user is responsible to keep track of the concepts ids and to make sure they are unique
type Idea struct {
	ID                             string                  `json:"id,omitempty"`
	EnglishHumanReadableExpression string                  `json:"english_human_readable_expression,omitempty"`
	Facts                          map[string]*Measurement `json:"facts,omitempty"`
}

func (idea Idea) FactCheck(conceptID string) *Measurement {
	return idea.Facts[conceptID]
}

// func (idea Idea) String() string {
// 	s := ""
// 	for concept, measurement := range idea.Facts {
// 		s += concept.String() + " -> " + measurement.String()
// 	}
// 	return s
// }

// IdeaSet is a set of related Ideas
type IdeaSet struct {
	ID    string  `json:"id,omitempty"`
	Ideas []*Idea `json:"ideas,omitempty"`
}

// func (ids IdeaSet) String() string {
// 	s := ""
// 	for _, id := range ids.Ideas {
// 		s += id.String() + ", "
// 	}
// 	return s
// }

func NewIdeaSet(id string, ideas []*Idea) IdeaSet {
	return IdeaSet{id, ideas}
}
