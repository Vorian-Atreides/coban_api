package common_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

//	"coban/api/src/controllers/common"
//	"coban/api/src/databases"
	"coban/api/src/databases"
	"coban/api/src/controllers/common"
)

type stationsTestSuite struct {
	suite.Suite
}

func TestStations(t *testing.T) {
	suite.Run(t, new(stationsTestSuite))
}

func (s *stationsTestSuite) Test01Get_Stations() {
	expectedStations := []databases.Station {
		databases.Station{ID:1, Name:"銀座線", Type:"metro"},
		databases.Station{ID:2, Name:"日比谷線", Type:"metro"},
		databases.Station{ID:3, Name:"千代田線", Type:"metro"},
		databases.Station{ID:4, Name:"南北線", Type:"metro"},
		databases.Station{ID:5, Name:"横須賀駅", Type:"train"},
		databases.Station{ID:6, Name:"大宮", Type:"train"},
	}

	stations := common.GetStations()
	s.Equal(expectedStations, stations)
}

func (s *stationsTestSuite) Test02Get_Station_ByValidID() {
	expectedStation := databases.Station{ID:1, Name:"銀座線", Type:"metro"}

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

func (s *stationsTestSuite) Test04Create_Station() {
	expectedStation := databases.Station{Name:"Shinjuku", Type:"train"}

	station, err := common.CreateStation(expectedStation.Name, expectedStation.Type)
	s.NoError(err)
	s.NotEqual(uint(0), station.ID)
	s.Equal(expectedStation.Name, station.Name)
	s.Equal(expectedStation.Type, station.Type)
}

func (s *stationsTestSuite) Test05CreateInvalid_Station() {
	station, err := common.CreateStation("Shinjuku", "train")
	s.Error(err, "STATION:: This station already exist.")
	s.Equal(uint(0), station.ID)

	station, err = common.CreateStation("Mejiro", "")
	s.Error(err, "STATION: The type is mandatory.")
	s.Equal(uint(0), station.ID)

	station, err = common.CreateStation("", "train")
	s.Error(err, "STATION: The name is mandatory.")
	s.Equal(uint(0), station.ID)
}

func (s *stationsTestSuite) Test06UpdateValid_Station_ByValidID() {
	var target databases.Station
	databases.DB.Where(databases.Station{Name:"Shinjuku"}).First(&target)

	expectedStation := databases.Station{ID:target.ID, Name:"Mejiro", Type:"metro"}
	station, err := common.UpdateStation(expectedStation.Name, expectedStation.Type, expectedStation.ID)
	s.NoError(err)
	s.Equal(expectedStation, station)
}

func (s *stationsTestSuite) Test07UpdateValid_Station_ByInvalidID() {
	_, err := common.UpdateStation("Ikubekuro", "train", 0)
	s.Error(err, "STATION: This station doesn't exist.")

	_, err = common.UpdateStation("Ikubekuro", "train", 10)
	s.Error(err, "STATION: This station doesn't exist.")
}

func (s *stationsTestSuite) Test08UpdateInvalid_Station_ByValidID() {
	var target databases.Station
	databases.DB.Where(databases.Station{Name:"Mejiro"}).First(&target)

	_, err := common.UpdateStation("銀座線", "metro", target.ID)
	s.Error(err, "STATION: This station already exist.")

	_, err = common.UpdateStation("", "metro", target.ID)
	s.Error(err, "STATION: The name is mandatory.")

	_, err = common.UpdateStation("銀座線", "", target.ID)
	s.Error(err, "STATION: The type is mandatory.")
}