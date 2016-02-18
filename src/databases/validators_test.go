package databases_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"coban/api/src/databases"
)

type validatorsTestSuite struct {
	suite.Suite
}

func TestValidators(t *testing.T) {
	suite.Run(t, new(validatorsTestSuite))
}

func (s *validatorsTestSuite) Test01_Valid_Address() {
	address := databases.Address{City: "Tokyo", CompanyID: 1,
		Street: "12 ginza", Zip: "1234567890"}

	s.NoError(address.IsValid())
}

func (s *validatorsTestSuite) Test02_Invalid_Address() {
	address := databases.Address{City: "Tokyo", CompanyID: 0,
		Street: "12 ginza", Zip: "1234567890"}
	s.Error(address.IsValid(), "ADDRESS: The company is mandatory.")

	address = databases.Address{City: "Tokyo", CompanyID: 10,
		Street: "12 ginza", Zip: "1234567890"}
	s.Error(address.IsValid(), "ADDRESS: This company doesn't exist.")

	address = databases.Address{City: "", CompanyID: 1,
		Street: "12 ginza", Zip: "1234567890"}
	s.Error(address.IsValid(), "ADDRESS: The city is mandatory.")

	address = databases.Address{City: "Tokyo", CompanyID: 1,
		Street: "12 ginza", Zip: ""}
	s.Error(address.IsValid(), "ADDRESS: The zip is mandatory.")

	address = databases.Address{City: "Tokyo", CompanyID: 1,
		Street: "", Zip: "1234567890"}
	s.Error(address.IsValid(), "ADDRESS: The street is mandatory.")

	address = databases.Address{City: "東京都", CompanyID: 1,
		Street: "千代田区丸の内二丁目1番1号明治安田生命ビル", Zip: "100-8307"}
	s.Error(address.IsValid(), "ADDRESS: This address already exist.")
}

func (s *validatorsTestSuite) Test03_Valid_Company() {
	company := databases.Company{Name: "Coban"}

	s.NoError(company.IsValid())
}

func (s *validatorsTestSuite) Test04_Invalid_Company() {
	company := databases.Company{Name: "アコム株式会社"}
	s.Error(company.IsValid(), "COMPANY: This company already exist.")

	company = databases.Company{Name: ""}
	s.Error(company.IsValid(), "COMPANY: The name is mandatory.")
}

func (s *validatorsTestSuite) Test05_Valid_Device() {
	device := databases.Device{IsPaired: false, UserID: 2}
	s.NoError(device.IsValid())
}

func (s *validatorsTestSuite) Test06_Invalid_Device() {
	device := databases.Device{IsPaired: false, UserID: 1}
	s.Error(device.IsValid(), "DEVICE: This device already exist.")

	device = databases.Device{IsPaired: false, UserID: 0}
	s.Error(device.IsValid(), "DEVICE: The user is mandatory.")

	device = databases.Device{IsPaired: false, UserID: 10}
	s.Error(device.IsValid(), "DEVICE: The user doesn't exist.")
}

func (s *validatorsTestSuite) Test07_Valid_Account() {
	account := databases.Account{Email: "siffer_g@coban.jp", Password: "Toto42",
		Scope: databases.ClientScope}
	s.NoError(account.IsValid(false))
}

func (s *validatorsTestSuite) Test08_Invalid_Account() {
	account := databases.Account{Email: "user@coban.jp", Password: "Toto42",
		Scope: databases.ClientScope}
	s.Error(account.IsValid(false), "ACCOUNT: This email is already used.")

	account = databases.Account{Email: "", Password: "Toto42",
		Scope: databases.ClientScope}
	s.Error(account.IsValid(false), "ACCOUNT: The email is mandatory.")

	account = databases.Account{Email: "user@coban.jp", Password: "",
		Scope: databases.ClientScope}
	s.Error(account.IsValid(false), "ACCOUNT: The password is mandatory.")

	account = databases.Account{Email: "user@coban.jp", Password: "Toto42",
		Scope: 0}
	s.Error(account.IsValid(false), "ACCOUNT: The scope is mandatory.")
}

func (s *validatorsTestSuite) Test09_Valid_User() {
	user := databases.User{FirstName: "Gaston", LastName: "Siffert",
		AccountID: 5, CompanyID: 1}
	s.NoError(user.IsValid())
}

func (s *validatorsTestSuite) Test10_Invalid_User() {
	user := databases.User{FirstName: "青木", LastName: "真琳",
		AccountID: 1, CompanyID: 1}
	s.Error(user.IsValid(), "USER: This user already exist.")

	user = databases.User{FirstName: "青木", LastName: "真琳",
		AccountID: 6, CompanyID: 0}
	s.Error(user.IsValid(), "USER: The company is mandatory.")

	user = databases.User{FirstName: "", LastName: "真琳",
		AccountID: 6, CompanyID: 1}
	s.Error(user.IsValid(), "USER: The first name is mandatory.")

	user = databases.User{FirstName: "青木", LastName: "",
		AccountID: 6, CompanyID: 1}
	s.Error(user.IsValid(), "USER: The last name is mandatory.")
}

// func (s *validatorsTestSuite) Test11_Valid_Station() {
// }
//
// func (s *validatorsTestSuite) Test12_Invalid_Station() {
// }

func (s *validatorsTestSuite) Test13_Valid_TransportType() {
	transportHistory := databases.TransportHistory{Date: time.Now(),
		Stock: 500, EntranceID: 1, ExitID: 2, UserID: 1}
	s.NoError(transportHistory.IsValid())
}

func (s *validatorsTestSuite) Test14_Invalid_TransportType() {
	dateTime, _ := time.Parse(time.RFC3339, "2016-01-10T06:30:00+00:00")

	transportHistory := databases.TransportHistory{Date: dateTime.UTC(),
		Stock: 500, EntranceID: 1, ExitID: 2, UserID: 1}
	s.Error(transportHistory.IsValid(),
		"TRANSPORT-HISTORY: This history already exist.")

	transportHistory = databases.TransportHistory{Date: dateTime.UTC(),
		Stock: 500, EntranceID: 0, ExitID: 2, UserID: 4}
	s.Error(transportHistory.IsValid(),
		"TRANSPORT-HISTORY: The entrance is mandatory.")

	transportHistory = databases.TransportHistory{Date: dateTime.UTC(),
		Stock: 500, EntranceID: 1, ExitID: 0, UserID: 4}
	s.Error(transportHistory.IsValid(),
		"TRANSPORT-HISTORY: The exit is mandatory.")

	transportHistory = databases.TransportHistory{Date: dateTime.UTC(),
		Stock: 500, EntranceID: 0, ExitID: 2, UserID: 3}
	s.Error(transportHistory.IsValid(),
		"TRANSPORT-HISTORY: This user doesn't have a device.")

	transportHistory = databases.TransportHistory{Date: dateTime.UTC(),
		Stock: 0, EntranceID: 1, ExitID: 2, UserID: 4}
	s.Error(transportHistory.IsValid(),
		"TRANSPORT-HISTORY: The stock is mandatory.")

	transportHistory = databases.TransportHistory{Date: time.Time{}.UTC(),
		Stock: 500, EntranceID: 1, ExitID: 2, UserID: 4}
	s.Error(transportHistory.IsValid(),
		"TRANSPORT-HISTORY: The date is mandatory.")

	transportHistory = databases.TransportHistory{Date: dateTime.UTC(),
		Stock: 500, EntranceID: 10, ExitID: 2, UserID: 4}
	s.Error(transportHistory.IsValid(),
		"TRANSPORT-HISTORY: This entrance doesn't exist.")

	transportHistory = databases.TransportHistory{Date: dateTime.UTC(),
		Stock: 500, EntranceID: 1, ExitID: 10, UserID: 4}
	s.Error(transportHistory.IsValid(),
		"TRANSPORT-HISTORY: This exit doesn't exist.")

	transportHistory = databases.TransportHistory{Date: dateTime.UTC(),
		Stock: 500, EntranceID: 1, ExitID: 2, UserID: 10}
	s.Error(transportHistory.IsValid(),
		"TRANSPORT-HISTORY: This user doesn't exist.")
}
