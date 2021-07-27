package ideas

import (
	"errors"
	"fmt"
	"sync"
)

type StepStatus int

const (
	Stopped StepStatus = iota
	Running
)

var ErrAlreadyStopped = errors.New("already stopped")
var ErrWorkflowFailed = errors.New("workflow arrived at a fail step")

// type GoNoGo interface {
// 	areWeGoodToGo() bool
// }

// type ALWAYS_GO struct{}

// func (a ALWAYS_GO) areWeGoodToGo() bool {
// 	return true
// }

// Step represents a single step in a workflow
type Step interface {
	ID() int64
	Title() string
	Status() StepStatus
	Reset() error
	StepForward() ([]Step, error)
	ForwardConnections() []Step
	String() string
}

type StartStep struct {
	sync.Mutex
	id     int64
	title  string
	status StepStatus
	next   []Step
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
	s.status = Running
	for _, s := range s.next {
		s.Reset()
	}
	s.Unlock()
	return nil
}

func (s *StartStep) Status() StepStatus {
	return s.status
}

func (s *StartStep) StepForward() ([]Step, error) {
	if s.status == Running {
		s.status = Stopped
		for _, step := range s.next {
			if step.Status() != Running {
				step.Reset()
			}
		}
		return s.next, nil
	}

	return []Step{}, ErrAlreadyStopped
}

func (s *StartStep) ForwardConnections() []Step {
	return s.next
}

// ProgressEvaluator accepts a set of ideas and decides if the step can progress to next step
type ProgressEvaluator func(ideas IdeaSet) (bool, error)

var alwaysProgressing = ProgressEvaluator(func(is IdeaSet) (bool, error) { return true, nil })

// TODO
type InquireyStep struct {
	acceptableIdeas IdeaSet
	canProgress     ProgressEvaluator
}

var rbcShapeStudy = InquireyStep{
	acceptableIdeas: IdeaSet{},
	canProgress:     alwaysProgressing,
}
