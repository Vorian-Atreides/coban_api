package administrations

import (
	"net/http"

	"coban/api/0.1/server/api/controllers/common"
	"coban/api/0.1/server/api/databases"
)

func GetAddresses(w http.ResponseWriter, r *http.Request) {
	addresses := common.GetAddresses()

	common.WriteBody(w, addresses)
}

func CreateAddress(w http.ResponseWriter, r *http.Request) {
	var address databases.Address

	address.FromBody(r)
	address, err := common.CreateAddress(address.Street, address.Zip, address.City, address.CompanyID)

	if err != nil {
		common.Error(w, err)
		return
	}
	common.WriteBody(w, address)
}