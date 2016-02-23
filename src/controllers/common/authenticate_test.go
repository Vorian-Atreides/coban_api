package common_test

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"coban/api/src/utils"
)

type authenticateTestSuite struct {
	suite.Suite
}

func TestAuthenticate(t *testing.T) {
	utils.InitTest()
	suite.Run(t, new(authenticateTestSuite))
}

func (s *authenticateTestSuite) Test01ClientValid_Authenticate() {
	json := `{
		"email":"user@coban.jp",
		"password":"user"
	}`

	url := fmt.Sprintf("%s/clients/authenticate", utils.Address)
	result, err := http.Post(url, "application/json", strings.NewReader(json))

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusOK, result.StatusCode)
}

func (s *authenticateTestSuite) Test02OfficeValid_Authenticate() {
	json := `{
		"email":"office@coban.jp",
		"password":"office"
	}`

	url := fmt.Sprintf("%s/offices/authenticate", utils.Address)
	result, err := http.Post(url, "application/json", strings.NewReader(json))

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusOK, result.StatusCode)
}

func (s *authenticateTestSuite) Test03AdminValid_Authenticate() {
	json := `{
		"email":"admin@coban.jp",
		"password":"admin"
	}`

	url := fmt.Sprintf("%s/administrations/authenticate", utils.Address)
	result, err := http.Post(url, "application/json", strings.NewReader(json))

	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusOK, result.StatusCode)
}

func (s *authenticateTestSuite) Test04RootValid_Authenticate() {
	json := `{
		"email":"root@coban.jp",
		"password":"root"
	}`

	url := fmt.Sprintf("%s/clients/authenticate", utils.Address)
	result, err := http.Post(url, "application/json", strings.NewReader(json))
	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusOK, result.StatusCode)

	url = fmt.Sprintf("%s/offices/authenticate", utils.Address)
	result, err = http.Post(url, "application/json", strings.NewReader(json))
	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusOK, result.StatusCode)

	url = fmt.Sprintf("%s/administrations/authenticate", utils.Address)
	result, err = http.Post(url, "application/json", strings.NewReader(json))
	s.NoError(err)
	s.NotNil(result)
	s.Equal(http.StatusOK, result.StatusCode)
}