package common

import (
	"fmt"
	"net/http"

	"coban/api/src/databases"
	"coban/api/src/utils"
)

type authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AuthenticateRequest(w http.ResponseWriter, r *http.Request) {
	var auth authentication

	err := databases.ReadBody(r, &auth)
	if err != nil {
		utils.Error(w, err)
		return
	}
	token, err := Authenticate(auth.Email, auth.Password)
	if err != nil {
		utils.Error(w, err)
		return
	}

	fmt.Fprint(w, token)
}
