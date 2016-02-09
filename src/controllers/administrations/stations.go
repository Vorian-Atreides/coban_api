package administrations

import (
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

func GetStations(w http.ResponseWriter, r *http.Request) {
	_, err := utils.CheckTokenAndScope(r, databases.IsAdmin)
	if err != nil {
		utils.Error(w, err)
		return
	}

	stations := common.GetStations()
	utils.WriteBody(w, stations)
}
