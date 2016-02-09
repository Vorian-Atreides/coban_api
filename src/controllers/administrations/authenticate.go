package administrations

import (
	"fmt"
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	var authentication common.Authentication

	err := databases.ReadBody(r, &authentication)
	if err != nil {
		utils.Error(w, err)
		return
	}
	token, err := common.Authenticate(authentication.Login, authentication.Password)
	if err != nil {
		utils.Error(w, err)
		return
	}

	fmt.Fprint(w, token)
}
