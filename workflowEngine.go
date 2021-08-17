package ideas

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

func NewWorkFlow(title string, startStep *StartStep) *WorkFlow {
	wf := &WorkFlow{
		id:        newID(),
		title:     title,
		status:    Stopped,
		startStep: startStep,
		next:      []Step{startStep},
		ideas:     []Idea{},
	}

	return wf
}

type WorkFlow struct {
	sync.Mutex
	id        int64
	title     string
	status    StepStatus
	startStep *StartStep
	next      []Step
	ideas     []Idea
}

func (wf *WorkFlow) MarshalJSON() ([]byte, error) {
	temp := struct {
		ID        int64      `json:"id"`
		Title     string     `json:"title"`
		Status    string     `json:"status"`
		StartStep *StartStep `json:"startStep"`
		Next      []Step     `json:"next"`
	}{
		ID:        wf.id,
		Title:     wf.title,
		Status:    wf.status.String(),
		StartStep: wf.startStep,
		Next:      wf.next,
	}
	return json.MarshalIndent(temp, "", "  ")
}

func (wf *WorkFlow) String() string {
	return fmt.Sprintf("Workflow: %s [%d]", wf.title, wf.id)
}

func (wf *WorkFlow) ID() int64 {
	return wf.id
}

func (wf *WorkFlow) Title() string {
	return wf.title
}

func (s *WorkFlow) Reset() error {
	s.Lock()
	defer s.Unlock()
	s.status = Running
	s.startStep.Reset()
	s.next = []Step{s.startStep}
	return nil
}

func (s *WorkFlow) Status() StepStatus {
	return s.status
}

func (s *WorkFlow) StepForward() ([]Step, []Idea, error) {
	nextSteps := []Step{}
	nextIdeas := []Idea{}
	if s.status == Running {
		for _, step := range s.next {
			next, outcomingIdeas, err := step.StepForward()
			if err != nil {
				return nextSteps, nextIdeas, err
			}
			for _, s := range next {
				s.Reset()
			}
			nextSteps = append(nextSteps, next...)
			nextIdeas = append(nextIdeas, outcomingIdeas...)
		}
		if len(nextSteps) == 0 {
			s.status = Stopped
		}
		s.next = nextSteps
		s.ideas = nextIdeas
		return s.next, nextIdeas, nil
	}

	return []Step{}, s.ideas, ErrAlreadyStopped(s.title)
}

func (s *WorkFlow) ForwardConnections() []Step {
	return s.next
}

func (wf *WorkFlow) validate() error {
	// RULES
	// Have a Start
	// Don't have orphan states
	// Don't have orphan set of states that you cannot reach from the start
	//
	if wf.startStep == nil {
		return errors.New("workflow must have a start step")
	}

	return nil
}

// Run taks StepForward after StepForward until the workflow is stopped
// and returns the last ste of steps (usually a StopStep or a FailStep),
// and ALL the accumulated ideas
// and an error which hopefuly is nil
func (wf *WorkFlow) Run() ([]Step, []Idea, error) {
	err := wf.Reset()
	if err != nil {
		return nil, nil, err
	}
	resultingIdeas := []Idea{}
	for wf.status != Stopped {
		next, nextIdeas, err := wf.StepForward()
		resultingIdeas = append(resultingIdeas, nextIdeas...)
		if err != nil || len(next) == 0 {
			wf.status = Stopped
			return next, resultingIdeas, err
		}
		wf.next = next
		wf.ideas = nextIdeas
	}
	return []Step{}, []Idea{}, nil
}

// TODO: Looks like our best bet is to save the code as a text file
//       and compile it at runtime and load it using go's plugin
//       OR
//       alternatively, we can use a JSON file to save the workflow
//       BUT not to support states that have func in them, like
//       create a specific limited factsheetStep that does simple logic
//       and drop support for serilization of generic conditional state
func (wf *WorkFlow) Serialize() ([]byte, error) {
	return []byte(""), fmt.Errorf("not implemented yet")
}

func (wf *WorkFlow) Deserialize(data []byte) error {
	return fmt.Errorf("not implemented yet")
}
