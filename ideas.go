package ideas

// Concept represents a bit of knowledge that can be queried form the user or from the machines
// an example concept is RBC count, an workflow can ask for RBC count from an algorith or from the user
type Concept struct {
	id                             int64
	englishHumanReadableExpression string
	englishDescription             string
}

func (c Concept) String() string {
	return c.englishHumanReadableExpression
}

// ConceptSet represents a group of related Concepts
type ConceptSet struct {
	id       int64
	concepts []*Concept
}

// MutuallyExclusiveConceptSet represents a group of Concepts where only one of them can be true at a time
// like macrocytosis, microcytosis and normocytosis
type MutuallyExclusiveConceptSet struct {
	ConceptSet
}

// Idea represents an instant of one or more Concepts being realized
// for example when we are talking about someones RBC count,
// basically we are associating the concept of RBC count with the measurment of their rbcs at a particular time
// like the Idea is this patient has an RBC count of 4.8M /µl at this timestamp
type Idea struct {
	id                             int64
	englishHumanReadableExpression string
	facts                          map[*Concept]Measurment
}

func (idea Idea) String() string {
	s := ""
	for concept, measurement := range idea.facts {
		s += concept.String() + " -> " + measurement.String() + "\n"
	}
	return s
}

// IdeaSet is a set of related Ideas
type IdeaSet struct {
	id    int64
	ideas []*Idea
}

func (ids IdeaSet) String() string {
	s := ""
	for _, id := range ids.ideas {
		s += id.String() + "\n"
	}
	return s
}
