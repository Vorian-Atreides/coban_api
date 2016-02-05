package common_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

//	"coban/api/src/controllers/common"
//	"coban/api/src/databases"
)

type accountsTestSuite struct {
	suite.Suite
}

func TestAccounts(t *testing.T) {
	suite.Run(t, new(accountsTestSuite))
}

func (s *accountsTestSuite) Test01Get_Accounts() {

}

func (s *accountsTestSuite) Test02Get_Account_ByValidID() {

}

func (s *accountsTestSuite) Test03Get_Account_ByInvalidID() {

}

func (s *accountsTestSuite) Test04Create_Account() {

}

func (s *accountsTestSuite) Test05CreateInvalid_Account() {

}

func (s *accountsTestSuite) Test06UpdateValid_Account_ByValidID() {

}

func (s *accountsTestSuite) Test07UpdateValid_Account_ByInvalidID() {

}

func (s *accountsTestSuite) Test08UpdateInvalid_Account_ByValidID() {

}

func (s *accountsTestSuite) Test09Delete_Account_ByValidID() {
}

func (s *accountsTestSuite) Test09Delete_Account_ByInvalidID() {

}