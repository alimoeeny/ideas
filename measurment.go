package ideas

import "fmt"

// TODO
type Conversion func(measurement *Measurment, targetUnit Unit) (interface{}, Unit)

type Unit struct {
	id     int64
	name   string
	short1 string
}

type Measurment struct {
	id        int64
	timestamp int64
	value     interface{}
	unit      *Unit
}

func (m Measurment) String() string {
	return fmt.Sprintf("%v [%v]", m.value, m.unit.name)
}
