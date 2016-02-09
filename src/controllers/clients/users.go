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

//{
//	"old-password":"password",
//	"password-1":"new_password",
//	"password-2":"new_password"
//}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	var update updatePassword
	databases.ReadBody(r, &update)

	user, err := utils.CheckTokenAndScope(r, databases.IsOffice)
	if err != nil {
		utils.Error(w, err)
		return
	}

	if update.Password1 != update.Password2 {
		utils.Error(w, errors.New("You haven't typed the same password."))
		return
	}

	if utils.HashPassword(update.OldPassword) != user.Account.Password {
		utils.Error(w, errors.New("Your previous password is wrong."))
		return
	}

	_, err = common.UpdateAccountPassword(update.Password1, user.AccountID)
	if err != nil {
		utils.Error(w, err)
		return
	}

	utils.WriteBody(w, user)
}
