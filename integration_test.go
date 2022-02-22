package ideas

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_integration_needMoreInfo(t *testing.T) {
	itIsRaining := Idea{ID: newStrID(), EnglishHumanReadableExpression: "It is raining", Facts: map[Concept]*Measurement{CONCEPT_WEATHER_IS_RAINING: {ID: newStrID(), Timestamp: time.Now().UnixNano(), Value: true}}}
	itIsNotRaining := Idea{ID: newStrID(), EnglishHumanReadableExpression: "It is not raining", Facts: map[Concept]*Measurement{CONCEPT_WEATHER_IS_RAINING: {ID: newStrID(), Timestamp: time.Now().UnixNano(), Value: false}}}
	needMoreInfo := Idea{ID: newStrID(), EnglishHumanReadableExpression: "Need More Info", Facts: map[Concept]*Measurement{CONCEPT_DATA_NEEDED: {ID: newStrID(), Timestamp: time.Now().UnixNano(), Value: true}}}
	factsheet := NewDictionaryFactsheet("110")
	isItRaining := NewFactConditionalStep("is it raining", factsheet, func(f Factsheet) ([]Step, error) {
		if f.CurrentValue("raining") == "YES" {
			return []Step{NewStopStep("it is raining", []Idea{itIsRaining})}, nil
		}
		if f.CurrentValue("raining") == "NO" {
			return []Step{NewStopStep("it is not raining", []Idea{itIsNotRaining})}, nil
		}
		return []Step{NewStopStep("need more info", []Idea{needMoreInfo})}, nil
	})
	ss := NewStartStep("Start", []Step{isItRaining})
	wrkflw := NewWorkFlow("is it raining", ss)
	if wrkflw.validate() != nil {
		t.Error("Workflow is not valid")
	}
	wrkflw.Reset()
	wrkflw.StepForward()
	wrkflw.StepForward()
	_, ideas, err := wrkflw.StepForward()
	if wrkflw.status != Stopped {
		t.Error("Workflow should be stopped")
	}
	if err != nil {
		t.Error("Expected no error")
		t.FailNow()
	}
	if len(ideas) != 1 {
		t.Error("Expected one idea")
		t.FailNow()
	}

	c, ok := ideas[0].Facts[CONCEPT_DATA_NEEDED]
	if !(ok && c.Value == true) {
		t.Error("Expected missing data")
	}
}

func Test_integration_giveMeIdeas(t *testing.T) {
	itIsRaining := Idea{ID: newStrID(), EnglishHumanReadableExpression: "It is raining", Facts: map[Concept]*Measurement{CONCEPT_WEATHER_IS_RAINING: {ID: newStrID(), Timestamp: time.Now().UnixNano(), Value: true}}}
	itIsNotRaining := Idea{ID: newStrID(), EnglishHumanReadableExpression: "It is not raining", Facts: map[Concept]*Measurement{CONCEPT_WEATHER_IS_RAINING: {ID: newStrID(), Timestamp: time.Now().UnixNano(), Value: false}}}
	factsheet := NewDictionaryFactsheet("110")
	isItRaining := NewFactConditionalStep("is it raining", factsheet, func(f Factsheet) ([]Step, error) {
		// NOTE: this is not a good way to do this, but for the purposes of this test it is fine
		//       since absense of the key in the factsheet is treated the same way as false value
		//       see Test_integration_needMoreInfo for a better implementation
		if f.CurrentValue("raining") == true {
			return []Step{NewStopStep("success", []Idea{itIsRaining})}, nil
		}
		return []Step{NewStopStep("failure", []Idea{itIsNotRaining})}, nil
	})
	ss := NewStartStep("Start", []Step{isItRaining})
	wrkflw := NewWorkFlow("is it raining", ss)
	if wrkflw.validate() != nil {
		t.Error("Workflow is not valid")
	}
	wrkflw.Reset()
	wrkflw.StepForward()
	wrkflw.StepForward()
	_, ideas, err := wrkflw.StepForward()
	if wrkflw.status != Stopped {
		t.Error("Workflow should be stopped")
	}
	if err != nil {
		t.Error("Expected no error")
		t.FailNow()
	}
	if len(ideas) != 1 {
		t.Error("Expected one idea")
		t.FailNow()
	}
	fmt.Printf("%#v\n", ideas)
	c, ok := ideas[0].Facts[CONCEPT_WEATHER_IS_RAINING]
	if !ok || c.Value == true {
		t.Error("Expected NOT raining")
	}
}

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
		goNoGo: func(f Factsheet) ([]Step, error) {
			if val, ok := f.CurrentValue("spherocyteModerate").(bool); ok && val {
				if val, ok := f.CurrentValue("low platelets").(bool); ok && val {
					return []Step{&StopStep{title: "success"}}, nil
				}
			}
			return []Step{&StopStep{title: "no sphero"}}, nil
		},
		facts: factsheet,
	}
	ss := NewStartStep("Start Sphericyteflow", []Step{moderateSpherocyte})
	wrkflw := NewWorkFlow("Spherocyte", ss)
	if wrkflw.validate() != nil {
		t.Error("Workflow is not valid")
	}
	wrkflw.Reset()
	wrkflw.StepForward()
	x, _, err := wrkflw.StepForward()
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
	x, _, err = wrkflw.StepForward()
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
	SpherocyteDetectorChan := make(chan Measurement)
	FragmentDetectorChan := make(chan Measurement)
	PlateletCounterChan := make(chan Measurement)
	Hemoglobin7DayDeltaChan := make(chan Measurement)
	HematocritChan := make(chan Measurement)
	HemoglobinConcentrationChan := make(chan Measurement)
	HaptoglobulinConcentrationChan := make(chan Measurement)
	DATChan := make(chan Measurement)
	LDHChan := make(chan Measurement)
	BillirubinChan := make(chan Measurement)
	MCVChan := make(chan Measurement)
	SickleCell2RBCRatioChan := make(chan Measurement)

	allSensorChannels := []chan Measurement{
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
		go func(idx int, c chan Measurement) {
			<-time.After(time.Second * time.Duration(1+rand.Int63n(5)))
			c <- Measurement{ID: newStrID(), Timestamp: time.Now().UnixNano(), Value: rand.Float64(), Unit: UNIT_MILLIMOLES}
			close(c)
			fmt.Printf("done with channel %d\n", idx)
		}(i, c)
	}
	waitForAllTheSensors := func(sensorChannels []chan Measurement) (bool, error) {
		agg := make(chan Measurement)
		for _, ch := range sensorChannels {
			go func(c chan Measurement) {
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
	ss := NewStartStep("Start Sphericyteflow", []Step{waitForAllSensorsToBeReady})
	wrkflw := NewWorkFlow("Spherocyte", ss)
	if wrkflw.validate() != nil {
		t.Error("Workflow is not valid")
	}
	wrkflw.Reset()
	counter := 0
	for wrkflw.status == Running {
		counter++
		_, _, err := wrkflw.StepForward()
		if err != nil {
			fmt.Printf(" %d -> Didn't expect this to fail %v", counter, err)
			t.FailNow()
		}
	}
}

func Test_integration_001(t *testing.T) {

	notEnoughBatteryToGetToDestination := IdeaSet{
		ID: newStrID(),
		Ideas: []*Idea{
			{
				ID:                             newStrID(),
				EnglishHumanReadableExpression: "Battery is at 100miles estimates range",
				Facts: map[Concept]*Measurement{
					batteryRangeEstimate: {
						ID:        newStrID(),
						Timestamp: time.Now().UnixNano(),
						Value:     100,
						Unit:      UNIT_MILES,
					},
				},
			},
			{
				ID:                             newStrID(),
				EnglishHumanReadableExpression: "Distance to destination is 200miles",
				Facts: map[Concept]*Measurement{
					distanceToDestination: {
						ID:        newStrID(),
						Timestamp: time.Now().UnixNano(),
						Value:     200,
						Unit:      UNIT_MILES,
					},
				},
			},
		},
	}

	fmt.Println(notEnoughBatteryToGetToDestination)

	step11 := &StopStep{}
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
