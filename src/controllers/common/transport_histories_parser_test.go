package common_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"coban/api/src/controllers/common"
)

type transportHistoriesParserTestSuite struct {
	suite.Suite
}

func TestTransportHistoriesParser(t *testing.T) {
	suite.Run(t, new(transportHistoriesParserTestSuite))
}

func (s *transportHistoriesParserTestSuite) Test01WithUnsuported_TypeMachine() {
	data := []string{
		"AQEAAiBI5SvkA5gPAAqgAA==",
		"AwEAAiBI5SvkA5gPAAqgAA==",
		"BQEAAiBI5SvkA5gPAAqgAA==",
		"CAEAAiBI5SvkA5gPAAqgAA==",
		"EQEAAiBI5SvkA5gPAAqgAA==",
		"FQEAAiBI5SvkA5gPAAqgAA==",
	}

	for _, sample := range data {
		_, err := common.ParseTransportHistory([]byte(sample))
		s.Error(err, "This type of machine isn't supported.")
	}
}

func (s *transportHistoriesParserTestSuite) Test02WithUnsuported_UseCase() {
	data := []string{
		"FgIAAiBI5SvkA5gPAAqgAA==",
		"FgMAAiBI5SvkA5gPAAqgAA==",
		"FgQAAiBI5SvkA5gPAAqgAA==",
		"FgUAAiBI5SvkA5gPAAqgAA==",
		"FggAAiBI5SvkA5gPAAqgAA==",
		"FhIAAiBI5SvkA5gPAAqgAA==",
	}

	for _, sample := range data {
		_, err := common.ParseTransportHistory([]byte(sample))
		s.Error(err, "This use case isn't supported.")
	}
}

func (s *transportHistoriesParserTestSuite) Test03WithUnsuported_Length() {
	data := []string{
		"FgIAAiBI5SvkA5gPAAqgAA",
		"FgMAAiBI5SvkA5gPAAqgAA==AA",
		"FgQAAiBI5SvkA5gPAA==",
		"FgUAAiBI5SvkA5gqgAA==",
		"FggAAiBI5SPAAqgAA==",
		"FhIAvkA5gPAAqgAA==",
	}

	for _, sample := range data {
		_, err := common.ParseTransportHistory([]byte(sample))
		s.Error(err, "This use case isn't supported.")
	}
}

func (s *transportHistoriesParserTestSuite) Test04WithUfound_Entrance() {
	data := []string{
		"FgEAAiBI/yvkA5gPAAqgAA==",
		"FgEAAiBIECvkA5gPAAqgAA==",
		"FgEAAiBIICvkA5gPAAqgAA==",
		"FgEAAjBIICvkA5gPAAqgAA==",
		"FgEAAkBIICvkA5gPAAqgAA==",
		"FgEAArFIICvkA5gPAAqgAA==",
	}

	for _, sample := range data {
		_, err := common.ParseTransportHistory([]byte(sample))
		s.Error(err, "There are an error with the entrance.")
	}
}

func (s *transportHistoriesParserTestSuite) Test04WithUfound_Exit() {
	data := []string{
		"FgEAArFIABDkA5gPAAqgAA==",
		"FgEAArFIACDkA5gPAAqgAA==",
		"FgEAArFIADDkA5gPAAqgAA==",
		"FgEAArFIAEDkA5gPAAqgAA==",
		"FgEAArFIAFDkA5gPAAqgAA==",
		"FgEAArFIAKDkA5gPAAqgAA==",
	}

	for _, sample := range data {
		_, err := common.ParseTransportHistory([]byte(sample))
		s.Error(err, "There are an error with the exit.")
	}
}

func (s *transportHistoriesParserTestSuite) Test05WithValid_Dates() {
	data := []string{
		"FgEAArFIAAEaIZgPAAqgAA==", // 2013-01-01
		"FgEAArFIAALkA5gPAAqgAA==", // 2010-05-22
		"FgEAArFIAAENUJgPAAqgAA==", // 2006-10-16
		"FgEAArFIAAZGUJgPAAqgAA==", // 2003-02-06
		"FgEAArFIAB95UJgPAAqgAA==", // 2015-11-25
		"FgEAArFIAB9UUJgPAAqgAA==", // 2015-10-20
	}

	date1, _ := time.Parse("2006-01-02", "2013-01-01")
	date2, _ := time.Parse("2006-01-02", "2010-05-22")
	date3, _ := time.Parse("2006-01-02", "2006-10-16")
	date4, _ := time.Parse("2006-01-02", "2003-02-06")
	date5, _ := time.Parse("2006-01-02", "2015-11-25")
	date6, _ := time.Parse("2006-01-02", "2015-10-20")

	dates := []time.Time{
		date1, date2, date3, date4, date5, date6,
	}

	for i, sample := range data {
		transportHistory, err := common.ParseTransportHistory([]byte(sample))
		s.NoError(err)
		s.Equal(dates[i], transportHistory.Date)
	}
}
