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

type accountsTestSuite struct {
	suite.Suite
}

func TestAccounts(t *testing.T) {
	utils.InitTest()
	suite.Run(t, new(accountsTestSuite))
}

func (s *accountsTestSuite) Test01GetAccounts_ValidUser() {
	expected := `[{"email":"user@coban.jp"},{"email":"office@coban.jp"},` +
		`{"email":"admin@coban.jp"},{"email":"root@coban.jp"},` +
		`{"email":"other@coban.jp"},{"email":"other2@coban.jp"}]`
	token, _ := utils.GenerateToken(3, databases.AdminScope)

	url := fmt.Sprintf("%s/administrations/accounts", utils.Address)
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

func (s *accountsTestSuite) Test02GetAccounts_InvalidUser() {
	token, _ := utils.GenerateToken(1, databases.ClientScope)

	url := fmt.Sprintf("%s/administrations/accounts", utils.Address)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, err := client.Do(request)
	defer result.Body.Close()

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusUnauthorized, result.StatusCode)
}
