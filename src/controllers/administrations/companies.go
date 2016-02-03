package administrations

import (
	"io/ioutil"
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/utils"
)

func GetCompanies(w http.ResponseWriter, r *http.Request) {
	addresses := common.GetCompanies()

	utils.WriteBody(w, addresses)
}

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.Error(w, err)
		return
	}

	company, err := common.CreateCompany(string(body))
	if err != nil {
		utils.Error(w, err)
		return
	}
	utils.WriteBody(w, company)
}