package gooption

// IsNone implements trait.Option.
func (g *GoOption[T]) IsNone() bool {
	return !g.isSome
}

// IsSome implements trait.Option.
func (g *GoOption[T]) IsSome() bool {
	return g.isSome
}

// Unwrap implements trait.Option.
func (g *GoOption[T]) Unwrap() T {
	if !g.isSome {
		panic("called `GoOption.Unwrap` on a `None` value")
	}
	return g.value
}

// UnwrapOr implements trait.Option.
func (g *GoOption[T]) UnwrapOr(defaultValue T) T {
	if g.isSome {
		return g.value
	}
	return defaultValue
}

// UnwrapOrDefault implements trait.Option.
func (g *GoOption[T]) UnwrapOrDefault() T {
	if g.isSome {
		return g.value
	}
	var zero T
	return zero
}

// UnwrapOrElse implements trait.Option.
func (g *GoOption[T]) UnwrapOrElse(f func() T) T {
	if g.isSome {
		return g.value
	}
	return f()
}
