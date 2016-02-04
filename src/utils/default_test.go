package utils_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"

	"coban/api/src/utils"
	"coban/api/src/databases"
)

type defaultTestSuite struct {
	suite.Suite
}

func TestDefault(t *testing.T) {
	suite.Run(t, new(defaultTestSuite))
}

func (s *defaultTestSuite) Test01CheckValidTokenAndScope() {
	var account databases.Account

	databases.DB.Where(databases.Account{Scope:databases.ClientScope}).First(&account)
	account.LoadRelated()

	tokenStr, _ := utils.GenerateToken(account.User.ID, account.Scope)
	request, _ := http.NewRequest("GET", "www.google.com", nil)
	request.Header.Add("Authorization", "Bearer " + tokenStr)

	user, err := utils.CheckTokenAndScope(request, databases.IsClient)
	s.NoError(err)
	s.Equal(account.User.ID, user.ID)
}

func (s *defaultTestSuite) Test02CheckUnauthorisedTokenAndScope() {
	var account databases.Account

	databases.DB.Where(databases.Account{Scope:databases.ClientScope}).First(&account)
	account.LoadRelated()

	tokenStr, _ := utils.GenerateToken(account.User.ID, account.Scope)
	request, _ := http.NewRequest("GET", "www.google.com", nil)
	request.Header.Add("Authorization", "Bearer " + tokenStr)

	_, err := utils.CheckTokenAndScope(request, databases.IsOffice)
	s.Error(err, "Unauthorised user.")
}

func (s *defaultTestSuite) Test03CheckUserNotFoundTokenAndScope() {
	tokenStr, _ := utils.GenerateToken(10, databases.ClientScope)
	request, _ := http.NewRequest("GET", "www.google.com", nil)
	request.Header.Add("Authorization", "Bearer " + tokenStr)

	_, err := utils.CheckTokenAndScope(request, databases.IsOffice)
	s.Error(err, "User not found.")
}