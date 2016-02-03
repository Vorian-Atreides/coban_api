package administrations

import (
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

func GetAddresses(w http.ResponseWriter, r *http.Request) {
	addresses := common.GetAddresses()

	utils.WriteBody(w, addresses)
}

func CreateAddress(w http.ResponseWriter, r *http.Request) {
	var address databases.Address

	address.FromBody(r)
	address, err := common.CreateAddress(address.Street, address.Zip, address.City, address.CompanyID)

	if err != nil {
		utils.Error(w, err)
		return
	}
	utils.WriteBody(w, address)
}