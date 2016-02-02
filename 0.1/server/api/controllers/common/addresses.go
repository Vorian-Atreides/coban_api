package common

import (
	"coban/api/0.1/server/api/databases"
)

func GetAddresses() []databases.Address {
	var addresses []databases.Address

	databases.DB.Find(&addresses)
	for i, _ := range addresses {
		addresses[i].LoadRelated()
	}

	return addresses
}

func GetAddressByID(id uint) databases.Address {
	var address databases.Address

	databases.DB.First(&address, id)
	address.LoadRelated()

	return address
}

func CreateAddress(street string, zip string, city string, companyID uint) (databases.Address, error) {
	address := databases.Address{Street:street, Zip:zip, City:city, CompanyID:companyID}

	if err := address.IsValid(true); err != nil {
		return address, err
	}
	databases.DB.Create(&address)
	return address, nil
}

func UpdateAddress(street string, zip string, city string, companyID uint, addressID uint) (databases.Address, error) {
	address := databases.Address{Street:street, Zip:zip, City:city, CompanyID:companyID, ID:addressID}

	if err := address.IsValid(false); err != nil {
		return address, err
	}
	databases.DB.Save(&address)
	return address, nil
}