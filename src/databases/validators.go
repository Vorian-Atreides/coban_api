package databases

import (
	"errors"
)

func buildError(err string) error {
	if err == "" {
		return nil
	}
	return errors.New(err)
}

//
// Address
//

func (address Address) IsValid() error {
	err := ""

	if address.Street == "" {
		err += "ADDRESS: The street is mandatory."
	}
	if address.Zip == "" {
		err += "ADDRESS: The zip is mandatory."
	}
	if address.City == "" {
		err += "ADDRESS: The city is mandatory."
	}
	if address.CompanyID == 0 {
		err += "ADDRESS: The company is mandatory."
	}

	var items []Address
	DB.Where(Address{
		City:address.City, Zip:address.Zip,
		Street:address.Street, CompanyID:address.CompanyID,
	}).Not(Address{ID:address.ID}).Find(&items)
	if len(items) > 0 {
		err += "ADDRESS: This address already exist."
	}

	return buildError(err)
}

//
// Company
//

func (company Company) IsValid() error {
	err := ""

	if company.Name == "" {
		err += "COMPANY: The name is mandatory."
	}

	var items []Company
	DB.Where(Company{Name:company.Name}).Not(Company{ID:company.ID}).Find(&items)
	if len(items) > 0 {
		err += "COMPANY: This company already exist."
	}

	return buildError(err)
}

//
// Device
//

func (device Device) IsValid() error {
	err := ""

	if device.UserID == 0 {
		err += "DEVICE: The user is mandatory."
	}

	var items []Device
	DB.Where(Device{UserID:device.UserID}).Not(Device{ID:device.ID}).Find(&items)
	if len(items) > 0 {
		err += "DEVICE: This device already exist."
	}

	return buildError(err)
}

//
// Account
//

func (account Account) IsValid() error {
	err := ""

	if account.Email == "" {
		err += "ACCOUNT: The email is mandatory."
	}
	if account.Password == "" {
		err += "ACCOUNT: The password is mandatory."
	}
	if account.Scope == 0 {
		err += "ACCOUNT: The scope is mandatory."
	}

	var items []Account
	DB.Where(Account{Email:account.Email}).Not(Account{ID:account.ID}) .Find(&items)
	if len(items) > 0 {
		err += "ACCOUNT: This email is already used."
	}

	return buildError(err)
}

//
// User
//

func (user User) IsValid() error {
	err := ""

	if user.FirstName == "" {
		err += "USER: The first name is mandatory."
	}
	if user.LastName == "" {
		err += "USER: The last name is mandatory."
	}
	if user.AccountID == 0 {
		err += "USER: The account is mandatory."
	}
	if user.CompanyID == 0 {
		err += "USER: The company is mandatory."
	}

	var items []User
	DB.Where(User{AccountID:user.AccountID}).Not(User{ID:user.ID}).Find(&items)
	if len(items) > 0 {
		err += "USER: This user already exist."
	}

	return buildError(err)
}

//
// Station
//

func (station Station) IsValid() error {
	err := ""

	if station.Name == "" {
		err += "STATION: The name is mandatory."
	}
	if station.Type == "" {
		err += "STATION: The type is mandatory."
	}

	var items []Station
	DB.Where(Station{Name:station.Name, Type:station.Type}).Not(User{ID:station.ID}).Find(&items)
	if len(items) > 0 {
		err += "STATION: This user already exist."
	}

	return buildError(err)
}

//
// TransportHistory
//

func (transportHistory TransportHistory) IsValid() error {
	err := ""

	if transportHistory.Date.IsZero() {
		err += "TRANSPORT-HISTORY: The date is mandatory."
	}
	if transportHistory.Expense <= 0 {
		err += "TRANSPORT-HISTORY: The expense is mandatory."
	}
	if transportHistory.Stock <= 0 {
		err += "TRANSPORT-HISTORY: The stock is mandatory."
	}
	if transportHistory.EntranceID == 0 {
		err += "TRANSPORT-HISTORY: The entrance is mandatory."
	}
	if transportHistory.ExitID == 0 {
		err += "TRANSPORT-HISTORY: The exit is mandatory."
	}
	if transportHistory.UserID == 0 {
		err += "TRANSPORT-HISTORY: The user is mandatory."
	}

	var items []TransportHistory
	DB.Where(TransportHistory{Date:transportHistory.Date, UserID:transportHistory.UserID}).
		Not(TransportHistory{ID:transportHistory.ID}).Find(&items)
	if len(items) > 0 {
		err += "TransportHistory: This history already exist."
	}

	return buildError(err)
}