package models

import (
	"fmt"

	u "../utils"
)

//
type CharacterSkill struct {
	ID          int    `json:"id"`
	AttributeID int    `json:"attribute_id"`
	Name        string `json:"name"`
}

func (cs *CharacterSkill) Validate() (map[string]interface{}, bool) {

	if cs.Name == "" {
		return u.Message(false, "Please enter all fields"), false
	}

	return u.Message(true, "success"), true
}

func GetCharacterSkills() []CharacterSkill {
	var css []CharacterSkill
	database := GetDB()
	rows, err := database.Query("SELECT id, name, attribute_id FROM character_skills")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var cs CharacterSkill
		err = rows.Scan(&cs.ID, &cs.Name, &cs.AttributeID)
		if err != nil {
			fmt.Println(err)
		}
		css = append(css, cs)
	}

	rows.Close() //good habit to close

	return css
}

func (cs *CharacterSkill) Create() map[string]interface{} {

	if res, ok := cs.Validate(); !ok {
		return res
	}

	statement, err := GetDB().Prepare("INSERT INTO `character_skills` (`name`,`attribute_id`) VALUES (?,?)")
	if err != nil {
		fmt.Println(err)
	}
	result, err := statement.Exec(cs.Name, cs.AttributeID)

	if err != nil {
		fmt.Println(err)
	}
	res := u.Message(true, "success")
	lastid, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	cs.ID = int(lastid)
	res["data"] = cs
	return res
}
