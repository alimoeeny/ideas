package ideas

import "fmt"

func (repo *ConceptsRepository) ConceptWithID(id string) *Concept {
	return (*repo).conceptsDict[id]
}

func (repo *ConceptsRepository) AcceptableUnitsForConceptWithID(id string) []Unit {
	units := []Unit{}
	uids := repo.conceptsToUnitsMap[id]
	for _, uid := range uids {
		units = append(units, *repo.unitsRepo.UnitWithID(uid))
	}
	return units
}

func (repo *ConceptsRepository) SetConcept(c Concept) error {
	if c.ID == "" {
		return fmt.Errorf("concept needs to have an non-empty id")
	}
	if (*repo).conceptsDict[c.ID] != nil {
		return fmt.Errorf("another concept with the same ID already exists")
	}
	// let it crash if they pass in a nil repo, it is danguarous to do anything else
	(*repo).conceptsDict[c.ID] = &c
	return nil
}

func (repo *ConceptsRepository) SetAvailableUnitsForConceptWithID(conceptID string, unitIDs ...string) error {
	if repo.conceptsToUnitsMap == nil {
		return fmt.Errorf("this repo is not initialized correctly")
	}
	if conceptID == "" {
		return fmt.Errorf("concept id cannot be empty or blank")
	}
	repo.conceptsToUnitsMap[conceptID] = uniqueArray(append(repo.conceptsToUnitsMap[conceptID], unitIDs...))
	return nil
}

func (repo *ConceptsRepository) All() map[string]*Concept {
	return repo.conceptsDict
}

// AllConceptsWithUnitIDs
// returns an arrayc of structs containing pairs of concepts along with arrays of ids for their units
func (repo *ConceptsRepository) AllConceptsWithUnitIDs() map[string]struct {
	Concept *Concept `json:"concept,omitempty"`
	UnitIDs []string `json:"unit_ids,omitempty"`
} {
	r := map[string]struct {
		Concept *Concept `json:"concept,omitempty"`
		UnitIDs []string `json:"unit_ids,omitempty"`
	}{}
	for _, c := range repo.conceptsDict {
		// units := []Unit{}
		// for _, unitID := range(repo.conceptsToUnitsMap[c.ID]) {
		// 	units = append(units, *repo.unitsRepo.UnitWithID(unitID))
		// }
		r[c.ID] = struct {
			Concept *Concept `json:"concept,omitempty"`
			UnitIDs []string `json:"unit_ids,omitempty"`
		}{c, repo.conceptsToUnitsMap[c.ID]}
	}
	return r
}

func (repo *ConceptsRepository) NewConcept(id string, expression string, short1 string, description string) (Concept, error) {
	if id == "" {
		id = newStrID()
	}
	c := Concept{id, expression, short1, description}
	err := repo.SetConcept(c)
	return c, err
}

// Concept represents a bit of knowledge that can be queried form the user or from the machines
// an example concept is RBC count, a workflow can ask for RBC count from an algorithm or from the user
type Concept struct {
	ID                             string `json:"id,omitempty"`
	EnglishHumanReadableExpression string `json:"english_human_readable_expression,omitempty"`
	Short1                         string `json:"short_1,omitempty"`
	EnglishDescription             string `json:"english_description,omitempty"`
}

// a map of CoceptIDs to Concepts
type ConceptsRepository struct {
	ID                 string
	conceptsDict       map[string]*Concept
	unitsRepo          *UnitsRepository
	conceptsToUnitsMap map[string][]string // mapping of concept ids to an array of unit ids
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

func (idea Idea) String(conceptMap ConceptsRepository) string {
	s := ""
	for conceptID, measurement := range idea.Facts {
		concept := conceptMap.ConceptWithID(conceptID)
		if concept != nil {
			s += concept.EnglishHumanReadableExpression + " -> " + measurement.String()
		} else {
			s += conceptID + " -> " + measurement.String()
		}
	}
	return s
}

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
