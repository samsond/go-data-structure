package stack

import "errors"

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	// implementation code
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}
	element := s.elements[len(s.elements)-1]
	return element, nil
}
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Size() (int, error) {
	return len(s.elements), nil
}
