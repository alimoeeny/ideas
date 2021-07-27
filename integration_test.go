package ideas

import (
	"fmt"
	"testing"
	"time"
)

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
