package ideas

import (
	"encoding/json"
	"fmt"
	"sync"
)

func NewMoreInfoNeededStep(title string, ideas []Idea) Step {
	return &MoreInfoNeededStep{
		id:     newID(),
		title:  title,
		ideas:  ideas,
		status: Stopped,
	}
}

// MoreInfoNeededStep is a terminal step that asks the user for more information
type MoreInfoNeededStep struct {
	sync.Mutex
	id     int64
	title  string
	status StepStatus
	ideas  []Idea
	next   []Step
}

func (ss *MoreInfoNeededStep) MarshalJSON() ([]byte, error) {
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

func (s *MoreInfoNeededStep) String() string {
	return fmt.Sprintf("MoreInfoNeededStep: %s [%d]", s.title, s.id)
}

func (s *MoreInfoNeededStep) ID() int64 {
	return s.id
}

func (s *MoreInfoNeededStep) Title() string {
	return s.title
}

func (s *MoreInfoNeededStep) Reset() error {
	s.Lock()
	defer s.Unlock()
	s.status = Running
	return nil
}

func (s *MoreInfoNeededStep) Status() StepStatus {
	return s.status
}

func (s *MoreInfoNeededStep) StepForward() ([]Step, []Idea, error) {
	if s.status == Running {
		s.status = Stopped
	}

	return []Step{}, s.ideas, nil
}

func (s *MoreInfoNeededStep) ForwardConnections() []Step {
	return s.next
}
