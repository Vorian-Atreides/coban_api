package offices

import (
	"errors"
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

func GetTransportHistories(w http.ResponseWriter, r *http.Request) {
	user, err := utils.CheckTokenAndScope(r, databases.IsOffice)
	if err != nil {
		utils.Error(w, err)
		return
	}

	var transportHistories []databases.TransportHistory
	databases.DB.Joins("left join users on users.id = transport_histories.user_id").
		Where(&databases.User{CompanyID: user.CompanyID}).
		Find(&transportHistories)
	for i := range transportHistories {
		transportHistories[i].LoadRelated()
	}

	utils.WriteBody(w, transportHistories)
}

func GetTransportHistoryByUser(w http.ResponseWriter, r *http.Request) {
	current, err := utils.CheckTokenAndScope(r, databases.IsOffice)
	if err != nil {
		utils.Error(w, err)
		return
	}

	id, err := utils.GetUINT64Parameter(r, "id")
	if err != nil {
		utils.Error(w, err)
		return
	}

	user, err := common.GetUserByID(uint(id))
	if err != nil {
		utils.Error(w, err)
		return
	}
	if user.CompanyID != current.CompanyID {
		utils.Error(w,
			errors.New("You don't have the right to access this user."))
		return
	}
	var transportHistories []databases.TransportHistory
	databases.DB.Where(databases.TransportHistory{UserID: user.ID}).
		Find(&transportHistories)
	for i := range transportHistories {
		transportHistories[i].LoadRelated()
	}

	utils.WriteBody(w, transportHistories)
}
