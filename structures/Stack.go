package structures

import "errors"

var (
	errEmptyStack = errors.New("stack is empty")
	errStackFull = errors.New("stack is full")
)
type Stack struct {
	list *LinkedList
}

// NewStack factory to generate new stacks
func NewStack() *Stack {
	return &Stack{list: NewLinkedList()}
}

// Size returns size of the stack
func (s *Stack) size() int {
	return s.list.Size()
}

func (s *Stack) Push(value interface{}) {
	s.list.Add(value)
}

func (s *Stack) isEmpty() bool {
	return s.list.IsEmpty()
}


func (s *Stack) Pull() (interface{}, error) {
	if s.isEmpty(){
		return nil, errEmptyStack
	}
	return s.list.RemoveAt(s.list.Size() - 1)
}

