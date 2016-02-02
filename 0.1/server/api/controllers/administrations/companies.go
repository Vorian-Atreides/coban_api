package administrations

import (
	"coban/0.1/server/api/controllers/common"
	"io/ioutil"
	"net/http"
)

func GetCompanies(w http.ResponseWriter, r *http.Request) {
	addresses := common.GetCompanies()

	common.WriteBody(w, addresses)
}

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		common.Error(w, err)
		return
	}

	company, err := common.CreateCompany(string(body))
	if err != nil {
		common.Error(w, err)
		return
	}
	common.WriteBody(w, company)
}