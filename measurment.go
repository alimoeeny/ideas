package ideas

import (
	"fmt"
	"reflect"
)

// TODO
type Conversion func(measurement *Measurement, targetUnit Unit) (interface{}, Unit)

type Unit struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Short1 string `json:"short_1,omitempty"`
}

type Measurement struct {
	ID        string      `json:"id,omitempty"`
	Timestamp int64       `json:"timestamp,omitempty"`
	Value     interface{} `json:"value,omitempty"`
	Unit      *Unit       `json:"unit,omitempty"`
}

// MeasurementRange is a struct that is primarily used to define ranges
// the lower and upper bounds are meant to only have a value and unit and not a timestamp or id
type MeasurementRange struct {
	LowerBound Measurement `json:"lower_bound,omitempty"`
	UpperBound Measurement `json:"upper_bound,omitempty"`
}

func (mr *MeasurementRange) Contains(m Measurement) (bool, error) {
	if mr.LowerBound.Unit == nil || mr.UpperBound.Unit == nil {
		return false, fmt.Errorf("MeasurementRange contains nil units")
	}
	if mr.LowerBound.Unit != m.Unit {
		return false, fmt.Errorf("lower bound unit %v does not match measurement unit %v", mr.LowerBound.Unit, m.Unit)
	}
	switch reflect.TypeOf(mr.LowerBound.Value).Kind() {
	case reflect.Int64:
		if m.Value.(int64) < mr.LowerBound.Value.(int64) {
			return false, nil
		}
		if m.Value.(int64) > mr.UpperBound.Value.(int64) {
			return false, nil
		}
		return true, nil
	// would this work on uints? or on non 64 archs?
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		if m.Value.(int) < mr.LowerBound.Value.(int) {
			return false, nil
		}
		if m.Value.(int) > mr.UpperBound.Value.(int) {
			return false, nil
		}
		return true, nil
	case reflect.Float32, reflect.Float64:
		if m.Value.(float64) < mr.LowerBound.Value.(float64) {
			return false, nil
		}
		if m.Value.(float64) > mr.UpperBound.Value.(float64) {
			return false, nil
		}
		return true, nil
	default:
		return false, fmt.Errorf("unsupported type %v", reflect.TypeOf(mr.LowerBound.Value).Kind())
	}
}

func (m Measurement) String() string {
	s := ""
	if m.Value != nil {
		s = fmt.Sprintf("%v", m.Value)
	} else {
		s = "N/A"
	}
	if m.Unit != nil {
		s = fmt.Sprintf("%v [%v]", s, m.Unit.Short1)
	}
	return s
}
