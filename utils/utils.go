package utils

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

//Message is exported
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

//Respond is exported
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(data)
}

func Roll(dice, sides int) int {
	total := 0
	min := 1

	for index := 0; index < dice; index++ {
		total = total + (rand.Intn(sides-min+1) + min)
	}

	return total
}
