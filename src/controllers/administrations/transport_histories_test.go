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

type transportHistoriesTestSuite struct {
	suite.Suite
}

func TestTransportHistories(t *testing.T) {
	utils.InitTest()
	suite.Run(t, new(transportHistoriesTestSuite))
}

func (s *transportHistoriesTestSuite) Test01GetTransportHistories_ValidUser() {
	expected := `[{"id":1,"date":"2016-01-10T06:30:00Z","stock":850,` +
		`"entrance":{"company":"試験","line":"試験","station":"端末試験用 V1"},` +
		`"exit":{"company":"試験","line":"試験","station":"端末試験用 V2"},` +
		`"user":{"id":1,"first-name":"青木","last-name":"真琳","account":null,"company":null}},` +
		`{"id":2,"date":"2016-01-10T14:10:00Z","stock":800,` +
		`"entrance":{"company":"試験","line":"試験","station":"端末試験用 V2"},` +
		`"exit":{"company":"試験","line":"試験","station":"端末試験用 V2-01"},` +
		`"user":{"id":1,"first-name":"青木","last-name":"真琳","account":null,"company":null}},` +
		`{"id":3,"date":"2016-01-10T22:45:00Z","stock":600,` +
		`"entrance":{"company":"試験","line":"試験","station":"端末試験用 V2-01"},` +
		`"exit":{"company":"試験","line":"試験","station":"端末試験用 V3-01-2"},` +
		`"user":{"id":1,"first-name":"青木","last-name":"真琳","account":null,"company":null}},` +
		`{"id":4,"date":"2016-02-06T04:30:00Z","stock":10000,` +
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
	token, _ := utils.GenerateToken(3, databases.AdminScope)

	url := fmt.Sprintf("%s/administrations/transport-histories", utils.Address)
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

	url := fmt.Sprintf("%s/administrations/transport-histories", utils.Address)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "Bearer"+" "+token)

	client := &http.Client{}
	result, err := client.Do(request)
	defer result.Body.Close()

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusUnauthorized, result.StatusCode)
}
