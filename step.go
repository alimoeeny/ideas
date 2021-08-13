package ideas

import (
	"errors"
	"fmt"
)

type StepStatus int

const (
	Stopped StepStatus = iota
	Running
)

func (st StepStatus) String() string {
	return fmt.Sprintf("%d", st)
}

var ErrAlreadyStopped = func(stepTitle string) error { return fmt.Errorf("%s is already stopped", stepTitle) }
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
	StepForward() ([]Step, []Idea, error)
	ForwardConnections() []Step
	String() string
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
