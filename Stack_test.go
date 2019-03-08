package Stack

import (
	"testing"
)

var s Stack

func TestNewStack(t *testing.T) {
	s := NewStack(1)
	if len(s.items) != 1 {
		t.Error("Error created stack of size")
	}

	if cap(s.items) != 1 {
		t.Error("Capacity of stack is incorrect")
	}
}

func TestStack_Push(t *testing.T) {
	s.Push(0)
	if s.items[0] != 0 {
		t.Error("Error pushing to stack")
	}

	s.Push(1)
	if s.items[1] != 1 {
		t.Error("Error pushing to stack")
	}
}

func TestIsEmpty(t *testing.T) {
	x := s.IsEmpty()
	if !x {
		t.Error("Stack should be empty")
	}

	s.Push(1)
	x = s.IsEmpty()
	if x {
		t.Error("Stack should have a value")
	}

}

func TestTop(t *testing.T) {
	s.Push(0)
	s.Push(1)
	x := s.Top()
	if x != 1 {
		t.Error("Error retrieving top")
	}
}

func TestPop(t *testing.T) {
	s.Push(0)
	s.Push(1)
	x := s.Pop()
	if x != 1 {
		t.Error("Error with pop, Value should be 1")
	}
	s.Push(2)
	x = s.Pop()
	if x != 2 {
		t.Error("Error with pop, Value should be 2")

	}
}