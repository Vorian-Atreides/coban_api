package administrations

import (
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

// GetAddresses get every addresses in the database
func GetAddresses(w http.ResponseWriter, r *http.Request) {
	_, status, err := utils.CheckTokenAndScope(r, databases.IsAdmin)
	if err != nil {
		utils.Error(w, err, status)
		return
	}

	addresses := common.GetAddresses()
	utils.WriteBody(w, addresses, http.StatusOK)
}
