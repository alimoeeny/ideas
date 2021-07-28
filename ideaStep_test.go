package ideas

import (
	"reflect"
	"testing"
)

func Test_IdeaStep(t *testing.T) {
	ss := &IdeaStep{}
	stepInterface := reflect.TypeOf((*Step)(nil)).Elem()
	if !reflect.TypeOf(ss).Implements(stepInterface) {
		t.Error("IdeaStep does not implement Step")
	}
}
