package offices

import (
	"net/http"
	"coban/api/src/databases"
	"coban/api/src/utils"
	"coban/api/src/controllers/common"
)

type creationCompany struct {
	Name			string	`json:"name"`

	Administrator	struct {
		FirstName	string	`json:"first-name"`
		LastName	string	`json:"last-name"`
		Email		string	`json:"email"`
		Password	string	`json:"password"`
 	} 						`json:"administrator"`
}

func GetCurrentCompany(r *http.Request) (databases.Company, error) {
	user, err := utils.CheckTokenAndScope(r, databases.IsOffice)
	if err != nil {
		return databases.Company{}, err
	}

	company, err := common.GetCompanyByID(user.CompanyID)
	if err != nil {
		// Should not be possible.
		return company, err
	}

	return company, nil
}

func GetCompany(w http.ResponseWriter, r *http.Request) {
	company, err := GetCurrentCompany(r)
	if err != nil {
		utils.Error(w, err)
		return
	}

	utils.WriteBody(w, company)
}

//
//{
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
	databases.ReadBody(r, &creation)

	company, err := common.CreateCompany(creation.Name)
	if err != nil {
		utils.Error(w, err)
		return
	}

	account, err := common.CreateAccount(creation.Administrator.Email, databases.OfficeScope,
		creation.Administrator.Password)
	if err != nil {
		common.DeleteCompany(company.ID)
		utils.Error(w, err)
		return
	}

	_, err = common.CreateUser(creation.Administrator.FirstName, creation.Administrator.LastName,
		account.ID, company.ID)
	if err != nil {
		common.DeleteAccount(account.ID)
		common.DeleteCompany(company.ID)
		utils.Error(w, err)
		return
	}
	company.LoadRelated()

	utils.WriteBody(w, company)
}