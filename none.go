package golibsoptions


type None[T any] struct{}

func NewNone[T any]() None[T] {
	return None[T]{}
}


//  content check
func (n None[T]) IsSome() bool {
	return false
}
func (n None[T]) IsNone() bool {
	return true
}

//  value access
func (n None[T]) Unwrap() T {
	panic("called Unwrap on a None value")
}
func (n None[T]) UnwrapOr(value T) T {
	return value
}

func (n None[T]) Expect(message string) T {
	panic(message)
}
