package ideas

import (
	"encoding/json"
	"testing"
)

func Test_Marshaling(t *testing.T) {
	step2Go := false
	step3Go := false
	step3 := WaitForItStep{
		id:     newID(),
		title:  "step3",
		goNoGo: func() (bool, error) { return step3Go, nil },
		next:   []Step{&StopStep{}},
	}
	step2 := WaitForItStep{
		id:     newID(),
		title:  "step2",
		goNoGo: func() (bool, error) { return step2Go, nil },
		next:   []Step{&step3},
	}
	wrkflw := &Workflow{
		id:        newID(),
		title:     "testing testing",
		status:    Stopped,
		startStep: &StartStep{id: newID(), title: "start", next: []Step{&step2}},
		next:      []Step{&StopStep{id: newID(), title: "end"}},
	}
	wrkflwStr, err := json.MarshalIndent(wrkflw, "", "  ")
	if err != nil {
		t.Errorf("Error marshaling workflow: %s", err)
	}
	t.Logf("Marshaled workflow: %s", wrkflwStr)
	wrkflw.Reset()
	wrkflwStr, err = json.MarshalIndent(wrkflw, "", "  ")
	if err != nil {
		t.Errorf("Error marshaling workflow: %s", err)
	}
	t.Logf("Marshaled workflow: %s", wrkflwStr)
	wrkflw.StepForward()
	wrkflwStr, err = json.MarshalIndent(wrkflw, "", "  ")
	if err != nil {
		t.Errorf("Error marshaling workflow: %s", err)
	}
	t.Logf("Marshaled workflow: %s", wrkflwStr)
}
