package utils_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"

	"coban/api/src/utils"
	"coban/api/src/databases"
)

type securityTestSuite struct {
	suite.Suite
}

func TestSecurity(t *testing.T) {
	suite.Run(t, new(securityTestSuite))
}

func (s *securityTestSuite) Test01_GenerateToken() {
	token, err := utils.GenerateToken(1, databases.ClientScope)
	s.NoError(err)
	s.NotNil(token)
}

func (s *securityTestSuite) Test02_ParseValidToken() {
	tokenStr, _ := utils.GenerateToken(1, databases.ClientScope)

	token, err := utils.ParseToken(tokenStr)
	s.NoError(err)
	s.Equal(true, token.Valid)
}

func (s *securityTestSuite) Test03_ParseValidTokenFromRequest() {
	request, _ := http.NewRequest("GET", "www.google.com", nil)
	tokenStr, _ := utils.GenerateToken(1, databases.ClientScope)
	request.Header.Add("Authorization", "Bearer " + tokenStr)

	token, err := utils.ParseTokenFromRequest(request)
	s.NoError(err)
	s.Equal(true, token.Valid)
}

func (s *securityTestSuite) Test04_ParseInvalidTokenFromRequest() {
	invalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
					"eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9." +
					"TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"

	request, _ := http.NewRequest("GET", "www.google.com", nil)
	request.Header.Add("Authorization", "Bearer " + invalidToken)

	token, err := utils.ParseTokenFromRequest(request)
	s.Error(err)
	s.Equal(false, token.Valid)
}

func (s *securityTestSuite) Test05_ParseInvalidToken() {
	invalidToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
					"eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9." +
					"TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"

	token, err := utils.ParseToken(invalidToken)
	s.Error(err)
	s.Equal(false, token.Valid)
}

func (s *securityTestSuite) Test06_HashPassword() {
	clearText := "hello alfred de la rigga"
	expectedHash := "3598257d8d26bc2218a4416ff33f857739cc8ab000429b2e0417c832ffd2" +
					"5107d30d7c0e84c21a08b47d5d2fc3369315adb8eb4e165f22e471e359ef80b2aafb"

	hash := utils.HashPassword(clearText)
	s.Equal(expectedHash, hash)
}