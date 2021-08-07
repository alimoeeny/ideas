package ideas

import (
	"reflect"
	"testing"
)

func Test_WorkflowIsAStepToo(t *testing.T) {
	wf := &WorkFlow{}
	stepInterface := reflect.TypeOf((*Step)(nil)).Elem()
	if !reflect.TypeOf(wf).Implements(stepInterface) {
		t.Error("workflow does not implement Step")
	}
}

func Test_WorkflowValidateStart(t *testing.T) {
	{
		wf := &WorkFlow{}
		if wf.validate().Error() != "workflow must have a start step" {
			t.Error("workflow should not be valid")
			t.FailNow()
		}
	}
	{
		wf := &WorkFlow{}
		wf.startStep = &StartStep{}
		if wf.validate() != nil {
			t.Error("an empty workflow is a valid workflow")
			t.FailNow()
		}
	}
}

func Test_WorkflowFundamentals(t *testing.T) {
	wf := &WorkFlow{}
	wf.startStep = &StartStep{next: []Step{&PipeStep{next: []Step{&StopStep{}}}}}
	if wf.validate() != nil {
		t.Error("a valid workflow should be valid")
		t.FailNow()
	}
	wf.Reset()
	if wf.status != Running {
		t.Error("a fresh workflow should have a status of running")
		t.FailNow()
	}
	wf.Run()
	if wf.status != Stopped {
		t.Error("a valid workflow should be stopped at the end of a run")
		t.FailNow()
	}
}

func Test_Workflow_and_Factsheet(t *testing.T) {
	wf := &WorkFlow{}
	factsheet := &DictionaryFactsheet{
		dic: map[string]interface{}{
			"name":   "test",
			"hight":  12,
			"weight": 170,
		},
	}
	fss := &FactConditionalStep{
		id:    newID(),
		title: "test",
		goNoGo: func(f Factsheet) ([]Step, error) {
			if val, ok := f.CurrentValue("hight").(int); ok && val > 170 {
				return []Step{&StopStep{}}, nil
			}
			return []Step{&FailStep{}}, nil
		},
		facts: factsheet,
	}
	wf.startStep = &StartStep{next: []Step{fss}}
	if wf.validate() != nil {
		t.Error("a valid workflow should be valid")
		t.FailNow()
	}
	wf.Reset()
	// next, err := wf.StepForward()
	// if next == nil || err != nil {
	// 	fmt.Printf("next is: %v and error is %v", next, err)
	// 	t.FailNow()
	// }
	err := wf.Run()
	//fmt.Printf("next is: %v and error is %v", next, err)
	if err != ErrWorkflowFailed {
		t.Error("a failed workflow should fail")
		t.FailNow()
	}
}
