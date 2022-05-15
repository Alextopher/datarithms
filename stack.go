package datarithims

// Implements a stack using a linked list.
type Stack[T any] struct {
	list Deque[T]
}

func NewListStack[T any]() *Stack[T] {
	return &Stack[T]{list: NewList[T]()}
}

func NewCircularBufferStack[T any]() *Stack[T] {
	return &Stack[T]{list: NewCircularBuffer[T]()}
}

func (s *Stack[T]) Push(value T) {
	s.list.PushBack(value)
}

func (s *Stack[T]) Pop() T {
	return s.list.PopBack()
}

func (s *Stack[T]) Peek() T {
	return s.list.PeekBack()
}
