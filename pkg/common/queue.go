package common

type Queue[T any] struct {
	inner   []T
	maxSize int
}

func MakeQueue[T any](max int) Queue[T] {
	return Queue[T]{make([]T, 0), max}
}

func (q Queue[T]) Add(item T) Queue[T] {
	// append copy to the queue
	q.inner = append(q.inner, item)
	// if the queue is over length, remove the first item
	if len(q.inner) > q.maxSize {
		q.inner = q.inner[1:]
	}
	return q
}

func (q Queue[T]) Get(index int) T {
	return q.inner[index]
}

func (q Queue[T]) Length() int {
	return len(q.inner)
}
