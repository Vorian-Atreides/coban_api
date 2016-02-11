package clients

import (
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

// GetCurrentCompany Get the company to whom belong the current user
func GetCurrentCompany(w http.ResponseWriter, r *http.Request) {
	user, status, err := utils.CheckTokenAndScope(r, databases.IsClient)
	if err != nil {
		utils.Error(w, err, status)
		return
	}

	company, err := common.GetCompanyByID(user.CompanyID)
	if err != nil {
		// Should not be possible.
		utils.Error(w, err, http.StatusBadRequest)
		return
	}

	utils.WriteBody(w, company, http.StatusOK)
}
