package utils_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"

	"coban/api/src/databases"
	"coban/api/src/utils"
)

type defaultTestSuite struct {
	suite.Suite
}

func TestDefault(t *testing.T) {
	suite.Run(t, new(defaultTestSuite))
}

func (s *defaultTestSuite) Test01CheckValid_TokenAndScope() {
	var account databases.Account

	databases.DB.Where(databases.Account{Scope: databases.ClientScope}).
		First(&account)
	account.LoadRelated()

	tokenStr, _ := utils.GenerateToken(account.User.ID, account.Scope)
	request, _ := http.NewRequest("GET", "www.google.com", nil)
	request.Header.Add("Authorization", "Bearer "+tokenStr)

	user, _, err := utils.CheckTokenAndScope(request, databases.IsClient)
	s.NoError(err)
	s.Equal(account.User.ID, user.ID)
}

func (s *defaultTestSuite) Test02CheckUnauthorised_TokenAndScope() {
	var account databases.Account

	databases.DB.Where(databases.Account{Scope: databases.ClientScope}).First(&account)
	account.LoadRelated()

	tokenStr, _ := utils.GenerateToken(account.User.ID, account.Scope)
	request, _ := http.NewRequest("GET", "www.google.com", nil)
	request.Header.Add("Authorization", "Bearer "+tokenStr)

	_, _, err := utils.CheckTokenAndScope(request, databases.IsOffice)
	s.Error(err, "Unauthorised user.")
}

func (s *defaultTestSuite) Test03CheckUserNotFound_TokenAndScope() {
	tokenStr, _ := utils.GenerateToken(10, databases.ClientScope)
	request, _ := http.NewRequest("GET", "www.google.com", nil)
	request.Header.Add("Authorization", "Bearer "+tokenStr)

	_, _, err := utils.CheckTokenAndScope(request, databases.IsOffice)
	s.Error(err, "User not found.")
}

type test struct {
	Name  string `json:"name"`
	Name2 string `json:"name2"`
}

func (s *defaultTestSuite) Test04ReadValid_Body() {
	body := `{
		"name":"Test1",
		"name2":"Test2"
	}`
	request, _ := http.NewRequest("POST", "www.google.com",
		bytes.NewBuffer([]byte(body)))
	var data test
	err := utils.ReadBody(request, &data)

	s.NoError(err)
	s.Equal("Test1", data.Name)
	s.Equal("Test2", data.Name2)
}

func (s *defaultTestSuite) Test04ReadInvalid_Body() {
	body := `{
		"name":"Test1"
		"name2":"Test2"
	}`
	request, _ := http.NewRequest("POST", "www.google.com",
		bytes.NewBuffer([]byte(body)))
	var data test
	err := utils.ReadBody(request, &data)

	s.Error(err)
}
