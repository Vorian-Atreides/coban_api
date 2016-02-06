package databases

import (
	"encoding/json"
	"net/http"
)

func ReadBody(r *http.Request, model interface{}) error {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(model)

	return err
}

//
// Address
//

func (address *Address) FromBody(r *http.Request) error {
	if err := ReadBody(r, address); err != nil {
		return err
	}
	return nil
}

//
// Company
//

func (company *Company) FromBody(r *http.Request) error {
	if err := ReadBody(r, company); err != nil {
		return err
	}
	return nil
}

//
// Device
//

func (device *Device) FromBody(r *http.Request) error {
	if err := ReadBody(r, device); err != nil {
		return err
	}
	return nil
}

//
// Account
//

func (account *Account) FromBody(r *http.Request) error {
	if err :=  ReadBody(r, account); err != nil {
		return err
	}
	return nil
}

//
// User
//

func (user *User) FromBody(r *http.Request) error {
	if err := ReadBody(r, user); err != nil {
		return err
	}
	user.AccountID = user.Account.ID
	user.CompanyID = user.Company.ID
	return nil
}

//
// Station
//

func (station *Station) FromBody(r *http.Request) error {
	if err := ReadBody(r, station); err != nil {
		return err
	}
	return nil
}

//
// TransportHistory
//

func (transportHistory *TransportHistory) FromBody(r *http.Request) error {
	if err := ReadBody(r, transportHistory); err != nil {
		return err
	}
	transportHistory.EntranceID = transportHistory.Entrance.ID
	transportHistory.ExitID = transportHistory.Exit.ID
	return nil
}
