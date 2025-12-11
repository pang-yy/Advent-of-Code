package utils

type Heap[T any] struct {
	heapArray []T
}

func NewHeap[T any]() *Heap[T] {
	return &Heap[T]{}
}
