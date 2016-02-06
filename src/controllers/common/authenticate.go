package common

import (
	"fmt"
	"net/http"

	"coban/api/src/databases"
	"coban/api/src/utils"
)

func AuthenticateRequest(w http.ResponseWriter, r *http.Request) {
	var authentication Authentication

	err := databases.ReadBody(r, &authentication)
	if err != nil {
		utils.Error(w, err)
		return
	}
	token, err := Authenticate(authentication.Login, authentication.Password)
	if err != nil {
		utils.Error(w, err)
		return
	}

	fmt.Fprint(w, token)
}
