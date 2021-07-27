package ideas

import "sync"

type IdeaStep struct {
	sync.Mutex
	id     int64
	title  string
	status StepStatus
	idea   Idea
	next   []Step
}

func (s *IdeaStep) ID() int64 {
	return s.id
}

func (s *IdeaStep) Title() string {
	return s.title
}

func (s *IdeaStep) Status() StepStatus {
	return s.status
}

func (s *IdeaStep) Reset() error {
	s.Lock()
	s.idea = Idea{}
	s.status = Running
	s.Unlock()
	return nil
}

func (s *IdeaStep) StepForward() ([]Step, error) {
	if s.status == Running {
		if s.idea.id != 0 {
			s.status = Stopped
			return s.next, nil
		} else {
			return nil, nil
		}
	}

	return []Step{}, ErrAlreadyStopped
}

func (s *IdeaStep) ForwardConnections() []Step {
	return s.next
}
