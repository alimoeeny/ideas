package ideas

import "fmt"

// TODO
type Conversion func(measurement *Measurment, targetUnit Unit) (interface{}, Unit)

type Unit struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Short1 string `json:"short_1,omitempty"`
}

type Measurment struct {
	ID        string      `json:"id,omitempty"`
	Timestamp int64       `json:"timestamp,omitempty"`
	Value     interface{} `json:"value,omitempty"`
	Unit      *Unit       `json:"unit,omitempty"`
}

func (m Measurment) String() string {
	return fmt.Sprintf("%v [%v]", m.Value, m.Unit.Name)
}
