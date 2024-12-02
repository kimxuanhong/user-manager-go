package list

import (
	"encoding/json"
	"fmt"
)

type Array[T any] struct {
	elements []T
}

func NewArray[T any]() *Array[T] {
	return &Array[T]{elements: make([]T, 0)}
}

func AsArray[T any](elements []T) *Array[T] {
	list := NewArray[T]()
	if elements != nil && len(elements) > 0 {
		list.elements = elements
	}
	return list
}

func (l *Array[T]) Add(element ...T) {
	l.elements = append(l.elements, element...)
}

func (l *Array[T]) Remove(index int) error {
	if index < 0 || index >= len(l.elements) {
		return fmt.Errorf("index out of range")
	}
	l.elements = append(l.elements[:index], l.elements[index+1:]...)
	return nil
}

func (l *Array[T]) RemoveIf(predicate func(T) bool) *Array[T] {
	var newElements []T
	for _, element := range l.elements {
		if !predicate(element) {
			newElements = append(newElements, element)
		}
	}
	l.elements = newElements
	return l
}

func (l *Array[T]) Get(index int) (T, error) {
	var zeroValue T
	if index < 0 || index >= len(l.elements) {
		return zeroValue, fmt.Errorf("index out of range")
	}
	return l.elements[index], nil
}

func (l *Array[T]) Size() int {
	return len(l.elements)
}

func (l *Array[T]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *Array[T]) Filter(predicate func(T) bool) *Array[T] {
	filteredList := &Array[T]{}
	for _, element := range l.elements {
		if predicate(element) {
			filteredList.Add(element)
		}
	}
	return filteredList
}

func Map[T, R any](arr *Array[T], mapper func(T) R) *Array[R] {
	mapped := &Array[R]{}
	for _, element := range arr.elements {
		mapped.Add(mapper(element))
	}
	return mapped
}

func (l *Array[T]) ForEach(action func(T)) {
	for _, element := range l.elements {
		action(element)
	}
}

func (l *Array[T]) Contains(element T, predicate func(T, T) bool) bool {
	for _, e := range l.elements {
		if predicate(e, element) {
			return true
		}
	}
	return false
}

func (l *Array[T]) First() (T, error) {
	var zeroValue T
	if l.IsEmpty() {
		return zeroValue, fmt.Errorf("array is empty")
	}
	return l.elements[0], nil
}

func (l *Array[T]) Last() (T, error) {
	var zeroValue T
	if l.IsEmpty() {
		return zeroValue, fmt.Errorf("array is empty")
	}
	return l.elements[l.Size()-1], nil
}

func (l *Array[T]) Slice() []T {
	return l.elements
}

func (l *Array[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.elements)
}

func (l *Array[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &l.elements)
}
