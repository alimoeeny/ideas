package ideas

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

func NewWaitForItStep(title string, timeout time.Duration, goNoGo func() (bool, error)) Step {
	return &WaitForItStep{
		id:      newID(),
		title:   title,
		timeout: timeout,
		goNoGo:  goNoGo,
	}
}

type WaitForItStep struct {
	sync.Mutex
	id          int64
	title       string
	status      StepStatus
	goNoGo      func() (bool, error)
	started     time.Time
	next        []Step
	timeout     time.Duration
	timeoutNext []Step
}

func (ss *WaitForItStep) MarshalJSON() ([]byte, error) {
	temp := struct {
		ID     int64  `json:"id,omitempty"`
		Title  string `json:"title,omitempty"`
		Status string `json:"status,omitempty"`
		//GoNoGo      func() (bool, error) `json:"go_no_go,omitempty"`
		Started     time.Time     `json:"started,omitempty"`
		Next        []Step        `json:"next,omitempty"`
		Timeout     time.Duration `json:"timeout,omitempty"`
		TimeoutNext []Step        `json:"timeout_next,omitempty"`
	}{
		ID:     ss.id,
		Title:  ss.title,
		Status: ss.status.String(),
		//GoNoGo:      ss.goNoGo,
		Started:     ss.started,
		Next:        ss.next,
		Timeout:     ss.timeout,
		TimeoutNext: ss.timeoutNext,
	}
	return json.MarshalIndent(temp, "", "  ")
}

func (s *WaitForItStep) String() string {
	return fmt.Sprintf("WaitForItStep: %s [%d]", s.title, s.id)
}

func (s *WaitForItStep) ID() int64 {
	return s.id
}

func (s *WaitForItStep) Title() string {
	return s.title
}

func (s *WaitForItStep) Reset() error {
	s.Lock()
	defer s.Unlock()
	if s.goNoGo == nil {
		return fmt.Errorf("no go/no go function provided")
	}
	s.status = Running
	s.started = time.Now()
	for _, s := range s.next {
		s.Reset()
	}

	return nil
}

func (s *WaitForItStep) Status() StepStatus {
	return s.status
}

func (s *WaitForItStep) StepForward() ([]Step, []Idea, error) {
	if s.status == Running {
		if s.goNoGo == nil {
			return nil, nil, fmt.Errorf("no go/no go function provided")
		}
		if s.timeout > 0 && time.Since(s.started) > s.timeout {
			return s.timeoutNext, nil, fmt.Errorf("timeout")
		}
		shallWe, err := s.goNoGo()
		if err != nil {
			return nil, nil, err
		}
		if shallWe {
			s.status = Stopped
			for _, step := range s.next {
				if step.Status() != Running {
					step.Reset()
				}
			}
			return s.next, nil, nil
		} else {
			return []Step{}, nil, nil
		}
	}

	return []Step{}, nil, ErrAlreadyStopped(s.title)
}

// func (s *WaitForItStep) StepForward() ([]Step, error) {
// 	if s.status == Running {
// 		select {
// 		case shallWe := <-s.goNoGo:
// 			if shallWe {
// 				s.status = Stopped
// 				for _, step := range s.next {
// 					if step.Status() != Running {
// 						step.Reset()
// 					}
// 				}
// 				return s.next, nil
// 			} else {
// 				return nil, nil
// 			}
// 		case err := <-s.errGo:
// 			return nil, err
// 		default:
// 			return nil, nil
// 		}
// 	}

// 	return []Step{}, ErrAlreadyStopped
// }

func (s *WaitForItStep) ForwardConnections() []Step {
	return s.next
}
