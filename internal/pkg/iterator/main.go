package iterator

type Collector[T any] interface {
	CreateIterator() Iterator[T]
}

type Iterator[T any] interface {
	HasNext() bool
	GetNext() *T
	Peek() *T
	LookAhead(int) *T
}

type Collection[T any] struct {
	values []*T
}

type CollectionIterator[T any] struct {
	index  int
	values []*T
}

func (i *Collection[T]) CreateIterator() Iterator[T] {
	return &CollectionIterator[T]{
		values: i.values,
	}
}

func (i *CollectionIterator[T]) HasNext() bool {
	if i.index < len(i.values) {
		return true
	}
	return false
}

func (i *CollectionIterator[T]) GetNext() *T {
	if i.HasNext() {
		num := i.values[i.index]
		i.index++
		return num
	}
	return nil
}

func (i *CollectionIterator[T]) Peek() *T {
	if i.HasNext() {
		return i.values[i.index]
	}
	return nil
}

func (i *CollectionIterator[T]) LookAhead(n int) *T {
	if i.index+n < len(i.values) {
		return i.values[i.index+n]
	}
	return nil
}

func NewCollection[T any](values []T) Collection[T] {
	ptr := make([]*T, len(values))
	for i, v := range values {
		ptr[i] = &v
	}
	return Collection[T]{
		values: ptr,
	}
}
