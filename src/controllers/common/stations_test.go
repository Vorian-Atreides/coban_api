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
	expectedStations := []databases.Station{
		databases.Station{ID: 1, LineCode: 0, StationCode: 0,
			Company: "試験", Line: "試験", Name: "端末試験用 V1"},
		databases.Station{ID: 2, LineCode: 0, StationCode: 1,
			Company: "試験", Line: "試験", Name: "端末試験用 V2"},
		databases.Station{ID: 3, LineCode: 0, StationCode: 2,
			Company: "試験", Line: "試験", Name: "端末試験用 V2-01"},
		databases.Station{ID: 4, LineCode: 0, StationCode: 3,
			Company: "試験", Line: "試験", Name: "端末試験用 V2-02"},
		databases.Station{ID: 5, LineCode: 0, StationCode: 4,
			Company: "試験", Line: "試験", Name: "端末試験用 V3-01-1"},
		databases.Station{ID: 6, LineCode: 0, StationCode: 5,
			Company: "試験", Line: "試験", Name: "端末試験用 V3-01-2"},
		databases.Station{ID: 7, LineCode: 0, StationCode: 6,
			Company: "試験", Line: "試験", Name: "端末試験用 V4-01"},
		databases.Station{ID: 8, LineCode: 0, StationCode: 7,
			Company: "試験", Line: "試験", Name: "端末試験用 V4-01-A"},
	}

	stations := common.GetStations(0)
	s.Equal(expectedStations, stations)
}

func (s *stationsTestSuite) Test04Get_Stations_Paginated() {
	expected := []databases.Station{}

	stations := common.GetStations(50)
	s.Equal(expected, stations)
}

func (s *stationsTestSuite) Test02Get_Station_ByValidID() {
	expectedStation := databases.Station{ID: 1, LineCode: 0,
		StationCode: 0, Company: "試験", Line: "試験", Name: "端末試験用 V1"}

	station, err := common.GetStationByID(expectedStation.ID)
	s.NoError(err)
	s.Equal(expectedStation, station)
}

func (s *stationsTestSuite) Test03Get_Station_ByInvalidID() {
	station, err := common.GetStationByID(0)
	s.Error(err, "This station doesn't exist.")
	s.Equal(uint(0), station.ID)

	station, err = common.GetStationByID(10)
	s.Error(err, "This station doesn't exist.")
	s.Equal(uint(0), station.ID)
}
