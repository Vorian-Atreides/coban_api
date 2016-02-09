package administrations

// @SubApi Addresses  [/administrations/addresses]
// @SubApi Allows you access to different features of the addresses [/administrations/addresses]

import (
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

func GetAddresses(w http.ResponseWriter, r *http.Request) {
	_, err := utils.CheckTokenAndScope(r, databases.IsAdmin)
	if err != nil {
		utils.Error(w, err)
		return
	}

	addresses := common.GetAddresses()
	utils.WriteBody(w, addresses)
}