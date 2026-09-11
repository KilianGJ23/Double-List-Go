package structure

import "fmt"

type DoubleNode[T any] struct {
	value    T
	previous *DoubleNode[T]
	next     *DoubleNode[T]
}

func NewDoubleNode[T any](value T) *DoubleNode[T] {
	return &DoubleNode[T]{
		value:    value,
		previous: nil,
		next:     nil,
	}
}

func (n *DoubleNode[T]) GetValue() T {
	return n.value
}

func (n *DoubleNode[T]) SetValue(value T) {
	n.value = value
}

func (n *DoubleNode[T]) GetPrevious() *DoubleNode[T] {
	return n.previous
}

func (n *DoubleNode[T]) SetPrevious(previous *DoubleNode[T]) {
	n.previous = previous
}

func (n *DoubleNode[T]) GetNext() *DoubleNode[T] {
	return n.next
}

func (n *DoubleNode[T]) SetNext(next *DoubleNode[T]) {
	n.next = next
}

func (n *DoubleNode[T]) String() string {
	if n == nil {
		return "<nil>"
	}
	if n.next != nil {
		return fmt.Sprintf("DoubleNode [value=%v, next=%v]", n.value, n.next.value)
	}
	return fmt.Sprintf("DoubleNode [value=%v, next=nil]", n.value)
}
