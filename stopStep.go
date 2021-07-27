package ideas

import (
	"fmt"
	"sync"
)

type StopStep struct {
	sync.Mutex
	id     int64
	title  string
	status StepStatus
	next   []Step
}

func (s *StopStep) String() string {
	return fmt.Sprintf("StopStep: %s [%d]", s.title, s.id)
}

func (s *StopStep) ID() int64 {
	return s.id
}

func (s *StopStep) Title() string {
	return s.title
}

func (s *StopStep) Reset() error {
	s.Lock()
	s.status = Running
	s.Unlock()
	return nil
}

func (s *StopStep) Status() StepStatus {
	return s.status
}

func (s *StopStep) StepForward() ([]Step, error) {
	if s.status == Running {
		s.status = Stopped
	}

	return []Step{}, nil
}

func (s *StopStep) ForwardConnections() []Step {
	return s.next
}
