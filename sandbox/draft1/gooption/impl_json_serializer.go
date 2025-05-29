package gooption

import "encoding/json"

// FromJSON implements trait.JSONSerializer.
func (g *GoOption[T]) FromJSON(s string) error {
	var dto GoOptionDTO[T]
	err := json.Unmarshal([]byte(s), &dto)

	if err != nil {
		return err
	}

	return g.FromDTO(dto)

}

// ToJSON implements trait.JSONSerializer.
func (g *GoOption[T]) ToJSON() (string, error) {
	dto, err := g.ToDTO()
	if err != nil {
		return "", err
	}

	data, err := json.Marshal(dto)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
