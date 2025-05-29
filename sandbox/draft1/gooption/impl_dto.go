package gooption




// GoOptionDTO is a Data Transfer Object for the GoOption type.
type GoOptionDTO[T any] struct {
	Value T      `json:"value"`
	Is    string `json:"is"`
}

const (
	SomeConst = "S"
	NoneConst = "N"
)

// FromDTO implements trait.DataTransferObject.
// It initializes the GoOption from a GoOptionDTO.
func (g *GoOption[T]) FromDTO(dto GoOptionDTO[T]) error {
	var zero T

	g.value = zero
	g.isSome = false

	isSome, err := isSome(dto.Is)
	if err != nil {
		return err
	}

	if isSome {
		g.value = dto.Value
		g.isSome = true
		return nil
	}

	return nil
}

// ToDTO implements trait.DataTransferObject.
// It converts the GoOption to a GoOptionDTO.
// If the GoOption is None, it returns a DTO with the zero value of T.
func (g *GoOption[T]) ToDTO() (GoOptionDTO[T], error) {
	var dto GoOptionDTO[T]

	if g.isSome {
		dto.Value = g.value
		dto.Is = SomeConst
		return dto, nil
	}

	var zero T
	dto.Value = zero
	dto.Is = NoneConst

	return dto, nil
}


// Helper function to determine if a string represents Some or None
func isSome(s string) (bool, error) {
	if s == SomeConst {
		return true, nil
	}

	if s == NoneConst {
		return false, nil
	}
	return false, ErrInvalidDTO
}
