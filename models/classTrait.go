package models

import (
	"fmt"

	u "../utils"
)

type ClassTrait struct {
	ID          int `json:"id"`
	ClassID     int `json:"class_id"`
	AttributeID int `json:"attribute_id"`
	Modifier    int `json:"modifier"`
}

func (ct *ClassTrait) Validate() (map[string]interface{}, bool) {

	if ct.ClassID == 0 || ct.AttributeID == 0 {
		return u.Message(false, "Please enter all fields"), false
	}

	return u.Message(true, "success"), true
}

func GetClassTraits() []ClassTrait {
	var cts []ClassTrait
	database := GetDB()
	rows, err := database.Query("SELECT * FROM class_traits")

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var ct ClassTrait
		err = rows.Scan(&ct.ID, &ct.ClassID, &ct.AttributeID, &ct.Modifier)
		if err != nil {
			fmt.Println(err)
		}

		cts = append(cts, ct)
	}

	rows.Close() //good habit to close

	return cts
}

func (ct *ClassTrait) Create() map[string]interface{} {
	if res, ok := ct.Validate(); !ok {
		return res
	}

	statement, err := GetDB().Prepare("INSERT INTO `class_traits` (`class_id`, `attribute_id`, `modifier`) VALUES (?,?,?)")
	if err != nil {
		fmt.Println(err)
	}
	result, err := statement.Exec(ct.ClassID, ct.AttributeID, ct.Modifier)

	if err != nil {
		fmt.Println(err)
	}
	res := u.Message(true, "success")
	lastid, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	ct.ID = int(lastid)
	res["data"] = ct
	return res
}
