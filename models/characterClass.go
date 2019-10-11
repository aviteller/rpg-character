package models

import (
	"fmt"

	u "../utils"
)

type CharacterClass struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	//Traits []ClassTrait
}

func (cc *CharacterClass) Validate() (map[string]interface{}, bool) {

	if cc.Title == "" {
		return u.Message(false, "Please enter all fields"), false
	}

	return u.Message(true, "success"), true
}

func GetCharacterClasses() []CharacterClass {
	var ccs []CharacterClass
	database := GetDB()
	rows, err := database.Query("SELECT id, title FROM class")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var cc CharacterClass
		err = rows.Scan(&cc.ID, &cc.Title)
		if err != nil {
			fmt.Println(err)
		}
		ccs = append(ccs, cc)
	}

	rows.Close() //good habit to close

	return ccs
}

func (cc *CharacterClass) Create() map[string]interface{} {

	if res, ok := cc.Validate(); !ok {
		return res
	}

	statement, err := GetDB().Prepare("INSERT INTO `class` (`title`) VALUES (?)")
	if err != nil {
		fmt.Println(err)
	}
	result, err := statement.Exec(cc.Title)

	if err != nil {
		fmt.Println(err)
	}
	res := u.Message(true, "success")
	lastid, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	cc.ID = int(lastid)
	res["class"] = cc
	return res
}
