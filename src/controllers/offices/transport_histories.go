package offices

import (
	"errors"
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

// GetTransportHistories get every transport histories from each employees
// working for the same company than the current user
func GetTransportHistories(w http.ResponseWriter, r *http.Request) {
	user, status, err := utils.CheckTokenAndScope(r, databases.IsOffice)
	if err != nil {
		utils.Error(w, err, status)
		return
	}

	offset, err := utils.GetPageOffset(r)
	begin, _ := utils.GetDateParameter(r, "begin")
	end, _ := utils.GetDateParameter(r, "end")

	transportHistories := common.GetTransportHistoriesByCompanyID(offset,
		begin, end, user.CompanyID)
	for i := range transportHistories {
		transportHistories[i].LoadRelated()
	}

	utils.WriteBody(w, transportHistories, http.StatusOK)
}

// GetTransportHistoryByUser get every transport histories for an user working
// for the same company than the current user
func GetTransportHistoryByUser(w http.ResponseWriter, r *http.Request) {
	current, status, err := utils.CheckTokenAndScope(r, databases.IsOffice)
	if err != nil {
		utils.Error(w, err, status)
		return
	}

	id, err := utils.GetUINT64Parameter(r, "id")
	if err != nil {
		utils.Error(w, err, http.StatusBadRequest)
		return
	}
	user, err := common.GetUserByID(uint(id))
	if err != nil {
		utils.Error(w, err, http.StatusBadRequest)
		return
	}
	if user.CompanyID != current.CompanyID {
		utils.Error(w,
			errors.New("You don't have the right to access this user."),
			http.StatusUnauthorized)
		return
	}

	offset, err := utils.GetPageOffset(r)
	begin, _ := utils.GetDateParameter(r, "begin")
	end, _ := utils.GetDateParameter(r, "end")

	transportHistories := common.GetTransportHistoriesByUserID(offset,
		begin, end, user.ID)
	for i := range transportHistories {
		transportHistories[i].LoadRelated()
	}

	utils.WriteBody(w, transportHistories, http.StatusOK)
}
