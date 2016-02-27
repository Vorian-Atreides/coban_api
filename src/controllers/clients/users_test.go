package clients_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

type usersTestSuite struct {
	suite.Suite
}

func TestUsers(t *testing.T) {
	utils.InitTest()
	suite.Run(t, new(usersTestSuite))
	common.UpdateAccountPassword("user", 1)
}

func (s *usersTestSuite) Test01GetCurrentUser_ValidUser() {
	expected := `{"id":1,"first-name":"青木","last-name":"真琳",` +
		`"account":{"email":"user@coban.jp"},` +
		`"company":{"id":1,"name":"アコム株式会社"},` +
		`"device":{"is-paired":false}}`
	token, _ := utils.GenerateToken(1, databases.ClientScope)

	url := fmt.Sprintf("%s/clients/users", utils.Address)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, _ := client.Do(request)
	defer result.Body.Close()

	body, err := ioutil.ReadAll(result.Body)
	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusOK, result.StatusCode)
	s.Equal(expected, string(body))
}

func (s *usersTestSuite) Test02GetCurrentUser_InvalidUser() {
	token, _ := utils.GenerateToken(2, databases.OfficeScope)

	url := fmt.Sprintf("%s/clients/users", utils.Address)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, _ := client.Do(request)
	defer result.Body.Close()

	s.Equal(http.StatusUnauthorized, result.StatusCode)
}

func (s *usersTestSuite) Test03UpdatePassword_ValidUserValidPassword() {
	data := `{"old-password":"user","password-1":"test","password-2":"test"}`
	token, _ := utils.GenerateToken(1, databases.ClientScope)

	url := fmt.Sprintf("%s/clients/users", utils.Address)
	request, _ := http.NewRequest("PUT", url, strings.NewReader(data))
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, _ := client.Do(request)
	defer result.Body.Close()

	s.Equal(http.StatusOK, result.StatusCode)
}

func (s *usersTestSuite) Test04UpdatePassword_ValidUserInvalidPassword() {
	token, _ := utils.GenerateToken(1, databases.ClientScope)

	data := `{"old-password":"user","password-1":"test","password-2":"test"}`
	url := fmt.Sprintf("%s/clients/users", utils.Address)
	request, _ := http.NewRequest("PUT", url, strings.NewReader(data))
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, _ := client.Do(request)
	defer result.Body.Close()
	s.Equal(http.StatusBadRequest, result.StatusCode)

	data = `{"old-password":"test","password-1":"pass","password-2":"pass2"}`
	request, _ = http.NewRequest("PUT", url, strings.NewReader(data))
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client = &http.Client{}
	result, _ = client.Do(request)
	defer result.Body.Close()
	s.Equal(http.StatusBadRequest, result.StatusCode)
}

func (s *usersTestSuite) Test05UpdatePassword_InvalidUserValidPassword() {
	data := `{"old-password":"user","password-1":"test","password-2":"test"}`
	token, _ := utils.GenerateToken(2, databases.OfficeScope)

	url := fmt.Sprintf("%s/clients/users", utils.Address)
	request, _ := http.NewRequest("PUT", url, strings.NewReader(data))
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, _ := client.Do(request)
	defer result.Body.Close()

	s.Equal(http.StatusUnauthorized, result.StatusCode)
}
