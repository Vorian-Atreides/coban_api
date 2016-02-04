package databases_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"coban/api/src/databases"
)

type relationsTestSuite struct {
	suite.Suite
}

func TestRelations(t *testing.T) {
	suite.Run(t, new(relationsTestSuite))
}

func (s *relationsTestSuite) Test01AddressRelations() {
	var address databases.Address

	databases.DB.First(&address)
	address.LoadRelated()

	s.NotNil(address.Company)
}

func (s *relationsTestSuite) Test02CompanyRelations() {
	var company databases.Company

	databases.DB.First(&company)
	company.LoadRelated()

	s.NotNil(company.Employees)
	s.NotNil(company.Addresses)
}

func (s *relationsTestSuite) Test03DeviceRelations() {
	var device databases.Device

	databases.DB.First(&device)
	device.LoadRelated()

	s.NotNil(device.User)
}

func (s *relationsTestSuite) Test04AccountRelations() {
	var account databases.Account

	databases.DB.First(&account)
	account.LoadRelated()

	s.NotNil(account.User)
}

func (s *relationsTestSuite) Test05UserRelations() {
	var user databases.User

	databases.DB.First(&user)
	user.LoadRelated()

	s.NotNil(user.Account)
	s.NotNil(user.Company)
	s.NotNil(user.Device)
}

func (s *relationsTestSuite) Test06StationRelations() {
	var station databases.Station

	databases.DB.First(&station)
	station.LoadRelated()

	s.True(true)
}

func (s *relationsTestSuite) Test07TransportHistoryRelations() {
	var transportHistory databases.TransportHistory

	databases.DB.First(&transportHistory)
	transportHistory.LoadRelated()

	s.NotNil(transportHistory.User)
	s.NotNil(transportHistory.Entrance)
	s.NotNil(transportHistory.Exit)
}