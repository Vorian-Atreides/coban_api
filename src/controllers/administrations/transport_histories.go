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
	offset, _ := utils.GetPageOffset(r)
	begin, _ := utils.GetDateParameter(r, "begin")
	end, _ := utils.GetDateParameter(r, "end")

	transportHistories := common.GetTransportHistories(offset, begin, end)
	utils.WriteBody(w, transportHistories, http.StatusOK)
}
