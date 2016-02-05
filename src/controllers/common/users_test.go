package common_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

//	"coban/api/src/controllers/common"
//	"coban/api/src/databases"
)

type usersTestSuite struct {
	suite.Suite
}

func TestUsers(t *testing.T) {
	suite.Run(t, new(usersTestSuite))
}

func (s *usersTestSuite) Test01Get_Users() {

}

func (s *usersTestSuite) Test02Get_User_ByValidID() {

}

func (s *usersTestSuite) Test03Get_User_ByInvalidID() {

}

func (s *usersTestSuite) Test04Create_User() {

}

func (s *usersTestSuite) Test05CreateInvalid_User() {

}

func (s *usersTestSuite) Test06UpdateValid_User_ByValidID() {

}

func (s *usersTestSuite) Test07UpdateValid_User_ByInvalidID() {

}

func (s *usersTestSuite) Test08UpdateInvalid_User_ByValidID() {

}

func (s *usersTestSuite) Test09Delete_User_ByValidID() {
}

func (s *usersTestSuite) Test09Delete_User_ByInvalidID() {

}