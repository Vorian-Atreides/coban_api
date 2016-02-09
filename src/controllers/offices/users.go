package offices

import (
	"errors"
	"math/rand"
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
	"log"
)

type creationUser 	struct {
	FirstName	string	`json:"first-name"`
	LastName	string	`json:"last-name"`
	Email		string	`json:"email"`
	Scope		string  `json:"scope"`
}

type updateUser 	struct {
	UserID		uint	`json:"user-id"`
	FirstName	string	`json:"first-name"`
	LastName	string	`json:"last-name"`
	Email		string	`json:"email"`
	Scope		string  `json:"scope"`
}

var scopes = map[string]byte {
	"Office": databases.OfficeScope,
	"Client": databases.ClientScope,
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	company, err := GetCurrentCompany(r)
	if err != nil {
		utils.Error(w, err)
		return
	}

	utils.WriteBody(w, company.Employees)
}

func getScope(scope string) (byte, error) {
	if value, ok := scopes[scope]; ok {
		return value, nil
	}
	return 0, errors.New("This scope isn't valid.")
}

func generateRandomPassword(length uint) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func sendPasswordEmail(password string, to string) {
	log.Println("Password: ", password, " to: ", to)
}

// Scope: "Office", "Client"
//
//{
//	"first-name":"Gaston",
//	"last-name":"Siffert",
//	"email":"gs060292@live.fr",
//	"scope":"Office"
//}
//

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	var creation creationUser
	databases.ReadBody(r, &creation)

	company, err := GetCurrentCompany(r)
	if err != nil {
		utils.Error(w, err)
		return
	}
	scope, err := getScope(creation.Scope)
	if err != nil {
		utils.Error(w, err)
		return
	}

	password := generateRandomPassword(20)
	account, err := common.CreateAccount(creation.Email, scope, password)
	if err != nil {
		utils.Error(w, err)
		return
	}

	user, err := common.CreateUser(creation.FirstName, creation.LastName, account.ID, company.ID)
	if err != nil {
		common.DeleteAccount(account.ID)
		utils.Error(w, err)
		return
	}
	sendPasswordEmail(password, account.Email)

	utils.WriteBody(w, user)
}

// Scope: "Office", "Client"
//
//{
//	"user-id":5
//	"first-name":"Gaston",
//	"last-name":"Siffert",
//	"email":"gs060292@live.fr",
//	"scope":"Office"
//}
//

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var update updateUser
	databases.ReadBody(r, &update)

	officer, err := utils.CheckTokenAndScope(r, databases.IsOffice)
	if err != nil {
		utils.Error(w, err)
		return
	}

	target, err := common.GetUserByID(update.UserID)
	if err != nil {
		utils.Error(w, err)
		return
	}
	if target.CompanyID != officer.CompanyID {
		utils.Error(w, errors.New("You don't have the right to modify this user."))
		return
	}

	scope, err := getScope(update.Scope)
	if err != nil {
		utils.Error(w, err)
		return
	}
	_, err = common.UpdateAccount(update.Email, scope, target.AccountID)
	if err != nil {
		utils.Error(w, err)
		return
	}
	user, err := common.UpdateUserByID(update.FirstName, update.LastName, target.CompanyID, target.ID)
	if err != nil {
		utils.Error(w, err)
		return
	}

	utils.WriteBody(w, user)
}