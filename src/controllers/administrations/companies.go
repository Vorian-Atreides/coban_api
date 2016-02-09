package administrations

// @SubApi Companies  [/administrations/companies]
// @SubApi Allows you access to different features of the companies [/administrations/companies]

import (
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

func GetCompanies(w http.ResponseWriter, r *http.Request) {
	_, err := utils.CheckTokenAndScope(r, databases.IsAdmin)
	if err != nil {
		utils.Error(w, err)
		return
	}

	companies := common.GetCompanies()
	utils.WriteBody(w, companies)
}