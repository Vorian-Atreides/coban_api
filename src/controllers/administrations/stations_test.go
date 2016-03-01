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

type stationsTestSuite struct {
	suite.Suite
}

func TestStations(t *testing.T) {
	utils.InitTest()
	suite.Run(t, new(stationsTestSuite))
}

func (s *stationsTestSuite) Test01GetStations_ValidUser() {
	expected := `[{"company":"試験","line":"試験","station":"端末試験用 V1"},` +
		`{"company":"試験","line":"試験","station":"端末試験用 V2"},` +
		`{"company":"試験","line":"試験","station":"端末試験用 V2-01"},` +
		`{"company":"試験","line":"試験","station":"端末試験用 V2-02"},` +
		`{"company":"試験","line":"試験","station":"端末試験用 V3-01-1"},` +
		`{"company":"試験","line":"試験","station":"端末試験用 V3-01-2"},` +
		`{"company":"試験","line":"試験","station":"端末試験用 V4-01"},` +
		`{"company":"試験","line":"試験","station":"端末試験用 V4-01-A"}]`
	token, _ := utils.GenerateToken(3, databases.AdminScope)

	url := fmt.Sprintf("%s/administrations/stations", utils.Address)
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

func (s *stationsTestSuite) Test02GetStations_InvalidUser() {
	token, _ := utils.GenerateToken(1, databases.ClientScope)

	url := fmt.Sprintf("%s/administrations/stations", utils.Address)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, err := client.Do(request)
	defer result.Body.Close()

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusUnauthorized, result.StatusCode)
}
