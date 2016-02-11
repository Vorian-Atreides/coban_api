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

// IsValid (Address) ensure the address is valid and can be serialised
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
		City: address.City, Zip: address.Zip,
		Street: address.Street, CompanyID: address.CompanyID,
	}).Not(Address{ID: address.ID}).Find(&items)
	if len(items) > 0 {
		err += "ADDRESS: This address already exist."
	}

	var company Company
	DB.First(&company, address.CompanyID)
	if company.ID == 0 {
		err += "ADDRESS: This company doesn't exist."
	}

	return buildError(err)
}

// IsValid (Company) ensure the company is valid and can be serialised
func (company Company) IsValid() error {
	err := ""

	if company.Name == "" {
		err += "COMPANY: The name is mandatory."
	}

	var items []Company
	DB.Where(Company{Name: company.Name}).Not(Company{ID: company.ID}).Find(&items)
	if len(items) > 0 {
		err += "COMPANY: This company already exist."
	}

	return buildError(err)
}

// IsValid (Device) ensure the device is valid and can be serialised
func (device Device) IsValid() error {
	err := ""

	if device.UserID == 0 {
		err += "DEVICE: The user is mandatory."
	} else {
		var user User
		DB.First(&user, device.UserID)
		if user.ID == 0 {
			err += "DEVICE: This user doesn't exist."
		}
	}

	var items []Device
	DB.Where(Device{UserID: device.UserID}).Not(Device{ID: device.ID}).Find(&items)
	if len(items) > 0 {
		err += "DEVICE: This device already exist."
	}

	return buildError(err)
}

// IsValid (Account) ensure the account is valid and can be serialised
func (account Account) IsValid(onlyPassword bool) error {
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

	if !onlyPassword {
		var items []Account
		DB.Where(Account{Email: account.Email}).Not(Account{ID: account.ID}).Find(&items)
		if len(items) > 0 {
			err += "ACCOUNT: This email is already used."
		}
	}

	return buildError(err)
}

// IsValid (User) ensure the user is valid and can be serialised
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
	} else {
		var account Account
		DB.First(&account, user.AccountID)
		if account.ID == 0 {
			err += "USER: This account doesn't exist."
		}
	}
	if user.CompanyID == 0 {
		err += "USER: The company is mandatory."
	} else {
		var company Company
		DB.First(&company, user.CompanyID)
		if company.ID == 0 {
			err += "USER: This company doesn't exist."
		}
	}

	var items []User
	DB.Where(User{AccountID: user.AccountID}).Not(User{ID: user.ID}).Find(&items)
	if len(items) > 0 {
		err += "USER: This user already exist."
	}

	return buildError(err)
}

// IsValid (Station) ensure the station is valid and can be serialised
func (station Station) IsValid() error {
	err := ""

	if station.Name == "" {
		err += "STATION: The name is mandatory."
	}
	// if station.Type == "" {
	// 	err += "STATION: The type is mandatory."
	// }

	// var items []Station
	// DB.Where(Station{Name: station.Name, Type: station.Type}).Not(User{ID: station.ID}).Find(&items)
	// if len(items) > 0 {
	// 	err += "STATION: This user already exist."
	// }

	return buildError(err)
}

// IsValid (TansportHistory) ensure the transport history is valid
// and can be serialised
func (transportHistory TransportHistory) IsValid() error {
	err := ""

	if transportHistory.Date.IsZero() {
		err += "TRANSPORT-HISTORY: The date is mandatory."
	} else {
		var items []TransportHistory
		DB.Where(TransportHistory{Date: transportHistory.Date,
			Stock:      transportHistory.Stock,
			EntranceID: transportHistory.EntranceID,
			ExitID:     transportHistory.ExitID,
			UserID:     transportHistory.UserID}).
			Not(TransportHistory{ID: transportHistory.ID}).Find(&items)
		if len(items) > 0 {
			err += "TRANSPORT-HISTORY: This history already exist."
		}
	}
	// if transportHistory.Expense <= 0 {
	// 	err += "TRANSPORT-HISTORY: The expense is mandatory."
	// }
	if transportHistory.Stock <= 0 {
		err += "TRANSPORT-HISTORY: The stock is mandatory."
	}

	if transportHistory.EntranceID == 0 {
		err += "TRANSPORT-HISTORY: The entrance is mandatory."
	} else {
		var entrance Station
		DB.First(&entrance, transportHistory.EntranceID)
		if entrance.ID == 0 {
			err += "TRANSPORT-HISTORY: This entrance doesn't exist."
		}
	}
	if transportHistory.ExitID == 0 {
		err += "TRANSPORT-HISTORY: The exit is mandatory."
	} else {
		var exit Station
		DB.First(&exit, transportHistory.ExitID)
		if exit.ID == 0 {
			err += "TRANSPORT-HISTORY: This exit doesn't exist."
		}
	}
	if transportHistory.UserID == 0 {
		err += "TRANSPORT-HISTORY: The user is mandatory."
	} else {
		var user User
		DB.First(&user, transportHistory.UserID)
		if user.ID == 0 {
			err += "TRANSPORT-HISTORY: This user doesn't exist."
		} else {
			user.LoadRelated()
			if user.Device == nil {
				err += "TRANSPORT-HISTORY: This user doesn't have a device."
			}
		}
	}
	return buildError(err)
}
