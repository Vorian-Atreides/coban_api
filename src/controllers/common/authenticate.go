package common

import (
	"fmt"
	"net/http"

	"coban/api/src/utils"
)

type authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthenticateRequest is used for generating a token for the user
func AuthenticateRequest(w http.ResponseWriter, r *http.Request) {
	var auth authentication

	err := utils.ReadBody(r, &auth)
	if err != nil {
		utils.Error(w, err, http.StatusBadRequest)
		return
	}

	token, err := Authenticate(auth.Email, auth.Password)
	if err != nil {
		utils.Error(w, err, http.StatusUnauthorized)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, token)
}
