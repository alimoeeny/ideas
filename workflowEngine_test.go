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
