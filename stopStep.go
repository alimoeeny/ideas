package ideas

import (
	"encoding/json"
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

func (ss *StopStep) MarshalJSON() ([]byte, error) {
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
	defer s.Unlock()
	s.status = Running
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
