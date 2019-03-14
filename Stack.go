package Stack

import (
	"fmt"
	"sync"
)

type items []interface {}

type Stack struct {
	items
	lock sync.Mutex
}

func NewStack(size int64) *Stack {
	return &Stack{
		items: make([]interface{}, size),
	}
}

func (s *Stack) Push (item interface{}) error {
	if item == nil {
		return fmt.Errorf("item had no value")
	}

	s.lock.Lock()
	s.items = append(s.items, item)
	s.lock.Unlock()
	return nil
}

func (s *Stack) Top() interface{} {
	var ret interface{}
	l := s.Length() - 1
	if l >= 0 {
		s.lock.Lock()
		ret = s.items[l]
		s.lock.Unlock()
	}
	return ret
}

func (s *Stack) Pop() interface{} {
	var ret interface{}
	l := s.Length() - 1
	if l >= 0 {
		s.lock.Lock()
		ret = s.items[l]
		s.items = append(s.items[:0], s.items[0:l]...)
		s.lock.Unlock()
	}
	return ret
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Length() int {
	return len(s.items)
}

func (s *Stack) ClearStack() {
	s.lock.Lock()
	s.items = nil
	s.lock.Unlock()
}