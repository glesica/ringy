package ringy

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestNew(t *testing.T) {
	t.Run("should provide a queue with the correct capacity", func(t *testing.T) {
		q, err := New[int](5)
		assert.NoError(t, err)
		assert.Equal(t, 5, q.Cap())
	})
}

func TestQueue_Add(t *testing.T) {
	t.Run("should add to a new, empty queue", func(t *testing.T) {
		q := &queue[int]{
			data: make([]int, 4), // capacity == 3
		}

		err := q.Add(10)
		assert.NoError(t, err)

		assert.Equal(t, 10, q.data[0])
		assert.Equal(t, 0, q.front)
		assert.Equal(t, 1, q.back)
	})

	t.Run("should add to a nearly full queue", func(t *testing.T) {
		q := &queue[int]{
			data: make([]int, 4), // capacity == 3
			back: 2,
		}

		err := q.Add(10)
		assert.NoError(t, err)

		assert.Equal(t, 10, q.data[2])
		assert.Equal(t, 0, q.front)
		assert.Equal(t, 3, q.back)
	})

	t.Run("should add to a queue with swapped front / back", func(t *testing.T) {
		q := &queue[int]{
			data:  make([]int, 4), // capacity == 3
			front: 3,
			back:  1,
		}

		err := q.Add(10)
		assert.NoError(t, err)

		assert.Equal(t, 10, q.data[1])
		assert.Equal(t, 3, q.front)
		assert.Equal(t, 2, q.back)
	})

	t.Run("should error on a full queue", func(t *testing.T) {
		q := &queue[int]{
			data: make([]int, 4), // capacity == 3
			back: 3,
		}

		err := q.Add(10)
		assert.IsError(t, err, QueueFull)
	})

	t.Run("should error on a full queue with swapped front / back", func(t *testing.T) {
		q := &queue[int]{
			data:  make([]int, 4), // capacity == 3
			front: 3,
			back:  2,
		}

		err := q.Add(10)
		assert.IsError(t, err, QueueFull)
	})
}

func TestQueue_Cap(t *testing.T) {
	t.Run("should compute capacity when >0", func(t *testing.T) {
		q := &queue[int]{
			data: make([]int, 4), // capacity == 3
		}

		assert.Equal(t, 3, q.Cap())
	})

	t.Run("should compute capacity when ==0", func(t *testing.T) {
		q := &queue[int]{
			data: make([]int, 1), // capacity == 0
		}

		assert.Equal(t, 0, q.Cap())
	})
}

func TestQueue_Len(t *testing.T) {
	t.Run("should compute length for empty queue", func(t *testing.T) {
		q := &queue[int]{
			data: make([]int, 4), // capacity == 3
		}

		assert.Equal(t, 0, q.Len())
	})

	t.Run("should compute length for queue with swapped front / back", func(t *testing.T) {
		q := &queue[int]{
			data:  make([]int, 4), // capacity == 3
			front: 3,
			back:  2,
		}

		assert.Equal(t, 3, q.Len())
	})
}

func TestQueue_Pop(t *testing.T) {
	t.Run("should pop a value from a queue with values", func(t *testing.T) {
		q := &queue[int]{
			data: []int{10, 20, 0, 0},
			back: 2,
		}

		value, err := q.Pop()
		assert.Equal(t, 10, value)
		assert.NoError(t, err)
		assert.Equal(t, 1, q.front)
	})

	t.Run("should pop a value from a queue with swapped front / back", func(t *testing.T) {
		q := &queue[int]{
			data:  []int{20, 0, 0, 10},
			front: 3,
			back:  1,
		}

		value, err := q.Pop()
		assert.Equal(t, 10, value)
		assert.NoError(t, err)
		assert.Equal(t, 0, q.front)
	})

	t.Run("should error on pop empty queue", func(t *testing.T) {
		q := &queue[int]{
			data: make([]int, 4), // capacity == 3
		}

		value, err := q.Pop()
		assert.Equal(t, 0, value)
		assert.IsError(t, err, QueueEmpty)
	})
}
