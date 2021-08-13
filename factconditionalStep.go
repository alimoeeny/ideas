package ideas

import (
	"fmt"
	"sync"
)

func NewFactConditionalStep(title string, factsheet Factsheet, goNoGo func(Factsheet) ([]Step, error)) Step {
	return &FactConditionalStep{
		id:     newID(),
		title:  title,
		status: Stopped,
		facts:  factsheet,
		goNoGo: goNoGo,
		next:   make([]Step, 0),
	}
}

type FactConditionalStep struct {
	sync.Mutex
	id     int64
	title  string
	status StepStatus
	goNoGo func(Factsheet) ([]Step, error)
	facts  Factsheet
	next   []Step
}

func (s *FactConditionalStep) String() string {
	return fmt.Sprintf("FactConditionalStep: %s [%d]", s.title, s.id)
}

func (s *FactConditionalStep) ID() int64 {
	return s.id
}

func (s *FactConditionalStep) Title() string {
	return s.title
}

func (s *FactConditionalStep) Reset() error {
	s.Lock()
	defer s.Unlock()
	s.status = Running
	err := s.facts.Reset()
	for _, s := range s.next {
		s.Reset()
	}
	return err
}

func (s *FactConditionalStep) Status() StepStatus {
	return s.status
}

func (s *FactConditionalStep) StepForward() ([]Step, []Idea, error) {
	if s.status == Running {
		next, err := s.goNoGo(s.facts)
		if err != nil {
			return nil, nil, err
		}
		if next != nil {
			s.status = Stopped
			for _, step := range next {
				if step.Status() != Running {
					step.Reset()
				}
			}
			return next, nil, nil
		}
	}

	return []Step{}, nil, ErrAlreadyStopped(s.title)
}

func (s *FactConditionalStep) ForwardConnections() []Step {
	return s.next
}
