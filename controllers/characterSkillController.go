package controllers

import (
	"encoding/json"
	"net/http"

	"../models"

	u "../utils"
)

var CreateCharacterSkill = func(w http.ResponseWriter, r *http.Request) {
	//user := r.Context().Value("user").(uint)
	cs := &models.CharacterSkill{}

	err := json.NewDecoder(r.Body).Decode(cs)
	if err != nil {
		u.Respond(w, u.Message(false, "error where decoding req"))
	}

	//contact.UserId = user
	res := cs.Create()
	u.Respond(w, res)
}

var GetCharacterSkills = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetCharacterSkills()

	res := u.Message(true, "success")
	if len(data) > 0 {

		res["data"] = data

	}

	u.Respond(w, res)

}
