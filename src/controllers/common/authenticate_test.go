package common_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"coban/api/src/utils"
	"fmt"
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