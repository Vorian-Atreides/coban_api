package clients

import (
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

func GetCurrentCompany(w http.ResponseWriter, r *http.Request) {
	user, err := utils.CheckTokenAndScope(r, databases.IsClient)
	if err != nil {
		utils.Error(w, err)
		return
	}

	company, err := common.GetCompanyByID(user.CompanyID)
	if err != nil {
		// Should not be possible.
		utils.Error(w, err)
		return
	}

	utils.WriteBody(w, company)
}
