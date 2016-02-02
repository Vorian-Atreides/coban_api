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

func (address Address) IsValid(forCreation bool) error {
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

	if forCreation {
		var other Address
		DB.Where(&Address {
			City:address.City, Street:address.Street,
			Zip:address.Zip, CompanyID:address.CompanyID }).Find(&other)
		if other.ID != 0 {
			err += "ADDRESS: This address already exist."
		}
	}

	return buildError(err)
}

//
// Company
//

func (company Company) IsValid(forCreation bool) error {
	err := ""

	if company.Name == "" {
		err += "COMPANY: The name is mandatory."
	}

	if forCreation {
		var other Company
		DB.Where(&Company{Name:company.Name}).Find(&other)
		if other.ID != 0 {
			err += "COMPANY: This company already exist."
		}
	}

	return buildError(err)
}

//
// Device
//

func (device Device) IsValid(forCreation bool) error {
	err := ""

	return buildError(err)
}

//
// Account
//

func (account Account) IsValid(forCreation bool) error {
	err := ""

	if account.Email == "" {
		err += "ACCOUNT: The email is mandatory."
	}
	if account.Password == "" {
		err += "ACCOUNT: The password is mandatory"
	}

	if forCreation {
		var other Account
		DB.Where(&Account{Email:account.Email}).Find(&other)
		if other.ID != 0 {
			err += "ACCOUNT: This email is already used."
		}
	}

	return buildError(err)
}

//
// User
//

func (user User) IsValid(forCreation bool) error {
	err := ""

	if user.FirstName == "" {
		err += "USER: The first name is mandatory."
	}
	if user.LastName == "" {
		err += "USER: The last name is mandatory."
	}

	if user.Account.ID == 0 {
		err += "USER: The account is mandatory."
	} else if er := user.Account.IsValid(false); er != nil {
		err += er.Error()
	}
	if user.Company.ID == 0 {
		err += "USER: The company is mandatory."
	} else if er := user.Company.IsValid(false); er != nil {
		err += er.Error()
	}

	return buildError(err)
}

//
// TransportType
//

func (transportType TransportType) IsValid(forCreation bool) error {
	err := ""

	if transportType.Name == "" {
		err += "TRANSPORT-TYPE: The name is mandatory."
	}

	if forCreation {
		var other TransportType
		DB.Where(&TransportType{Name:transportType.Name}).Find(&other)
		if other.ID != 0 {
			err += "TRANSPORT-TYPE: This type already exist."
		}
	}

	return buildError(err)
}

//
// Station
//

func (station Station) IsValid(forCreation bool) error {
	err := ""

	if station.Name == "" {
		err += "STATION: The name is mandatory."
	}

	if station.Type.ID == 0 {
		err += "STATION: The type is mandatory."
	} else if er := station.Type.IsValid(false); er != nil {
		err += er.Error()
	}

	return buildError(err)
}

//
// TransportHistory
//

func (transportHistory TransportHistory) IsValid(forCreation bool) error {
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

	if transportHistory.Entrance.ID == 0 {
		err += "TRANSPORT-HISTORY: The entrance is mandatory."
	} else if er := transportHistory.Entrance.IsValid(false); er != nil {
		err += er.Error()
	}
	if transportHistory.Exit.ID == 0 {
		err += "TRANSPORT-HISTORY: The exit is mandatory."
	} else if er := transportHistory.Exit.IsValid(false); er != nil {
		err += er.Error()
	}

	return buildError(err)
}