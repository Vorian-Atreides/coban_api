package offices

import (
	"errors"
	"log"
	"math/rand"
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

type bodyUser struct {
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	Email     string `json:"email"`
	Scope     string `json:"scope"`
}

var scopes = map[string]byte{
	"Office": databases.OfficeScope,
	"Client": databases.ClientScope,
	"Both":   databases.ClientScope | databases.OfficeScope,
}

// GetEmployees Get every employees belonging to the same company than the user
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	company, status, err := GetCurrentCompany(r)
	if err != nil {
		utils.Error(w, err, status)
		return
	}

	offset, err := utils.GetPageOffset(r)
	var employees []databases.User
	databases.DB.Model(&company).Related(&employees).
		Offset(offset).Limit(utils.PageSize)

	utils.WriteBody(w, employees, http.StatusOK)
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

// AddEmployee Scope: "Office", "Client"
//
//{
//	"first-name":"Gaston",
//	"last-name":"Siffert",
//	"email":"gs060292@live.fr",
//	"scope":"Office"
//}
//
func AddEmployee(w http.ResponseWriter, r *http.Request) {
	var creation bodyUser
	utils.ReadBody(r, &creation)

	company, status, err := GetCurrentCompany(r)
	if err != nil {
		utils.Error(w, err, status)
		return
	}
	scope, err := getScope(creation.Scope)
	if err != nil {
		utils.Error(w, err, http.StatusBadRequest)
		return
	}

	password := generateRandomPassword(20)
	account, err := common.CreateAccount(creation.Email, scope, password)
	if err != nil {
		utils.Error(w, err, http.StatusBadRequest)
		return
	}

	user, err := common.CreateUser(creation.FirstName, creation.LastName, account.ID, company.ID)
	if err != nil {
		common.DeleteAccount(account.ID)
		utils.Error(w, err, http.StatusBadRequest)
		return
	}
	sendPasswordEmail(password, account.Email)

	utils.WriteBody(w, user, http.StatusCreated)
}

// UpdateEmployee Scope: "Office", "Client"
//
//{
//	"first-name":"Gaston",
//	"last-name":"Siffert",
//	"email":"gs060292@live.fr",
//	"scope":"Office"
//}
//
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var update bodyUser
	utils.ReadBody(r, &update)

	id, err := utils.GetUINT64Parameter(r, "id")
	if err != nil {
		utils.Error(w, err, http.StatusBadRequest)
		return
	}

	officer, status, err := utils.CheckTokenAndScope(r, databases.IsOffice)
	if err != nil {
		utils.Error(w, err, status)
		return
	}

	target, err := common.GetUserByID(uint(id))
	if err != nil {
		utils.Error(w, err, http.StatusBadRequest)
		return
	}
	if target.CompanyID != officer.CompanyID {
		utils.Error(w,
			errors.New("You don't have the right to modify this user."),
			http.StatusUnauthorized)
		return
	}

	scope, err := getScope(update.Scope)
	if err != nil {
		utils.Error(w, err, http.StatusBadRequest)
		return
	}
	_, err = common.UpdateAccount(update.Email, scope, target.AccountID)
	if err != nil {
		utils.Error(w, err, http.StatusBadRequest)
		return
	}
	user, err := common.UpdateUserByID(update.FirstName, update.LastName, target.CompanyID, target.ID)
	if err != nil {
		utils.Error(w, err, http.StatusBadRequest)
		return
	}

	utils.WriteBody(w, user, http.StatusOK)
}
