package clients_test

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

type transportHistoriesTestSuite struct {
	suite.Suite
}

func TestTransportHistories(t *testing.T) {
	utils.InitTest()
	suite.Run(t, new(transportHistoriesTestSuite))
	deleteCreatedTransportHistory()
}

func deleteCreatedTransportHistory() {
	var transport databases.TransportHistory
	databases.DB.Find(&transport, 7)
	databases.DB.Delete(transport)
}

func (s *transportHistoriesTestSuite) Test01GetTransportHistories_ValidUser() {
	expected := `[{"id":1,"date":"2016-01-10T06:30:00Z","stock":850,`+
		`"entrance":{"company":"試験","line":"試験","station":"端末試験用 V1"},`+
		`"exit":{"company":"試験","line":"試験","station":"端末試験用 V2"},`+
		`"user":{"id":1,"first-name":"青木","last-name":"真琳","account":null,"company":null}},`+
		`{"id":2,"date":"2016-01-10T14:10:00Z","stock":800,`+
		`"entrance":{"company":"試験","line":"試験","station":"端末試験用 V2"},`+
		`"exit":{"company":"試験","line":"試験","station":"端末試験用 V2-01"},`+
		`"user":{"id":1,"first-name":"青木","last-name":"真琳","account":null,"company":null}},`+
		`{"id":3,"date":"2016-01-10T22:45:00Z","stock":600,`+
		`"entrance":{"company":"試験","line":"試験","station":"端末試験用 V2-01"},`+
		`"exit":{"company":"試験","line":"試験","station":"端末試験用 V3-01-2"},`+
		`"user":{"id":1,"first-name":"青木","last-name":"真琳","account":null,"company":null}}]`

	token, _ := utils.GenerateToken(1, databases.ClientScope)
	url := fmt.Sprintf("%s/clients/transport-histories", utils.Address)
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

func (s *transportHistoriesTestSuite) Test02GetTransportHistories_InvalidUser() {
	token, _ := utils.GenerateToken(2, databases.OfficeScope)
	url := fmt.Sprintf("%s/clients/transport-histories", utils.Address)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", "Bearer" + " " + token)

	client := &http.Client{}
	result, _ := client.Do(request)
	defer result.Body.Close()

	s.Equal(http.StatusUnauthorized, result.StatusCode)
}

func (s *transportHistoriesTestSuite) Test03TransportHistories_InvalidAdd() {
	data := `["AQEAAiBI5SvkPpgPAAqgAA==",
		"BQEAAiBI5SvkPpgPAAqgAA==",
		"CAEAAiBI5SvkPpgPAAqgAA==",
		"EgEAAiBI5SvkPpgPAAqgAA==",
		"FQEAAiBI5SvkPpgPAAqgAA=="]`

	token, _ := utils.GenerateToken(1, databases.ClientScope)
	url := fmt.Sprintf("%s/clients/transport-histories", utils.Address)
	request, _ := http.NewRequest("POST", url, strings.NewReader(data))
	request.Header.Set("Authorization", "Bearer" + " " + token)

	client := &http.Client{}
	result, _ := client.Do(request)
	defer result.Body.Close()

	s.Equal(http.StatusBadRequest, result.StatusCode)
}

func (s *transportHistoriesTestSuite) Test03TransportHistories_InvalidAddWithOneValid() {
	expected := `[{"id":7,"date":"2013-01-01T00:00:00Z","stock":10000,`+
		`"entrance":{"company":"試験","line":"試験","station":"端末試験用 V1"},`+
		`"exit":{"company":"試験","line":"試験","station":"端末試験用 V2"},`+
		`"user":{"id":1,"first-name":"青木","last-name":"真琳","account":null,"company":null}}]`
	data := `["AQEAAiBI5SvkPpgPAAqgAA==",
		"BQEAAiBI5SvkPpgPAAqgAA==",
		"FgEAAhohAAAAARAnAAqgAA==",
		"EgEAAiBI5SvkPpgPAAqgAA==",
		"FQEAAiBI5SvkPpgPAAqgAA=="]`

	token, _ := utils.GenerateToken(1, databases.ClientScope)
	url := fmt.Sprintf("%s/clients/transport-histories", utils.Address)
	request, _ := http.NewRequest("POST", url, strings.NewReader(data))
	request.Header.Set("Authorization", "Bearer" + " " + token)

	client := &http.Client{}
	result, _ := client.Do(request)
	defer result.Body.Close()

	body, _ := ioutil.ReadAll(result.Body)
	s.Equal(http.StatusBadRequest, result.StatusCode)
	s.Equal(expected, string(body))
}