package common

import (
	"errors"

	"coban/api/0.1/server/api/databases"
	"coban/api/0.1/server/api/utils"
)

const (
	ClientScope 	= 1
	OfficeScope		= 1 << 1
	AdminScope 		= 1 << 2
)

type IsScope func(byte) bool

func IsClient(scopes byte) bool {
	return scopes & ClientScope == ClientScope
}

func IsOffice(scopes byte) bool {
	return scopes & OfficeScope == OfficeScope
}

func IsAdmin(scopes byte) bool {
	return scopes & AdminScope == AdminScope
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

func GetAccounts() []databases.Account {
	var accounts []databases.Account

	databases.DB.Find(&accounts)
	for i, _ := range accounts {
		accounts[i].LoadRelated()
	}

	return accounts
}

func GetAccountByID(id uint) databases.Account {
	var account databases.Account

	databases.DB.First(&account, id)
	account.LoadRelated()

	return account
}

func CreateAccount(email string, scope byte, password string) (databases.Account, error) {
	account := databases.Account{Email:email, Scope:scope, Password: utils.HashPassword(password)}

	if err := account.IsValid(true); err != nil {
		return account, err
	}
	databases.DB.Save(&account)

	return account, nil
}

func UpdateAccount(email string, scope byte, id uint) (databases.Account, error) {
	account := databases.Account{Email:email, Scope:scope, ID:id}

	if err := account.IsValid(false); err != nil {
		return account, err
	}

	return account, nil
}

func UpdateAccountPassword(password string, id uint) (databases.Account, error) {
	var account databases.Account

	databases.DB.First(&account, id)
	account.Password = utils.HashPassword(password)
	if err := account.IsValid(false); err != nil {
		return account, err
	}
	databases.DB.Update(&account)

	return account, nil
}

func DeleteAccount(id uint) error {
	var account databases.Account

	databases.DB.First(&account, id)
	databases.DB.Delete(&account)

	if account.ID != 0 {
		return errors.New("This account can't be deleted.")
	}

	return nil
}