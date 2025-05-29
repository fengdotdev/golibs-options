package gooption

// Clone implements trait.Cloneable.
func (g *GoOption[T]) Clone() GoOption[T] {
	return GoOption[T]{
		value:  g.value,
		isSome: g.isSome,
	}
}
