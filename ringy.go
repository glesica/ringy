package ringy

type Queue[T any] interface {
	Add(T) error
	Cap() int
	Len() int
	Pop() (T, error)
}

func New[T any](cap int) (Queue[T], error) {
	return &queue[T]{
		data: make([]T, cap+1),
	}, nil
}

type queue[T any] struct {
	data  []T
	front int
	back  int
}

func (q *queue[T]) Add(value T) error {
	if q.Len() == q.Cap() {
		return QueueFull
	}

	// q.back is always empty
	q.data[q.back] = value

	// wrap q.back
	q.back = (q.back + 1) % len(q.data)

	return nil
}

func (q *queue[T]) Cap() int {
	return len(q.data) - 1
}

func (q *queue[T]) Len() int {
	if q.front <= q.back {
		return q.back - q.front
	}

	return len(q.data) - q.front + q.back
}

func (q *queue[T]) Pop() (T, error) {
	if q.Len() == 0 {
		return *new(T), QueueEmpty
	}

	value := q.data[q.front]

	// wrap q.front
	q.front = (q.front + 1) % len(q.data)

	return value, nil
}
