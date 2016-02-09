package administrations

// @SubApi Accounts  [/administrations/accounts]
// @SubApi Allows you access to different features of the accounts [/administrations/accounts]

import (
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

// @Title Get Accounts
// @Description Get every accounts
// @Accept json
// @Success 200 {array} common.Account
// @Failure 401 {array} error
// @Resource /accounts
// @Router /administrations/accounts [get]

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	_, err := utils.CheckTokenAndScope(r, databases.IsAdmin)
	if err != nil {
		utils.Error(w, err)
		return
	}

	accounts := common.GetAccounts()
	utils.WriteBody(w, accounts)
}