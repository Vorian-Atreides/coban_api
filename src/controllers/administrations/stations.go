package administrations

import (
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

// GetStations get every stations in the database
func GetStations(w http.ResponseWriter, r *http.Request) {
	_, status, err := utils.CheckTokenAndScope(r, databases.IsAdmin)
	if err != nil {
		utils.Error(w, err, status)
		return
	}
	offset, err := utils.GetPageOffset(r)
	stations := common.GetStations(offset)
	utils.WriteBody(w, stations, http.StatusOK)
}
