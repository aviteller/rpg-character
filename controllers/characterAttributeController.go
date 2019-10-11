package controllers

import (
	"encoding/json"
	"net/http"

	"../models"

	u "../utils"
)

var CreateCharacterAttribute = func(w http.ResponseWriter, r *http.Request) {
	//user := r.Context().Value("user").(uint)
	ca := &models.CharacterAttribute{}

	err := json.NewDecoder(r.Body).Decode(ca)
	if err != nil {
		u.Respond(w, u.Message(false, "error where decoding req"))
	}

	//contact.UserId = user
	res := ca.Create()
	u.Respond(w, res)
}

var GetCharacterAttributes = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetCharacterAttributes()

	res := u.Message(true, "success")
	if len(data) > 0 {

		res["data"] = data

	}

	u.Respond(w, res)

}
