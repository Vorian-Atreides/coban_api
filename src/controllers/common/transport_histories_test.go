package common_test

import (
	"testing"
	"time"

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

func (s *transportHistoriesTestSuite) Test01Get_TransportHistories() {
	dateTime1, _ := time.Parse(time.RFC3339, "2016-01-10T06:30:00+00:00")
	dateTime2, _ := time.Parse(time.RFC3339, "2016-01-10T14:10:00+00:00")
	dateTime3, _ := time.Parse(time.RFC3339, "2016-01-10T22:45:00+00:00")
	dateTime4, _ := time.Parse(time.RFC3339, "2016-02-06T04:30:00+00:00")
	dateTime5, _ := time.Parse(time.RFC3339, "2016-02-06T12:25:00+00:00")
	dateTime6, _ := time.Parse(time.RFC3339, "2016-02-06T18:55:00+00:00")

	stations := []*databases.Station{
		&databases.Station{ID:1, Name:"銀座線", Type:"metro"},
		&databases.Station{ID:2, Name:"日比谷線", Type:"metro"},
		&databases.Station{ID:3, Name:"千代田線", Type:"metro"},
		&databases.Station{ID:4, Name:"南北線", Type:"metro"},
		&databases.Station{ID:5, Name:"横須賀駅", Type:"train"},
		&databases.Station{ID:6, Name:"大宮", Type:"train"},
	}

	users := []*databases.User {
		&databases.User{ID:1, FirstName:"青木", LastName:"真琳", AccountID:1, CompanyID:1},
		&databases.User{ID:4, FirstName:"徳川", LastName:"家康", AccountID:4, CompanyID:4},
	}

	expectedTransportHistories := []databases.TransportHistory{
		databases.TransportHistory{ID:1, Date:dateTime1.UTC(), Stock:850, Expense:150, EntranceID:1, ExitID:2, UserID:1,
		Entrance:stations[0], Exit:stations[1], User:users[0]},
		databases.TransportHistory{ID:2, Date:dateTime2.UTC(), Stock:800, Expense:50, EntranceID:2, ExitID:3, UserID:1,
			Entrance:stations[1], Exit:stations[2], User:users[0]},
		databases.TransportHistory{ID:3, Date:dateTime3.UTC(), Stock:600, Expense:200, EntranceID:3, ExitID:6, UserID:1,
			Entrance:stations[2], Exit:stations[5], User:users[0]},
		databases.TransportHistory{ID:4, Date:dateTime4.UTC(), Stock:10000, Expense:500, EntranceID:5, ExitID:6, UserID:4,
			Entrance:stations[4], Exit:stations[5], User:users[1]},
		databases.TransportHistory{ID:5, Date:dateTime5.UTC(), Stock:8000, Expense:2000, EntranceID:6, ExitID:1, UserID:4,
			Entrance:stations[5], Exit:stations[0], User:users[1]},
		databases.TransportHistory{ID:6, Date:dateTime6.UTC(), Stock:7500, Expense:500, EntranceID:1, ExitID:3, UserID:4,
			Entrance:stations[0], Exit:stations[2], User:users[1]},
	}

	transportHistories := common.GetTransportHistories()
	s.Equal(expectedTransportHistories, transportHistories)
}

func (s *transportHistoriesTestSuite) Test02Get_TransportHistory_ByValidID() {
	dateTime, _ := time.Parse(time.RFC3339, "2016-01-10T06:30:00+00:00")
	expected := databases.TransportHistory{ID:1, Date:dateTime.UTC(), Stock:850, Expense:150,
		UserID:1, EntranceID:1, ExitID:2,
		User: &databases.User{ID:1, FirstName:"青木", LastName:"真琳", AccountID:1, CompanyID:1},
		Entrance: &databases.Station{ID:1, Name:"銀座線", Type:"metro"},
		Exit: &databases.Station{ID:2, Name:"日比谷線", Type:"metro"},
	}

	transportHistory, err := common.GetTransportHistoryByID(expected.ID)
	s.NoError(err)
	s.Equal(expected, transportHistory)
}

func (s *transportHistoriesTestSuite) Test03Get_TransportHistory_ByInvalidID() {
	transportHistory, err := common.GetTransportHistoryByID(0)
	s.Error(err, "This history doesn't exist.")
	s.Equal(uint(0), transportHistory.ID)

	transportHistory, err = common.GetTransportHistoryByID(10)
	s.Error(err, "This history doesn't exist.")
	s.Equal(uint(0), transportHistory.ID)
}

func (s *transportHistoriesTestSuite) Test04Create_TransportHistory() {
	dateTime, _ := time.Parse(time.RFC3339, "2016-02-14T10:30:00+00:00")
	expected := databases.TransportHistory{Date:dateTime, UserID:1, EntranceID:1, ExitID:2, Stock:800, Expense:100}

	transportHistory, err := common.CreateTransportHistory(expected.Date, expected.Stock, expected.Expense,
		expected.EntranceID, expected.ExitID, expected.UserID)
	s.NoError(err)
	s.NotEqual(uint(0), transportHistory.ID)
	s.Equal(expected.Date, transportHistory.Date)
	s.Equal(expected.Stock, transportHistory.Stock)
	s.Equal(expected.Expense, transportHistory.Expense)
	s.Equal(expected.EntranceID, transportHistory.EntranceID)
	s.Equal(expected.ExitID, transportHistory.ExitID)
	s.Equal(expected.UserID, transportHistory.UserID)
}

func (s *transportHistoriesTestSuite) Test05CreateInvalid_TransportHistory() {
	dateTime, _ := time.Parse(time.RFC3339, "2016-02-14T10:30:00+00:00")

	_, err := common.CreateTransportHistory(dateTime, 800, 100, 1, 2, 1)
	s.Error(err, "TRANSPORT-HISTORY: This history already exist.")

	dateTime, _ = time.Parse(time.RFC3339, "2016-02-14T12:30:00+00:00")
	_, err = common.CreateTransportHistory(dateTime, 800, 100, 1, 2, 0)
	s.Error(err, "TRANSPORT-HISTORY: The user is mandatory.")

	_, err = common.CreateTransportHistory(dateTime, 800, 100, 1, 2, 10)
	s.Error(err, "TRANSPORT-HISTORY: This user doesn't exist.")

	_, err = common.CreateTransportHistory(dateTime, 800, 100, 1, 0, 1)
	s.Error(err, "TRANSPORT-HISTORY: The exit is mandatory.")

	_, err = common.CreateTransportHistory(dateTime, 800, 100, 1, 10, 1)
	s.Error(err, "TRANSPORT-HISTORY: This exit doesn't exist.")

	_, err = common.CreateTransportHistory(dateTime, 800, 100, 0, 2, 1)
	s.Error(err, "TRANSPORT-HISTORY: The entrance is mandatory.")

	_, err = common.CreateTransportHistory(dateTime, 800, 100, 10, 2, 1)
	s.Error(err, "TRANSPORT-HISTORY: This entrance doesn't exist.")

	_, err = common.CreateTransportHistory(dateTime, 800, 0, 1, 2, 1)
	s.Error(err, "TRANSPORT-HISTORY: The expense is mandatory.")

	_, err = common.CreateTransportHistory(dateTime, 0, 100, 1, 2, 1)
	s.Error(err, "TRANSPORT-HISTORY: The stock is mandatory.")

	_, err = common.CreateTransportHistory(time.Time{}, 800, 100, 1, 2, 1)
	s.Error(err, "TRANSPORT-HISTORY: The date is mandatory.")
}