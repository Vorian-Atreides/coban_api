package common

import (
	"errors"

	"coban/api/src/databases"
	"coban/api/src/utils"
)

// GetUsers get every users in the database
func GetUsers(offset uint) []databases.User {
	var users []databases.User

	databases.DB.Offset(offset).Limit(utils.PageSize).Find(&users)
	for i := range users {
		users[i].LoadRelated()
	}

	return users
}

// GetUserByID get an user by its ID
func GetUserByID(id uint) (databases.User, error) {
	var user databases.User

	databases.DB.First(&user, id)
	if user.ID == 0 {
		return user, errors.New("This user doesn't exist.")
	}
	user.LoadRelated()

	return user, databases.DB.Error
}

// CreateUser try to create a new user
func CreateUser(firstName string, lastName string, accountID uint,
	companyID uint) (databases.User, error) {
	user := databases.User{FirstName: firstName, LastName: lastName,
		CompanyID: companyID, AccountID: accountID}
	if err := user.IsValid(); err != nil {
		return user, err
	}
	databases.DB.Save(&user)
	user.LoadRelated()

	return user, databases.DB.Error
}

// UpdateUserByID try to update an user
func UpdateUserByID(firstName string, lastName string,
	companyID uint, id uint) (databases.User, error) {
	var user databases.User

	databases.DB.First(&user, id)
	if user.ID == 0 {
		return user, errors.New("This user doesn't exist.")
	}
	user.FirstName = firstName
	user.LastName = lastName
	user.CompanyID = companyID
	if err := user.IsValid(); err != nil {
		return user, err
	}
	databases.DB.Save(&user)

	return user, databases.DB.Error
}

// DeleteUser try to delete an user
func DeleteUser(id uint) error {
	var user databases.User
	databases.DB.First(&user, id)
	if user.ID == 0 {
		return errors.New("This user doesn't exist.")
	}
	if err := DeleteAccount(user.AccountID); err != nil {
		return err
	}
	databases.DB.Delete(user)

	return databases.DB.Error
}
