package common

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
)

type accountsTestSuite struct {
	suite.Suite
}

func TestAccounts(t *testing.T) {
	suite.Run(t, new(accountsTestSuite))
}

func (s *stationsTestSuite) Test01Get_Accounts() {

}

func (s *stationsTestSuite) Test02Get_Account_ByValidID() {

}

func (s *stationsTestSuite) Test03Get_Account_ByInvalidID() {

}

func (s *stationsTestSuite) Test04Create_Account() {

}

func (s *stationsTestSuite) Test05CreateInvalid_Account() {

}

func (s *stationsTestSuite) Test06UpdateValid_Account_ByValidID() {

}

func (s *stationsTestSuite) Test07UpdateValid_Account_ByInvalidID() {

}

func (s *stationsTestSuite) Test08UpdateInvalid_Account_ByValidID() {

}

func (s *stationsTestSuite) Test09Delete_Account_ByValidID() {
}

func (s *stationsTestSuite) Test09Delete_Account_ByInvalidID() {

}