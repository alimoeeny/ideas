package ideas

func NewConcept(id string, expression string, description string) Concept {
	return Concept{id, expression, description}
}

// Concept represents a bit of knowledge that can be queried form the user or from the machines
// an example concept is RBC count, a workflow can ask for RBC count from an algorithm or from the user
type Concept struct {
	id                             string
	englishHumanReadableExpression string
	englishDescription             string
}

func (c Concept) String() string {
	return c.englishHumanReadableExpression
}

// ConceptSet represents a group of related Concepts
type ConceptSet struct {
	ID                             string
	englishHumanReadableExpression string
	englishDescription             string
	concepts                       []Concept
}

// MutuallyExclusiveConceptSet represents a group of Concepts where only one of them can be true at a time
// like macrocytosis, microcytosis and normocytosis
type MutuallyExclusiveConceptSet struct {
	ConceptSet
}

func NewIdea(description string, facts map[Concept]*Measurement) Idea {
	return Idea{newStrID(), description, facts}
}

// Idea represents an instant of one or more Concepts being realized
// for example when we are talking about someones RBC count at a specific time,
// basically we are associating the concept of RBC count with the measurement of their rbcs at a particular time
// like the Idea is this patient has an RBC count of 4.8M /µl at this timestamp
type Idea struct {
	id                             string
	englishHumanReadableExpression string
	facts                          map[Concept]*Measurement
}

func (idea Idea) FactCheck(concept Concept) *Measurement {
	return idea.facts[concept]
}

func (idea Idea) Facts() map[Concept]*Measurement {
	return idea.facts
}

func (idea Idea) String() string {
	s := ""
	for concept, measurement := range idea.facts {
		s += concept.String() + " -> " + measurement.String()
	}
	return s
}

// IdeaSet is a set of related Ideas
type IdeaSet struct {
	id    string
	ideas []*Idea
}

func (ids IdeaSet) ID() string {
	return ids.id
}

func (ids IdeaSet) String() string {
	s := ""
	for _, id := range ids.ideas {
		s += id.String() + ", "
	}
	return s
}
