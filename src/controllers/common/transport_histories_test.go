package common

import (
	"testing"

	"github.com/stretchr/testify/suite"

//	"coban/api/src/controllers/common"
//	"coban/api/src/databases"
)

type transportHistoriesTestSuite struct {
	suite.Suite
}

func TestTransportHistories(t *testing.T) {
	suite.Run(t, new(transportHistoriesTestSuite))
}

func (s *transportHistoriesTestSuite) Test01Get_TransportHistories() {

}

func (s *transportHistoriesTestSuite) Test02Get_TransportHistory_ByValidID() {

}

func (s *transportHistoriesTestSuite) Test03Get_TransportHistory_ByInvalidID() {

}

func (s *transportHistoriesTestSuite) Test04Create_TransportHistory() {

}

func (s *transportHistoriesTestSuite) Test05CreateInvalid_TransportHistory() {

}

func (s *transportHistoriesTestSuite) Test06UpdateValid_TransportHistory_ByValidID() {

}

func (s *transportHistoriesTestSuite) Test07UpdateValid_TransportHistory_ByInvalidID() {

}

func (s *transportHistoriesTestSuite) Test08UpdateInvalid_TransportHistory_ByValidID() {

}

func (s *transportHistoriesTestSuite) Test09Delete_TransportHistory_ByValidID() {
}

func (s *transportHistoriesTestSuite) Test09Delete_TransportHistory_ByInvalidID() {

}
