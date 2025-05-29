package gooption

// GoOption  constructors

// NewGOption creates a new GoOption with the given value and isSome flag.
func NewGOption[T any](value T, isSome bool) *GoOption[T] {
	return &GoOption[T]{
		value:  value,
		isSome: isSome,
	}
}

// Some creates a GoOption that holds a value of type T.
func Some[T any](value T) Option[T] {
	return &GoOption[T]{
		value:  value,
		isSome: true,
	}
}

// None creates a GoOption that does not hold a value (None).
func None[T any]() Option[T] {
	var zero T
	return &GoOption[T]{
		value:  zero, // zero value of T
		isSome: false,
	}
}

// DTO constructors

// NewDTO creates a GoOptionDTO with the provided value and a flag indicating if it is Some or None.
// This function is used to create a DTO for the GoOption type.
// It sets the `Is` field to "S" for Some and "N" for None.
// This is useful for serialization and deserialization of the GoOption type.
func NewDTO[T any](value T, isSome bool) GoOptionDTO[T] {
	var dto GoOptionDTO[T]
	dto.Value = value

	if isSome {
		dto.Is = SomeConst
	} else {
		dto.Is = NoneConst
	}

	return dto
}

// SomeDTO creates a GoOptionDTO with a value and marks it as Some.
func SomeDTO[T any](value T) GoOptionDTO[T] {
	return NewDTO(value, true)
}

// NoneDTO creates a GoOptionDTO with the zero value of T and marks it as None.
func NoneDTO[T any]() GoOptionDTO[T] {
	var zero T
	return NewDTO(zero, false)
}
