package pool

import (
	"container/list"
	"sync"
)

type Stack struct {
	mux sync.Mutex
	l   *list.List
}

func NewStack() *Stack {
	s := &Stack{
		l:   list.New(),
	}

	s.l.Init()
	return s
}

func (s *Stack) Push(fun func(string) error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.l.PushFront(fun)
}

func (s *Stack) Pull() func(string) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	f := s.l.Back()
	if f != nil {
		s.l.Remove(f)
		return f.Value.(func(string) error)
	}

	return nil
}

func (s *Stack) Len() int {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.l.Len()
}
