package models

//Character is
type Character struct {
	Name       string `json:"name"`
	Class      CharacterClass
	Attrubutes []CharacterAttribute
	// Skills     []CharacterSkill
	// History    []CharacterHistory
}
