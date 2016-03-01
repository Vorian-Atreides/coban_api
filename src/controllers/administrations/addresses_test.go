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

type addressesTestSuite struct {
	suite.Suite
}

func TestAddresses(t *testing.T) {
	utils.InitTest()
	suite.Run(t, new(addressesTestSuite))
}

func (s *accountsTestSuite) Test01GetAddresses_ValidUser() {
	expected := `[{"id":1,"zip":"100-8307","street":"千代田区丸の内二丁目1番1号明治安田生命ビル","city":"東京都","company-id":1},` +
		`{"id":2,"zip":"23106-1","street":"中区栄3-14-12","city":"名古屋市","company-id":2},` +
		`{"id":3,"zip":"140-0002","street":"品川区東品川二丁目5番8号","city":"東京都","company-id":3},` +
		`{"id":4,"zip":"101-8101","street":"千代田区神田神保町1丁目105番地","city":"東京都","company-id":4},` +
		`{"id":5,"zip":"111-8081","street":"台東区駒形1丁目4-8","city":"東京都","company-id":5},` +
		`{"id":6,"zip":"111-8081","street":"東京都台東区駒形2-5-4","city":"東京都","company-id":5}]`
	token, _ := utils.GenerateToken(3, databases.AdminScope)

	url := fmt.Sprintf("%s/administrations/addresses", utils.Address)
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

func (s *accountsTestSuite) Test02GetAddresses_InvalidUser() {
	token, _ := utils.GenerateToken(1, databases.ClientScope)

	url := fmt.Sprintf("%s/administrations/addresses", utils.Address)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, err := client.Do(request)
	defer result.Body.Close()

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusUnauthorized, result.StatusCode)
}
