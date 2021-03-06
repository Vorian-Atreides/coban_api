package common

import (
	"errors"

	"coban/api/src/databases"
	"coban/api/src/utils"
)

// Authenticate try to authenticate an user with its email and its password
func Authenticate(email string, password string) (string, error) {
	var account databases.Account
	databases.DB.Where(&databases.Account{Email: email}).First(&account)

	if account.ID == 0 {
		return "", errors.New("This account doesn't exist.")
	}
	if account.Password != utils.HashPassword(password) {
		return "", errors.New("The credentials are invalid.")
	}
	return utils.GenerateToken(account.ID, account.Scope)
}

// GetAccounts get every accounts in the database
func GetAccounts(offset int) []databases.Account {
	var accounts []databases.Account

	databases.DB.Offset(offset).Limit(utils.PageSize).Find(&accounts)
	for i := range accounts {
		accounts[i].LoadRelated()
	}

	return accounts
}

// GetAccountByID Get an account by its ID
func GetAccountByID(id uint) (databases.Account, error) {
	var account databases.Account

	databases.DB.First(&account, id)
	if account.ID == 0 {
		return account, errors.New("This account doesn't exist.")
	}
	account.LoadRelated()

	return account, databases.DB.Error
}

// CreateAccount try to create a new account
func CreateAccount(email string, scope byte,
	password string) (databases.Account, error) {
	account := databases.Account{Email: email, Scope: scope,
		Password: utils.HashPassword(password)}

	if err := account.IsValid(false); err != nil {
		return account, err
	}
	databases.DB.Save(&account)

	return account, databases.DB.Error
}

// UpdateAccount try to update an account
func UpdateAccount(email string, scope byte,
	id uint) (databases.Account, error) {
	var account databases.Account

	databases.DB.First(&account, id)
	if account.ID == 0 {
		return account, errors.New("This account doesn't exist.")
	}
	account.Email = email
	account.Scope = scope
	if err := account.IsValid(false); err != nil {
		return account, err
	}
	databases.DB.Save(&account)

	return account, databases.DB.Error
}

// UpdateAccountPassword try to update th password from an account
func UpdateAccountPassword(password string, id uint) (databases.Account, error) {
	var account databases.Account

	databases.DB.First(&account, id)
	if account.ID == 0 {
		return account, errors.New("This account doesn't exist.")
	}
	account.Password = password
	if err := account.IsValid(true); err != nil {
		return account, err
	}
	account.Password = utils.HashPassword(account.Password)
	databases.DB.Save(&account)

	return account, databases.DB.Error
}

// DeleteAccount try to delete an account
func DeleteAccount(id uint) error {
	var account databases.Account

	databases.DB.First(&account, id)
	if account.ID == 0 {
		return errors.New("This account doesn't exist.")
	}
	databases.DB.Delete(&account)

	return databases.DB.Error
}
