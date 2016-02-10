package offices

import (
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

func GetAddresses(w http.ResponseWriter, r *http.Request) {
	company, err := GetCurrentCompany(r)
	if err != nil {
		utils.Error(w, err)
		return
	}
	company.LoadRelated()

	utils.WriteBody(w, company.Addresses)
}

func AddAddress(w http.ResponseWriter, r *http.Request) {
	user, err := utils.CheckTokenAndScope(r, databases.IsOffice)
	if err != nil {
		utils.Error(w, err)
		return
	}
	var data databases.Address
	data.FromBody(r)

	address, err := common.CreateAddress(data.Street, data.Zip,
		data.City, user.CompanyID)
	if err != nil {
		utils.Error(w, err)
		return
	}

	utils.WriteBody(w, address)
}
