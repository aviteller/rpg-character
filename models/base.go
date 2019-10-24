package models

import (
	"database/sql"
	"fmt"
)

func InitTables() {
	GetDB().Exec("CREATE TABLE IF NOT EXISTS class_traits (id INTEGER PRIMARY KEY autoincrement, class_id int, attribute_id int, modifier int)")
	//INSERT INTO `class` (title) VALUES ("Warrior"),("Paladin"),("Barbarian"),("Monk"),("Ranger"),("Rogue"),("Mage"),("Healer"),("Driud");
	GetDB().Exec("CREATE TABLE IF NOT EXISTS class (id INTEGER PRIMARY KEY autoincrement, title text)")
	GetDB().Exec("CREATE TABLE IF NOT EXISTS character_attributes (id INTEGER PRIMARY KEY autoincrement, name text)")
	GetDB().Exec("CREATE TABLE IF NOT EXISTS character_skills (id INTEGER PRIMARY KEY autoincrement, name text, attribute_id int)")
}

//GetDB opens connection to sqlite db
func GetDB() *sql.DB {
	//:databaselocked.sqlite?cache=shared&mode=rwc
	database, err := sql.Open("sqlite3", "./rpg.db")
	if err != nil {
		fmt.Println(err)
	}

	return database
}
