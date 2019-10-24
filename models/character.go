package models

//Character is
type Character struct {
	Name       string `json:"name"`
	Class      CharacterClass
	Attrubutes []CharacterAttribute
	Skills     []CharacterSkill
	// History    []CharacterHistory
}

func GenerateRandomCharacter() Character {
	var rchar Character

	rchar.Name = "hello"

	return rchar
}
