package clients_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"

	"coban/api/src/databases"
	"coban/api/src/utils"
)

type usersTestSuite struct {
	suite.Suite
}

func TestUsers(t *testing.T) {
	utils.InitTest()
	suite.Run(t, new(usersTestSuite))
}

func (s *usersTestSuite) Test01GetCurrentUser_ValidUser() {
	expected := `{"id":1,"first-name":"青木","last-name":"真琳",`+
		`"account":{"email":"user@coban.jp"},`+
		`"company":{"id":1,"name":"アコム株式会社"},`+
		`"device":{"is-paired":false}}`
	token, _ := utils.GenerateToken(1, databases.ClientScope)

	url := fmt.Sprintf("%s/clients/users", utils.Address)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "Bearer" + " " + token)

	client := &http.Client{}
	result, _ := client.Do(request)
	defer result.Body.Close()

	body, err := ioutil.ReadAll(result.Body)
	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusOK, result.StatusCode)
	s.Equal(expected, string(body))
}