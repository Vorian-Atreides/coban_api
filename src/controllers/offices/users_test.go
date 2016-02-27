package offices_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"coban/api/src/databases"
	"coban/api/src/utils"
)

type userTestSuite struct {
	suite.Suite
}

func TestUser(t *testing.T) {
	utils.InitTest()
	suite.Run(t, new(userTestSuite))
}

func (s *userTestSuite) Test01GetEmployees_ValidUser() {
	expected := `[{"id":4,"first-name":"徳川","last-name":"家康",` +
		`"account":null,"company":null}]`
	token, _ := utils.GenerateToken(4, databases.OfficeScope)

	url := fmt.Sprintf("%s/offices/users", utils.Address)
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

func (s *userTestSuite) Test02GetEmployees_InvalidUser() {
	token, _ := utils.GenerateToken(1, databases.ClientScope)

	url := fmt.Sprintf("%s/offices/users", utils.Address)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, err := client.Do(request)
	defer result.Body.Close()

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusUnauthorized, result.StatusCode)
}

func (s *userTestSuite) Test03CreateEmployees_ValidOfficerValidUser() {
	expected := `{"id":7,"first-name":"Gaston","last-name":"Siffert",` +
		`"account":{"email":"gs060292@live.fr"},` +
		`"company":{"id":4,"name":"旭化成株式会社"},"device":{"is-paired":false}}`
	data := `{"first-name":"Gaston","last-name":"Siffert",` +
		`"email":"gs060292@live.fr","scope":"Office"}`
	token, _ := utils.GenerateToken(4, databases.OfficeScope)

	url := fmt.Sprintf("%s/offices/users", utils.Address)
	request, _ := http.NewRequest("POST", url, strings.NewReader(data))
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, _ := client.Do(request)
	defer result.Body.Close()

	body, err := ioutil.ReadAll(result.Body)
	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusCreated, result.StatusCode)
	s.Equal(expected, string(body))
}

func (s *userTestSuite) Test04CreateEmployees_InvalidOfficerValidUser() {
	data := `{"first-name":"Gaston","last-name":"Siffert",` +
		`"email":"gs060292@live.fr","scope":"Office"}`
	token, _ := utils.GenerateToken(1, databases.ClientScope)

	url := fmt.Sprintf("%s/offices/users", utils.Address)
	request, _ := http.NewRequest("POST", url, strings.NewReader(data))
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, err := client.Do(request)
	defer result.Body.Close()

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusUnauthorized, result.StatusCode)
}

func (s *userTestSuite) Test05CreateEmployees_ValidOfficerInvalidUser() {
	data := `{"first-name":"Gaston","last-name":"Siffert",scope":"Office"}`
	token, _ := utils.GenerateToken(4, databases.OfficeScope)

	url := fmt.Sprintf("%s/offices/users", utils.Address)
	request, _ := http.NewRequest("POST", url, strings.NewReader(data))
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, err := client.Do(request)
	defer result.Body.Close()

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusBadRequest, result.StatusCode)
}

func (s *userTestSuite) Test06UpdateEmployees_ValidOfficerValidUser() {
	expected := `{"id":7,"first-name":"Alfred","last-name":"Dupont",` +
		`"account":null,"company":null}`
	data := `{"first-name":"Alfred","last-name":"Dupont",` +
		`"email":"gs060292@live.fr","scope":"Office"}`
	token, _ := utils.GenerateToken(4, databases.OfficeScope)

	url := fmt.Sprintf("%s/offices/users/7", utils.Address)
	request, _ := http.NewRequest("PUT", url, strings.NewReader(data))
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

func (s *userTestSuite) Test07UpdateEmployees_InvalidOfficerValidUser() {
	data := `{"first-name":"Alfred2","last-name":"Dupont2",` +
		`"email":"gs060292@live.fr","scope":"Office"}`
	token, _ := utils.GenerateToken(1, databases.ClientScope)

	url := fmt.Sprintf("%s/offices/users/4", utils.Address)
	request, _ := http.NewRequest("PUT", url, strings.NewReader(data))
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, err := client.Do(request)
	defer result.Body.Close()

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusUnauthorized, result.StatusCode)
}

func (s *userTestSuite) Test08UpdateEmployees_ValidOfficerInvalidUser() {
	data := `{"first-name":"Alfred2","last-name":"Dupont2",scope":"Office"}`
	token, _ := utils.GenerateToken(4, databases.OfficeScope)

	url := fmt.Sprintf("%s/offices/users/1", utils.Address)
	request, _ := http.NewRequest("PUT", url, strings.NewReader(data))
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, err := client.Do(request)
	defer result.Body.Close()

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusUnauthorized, result.StatusCode)
}

func (s *userTestSuite) Test09DeleteEmployee_InvalidOfficerValidUser() {
	token, _ := utils.GenerateToken(1, databases.ClientScope)

	url := fmt.Sprintf("%s/offices/users/1", utils.Address)
	request, _ := http.NewRequest("DELETE", url, nil)
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, err := client.Do(request)
	defer result.Body.Close()

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusUnauthorized, result.StatusCode)
}

func (s *userTestSuite) Test10DeleteEmployee_ValidOfficerInvalidUser() {
	token, _ := utils.GenerateToken(4, databases.OfficeScope)

	url := fmt.Sprintf("%s/offices/users/1", utils.Address)
	request, _ := http.NewRequest("DELETE", url, nil)
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, err := client.Do(request)
	defer result.Body.Close()

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusUnauthorized, result.StatusCode)
}

func (s *userTestSuite) Test11DeleteEmployee_ValidOfficerValidUser() {
	token, _ := utils.GenerateToken(4, databases.OfficeScope)

	url := fmt.Sprintf("%s/offices/users/4", utils.Address)
	request, _ := http.NewRequest("DELETE", url, nil)
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, err := client.Do(request)
	defer result.Body.Close()

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusNoContent, result.StatusCode)
}
