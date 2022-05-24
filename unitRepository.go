package ideas

import (
	"fmt"
	"math/rand"
)

var DEFAULT_UNIT_REPO = UnitsRepository{
	ID:   "nogobbaoqj",
	dict: map[string]*Unit{},
}

func init() {
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_UNITLESS)
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_RATIO)
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_PERCENTAGE)

	DEFAULT_UNIT_REPO.SetUnit(*UNIT_IDENTIFIER)

	// Categorical
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_CATEGORICAL)
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_POSITIVE_NEGATIVE_MISSING)

	DEFAULT_UNIT_REPO.SetUnit(*UNIT_COUNT)

	DEFAULT_UNIT_REPO.SetUnit(*UNIT_COUNT_PER_LITER)
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_COUNT_MILLION_PER_LITER)
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_COUNT_BILLION_PER_LITER)
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_COUNT_TRILLION_PER_LITER)

	DEFAULT_UNIT_REPO.SetUnit(*UNIT_COUNTS_PER_MILLILITER)

	DEFAULT_UNIT_REPO.SetUnit(*UNIT_COUNT_PER_UL)
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_COUNT_MILLION_PER_UL)

	// Length
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_MILES)

	DEFAULT_UNIT_REPO.SetUnit(*UNIT_MILLIMOLES)

	DEFAULT_UNIT_REPO.SetUnit(*UNIT_GRAMS_PER_LITER)
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_GRAMS_PER_DECILITER)

	// Volume
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_LITER)
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_DECILITER)
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_CENTILITER)
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_MILLILITER)
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_MICROLITER)
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_NANOLITER)
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_PICOLITER)
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_FEMTOLITER)

	// Wight
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_GRAM)
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_KILOGRAM)
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_MILLIGRAM)
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_MICROGRAM)
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_NANOGRAM)
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_PICOGRAM)
	DEFAULT_UNIT_REPO.SetUnit(*UNIT_FEMTOGRAM)

}

type UnitsRepository struct {
	ID   string
	dict map[string]*Unit
}

func (repo *UnitsRepository) UnitWithID(id string) *Unit {
	return (*repo).dict[id]
}

func (repo *UnitsRepository) SetUnit(c Unit) error {
	if c.ID == "" {
		return fmt.Errorf("unit needs to have an non-enpty id")
	}
	if (*repo).dict[c.ID] != nil {
		return fmt.Errorf("another unit with the same ID already exists")
	}
	// let it crash if they pass in a nil repo, it is danguarous to do anything else
	(*repo).dict[c.ID] = &c
	return nil
}

func (repo *UnitsRepository) All() map[string]*Unit {
	return repo.dict
}

func (repo *UnitsRepository) NewUnit(id string, name string, short_01 string) (Unit, error) {
	if id == "" {
		id = newStrID()
	}
	u := Unit{id, name, short_01}
	err := repo.SetUnit(u)
	return u, err
}

func (repo *UnitsRepository) RandomUnit() *Unit {
	if repo.dict == nil || len(repo.dict) < 1 {
		return &Unit{}
	}
	ri := rand.Int31n(int32(len(repo.dict)))
	counter := int32(0)
	for _, v := range repo.dict {
		if counter >= ri {
			return v
		}
		counter += 1
	}
	return &Unit{}
}
