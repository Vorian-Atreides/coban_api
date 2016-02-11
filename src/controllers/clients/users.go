package clients

import (
	"errors"
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

type updatePassword struct {
	OldPassword string `json:"old-password"`
	Password1   string `json:"password-1"`
	Password2   string `json:"password-2"`
}

// GetCurrentUser get the information about the current user
func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	user, status, err := utils.CheckTokenAndScope(r, databases.IsClient)
	if err != nil {
		utils.Error(w, err, status)
		return
	}
	user.LoadRelated()

	utils.WriteBody(w, user, http.StatusOK)
}

// UpdatePassword Update the password for the current user
//{
//	"old-password":"password",
//	"password-1":"new_password",
//	"password-2":"new_password"
//}
//
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	var update updatePassword
	utils.ReadBody(r, &update)

	user, status, err := utils.CheckTokenAndScope(r, databases.IsClient)
	if err != nil {
		utils.Error(w, err, status)
		return
	}

	if update.Password1 != update.Password2 {
		utils.Error(w,
			errors.New("You haven't typed the same password."),
			http.StatusBadRequest)
		return
	}

	if utils.HashPassword(update.OldPassword) != user.Account.Password {
		utils.Error(w,
			errors.New("Your previous password is wrong."),
			http.StatusBadRequest)
		return
	}

	_, err = common.UpdateAccountPassword(update.Password1, user.AccountID)
	if err != nil {
		utils.Error(w, err, http.StatusBadRequest)
		return
	}

	utils.WriteBody(w, user, http.StatusOK)
}
