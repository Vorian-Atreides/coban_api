package databases

import (

)

//
// Address
//

func (address *Address) LoadRelated() {
	DB.Model(address).Related(&address.City)
}

//
// Company
//

func (company *Company) LoadRelated() {
	DB.Model(company).Related(&company.Addresses)
	DB.Model(company).Related(&company.Employees)

	for i, _ := range company.Addresses {
		company.Addresses[i].LoadRelated()
	}
	for i, _ := range company.Employees {
		company.Employees[i].LoadRelated()
	}
}

//
// Device
//

func (device *Device) LoadRelated() {
}

//
// Account
//

func (account *Account) LoadRelated() {
}

//
// User
//

func (user *User) LoadRelated() {
	DB.Model(user).Related(&user.Account)
	DB.Model(user).Related(&user.Company)
	DB.Model(user).Related(&user.Device)

	user.Account.LoadRelated()
	user.Device.LoadRelated()
}

//
// Station
//

func (station *Station) LoadRelated() {
	DB.Model(station).Related(&station.Type)
}

//
// TransportType
//

func (transportHistory *TransportHistory) LoadRelated() {
	DB.Model(transportHistory).Related(&transportHistory.Entrance)
	DB.Model(transportHistory).Related(&transportHistory.Exit)

	transportHistory.Entrance.LoadRelated()
	transportHistory.Exit.LoadRelated()
}