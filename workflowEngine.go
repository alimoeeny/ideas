package ideas

import (
	"errors"
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

func (wf *WorkFlow) ID() int64 {
	return wf.id
}

func (wf *WorkFlow) Title() string {
	return wf.title
}

func (s *WorkFlow) Reset() error {
	s.Lock()
	s.status = Running
	s.Unlock()
	return nil
}

func (s *WorkFlow) Status() StepStatus {
	return s.status
}

func (s *WorkFlow) StepForward() ([]Step, error) {
	if s.status == Running {
		s.status = Stopped
		return s.next, nil
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

func (wf *WorkFlow) TakeStep() error {
	// TODO
	return nil
}

func (wf *WorkFlow) Run() error {
	for wf.status != Stopped {
		wf.StepForward()
	}
	return nil
}
