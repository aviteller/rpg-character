package controllers

import (
	"encoding/json"
	"net/http"

	"../models"

	u "../utils"
)

var CreateCharacterClass = func(w http.ResponseWriter, r *http.Request) {
	//user := r.Context().Value("user").(uint)
	cc := &models.CharacterClass{}

	err := json.NewDecoder(r.Body).Decode(cc)
	if err != nil {
		u.Respond(w, u.Message(false, "error where decoding req"))
	}

	//contact.UserId = user
	res := cc.Create()
	u.Respond(w, res)
}

var GetCharacterClasses = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetCharacterClasses()

	res := u.Message(true, "success")
	if len(data) > 0 {

		res["data"] = data

	}

	u.Respond(w, res)

}
