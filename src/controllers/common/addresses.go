package common

import (
	"errors"

	"coban/api/src/databases"
)

func GetAddresses() []databases.Address {
	var addresses []databases.Address

	databases.DB.Find(&addresses)
	for i, _ := range addresses {
		addresses[i].LoadRelated()
	}

	return addresses
}

func GetAddressByID(id uint) (databases.Address, error) {
	var address databases.Address

	databases.DB.First(&address, id)
	if address.ID == 0 {
		return address, errors.New("This address doesn't exist.")
	}
	address.LoadRelated()

	return address, nil
}

func CreateAddress(street string, zip string, city string, companyID uint) (databases.Address, error) {
	address := databases.Address{Street:street, Zip:zip, City:city, CompanyID:companyID}

	if err := address.IsValid(); err != nil {
		return address, err
	}
	databases.DB.Create(&address)

	return address, databases.DB.Error
}

func UpdateAddress(street string, zip string, city string, companyID uint, id uint) (databases.Address, error) {
	address := databases.Address{Street:street, Zip:zip, City:city, CompanyID:companyID, ID:id}

	var existingAddress databases.Address
	databases.DB.First(&existingAddress, id)
	if existingAddress.ID == 0 {
		return address, errors.New("This address doesn't exist.")
	}
	if err := address.IsValid(); err != nil {
		return address, err
	}
	databases.DB.Save(&address)

	return address, databases.DB.Error
}

func DeleteAddressByID(id uint) error {
	var address databases.Address

	databases.DB.First(&address, id)
	if address.ID == 0 {
		return errors.New("This address doesn't exist")
	}
	databases.DB.Delete(&address)

	return databases.DB.Error
}