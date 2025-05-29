package gooption

import (
	"reflect"
)

// AreEqual implements trait.Comparable.
func (g *GoOption[T]) AreEqual(other any) bool {

	otherOption, ok := other.(*GoOption[T])
	if !ok {
		return false
	}

	valuesAreEqual := reflect.DeepEqual(g.value, otherOption.value)

	bothAreNone := !g.isSome && !otherOption.isSome // case where both are None but not equal zero values

	return valuesAreEqual || bothAreNone
}

// AreNotEqual implements trait.Comparable.
func (g *GoOption[T]) AreNotEqual(other any) bool {
	return !g.AreEqual(other)
}

// AreNotSameType implements trait.Comparable.
func (g *GoOption[T]) AreNotSameType(other any) bool {

	otherType := reflect.TypeOf(other)
	gType := reflect.TypeOf(g)

	areNotSameType := otherType != gType

	return areNotSameType
}

// AreSameType implements trait.Comparable.
func (g *GoOption[T]) AreSameType(other any) bool {
	otherType := reflect.TypeOf(other)
	gType := reflect.TypeOf(g)

	areSameType := otherType == gType

	return areSameType

}
