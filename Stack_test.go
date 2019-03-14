package Stack

import (
	"testing"
)

var s Stack

func TestNewStack(t *testing.T) {
	s := NewStack(1)
	if s.Length() != 1 {
		t.Error("Error created stack of size")
	}
}

func TestPush(t *testing.T) {

	if err := s.Push(nil); err == nil {
		t.Error("Should have recieved an error for adding nil value")
	}

	s.Push(0)
	if s.Top() != 0 {
		t.Error("Error pushing to stack")
	}

	s.Push(1)
	if s.Top() != 1 {
		t.Error("Error pushing to stack")
	}
}

func TestPop(t *testing.T) {
	s.Push(0)
	s.Push(1)

	if x := s.Pop(); x != 1 {
		t.Error("Error with pop, Value should be 1")
	}
	s.Push(2)

	if x := s.Pop(); x != 2 {
		t.Error("Error with pop, Value should be 2")

	}

	if x := s.Pop(); x != 0 {
		t.Error("Error with pop, Value should be 0")
	}
}

func TestTop(t *testing.T) {
	s.Push(0)
	s.Push(1)

	if x := s.Top(); x != 1 {
		t.Error("Error retrieving top")
	}
}

func TestClearStack(t *testing.T) {
	s.ClearStack()
	if !s.IsEmpty() {
		t.Error("Stack should be empty now")
	}
}

func TestIsEmpty(t *testing.T) {
	s.ClearStack()
	if x := s.IsEmpty(); !x {
		t.Error("Stack should be empty")
	}

	s.Push(1)
	if x := s.IsEmpty(); x {
		t.Error("Stack should have a value")
	}

}

func TestLength(t *testing.T) {
	s.ClearStack()

	if x := s.Length(); x != 0 {
		t.Error("Stack should be empty")
	}

	s.Push(1)
	if x := s.Length(); x != 1 {
		t.Error("Length should be one")
	}
}
