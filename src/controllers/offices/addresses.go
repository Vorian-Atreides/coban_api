package offices

import (
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

// GetAddresses get the addresses belonging to the current user's company
func GetAddresses(w http.ResponseWriter, r *http.Request) {
	company, status, err := GetCurrentCompany(r)
	if err != nil {
		utils.Error(w, err, status)
		return
	}

	offset, err := utils.GetPageOffset(r)
	var addresses []databases.Address
	databases.DB.Model(&company).Related(&addresses).
		Offset(offset).Limit(utils.PageSize)

	utils.WriteBody(w, addresses, http.StatusOK)
}

// AddAddress add an address to the current user's company
func AddAddress(w http.ResponseWriter, r *http.Request) {
	user, status, err := utils.CheckTokenAndScope(r, databases.IsOffice)
	if err != nil {
		utils.Error(w, err, status)
		return
	}
	var data databases.Address
	utils.ReadBody(r, &data)

	address, err := common.CreateAddress(data.Street, data.Zip,
		data.City, user.CompanyID)
	if err != nil {
		utils.Error(w, err, http.StatusBadRequest)
		return
	}

	utils.WriteBody(w, address, http.StatusCreated)
}
