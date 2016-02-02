package common

import (
	"errors"

	"coban/api/0.1/server/api/databases"
	"coban/api/0.1/server/api/utils"
)

const (
	client 	= 1
	office	= 1 << 1
	admin 	= 1 << 2
)

func IsClient(scopes byte) bool {
	return scopes & client == client
}

func IsOffice(scopes byte) bool {
	println(scopes)
	return scopes & office == office
}

func IsAdmin(scopes byte) bool {
	return scopes & admin == admin
}

func BuildScope(scopes ...byte) byte {
	var scope byte

	for _, val := range scopes {
		scope |= val
	}
	return scope
}

func Authenticate(email string, password string, scope byte) (string, error) {
	var account databases.Account
	databases.DB.Where(&databases.Account{Email:email}).First(&account)

	if account.ID != 0 && account.Password == utils.HashPassword(password) {
		return utils.GenerateToken(account.ID, scope)
	}
	return "", errors.New("This account doesn't exist.")
}

func CreateAccount(email string, password string) (databases.Account, error) {
	account := databases.Account{Email:email, Password: utils.HashPassword(password)}

	if err := account.IsValid(true); err != nil {
		return account, err
	}
	databases.DB.Save(&account)
	return account, nil
}