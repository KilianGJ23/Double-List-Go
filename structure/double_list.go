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

func (l *DoubleList[T]) Add(e T) bool {
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

func (l *DoubleList[T]) Remove(o T) bool {
	aux := l.head

	for aux != nil {
		if reflect.DeepEqual(o, aux.GetValue()) {
			if aux.GetPrevious() == nil {
				l.head = aux.GetNext()
				if l.head != nil {
					l.head.SetPrevious(nil)
				}
			} else {
				aux.GetPrevious().SetNext(aux.GetNext())
				if aux.GetNext() != nil {
					aux.GetNext().SetPrevious(aux.GetPrevious())
				}
			}
			return true
		}
		aux = aux.GetNext()
	}

	return false
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
