package administrations

import (
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

// GetTransportHistories get every transport histories in the database
func GetTransportHistories(w http.ResponseWriter, r *http.Request) {
	_, status, err := utils.CheckTokenAndScope(r, databases.IsAdmin)
	if err != nil {
		utils.Error(w, err, status)
		return
	}
	offset, err := utils.GetPageOffset(r)
	transportHistories := common.GetTransportHistories(offset)
	utils.WriteBody(w, transportHistories, http.StatusOK)
}
