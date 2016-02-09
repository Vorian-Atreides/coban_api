package administrations

// @SubApi Transport Histories  [/administrations/transport-histories]
// @SubApi Allows you access to different features of the transport histories [/administrations/transport-histories]

import (
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

func GetTransportHistories(w http.ResponseWriter, r *http.Request) {
	_, err := utils.CheckTokenAndScope(r, databases.IsAdmin)
	if err != nil {
		utils.Error(w, err)
		return
	}

	transportHistories := common.GetTransportHistories()
	utils.WriteBody(w, transportHistories)
}
