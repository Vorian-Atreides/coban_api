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

type transportHistoriesTestSuite struct {
	suite.Suite
}

func TestTransportHistories(t *testing.T) {
	utils.InitTest()
	suite.Run(t, new(transportHistoriesTestSuite))
}

func (s *transportHistoriesTestSuite) Test01GetTransportHistories_ValidUser() {
	expected := `[{"id":4,"date":"2016-02-06T04:30:00Z","stock":10000,` +
		`"entrance":{"company":"試験","line":"試験","station":"端末試験用 V3-01-1"},` +
		`"exit":{"company":"試験","line":"試験","station":"端末試験用 V3-01-2"},` +
		`"user":{"id":4,"first-name":"徳川","last-name":"家康","account":null,"company":null}},` +
		`{"id":5,"date":"2016-02-06T12:25:00Z","stock":8000,` +
		`"entrance":{"company":"試験","line":"試験","station":"端末試験用 V3-01-2"},` +
		`"exit":{"company":"試験","line":"試験","station":"端末試験用 V1"},` +
		`"user":{"id":4,"first-name":"徳川","last-name":"家康","account":null,"company":null}},` +
		`{"id":6,"date":"2016-02-06T18:55:00Z","stock":7500,` +
		`"entrance":{"company":"試験","line":"試験","station":"端末試験用 V1"},` +
		`"exit":{"company":"試験","line":"試験","station":"端末試験用 V2-01"},` +
		`"user":{"id":4,"first-name":"徳川","last-name":"家康","account":null,"company":null}}]`
	token, _ := utils.GenerateToken(4, databases.OfficeScope)

	url := fmt.Sprintf("%s/offices/transport-histories", utils.Address)
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

func (s *transportHistoriesTestSuite) Test02GetTransportHistories_InvalidUser() {
	token, _ := utils.GenerateToken(1, databases.ClientScope)

	url := fmt.Sprintf("%s/offices/transport-histories", utils.Address)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, err := client.Do(request)
	defer result.Body.Close()

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusUnauthorized, result.StatusCode)
}

func (s *transportHistoriesTestSuite) Test03GetTransportHistories_ValidOfficerValidUser() {
	expected := `[{"id":4,"date":"2016-02-06T04:30:00Z","stock":10000,` +
		`"entrance":{"company":"試験","line":"試験","station":"端末試験用 V3-01-1"},` +
		`"exit":{"company":"試験","line":"試験","station":"端末試験用 V3-01-2"},` +
		`"user":{"id":4,"first-name":"徳川","last-name":"家康","account":null,"company":null}},` +
		`{"id":5,"date":"2016-02-06T12:25:00Z","stock":8000,` +
		`"entrance":{"company":"試験","line":"試験","station":"端末試験用 V3-01-2"},` +
		`"exit":{"company":"試験","line":"試験","station":"端末試験用 V1"},` +
		`"user":{"id":4,"first-name":"徳川","last-name":"家康","account":null,"company":null}},` +
		`{"id":6,"date":"2016-02-06T18:55:00Z","stock":7500,` +
		`"entrance":{"company":"試験","line":"試験","station":"端末試験用 V1"},` +
		`"exit":{"company":"試験","line":"試験","station":"端末試験用 V2-01"},` +
		`"user":{"id":4,"first-name":"徳川","last-name":"家康","account":null,"company":null}}]`
	token, _ := utils.GenerateToken(4, databases.OfficeScope)

	url := fmt.Sprintf("%s/offices/transport-histories/4", utils.Address)
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

func (s *transportHistoriesTestSuite) Test04GetTransportHistories_ValidOfficerInvalidUser() {
	token, _ := utils.GenerateToken(4, databases.OfficeScope)

	url := fmt.Sprintf("%s/offices/transport-histories/1", utils.Address)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, err := client.Do(request)
	defer result.Body.Close()

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusUnauthorized, result.StatusCode)
}

func (s *transportHistoriesTestSuite) Test04GetTransportHistories_InvalidOfficerValidUser() {
	token, _ := utils.GenerateToken(1, databases.ClientScope)

	url := fmt.Sprintf("%s/offices/transport-histories/4", utils.Address)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, err := client.Do(request)
	defer result.Body.Close()

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusUnauthorized, result.StatusCode)
}
