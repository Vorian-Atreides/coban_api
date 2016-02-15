package common

import (
	"errors"
	"log"

	"coban/api/src/databases"
	"coban/api/src/utils"
)

// GetCompanies get every companies in the database
func GetCompanies(offset uint) []databases.Company {
	var companies []databases.Company

	databases.DB.Offset(offset).Limit(utils.PageSize).Find(&companies)
	for i := range companies {
		companies[i].LoadRelated()
	}

	return companies
}

// GetCompanyByID get a company by its ID
func GetCompanyByID(id uint) (databases.Company, error) {
	var company databases.Company

	databases.DB.First(&company, id)
	if company.ID == 0 {
		return company, errors.New("This company doesn't exist.")
	}
	company.LoadRelated()

	return company, nil
}

// CreateCompany try to create a new company
func CreateCompany(name string) (databases.Company, error) {
	company := databases.Company{Name: name}

	if err := company.IsValid(); err != nil {
		return company, err
	}
	databases.DB.Save(&company)

	return company, databases.DB.Error
}

// UpdateCompany try to update a company
func UpdateCompany(name string, id uint) (databases.Company, error) {
	company := databases.Company{Name: name, ID: id}

	var existingCompany databases.Company
	databases.DB.First(&existingCompany, id)
	if existingCompany.ID == 0 {
		return company, errors.New("This company doesn't exist.")
	}
	if err := company.IsValid(); err != nil {
		return company, err
	}
	databases.DB.Save(&company)

	return company, databases.DB.Error
}

// DeleteCompany try to delete a company
func DeleteCompany(id uint) error {
	var company databases.Company

	databases.DB.First(&company, id)
	if company.ID == 0 {
		return errors.New("This company doesn't exist.")
	}
	company.LoadRelated()

	for i := range company.Employees {
		err := DeleteUser(company.Employees[i].ID)
		if err != nil {
			log.Println(err)
		}
	}
	for i := range company.Addresses {
		company.Addresses[i].LoadRelated()
	}

	databases.DB.Delete(&company)

	return databases.DB.Error
}
