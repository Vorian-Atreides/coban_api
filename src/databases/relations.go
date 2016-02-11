package databases

// LoadRelated (Address) Instatiate the related objects
func (address *Address) LoadRelated() {
	address.Company = new(Company)

	DB.Model(address).Related(address.Company)
}

// LoadRelated (Company) Instatiate the related objects
func (company *Company) LoadRelated() {
	DB.Model(company).Related(&company.Addresses)
	DB.Model(company).Related(&company.Employees)
}

// LoadRelated (Device) Instatiate the related objects
func (device *Device) LoadRelated() {
	device.User = new(User)

	DB.Model(device).Related(device.User)
}

// LoadRelated (Account) Instatiate the related objects
func (account *Account) LoadRelated() {
	account.User = new(User)

	DB.Model(account).Related(account.User)
}

// LoadRelated (User) Instatiate the related objects
func (user *User) LoadRelated() {
	user.Account = new(Account)
	user.Company = new(Company)
	user.Device = new(Device)

	DB.Model(user).Related(user.Account)
	DB.Model(user).Related(user.Company)
	DB.Model(user).Related(user.Device)
}

// LoadRelated (Station) Instatiate the related objects
func (station *Station) LoadRelated() {
}

// LoadRelated (TransportHistoy) Instatiate the related objects
func (transportHistory *TransportHistory) LoadRelated() {
	transportHistory.Entrance = new(Station)
	transportHistory.Exit = new(Station)
	transportHistory.User = new(User)

	DB.Model(transportHistory).Related(transportHistory.Entrance, "EntranceID")
	DB.Model(transportHistory).Related(transportHistory.Exit, "ExitID")
	DB.Model(transportHistory).Related(transportHistory.User)
}
