package structure

import (
	"fmt"
	"reflect"
	"strings"
)

type DoubleList[T any] struct {
	head *DoubleNode[T]
}

func NewDoubleList[T any]() *DoubleList[T] {
	return &DoubleList[T]{
		head: nil,
	}
}

func (l *DoubleList[T]) AddFirst(e T) bool {
	newNode := NewDoubleNode(e)

	if l.head == nil {
		l.head = newNode
	} else {
		newNode.SetNext(l.head)
		l.head.SetPrevious(newNode)
		l.head = newNode
	}

	return true
}

func (l *DoubleList[T]) AddFinal(e T) bool {
	newNode := NewDoubleNode(e)

	if l.head == nil {
		l.head = newNode
	} else {
		aux := l.head
		for aux.GetNext() != nil {
			aux = aux.GetNext()
		}
		newNode.SetPrevious(aux)
		aux.SetNext(newNode)
	}

	return true
}

func (l *DoubleList[T]) RemoveFirst() bool {
	if l.head == nil {
		return false
	}

	if l.head.GetNext() == nil {
		l.head = nil
		return true
	}

	next := l.head.GetNext()
	next.SetPrevious(nil)
	l.head.SetNext(nil)
	l.head = next

	return true
}

func (l *DoubleList[T]) RemoveFinal() bool {
	if l.head == nil {
		return false
	}

	if l.head.GetNext() == nil {
		l.head = nil
		return true
	}

	aux := l.head
	for aux.GetNext() != nil {
		aux = aux.GetNext()
	}

	aux.GetPrevious().SetNext(nil)
	aux.SetPrevious(nil)

	return true
}

func (l *DoubleList[T]) Size() int {
	i := 0
	aux := l.head
	for aux != nil {
		i++
		aux = aux.GetNext()
	}
	return i
}

func (l *DoubleList[T]) Contains(o T) bool {
	aux := l.head
	for aux != nil {
		if reflect.DeepEqual(o, aux.GetValue()) {
			return true
		}
		aux = aux.GetNext()
	}
	return false
}

func (l *DoubleList[T]) IsEmpty() bool {
	return l.head == nil
}

func (l *DoubleList[T]) String() string {
	if l.head == nil {
		return "DoubleList [head=nil]"
	}

	var sb strings.Builder
	sb.WriteString("DoubleList [head=")
	aux := l.head
	for aux != nil {
		sb.WriteString(fmt.Sprintf("[%v]", aux.GetValue()))
		if aux.GetNext() != nil {
			sb.WriteString(" <-> ")
		}
		aux = aux.GetNext()
	}
	sb.WriteString("]")
	return sb.String()
}
