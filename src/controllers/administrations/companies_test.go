package administrations_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

func (s *companiesTestSuite) Test01GetCompanies_ValidUser() {
	expected := `[{"id":1,"name":"アコム株式会社"},` +
		`{"id":2,"name":"株式会社愛知銀行"},` +
		`{"id":3,"name":"AOCホールディングス株式会社"},` +
		`{"id":4,"name":"旭化成株式会社"},` +
		`{"id":5,"name":"株式会社バンダイ"}]`
	token, _ := utils.GenerateToken(3, databases.AdminScope)

	url := fmt.Sprintf("%s/administrations/companies", utils.Address)
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

func (s *companiesTestSuite) Test02GetCompanies_InvalidUser() {
	token, _ := utils.GenerateToken(1, databases.ClientScope)

	url := fmt.Sprintf("%s/administrations/companies", utils.Address)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, err := client.Do(request)
	defer result.Body.Close()

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusUnauthorized, result.StatusCode)
}
