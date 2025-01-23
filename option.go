package golibsoptions


type Option[T any] interface {
	OptionContentCheck[T]
	OptionValueAccess[T]
	// OptionConvertions[T]

	// transform

	// conditionals

	// filters

	// iterators
}



type OptionContentCheck[T any] interface {
	// content check
	IsSome() bool
	IsNone() bool
}


type OptionValueAccess[T any] interface {
	// value access
	Unwrap() T
	UnwrapOr(T) T
	Expect(string) T
}

type OptionConvertions[T any] interface {
	// convertions
	Ok() Option[struct{}]
	OkOr(struct{}) Option[struct{}]
	OkOrElse(func() struct{}) Option[struct{}]
}