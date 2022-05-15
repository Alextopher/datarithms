package datarithims

// A Double-ended-queue that is implemented using a circular dynamic buffer.
type CircularBuffer[T any] struct {
	buffer []T
	head   int
	tail   int
	size   int
}

// Creates a new empty CircularBuffer.
func NewCircularBuffer[T any]() *CircularBuffer[T] {
	return &CircularBuffer[T]{buffer: make([]T, 0), head: 0, tail: 0, size: 0}
}

// Pushes a value onto the front of the CircularBuffer.
func (d *CircularBuffer[T]) PushFront(value T) {
	if d.size == len(d.buffer) {
		d.resize(d.size * 2)
	}

	d.head = (d.head - 1) % len(d.buffer)
	d.buffer[d.head] = value
	d.size++
}

func (d *CircularBuffer[T]) PushBack(value T) {
	if d.size == len(d.buffer) {
		d.resize(d.size * 2)
	}

	d.buffer[d.tail] = value
	d.tail = (d.tail + 1) % len(d.buffer)
	d.size++
}

func (d *CircularBuffer[T]) PopFront() T {
	if d.size == 0 {
		panic("CircularBuffer is empty")
	}

	value := d.buffer[d.head]
	d.head = (d.head + 1) % len(d.buffer)
	d.size--

	if d.size < len(d.buffer)/4 && len(d.buffer)/2 != 0 {
		d.resize(len(d.buffer) / 2)
	}

	return value
}

func (d *CircularBuffer[T]) PopBack() T {
	if d.size == 0 {
		panic("CircularBuffer is empty")
	}

	value := d.buffer[d.tail]
	d.tail = (d.tail - 1) % len(d.buffer)
	d.size--

	if d.size < len(d.buffer)/4 && len(d.buffer)/2 != 0 {
		d.resize(len(d.buffer) / 2)
	}

	return value
}

func (d *CircularBuffer[T]) PeekFront() T {
	if d.size == 0 {
		panic("CircularBuffer is empty")
	}

	return d.buffer[d.head]
}

func (d *CircularBuffer[T]) PeekBack() T {
	if d.size == 0 {
		panic("CircularBuffer is empty")
	}

	return d.buffer[d.tail]
}

func (d *CircularBuffer[T]) IsEmpty() bool {
	return d.size == 0
}

func (d *CircularBuffer[T]) Size() int {
	return d.size
}

func (d *CircularBuffer[T]) resize(newSize int) {
	newBuffer := make([]T, newSize)

	if d.head < d.tail {
		copy(newBuffer, d.buffer[d.head:d.tail])
	} else {
		copy(newBuffer, d.buffer[d.head:])
		copy(newBuffer[len(d.buffer)-d.head:], d.buffer[:d.tail])
	}

	d.buffer = newBuffer
	d.head = 0
	d.tail = d.size
}
