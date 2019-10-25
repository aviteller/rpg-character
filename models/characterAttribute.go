package models

import (
	"fmt"

	u "../utils"
)

//CharacterAttribute is
type CharacterAttribute struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (ca *CharacterAttribute) Validate() (map[string]interface{}, bool) {

	if ca.Name == "" {
		return u.Message(false, "Please enter all fields"), false
	}

	return u.Message(true, "success"), true
}

func GetCharacterAttributes() []CharacterAttribute {
	var cas []CharacterAttribute
	database := GetDB()
	rows, err := database.Query("SELECT * FROM character_attributes")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var ca CharacterAttribute
		err = rows.Scan(&ca.ID, &ca.Name)

		if err != nil {
			fmt.Println(err)
		}
		cas = append(cas, ca)
	}

	rows.Close() //good habit to close

	return cas
}

func (ca *CharacterAttribute) Create() map[string]interface{} {
	if res, ok := ca.Validate(); !ok {
		return res
	}

	statement, err := GetDB().Prepare("INSERT INTO `character_attributes` (`name`) VALUES (?)")
	if err != nil {
		fmt.Println(err)
	}
	result, err := statement.Exec(ca.Name)

	if err != nil {
		fmt.Println(err)
	}
	res := u.Message(true, "success")
	lastid, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	ca.ID = int(lastid)
	res["data"] = ca
	return res
}
