package ideas

import (
	"testing"
)

func Test_MeasurementString(t *testing.T) {
	{
		m := Measurement{}
		if m.String() != "N/A" {
			t.Errorf("Measurement should be N/A but it is %s", m.String())
		}
	}
	{
		m := Measurement{ID: "123"}
		if m.String() != "N/A" {
			t.Errorf("Measurement should be N/A but it is %s", m.String())
		}
	}
	{
		m := Measurement{ID: "123", Value: 110}
		if m.String() != "110" {
			t.Errorf("Measurement should be 110 but it is %s", m.String())
		}
	}
	{
		m := Measurement{ID: "123", Value: 110, Unit: &Unit{Short1: "ml"}}
		if m.String() != "110 [ml]" {
			t.Errorf("Measurement should be 110 [] but it is %s", m.String())
		}
	}
}

func Test_MeasurmentRange(t *testing.T) {
	{
		r := MeasurementRange{
			LowerBound: Measurement{Value: 110, Unit: UNIT_COUNT},
			UpperBound: Measurement{Value: 120, Unit: UNIT_COUNT},
		}

		if x, err := r.Contains(Measurement{Value: 115, Unit: UNIT_COUNT}); x != true || err != nil {
			t.Errorf("Range should contain 115 but it does not or err is not nil %v", err)
		}
	}
}
