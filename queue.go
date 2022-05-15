package datarithims

// A Queue that is implemented using a linked list.
type Queue[T any] struct {
	list *List[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{list: NewList[T]()}
}

func (q *Queue[T]) Push(value T) {
	q.list.PushBack(value)
}

func (q *Queue[T]) Pop() T {
	return q.list.PopFront()
}

func (q *Queue[T]) Peek() T {
	return q.list.PeekFront()
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}
