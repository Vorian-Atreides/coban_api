package common

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
)

type transportHistoriesTestSuite struct {
	suite.Suite
}

func TestTransportHistories(t *testing.T) {
	suite.Run(t, new(transportHistoriesTestSuite))
}

func (s *stationsTestSuite) Test01Get_TransportHistories() {

}

func (s *stationsTestSuite) Test02Get_TransportHistory_ByValidID() {

}

func (s *stationsTestSuite) Test03Get_TransportHistory_ByInvalidID() {

}

func (s *stationsTestSuite) Test04Create_TransportHistory() {

}

func (s *stationsTestSuite) Test05CreateInvalid_TransportHistory() {

}

func (s *stationsTestSuite) Test06UpdateValid_TransportHistory_ByValidID() {

}

func (s *stationsTestSuite) Test07UpdateValid_TransportHistory_ByInvalidID() {

}

func (s *stationsTestSuite) Test08UpdateInvalid_TransportHistory_ByValidID() {

}

func (s *stationsTestSuite) Test09Delete_TransportHistory_ByValidID() {
}

func (s *stationsTestSuite) Test09Delete_TransportHistory_ByInvalidID() {

}
