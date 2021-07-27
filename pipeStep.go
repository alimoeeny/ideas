package ideas

import "sync"

type PipeStep struct {
	sync.Mutex
	id     int64
	title  string
	status StepStatus
	next   []Step
}

func (s *PipeStep) ID() int64 {
	return s.id
}

func (s *PipeStep) Title() string {
	return s.title
}

func (s *PipeStep) Reset() error {
	s.Lock()
	s.status = Running
	s.Unlock()
	return nil
}

func (s *PipeStep) Status() StepStatus {
	return s.status
}

func (s *PipeStep) StepForward() ([]Step, error) {
	if s.status == Running {
		s.status = Stopped
		return s.next, nil
	}

	return []Step{}, ErrAlreadyStopped
}

func (s *PipeStep) ForwardConnections() []Step {
	return s.next
}
