package databases_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"coban/api/src/databases"
	"time"
)

type validatorsTestSuite struct {
	suite.Suite
}

func TestValidators(t *testing.T) {
	suite.Run(t, new(validatorsTestSuite))
}

func (s *validatorsTestSuite) Test01ValidAddress() {
	address := databases.Address{City:"Tokyo", CompanyID:1, Street:"12 ginza", Zip:"1234567890"}

	s.NoError(address.IsValid())
}

func (s *validatorsTestSuite) Test02InvalidAddress() {
	address := databases.Address{City:"Tokyo", CompanyID:0, Street:"12 ginza", Zip:"1234567890"}
	s.Error(address.IsValid(), "ADDRESS: The company is mandatory.")

	address = databases.Address{City:"Tokyo", CompanyID:10, Street:"12 ginza", Zip:"1234567890"}
	s.Error(address.IsValid(), "ADDRESS: This company doesn't exist.")

	address = databases.Address{City:"", CompanyID:1, Street:"12 ginza", Zip:"1234567890"}
	s.Error(address.IsValid(), "ADDRESS: The city is mandatory.")

	address = databases.Address{City:"Tokyo", CompanyID:1, Street:"12 ginza", Zip:""}
	s.Error(address.IsValid(), "ADDRESS: The zip is mandatory.")

	address = databases.Address{City:"Tokyo", CompanyID:1, Street:"", Zip:"1234567890"}
	s.Error(address.IsValid(), "ADDRESS: The street is mandatory.")

	address = databases.Address{City:"東京都", CompanyID:1,
		Street:"千代田区丸の内二丁目1番1号明治安田生命ビル", Zip:"100-8307"}
	s.Error(address.IsValid(), "ADDRESS: This address already exist.")
}

func (s *validatorsTestSuite) Test03ValidCompany() {
	company := databases.Company{Name:"Coban"}

	s.NoError(company.IsValid())
}

func (s *validatorsTestSuite) Test04InvalidCompany() {
	company := databases.Company{Name: "アコム株式会社"}
	s.Error(company.IsValid(), "COMPANY: This company already exist.")

	company = databases.Company{Name: ""}
	s.Error(company.IsValid(), "COMPANY: The name is mandatory.")
}

func (s *validatorsTestSuite) Test05ValidDevice() {
	device := databases.Device{IsPaired:false, UserID:2}
	s.NoError(device.IsValid())
}

func (s *validatorsTestSuite) Test06InvalidDevice() {
	device := databases.Device{IsPaired:false, UserID:1}
	s.Error(device.IsValid(), "DEVICE: This device already exist.")

	device = databases.Device{IsPaired:false, UserID:0}
	s.Error(device.IsValid(), "DEVICE: The user is mandatory.")

	device = databases.Device{IsPaired:false, UserID:10}
	s.Error(device.IsValid(), "DEVICE: The user doesn't exist.")
}

func (s *validatorsTestSuite) Test07ValidAccount() {
	account := databases.Account{Email:"gaston@coban.jp", Password:"Toto42", Scope:databases.ClientScope}
	s.NoError(account.IsValid())
}

func (s *validatorsTestSuite) Test08InvalidAccount() {
	account := databases.Account{Email:"user@coban.jp", Password:"Toto42", Scope:databases.ClientScope}
	s.Error(account.IsValid(), "ACCOUNT: This email is already used.")

	account = databases.Account{Email:"", Password:"Toto42", Scope:databases.ClientScope}
	s.Error(account.IsValid(), "ACCOUNT: The email is mandatory.")

	account = databases.Account{Email:"user@coban.jp", Password:"", Scope:databases.ClientScope}
	s.Error(account.IsValid(), "ACCOUNT: The password is mandatory.")

	account = databases.Account{Email:"user@coban.jp", Password:"Toto42", Scope:0}
	s.Error(account.IsValid(), "ACCOUNT: The scope is mandatory.")
}

func (s *validatorsTestSuite) Test09ValidUser() {
	account := databases.Account{Email:"gaston@coban.jp", Password:"Toto42", Scope:databases.ClientScope}
	databases.DB.Save(&account)
	user := databases.User{FirstName:"Gaston", LastName:"Siffert", AccountID:account.ID, CompanyID:1}
	s.NoError(user.IsValid())
}

func (s *validatorsTestSuite) Test10InvalidUser() {
	user := databases.User{FirstName:"青木", LastName:"真琳", AccountID:1, CompanyID:1}
	s.Error(user.IsValid(), "USER: This user already exist.")

	account := databases.Account{Email:"gaston.siffert@coban.jp", Password:"Toto42", Scope:databases.ClientScope}
	databases.DB.Save(&account)
	user = databases.User{FirstName:"青木", LastName:"真琳", AccountID:account.ID, CompanyID:0}
	s.Error(user.IsValid(), "USER: The company is mandatory.")

	user = databases.User{FirstName:"", LastName:"真琳", AccountID:account.ID, CompanyID:1}
	s.Error(user.IsValid(), "USER: The first name is mandatory.")

	user = databases.User{FirstName:"青木", LastName:"", AccountID:account.ID, CompanyID:1}
	s.Error(user.IsValid(), "USER: The last name is mandatory.")
}

func (s *validatorsTestSuite) Test11ValidStation() {
	station := databases.Station{Name:"Ginza", Type:"Metro"}
	s.NoError(station.IsValid())
}

func (s *validatorsTestSuite) Test12InvalidStation() {
	station := databases.Station{Name:"銀座線", Type:"Metro"}
	s.Error(station.IsValid(), "STATION: This station already exist.")

	station = databases.Station{Name:"", Type:"Metro"}
	s.Error(station.IsValid(), "STATION: The name is mandatory.")

	station = databases.Station{Name:"Ginza", Type:""}
	s.Error(station.IsValid(), "STATION: The type is mandatory.")
}

func (s *validatorsTestSuite) Test13ValidTransportType() {
	transportHistory := databases.TransportHistory{Date:time.Now(),
		Stock:500, Expense:100, EntranceID:1, ExitID:2, UserID:1}
	s.NoError(transportHistory.IsValid())
}

func (s *validatorsTestSuite) Test14InvalidTransportType() {
	datetime, _ := time.Parse(time.RFC3339, "2016-01-10 06:30:00")

	transportHistory := databases.TransportHistory{Date:datetime,
		Stock:500, Expense:100, EntranceID:1, ExitID:2, UserID:1}
	s.Error(transportHistory.IsValid(), "TransportHistory: This history already exist.")

	transportHistory = databases.TransportHistory{Date:datetime,
		Stock:500, Expense:100, EntranceID:0, ExitID:2, UserID:4}
	s.Error(transportHistory.IsValid(), "TransportHistory: The entrance is mandatory.")

	transportHistory = databases.TransportHistory{Date:datetime,
		Stock:500, Expense:100, EntranceID:1, ExitID:0, UserID:4}
	s.Error(transportHistory.IsValid(), "TransportHistory: The exit is mandatory.")

	transportHistory = databases.TransportHistory{Date:datetime,
		Stock:500, Expense:100, EntranceID:0, ExitID:2, UserID:3}
	s.Error(transportHistory.IsValid(), "TransportHistory: This user dpesn't have a device.")

	transportHistory = databases.TransportHistory{Date:datetime,
		Stock:500, Expense:0, EntranceID:0, ExitID:2, UserID:4}
	s.Error(transportHistory.IsValid(), "TransportHistory: The expense is mandatory.")

	transportHistory = databases.TransportHistory{Date:datetime,
		Stock:0, Expense:100, EntranceID:0, ExitID:2, UserID:4}
	s.Error(transportHistory.IsValid(), "TransportHistory: The stock is mandatory.")

	transportHistory = databases.TransportHistory{Date:time.Time{},
		Stock:500, Expense:100, EntranceID:0, ExitID:2, UserID:4}
	s.Error(transportHistory.IsValid(), "TransportHistory: The date is mandatory.")
}