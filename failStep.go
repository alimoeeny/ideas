package ideas

import (
	"fmt"
	"sync"
)

type FailStep struct {
	sync.Mutex
	id     int64
	title  string
	status StepStatus
	next   []Step
}

func (s *FailStep) String() string {
	return fmt.Sprintf("FailStep: %s [%d]", s.title, s.id)
}

func (s *FailStep) ID() int64 {
	return s.id
}

func (s *FailStep) Title() string {
	return s.title
}

func (s *FailStep) Reset() error {
	s.Lock()
	defer s.Unlock()
	s.status = Running
	return nil
}

func (s *FailStep) Status() StepStatus {
	return s.status
}

func (s *FailStep) StepForward() ([]Step, error) {
	return []Step{}, ErrWorkflowFailed
}

func (s *FailStep) ForwardConnections() []Step {
	return []Step{}
}
