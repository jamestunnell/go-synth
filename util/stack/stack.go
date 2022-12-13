package stack

import "github.com/rs/zerolog/log"

type Stack[T any] struct {
	elems []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		elems: []T{},
	}
}

func (s *Stack[T]) Push(el T) {
	s.elems = append(s.elems, el)
}

func (s *Stack[T]) Pop() (T, bool) {
	n := len(s.elems)

	if n == 0 {
		return Zero[T](), false
	}

	val := s.elems[n-1]

	s.elems = s.elems[:n-1]

	return val, true
}

func (s *Stack[T]) Top() T {
	if len(s.elems) == 0 {
		log.Fatal().Msg("stack is empty")
	}

	return s.elems[0]
}

func Zero[T any]() T {
	var val T

	return val
}
