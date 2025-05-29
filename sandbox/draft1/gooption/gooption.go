package gooption

import "github.com/fengdotdev/golibs-traits/trait"

// implementation of the Option trait for GoOption
var _ trait.Option[string] = &GoOption[string]{}

// implementation of the DataTransferObject trait for GoOptionDTO
var _ trait.DataTransferObject[GoOptionDTO[string]] = &GoOption[string]{}

// implementation of the Cloneable trait for GoOption
var _ trait.Cloneable[GoOption[string]] = &GoOption[string]{}

var _ trait.Comparable = &GoOption[string]{}

// implementation of the JSONSerializer trait for GoOption
var _ trait.JSONSerializer = &GoOption[string]{}

// GoOption is a generic type that implements the Option trait.
// It can hold a value of type T or be None.
// It is used to represent optional values in a type-safe manner.

type GoOption[T any] struct {
	value  T
	isSome bool
}

