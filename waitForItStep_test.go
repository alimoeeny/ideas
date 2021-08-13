package ideas

import (
	"strings"
	"testing"
	"time"
)

func Test_WaitForItStep_StuckForEver(t *testing.T) {
	{
		wfis := WaitForItStep{
			id:     newID(),
			goNoGo: func() (bool, error) { return false, nil },
		}
		err := wfis.Reset()
		if err != nil {
			t.Errorf("Error: %s", err)
			t.FailNow()
		}
		next, outcomingIdeas, err := wfis.StepForward()
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if len(next) > 0 {
			t.FailNow()
		}
		if len(outcomingIdeas) > 0 {
			t.Error("DIdn't expect any new ideas coming out of wait for it")
			t.FailNow()
		}
	}
}

func Test_WaitForItStep_GoFromGetGo(t *testing.T) {
	{
		wfis := WaitForItStep{
			id:     newID(),
			goNoGo: func() (bool, error) { return true, nil },
			next: []Step{
				&PipeStep{next: []Step{&StopStep{}}},
				&PipeStep{next: []Step{&StopStep{}}},
			},
		}
		err := wfis.Reset()
		if err != nil {
			t.Errorf("Error: %s", err)
			t.FailNow()
		}
		next, _, err := wfis.StepForward()
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if len(next) < 2 {
			t.FailNow()
		}
	}
}

func Test_WaitForItStep_ActuallyWait(t *testing.T) {
	{
		startedAt := time.Now()
		wfis := WaitForItStep{
			id: newID(),
			goNoGo: func() (bool, error) {
				return time.Since(startedAt) > time.Second, nil
			},
			next: []Step{
				&PipeStep{next: []Step{&StopStep{}}},
				&PipeStep{next: []Step{&StopStep{}}},
			},
		}
		err := wfis.Reset()
		if err != nil {
			t.Errorf("Error: %s", err)
			t.FailNow()
		}
		next, _, err := wfis.StepForward()
		for len(next) < 1 && err == nil {
			next, _, err = wfis.StepForward()
		}
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if len(next) < 2 {
			t.Errorf("Expected 2 steps, got %d", len(next))
			t.FailNow()
		}
		if time.Since(startedAt) < time.Second {
			t.Errorf("Expected to wait for at least 1 second, but waited for %s", time.Since(startedAt))
			t.FailNow()
		}
	}
}

func Test_WaitForItStep_Timeout(t *testing.T) {
	{
		startedAt := time.Now()
		wfis := WaitForItStep{
			id:          newID(),
			timeout:     time.Second,
			timeoutNext: []Step{&StopStep{}},
			goNoGo: func() (bool, error) {
				return time.Since(startedAt) > 3*time.Second, nil
			},
			next: []Step{
				&PipeStep{next: []Step{&StopStep{}}},
				&PipeStep{next: []Step{&StopStep{}}},
			},
		}
		err := wfis.Reset()
		if err != nil {
			t.Errorf("Error: %s", err)
			t.FailNow()
		}
		next, _, err := wfis.StepForward()
		for len(next) < 1 && err == nil {
			next, _, err = wfis.StepForward()
		}
		if !strings.Contains(err.Error(), "timeout") {
			t.Errorf("Expected timeout error: %s", err)
		}
		if len(next) != 1 {
			t.Errorf("Expected 1 steps, got %d", len(next))
			t.FailNow()
		}
		if time.Since(startedAt) > 2*time.Second {
			t.Errorf("Expected to timeout after 1 second, but waited for %s", time.Since(startedAt))
			t.FailNow()
		}
	}
}

func Test_WaitForItStep_WaitAsync(t *testing.T) {
	{
		startedAt := time.Now()
		outChan := make(chan int)
		errChan := make(chan error)

		wfis := WaitForItStep{
			id: newID(),
			goNoGo: func() (bool, error) {
				job := func(out chan int, errChan chan error) {}
				go job(outChan, errChan)
				return time.Since(startedAt) > time.Second, nil
			},
			next: []Step{
				&PipeStep{next: []Step{&StopStep{}}},
				&PipeStep{next: []Step{&StopStep{}}},
			},
		}
		err := wfis.Reset()
		if err != nil {
			t.Errorf("Error: %s", err)
			t.FailNow()
		}
		next, _, err := wfis.StepForward()
		for len(next) < 1 && err == nil {
			next, _, err = wfis.StepForward()
		}
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if len(next) < 2 {
			t.Errorf("Expected 2 steps, got %d", len(next))
			t.FailNow()
		}
		if time.Since(startedAt) < time.Second {
			t.Errorf("Expected to wait for at least 1 second, but waited for %s", time.Since(startedAt))
			t.FailNow()
		}
	}
}
