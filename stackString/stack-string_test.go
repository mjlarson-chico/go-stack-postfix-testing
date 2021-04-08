package stackString

import "testing"

func TestIsEmpty(t *testing.T) {
	s := NewStack()
	emptyResult := s.IsEmpty()

	if !emptyResult {
		t.Error("IsEmpty true failed")
	} else {
		t.Log("IsEmpty true success")
	}

	s.Push("l")

	notEmptyResult := s.IsEmpty()

	if notEmptyResult {
		t.Error("IsEmpty false failed")
	} else {
		t.Log("IsEmpty false success")
	}
}
