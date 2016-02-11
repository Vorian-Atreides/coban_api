package offices

import (
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

type creationCompany struct {
	Name string `json:"name"`

	Administrator struct {
		FirstName string `json:"first-name"`
		LastName  string `json:"last-name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	} `json:"administrator"`
}

// GetCurrentCompany Get the company to whom belong the current user
func GetCurrentCompany(r *http.Request) (databases.Company, int, error) {
	user, status, err := utils.CheckTokenAndScope(r, databases.IsOffice)
	if err != nil {
		return databases.Company{}, status, err
	}

	company, err := common.GetCompanyByID(user.CompanyID)
	if err != nil {
		// Should not be possible.
		return company, http.StatusBadRequest, err
	}

	return company, 0, nil
}

// GetCompany Get the company belonging to the current user
func GetCompany(w http.ResponseWriter, r *http.Request) {
	company, status, err := GetCurrentCompany(r)
	if err != nil {
		utils.Error(w, err, status)
		return
	}

	utils.WriteBody(w, company, http.StatusOK)
}

// CreateCompany {
//	"name":"Coban2",
//	"administrator":
//	{
//		"first-name":"Gaston",
//		"last-name":"Siffert",
//		"email":"gs060292@live.fr",
//		"password":"password"
//	}
//}
//
func CreateCompany(w http.ResponseWriter, r *http.Request) {
	var creation creationCompany
	utils.ReadBody(r, &creation)

	company, err := common.CreateCompany(creation.Name)
	if err != nil {
		utils.Error(w, err, http.StatusBadRequest)
		return
	}

	account, err := common.CreateAccount(creation.Administrator.Email, databases.OfficeScope,
		creation.Administrator.Password)
	if err != nil {
		common.DeleteCompany(company.ID)
		utils.Error(w, err, http.StatusBadRequest)
		return
	}

	_, err = common.CreateUser(creation.Administrator.FirstName, creation.Administrator.LastName,
		account.ID, company.ID)
	if err != nil {
		common.DeleteAccount(account.ID)
		common.DeleteCompany(company.ID)
		utils.Error(w, err, http.StatusBadRequest)
		return
	}
	company.LoadRelated()

	utils.WriteBody(w, company, http.StatusCreated)
}
