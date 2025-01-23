package golibsoptions_test

import (
	"testing"

	option "github.com/fengdotdev/golibs-options"
	assert "github.com/fengdotdev/golibs-testing/assert"
)

func TestOptionInt(t *testing.T) {

	//some
	var number int = 10
	var someNumber option.Option[int] = option.NewSome(number)

	assert.True(t, someNumber.IsSome())

	//none

	var optionNumberNone option.Option[int] = option.NewNone[int]()

	assert.False(t, optionNumberNone.IsSome())
	assert.Equal(t, 0, optionNumberNone.UnwrapOr(0))

}

func TestOptionString(t *testing.T) {
	var name string = "John"
	var optionName option.Option[string] = option.NewSome(name)

	assert.True(t, optionName.IsSome())
	assert.Equal(t, name, optionName.Unwrap())

	//none

	var optionNameNone option.Option[string] = option.NewNone[string]()
	assert.False(t, optionNameNone.IsSome())
	assert.Equal(t, "", optionNameNone.UnwrapOr(""))

}

type PersonWithMayDog struct {
	Name string
	Dog  option.Option[string]
}

func TestOptionAsField(t *testing.T) {

	// person with dog
	doggo := "Rex"

	var personWithDog PersonWithMayDog = PersonWithMayDog{
		Name: "John",
		Dog:  option.NewSome(doggo),
	}

	assert.True(t, personWithDog.Dog.IsSome())
	assert.Equal(t, doggo, personWithDog.Dog.Unwrap())

	// person without dog
	var personWithoutDog PersonWithMayDog = PersonWithMayDog{
		Name: "John",
		Dog:  option.NewNone[string](),
	}

	assert.False(t, personWithoutDog.Dog.IsSome())
	assert.Equal(t, "", personWithoutDog.Dog.UnwrapOr(""))

}
