package ideas

import "testing"

func Test_ConceptsRepository(t *testing.T) {
	{
		r := ConceptsRepository{ID: "", conceptsDict: map[string]*Concept{}}
		c := r.ConceptWithID("abcd")
		if c != nil {
			t.FailNow()
		}
	}

	{
		r := ConceptsRepository{ID: "", conceptsDict: map[string]*Concept{}}
		c := r.ConceptWithID("abcd")
		if c != nil {
			t.FailNow()
		}
		err := r.SetConcept(Concept{"abcd", "b", "short", "c"})
		if err != nil {
			t.FailNow()
		}
		c = r.ConceptWithID("abcd")
		if c.ID != "abcd" || c.EnglishDescription != "c" || c.EnglishHumanReadableExpression != "b" || c.Short1 != "short" {
			t.FailNow()
		}
	}
}

func Test_DefaultConceptsRepository(t *testing.T) {
	c := DEFAULT_CONCEPT_REPO.ConceptWithID("some id that does not exist")
	if c != nil {
		t.FailNow()
	}
	c = DEFAULT_CONCEPT_REPO.ConceptWithID(CONCEPT_NOT_DETECTED.ID)
	if c.ID != CONCEPT_NOT_DETECTED.ID || c.EnglishDescription != CONCEPT_NOT_DETECTED.EnglishDescription || c.EnglishHumanReadableExpression != CONCEPT_NOT_DETECTED.EnglishHumanReadableExpression || c.Short1 != CONCEPT_NOT_DETECTED.Short1 {
		t.FailNow()
	}

	us := DEFAULT_CONCEPT_REPO.AcceptableUnitsForConceptWithID("some id that does not exist")
	if len(us) != 0 {
		t.FailNow()
	}

}
