package databases

import (
	"encoding/json"
	"net/http"
)

func readBody(r *http.Request, model interface{}) error {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(model)

	return err
}

//
// Address
//

func (address *Address) FromBody(r *http.Request) error {
	if err := readBody(r, address); err != nil {
		return err
	}
	return nil
}

//
// Company
//

func (company *Company) FromBody(r *http.Request) error {
	if err := readBody(r, company); err != nil {
		return err
	}
	return nil
}

//
// Device
//

func (device *Device) FromBody(r *http.Request) error {
	if err := readBody(r, device); err != nil {
		return err
	}
	return nil
}

//
// Account
//

func (account *Account) FromBody(r *http.Request) error {
	if err :=  readBody(r, account); err != nil {
		return err
	}
	return nil
}

//
// User
//

func (user *User) FromBody(r *http.Request) error {
	if err := readBody(r, user); err != nil {
		return err
	}
	user.AccountID = user.Account.ID
	user.CompanyID = user.Company.ID
	return nil
}

//
// TransportType
//

func (transportType *TransportType) FromBody(r *http.Request) error {
	if err := readBody(r, transportType); err != nil {
		return err
	}
	return nil
}

//
// Station
//

func (station *Station) FromBody(r *http.Request) error {
	if err := readBody(r, station); err != nil {
		return err
	}
	station.TypeID = station.Type.ID
	return nil
}

//
// TransportHistory
//

func (transportHistory *TransportHistory) FromBody(r *http.Request) error {
	if err := readBody(r, transportHistory); err != nil {
		return err
	}
	transportHistory.EntranceID = transportHistory.Entrance.ID
	transportHistory.ExitID = transportHistory.Exit.ID
	return nil
}
