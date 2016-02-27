package offices_test

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

func (s *addressesTestSuite) Test01GetAddress_ValidUser() {
	expected := `[{"id":2,"zip":"23106-1","street":"中区栄3-14-12",` +
		`"city":"名古屋市","company-id":2}]`
	token, _ := utils.GenerateToken(2, databases.OfficeScope)

	url := fmt.Sprintf("%s/offices/addresses", utils.Address)
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

func (s *addressesTestSuite) Test02GetAddress_InvalidUser() {
	token, _ := utils.GenerateToken(1, databases.ClientScope)

	url := fmt.Sprintf("%s/offices/addresses", utils.Address)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, _ := client.Do(request)
	defer result.Body.Close()

	s.NotNil(result)
	s.Equal(http.StatusUnauthorized, result.StatusCode)
}
