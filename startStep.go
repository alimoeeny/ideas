package ideas

import (
	"encoding/json"
	"fmt"
	"sync"
)

func NewStartStep(title string, nextStep []Step) *StartStep {
	return &StartStep{
		id:     newID(),
		title:  title,
		status: Stopped,
		next:   nextStep,
	}
}

type StartStep struct {
	sync.Mutex
	id     int64
	title  string
	status StepStatus
	next   []Step
}

func (ss *StartStep) MarshalJSON() ([]byte, error) {
	temp := struct {
		ID     int64  `json:"id"`
		Title  string `json:"title"`
		Status string `json:"status"`
		Next   []Step `json:"next"`
	}{
		ID:     ss.id,
		Title:  ss.title,
		Status: ss.status.String(),
		Next:   ss.next,
	}
	return json.MarshalIndent(temp, "", "  ")
}

func (s *StartStep) String() string {
	return fmt.Sprintf("StartStep: %s [%d]", s.title, s.id)
}

func (s *StartStep) ID() int64 {
	return s.id
}

func (s *StartStep) Title() string {
	return s.title
}

func (s *StartStep) Reset() error {
	s.Lock()
	defer s.Unlock()
	s.status = Running
	for _, s := range s.next {
		s.Reset()
	}
	return nil
}

func (s *StartStep) Status() StepStatus {
	return s.status
}

func (s *StartStep) StepForward() ([]Step, []Idea, error) {
	if s.status == Running {
		s.status = Stopped
		for _, step := range s.next {
			if step.Status() != Running {
				step.Reset()
			}
		}
		return s.next, []Idea{}, nil
	}

	return []Step{}, []Idea{}, ErrAlreadyStopped(s.title)
}

func (s *StartStep) ForwardConnections() []Step {
	return s.next
}
