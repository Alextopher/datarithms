// Implementations of a doubly linked list.
package datarithims

type Node[T any] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

// List[T] is a doubly linked list.
// It implements the dequeue interface.
type List[T any] struct {
	// The head of the list is the fist element
	Head *Node[T]
	// The tail of the list is the last element
	Tail *Node[T]
	// The number of elements in the list
	Size int
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) IsEmpty() bool {
	return l.Size == 0
}

func (l *List[T]) PushFront(value T) {
	node := &Node[T]{Value: value}

	if l.IsEmpty() {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head
		l.Head.Prev = node
		l.Head = node
	}

	l.Size++
}

func (l *List[T]) PushBack(value T) {
	node := &Node[T]{Value: value}

	if l.IsEmpty() {
		l.Head = node
		l.Tail = node
	} else {
		node.Prev = l.Tail
		l.Tail.Next = node
		l.Tail = node
	}

	l.Size++
}

func (l *List[T]) PopFront() T {
	if l.IsEmpty() {
		panic("List is empty")
	}

	value := l.Head.Value
	l.Head = l.Head.Next

	if l.Head == nil {
		l.Tail = nil
	} else {
		l.Head.Prev = nil
	}

	l.Size--

	return value
}

func (l *List[T]) PopBack() T {
	if l.IsEmpty() {
		panic("List is empty")
	}

	value := l.Tail.Value
	l.Tail = l.Tail.Prev

	if l.Tail == nil {
		l.Head = nil
	} else {
		l.Tail.Next = nil
	}

	l.Size--

	return value
}

func (l *List[T]) PeekFront() T {
	if l.IsEmpty() {
		panic("List is empty")
	}

	return l.Head.Value
}

func (l *List[T]) PeekBack() T {
	if l.IsEmpty() {
		panic("List is empty")
	}

	return l.Tail.Value
}

func FindFront[T comparable](list *List[T], value T) *Node[T] {
	for node := list.Head; node != nil; node = node.Next {
		if node.Value == value {
			return node
		}
	}

	return nil
}

func FindBack[T comparable](list *List[T], value T) *Node[T] {
	for node := list.Tail; node != nil; node = node.Prev {
		if node.Value == value {
			return node
		}
	}

	return nil
}
