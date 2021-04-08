package stackString

import (
	"sync"
)

type Stack struct {
	data []string
	lock sync.Mutex
}

func NewStack() *Stack {
	s := new(Stack)
	s.data = []string{}
	return s
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.lock.Lock()
	s.data = s.data[:len(s.data)-1]
	s.lock.Unlock()
}

func (s *Stack) Push(item string) {
	s.lock.Lock()
	s.data = append(s.data, item)
	s.lock.Unlock()
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Top() string {
	return s.data[len(s.data)-1]
}
