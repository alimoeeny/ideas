package ideas

import "fmt"

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
