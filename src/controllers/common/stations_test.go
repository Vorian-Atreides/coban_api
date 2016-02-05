package common_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
)

type stationsTestSuite struct {
	suite.Suite
}

func TestStations(t *testing.T) {
	suite.Run(t, new(stationsTestSuite))
}

func (s *stationsTestSuite) Test01Get_Stations() {

}

func (s *stationsTestSuite) Test02Get_Station_ByValidID() {

}

func (s *stationsTestSuite) Test03Get_Station_ByInvalidID() {

}

func (s *stationsTestSuite) Test04Create_Station() {

}

func (s *stationsTestSuite) Test05CreateInvalid_Station() {

}

func (s *stationsTestSuite) Test06UpdateValid_Station_ByValidID() {

}

func (s *stationsTestSuite) Test07UpdateValid_Station_ByInvalidID() {

}

func (s *stationsTestSuite) Test08UpdateInvalid_Station_ByValidID() {

}

func (s *stationsTestSuite) Test09Delete_Station_ByValidID() {
}

func (s *stationsTestSuite) Test09Delete_Station_ByInvalidID() {

}