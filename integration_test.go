package ideas

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_integration_spherocyte(t *testing.T) {
	// no spherocytes, no low platelets
	factsheet := &DictionaryFactsheet{
		dic: map[string]interface{}{
			"spherocyteModerate": false,
		},
	}
	moderateSpherocyte := &FactConditionalStep{
		id:     123,
		title:  "Spherocyte",
		status: Stopped,
		goNoGo: func(f FactSheet) ([]Step, error) {
			if val, ok := f.CurrentValue("spherocyteModerate").(bool); ok && val {
				if val, ok := f.CurrentValue("low platelets").(bool); ok && val {
					return []Step{&StopStep{title: "success"}}, nil
				}
			}
			return []Step{&StopStep{title: "no sphero"}}, nil
		},
		facts: factsheet,
	}
	ss := NewStartStep("Start Sphericyteflow", moderateSpherocyte)
	wrkflw := NewWorkFlow("Spherocyte", ss)
	if wrkflw.validate() != nil {
		t.Error("Workflow is not valid")
	}
	wrkflw.Reset()
	wrkflw.StepForward()
	x, err := wrkflw.StepForward()
	wrkflw.StepForward()
	if wrkflw.status != Stopped {
		t.Error("expected to have stopped")
	}
	if err != nil {
		t.Error("expected no error")
	}
	if len(x) != 1 || x[0].Title() != "no sphero" {
		t.Errorf("expected no sphero but got %v", x)
	}
	// yes spherocytes, yes fragments
	factsheet.SetValue("spherocyteModerate", true)
	factsheet.SetValue("low platelets", true)
	wrkflw.Reset()
	wrkflw.StepForward()
	x, err = wrkflw.StepForward()
	wrkflw.StepForward()
	if wrkflw.status != Stopped {
		t.Error("expected to have stopped")
	}
	if err != nil {
		t.Error("expected no error")
	}
	if len(x) != 1 || x[0].Title() != "success" {
		t.Errorf("expected success but got %v", x)
	}
}

func Test_integration_wait_for_channels(t *testing.T) {
	SpherocyteDetectorChan := make(chan Measurment)
	FragmentDetectorChan := make(chan Measurment)
	PlateletCounterChan := make(chan Measurment)
	Hemoglobin7DayDeltaChan := make(chan Measurment)
	HematocritChan := make(chan Measurment)
	HemoglobinConcentrationChan := make(chan Measurment)
	HaptoglobulinConcentrationChan := make(chan Measurment)
	DATChan := make(chan Measurment)
	LDHChan := make(chan Measurment)
	BillirubinChan := make(chan Measurment)
	MCVChan := make(chan Measurment)
	SickleCell2RBCRatioChan := make(chan Measurment)

	allSensorChannels := []chan Measurment{
		SpherocyteDetectorChan,
		FragmentDetectorChan,
		PlateletCounterChan,
		Hemoglobin7DayDeltaChan,
		HematocritChan,
		HemoglobinConcentrationChan,
		HaptoglobulinConcentrationChan,
		DATChan,
		LDHChan,
		BillirubinChan,
		MCVChan,
		SickleCell2RBCRatioChan,
	}
	for i, c := range allSensorChannels {
		go func(idx int, c chan Measurment) {
			<-time.After(time.Second * time.Duration(1+rand.Int63n(5)))
			c <- Measurment{id: newID(), timestamp: time.Now().UnixNano(), value: rand.Float64(), unit: UNIT_MILLIMOLES}
			close(c)
			fmt.Printf("done with channel %d\n", idx)
		}(i, c)
	}
	waitForAllTheSensors := func(sensorChannels []chan Measurment) (bool, error) {
		agg := make(chan Measurment)
		for _, ch := range sensorChannels {
			go func(c chan Measurment) {
				for msg := range c {
					agg <- msg
				}
			}(ch)
		}
		counter := 0
		for counter < len(sensorChannels) {
			msg := <-agg
			fmt.Println("received ", msg)
			counter++
		}
		return true, nil
	}

	waitForAllSensorsToBeReady := NewWaitForItStep("Wait for all the sensors", time.Hour, func() (bool, error) { return waitForAllTheSensors(allSensorChannels) })
	ss := NewStartStep("Start Sphericyteflow", waitForAllSensorsToBeReady)
	wrkflw := NewWorkFlow("Spherocyte", ss)
	if wrkflw.validate() != nil {
		t.Error("Workflow is not valid")
	}
	wrkflw.Reset()
	counter := 0
	for wrkflw.status == Running {
		counter++
		_, err := wrkflw.StepForward()
		if err != nil {
			fmt.Printf(" %d -> Didn't expect this to fail %v", counter, err)
			t.FailNow()
		}
	}
}

func Test_integration_001(t *testing.T) {

	notEnoughBatteryToGetToDestination := IdeaSet{
		id: newID(),
		ideas: []*Idea{
			{
				id:                             newID(),
				englishHumanReadableExpression: "Battery is at 100miles estimates range",
				facts: map[*Concept]Measurment{
					&batteryRangeEstimate: {
						id:        newID(),
						timestamp: time.Now().UnixNano(),
						value:     100,
						unit:      UNIT_MILES,
					},
				},
			},
			{
				id:                             newID(),
				englishHumanReadableExpression: "Distance to destination is 200miles",
				facts: map[*Concept]Measurment{
					&distanceToDestination: {
						id:        newID(),
						timestamp: time.Now().UnixNano(),
						value:     200,
						unit:      UNIT_MILES,
					},
				},
			},
		},
	}

	fmt.Println(notEnoughBatteryToGetToDestination)

	step11 := &IdeaStep{}
	// step12 := &Step{}
	wrkfklw := WorkFlow{
		id:     newID(),
		status: Stopped,
		startStep: &StartStep{
			title:  "Start",
			id:     newID(),
			status: Stopped,
			next: []Step{
				step11,
			},
		},
	}

	if wrkfklw.validate() != nil {
		t.FailNow()
	}
}
