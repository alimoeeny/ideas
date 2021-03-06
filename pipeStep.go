package ideas

import (
	"fmt"
	"sync"
)

type PipeStep struct {
	sync.Mutex
	id     int64
	title  string
	status StepStatus
	next   []Step
	ideas  []Idea
}

func (s *PipeStep) String() string {
	return fmt.Sprintf("PipeStep: %s [%d]", s.title, s.id)
}

func (s *PipeStep) ID() int64 {
	return s.id
}

func (s *PipeStep) Title() string {
	return s.title
}

func (s *PipeStep) Reset() error {
	s.Lock()
	defer s.Unlock()
	s.status = Running
	return nil
}

func (s *PipeStep) Status() StepStatus {
	return s.status
}

func (s *PipeStep) StepForward() ([]Step, []Idea, error) {
	if s.status == Running {
		s.status = Stopped
		for _, step := range s.next {
			if step.Status() != Running {
				step.Reset()
			}
		}
		return s.next, s.ideas, nil
	}

	return []Step{}, s.ideas, ErrAlreadyStopped(s.title)
}

func (s *PipeStep) ForwardConnections() []Step {
	return s.next
}
