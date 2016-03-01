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

type usersTestSuite struct {
	suite.Suite
}

func TestUsers(t *testing.T) {
	utils.InitTest()
	suite.Run(t, new(usersTestSuite))
}

func (s *usersTestSuite) Test01GetUsers_ValidUser() {
	expected := `[{"id":1,"first-name":"青木","last-name":"真琳",` +
		`"account":{"email":"user@coban.jp"},` +
		`"company":{"id":1,"name":"アコム株式会社"},"device":{"is-paired":false}},` +
		`{"id":2,"first-name":"織田","last-name":"信長",` +
		`"account":{"email":"office@coban.jp"},` +
		`"company":{"id":2,"name":"株式会社愛知銀行"},"device":{"is-paired":false}},` +
		`{"id":3,"first-name":"豊臣","last-name":"秀吉",` +
		`"account":{"email":"admin@coban.jp"},` +
		`"company":{"id":3,"name":"AOCホールディングス株式会社"},` +
		`"device":{"is-paired":false}},` +
		`{"id":4,"first-name":"徳川","last-name":"家康",` +
		`"account":{"email":"root@coban.jp"},` +
		`"company":{"id":4,"name":"旭化成株式会社"},"device":{"is-paired":true}}]`
	token, _ := utils.GenerateToken(3, databases.AdminScope)

	url := fmt.Sprintf("%s/administrations/users", utils.Address)
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

func (s *usersTestSuite) Test02GetUsers_InvalidUser() {
	token, _ := utils.GenerateToken(1, databases.ClientScope)

	url := fmt.Sprintf("%s/administrations/users", utils.Address)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, err := client.Do(request)
	defer result.Body.Close()

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusUnauthorized, result.StatusCode)
}
