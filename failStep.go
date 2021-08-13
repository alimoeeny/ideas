package ideas

import (
	"fmt"
	"sync"
)

func NewFailStep(title string, ideas []Idea) Step {
	return &FailStep{
		id:     newID(),
		title:  title,
		status: Stopped,
		ideas:  ideas,
	}
}

type FailStep struct {
	sync.Mutex
	id     int64
	title  string
	status StepStatus
	ideas  []Idea
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

func (s *FailStep) StepForward() ([]Step, []Idea, error) {
	return []Step{}, s.ideas, ErrWorkflowFailed
}

func (s *FailStep) ForwardConnections() []Step {
	return []Step{}
}
