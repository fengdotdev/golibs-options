package golibsoptions

type Some[T any] struct {
	value T
}

func NewSome[T any](value T) Some[T] {
	return Some[T]{value: value}
}

// content check
func (s Some[T]) IsSome() bool {
	return true
}

func (s Some[T]) IsNone() bool {
	return false
}

// value access
func (s Some[T]) Unwrap() T {
	return s.value
}

func (s Some[T]) UnwrapOr(_ T) T {
	return s.value
}

func (s Some[T]) Expect(_ string) T {
	return s.value
}

// convertions
func (s Some[T]) Ok() Option[struct{}] {
	return NewSome(struct{}{})
}
