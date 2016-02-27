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

type companiesTestSuite struct {
	suite.Suite
}

func TestCompanies(t *testing.T) {
	utils.InitTest()
	suite.Run(t, new(companiesTestSuite))
}

func (s *companiesTestSuite) Test01GetCurrentCompany_ValidUser() {
	expected := `{"id":2,"name":"株式会社愛知銀行"}`
	token, _ := utils.GenerateToken(2, databases.OfficeScope)

	url := fmt.Sprintf("%s/offices/companies", utils.Address)
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

func (s *companiesTestSuite) Test02GetCurentCompany_InvalidUser() {
	token, _ := utils.GenerateToken(1, databases.ClientScope)

	url := fmt.Sprintf("%s/offices/companies", utils.Address)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, _ := client.Do(request)
	defer result.Body.Close()

	s.NotNil(result)
	s.Equal(http.StatusUnauthorized, result.StatusCode)
}

func (s *companiesTestSuite) Test03Create_ValidCompany() {
	expected := `{"id":7,"name":"Coban"}`
	data := `{"name":"Coban",` +
		`"administrator":{"first-name":"Tatsuya", "last-name":"Zembutsu",` +
		`"email":"tatsuya.zembutsu@coban.co.jp", "password":"coban-password"}}`

	url := fmt.Sprintf("%s/offices/companies", utils.Address)
	result, _ := http.Post(url, "application/json", strings.NewReader(data))
	defer result.Body.Close()

	body, err := ioutil.ReadAll(result.Body)
	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusCreated, result.StatusCode)
	s.Equal(expected, string(body))
}

func (s *companiesTestSuite) Test04Create_InvalidCompany() {
	data := `{"name":"Coban",` +
		`"administrator":{"first-name":"Tatsuya", "last-name":"Zembutsu",` +
		`"password":"coban-password"}}`

	url := fmt.Sprintf("%s/offices/companies", utils.Address)
	result, err := http.Post(url, "application/json", strings.NewReader(data))
	defer result.Body.Close()

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusBadRequest, result.StatusCode)
}

func (s *companiesTestSuite) Test05DeleteCurrentCompany_ValidUser() {
	token, _ := utils.GenerateToken(6, databases.OfficeScope)

	url := fmt.Sprintf("%s/offices/companies", utils.Address)
	request, _ := http.NewRequest("DELETE", url, nil)
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, err := client.Do(request)
	defer result.Body.Close()

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusNoContent, result.StatusCode)
}

func (s *companiesTestSuite) Test06DeleteCurrentCompany_InvalidUser() {
	token, _ := utils.GenerateToken(1, databases.ClientScope)

	url := fmt.Sprintf("%s/offices/companies", utils.Address)
	request, _ := http.NewRequest("DELETE", url, nil)
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, err := client.Do(request)
	defer result.Body.Close()

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusUnauthorized, result.StatusCode)
}
