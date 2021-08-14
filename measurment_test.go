package ideas

import "testing"

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
