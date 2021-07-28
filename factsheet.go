package ideas

import (
	"fmt"
	"sync"
)

type FactSheet interface {
	CurrentValue(key string) interface{}
	Reset() error
}

type DictionaryFactsheet struct {
	dic map[string]interface{}
}

func (dfs *DictionaryFactsheet) CurrentValue(key string) interface{} {
	return dfs.dic[key]
}

func (dfs *DictionaryFactsheet) Reset() error {
	return nil
}

type FactConditionalStep struct {
	sync.Mutex
	id     int64
	title  string
	status StepStatus
	goNoGo func(FactSheet) ([]Step, error)
	facts  FactSheet
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

func (s *FactConditionalStep) StepForward() ([]Step, error) {
	if s.status == Running {
		next, err := s.goNoGo(s.facts)
		if err != nil {
			return nil, err
		}
		if next != nil {
			s.status = Stopped
			for _, step := range next {
				if step.Status() != Running {
					step.Reset()
				}
			}
			return next, nil
		}
	}

	return []Step{}, ErrAlreadyStopped
}

func (s *FactConditionalStep) ForwardConnections() []Step {
	return s.next
}
