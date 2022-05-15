package datarithims

// Deque is a double-ended queue
type Deque[T any] interface {
	PushFront(value T)
	PushBack(value T)
	PopFront() T
	PopBack() T
	PeekFront() T
	PeekBack() T
	IsEmpty() bool
}
