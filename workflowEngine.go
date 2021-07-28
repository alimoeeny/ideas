package ideas

import (
	"errors"
	"fmt"
	"sync"
)

type WorkFlow struct {
	sync.Mutex
	id        int64
	title     string
	status    StepStatus
	startStep *StartStep
	next      []Step
}

func (wf *WorkFlow) String() string {
	return fmt.Sprintf("Workflow: %s [%d]", wf.title, wf.id)
}

func (wf *WorkFlow) ID() int64 {
	return wf.id
}

func (wf *WorkFlow) Title() string {
	return wf.title
}

func (s *WorkFlow) Reset() error {
	s.Lock()
	defer s.Unlock()
	s.status = Running
	s.startStep.Reset()
	s.next = []Step{s.startStep}
	return nil
}

func (s *WorkFlow) Status() StepStatus {
	return s.status
}

func (s *WorkFlow) StepForward() ([]Step, error) {
	nextSteps := []Step{}
	if s.status == Running {
		for _, step := range s.next {
			next, err := step.StepForward()
			if err != nil {
				return nextSteps, err
			}
			for _, s := range next {
				s.Reset()
			}
			nextSteps = append(nextSteps, next...)
		}
		if len(nextSteps) == 0 {
			s.status = Stopped
		}
		return nextSteps, nil
	}

	return []Step{}, ErrAlreadyStopped
}

func (s *WorkFlow) ForwardConnections() []Step {
	return s.next
}

func (wf *WorkFlow) validate() error {
	// RULES
	// Have a Start
	// Don't have orphan states
	// Don't have orphan set of states that you cannot reach from the start
	//
	if wf.startStep == nil {
		return errors.New("workflow must have a start step")
	}

	return nil
}

// func (wf *WorkFlow) TakeStep() error {
// 	// TODO
// 	return nil
// }

func (wf *WorkFlow) Run() error {
	for wf.status != Stopped {
		next, err := wf.StepForward()
		if err != nil || len(next) == 0 {
			wf.status = Stopped
			return err
		}
		wf.next = next
	}
	return nil
}
