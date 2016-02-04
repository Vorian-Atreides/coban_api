package databases

import (
)

//
// Address
//

func (address *Address) LoadRelated() {
	address.Company = new(Company)

	DB.Model(address).Related(address.Company)
}

//
// Company
//

func (company *Company) LoadRelated() {
	DB.Model(company).Related(&company.Addresses)
	DB.Model(company).Related(&company.Employees)
}

//
// Device
//

func (device *Device) LoadRelated() {
	device.User = new(User)

	DB.Model(device).Related(device.User)
}

//
// Account
//

func (account *Account) LoadRelated() {
	account.User = new(User)

	DB.Model(account).Related(account.User)
}

//
// User
//

func (user *User) LoadRelated() {
	user.Account = new(Account)
	user.Company = new(Company)
	user.Device = new(Device)

	DB.Model(user).Related(user.Account)
	DB.Model(user).Related(user.Company)
	DB.Model(user).Related(user.Device)
}

//
// Station
//

func (station *Station) LoadRelated() {
}

//
// TransportType
//

func (transportHistory *TransportHistory) LoadRelated() {
	transportHistory.Entrance = new(Station)
	transportHistory.Exit = new(Station)
	transportHistory.User = new(User)

	DB.Model(transportHistory).Related(transportHistory.Entrance)
	DB.Model(transportHistory).Related(transportHistory.Exit)
	DB.Model(transportHistory).Related(transportHistory.User)
}