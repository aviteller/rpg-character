package controllers

import (
	"encoding/json"
	"net/http"

	"../models"

	u "../utils"
)

var CreateClassTrait = func(w http.ResponseWriter, r *http.Request) {
	//user := r.Context().Value("user").(uint)
	ct := &models.ClassTrait{}

	err := json.NewDecoder(r.Body).Decode(ct)
	if err != nil {
		u.Respond(w, u.Message(false, "error where decoding req"))
	}

	//contact.UserId = user
	res := ct.Create()
	u.Respond(w, res)
}

var GetClassTraits = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetClassTraits()

	res := u.Message(true, "success")
	if len(data) > 0 {

		res["data"] = data

	}

	u.Respond(w, res)

}
