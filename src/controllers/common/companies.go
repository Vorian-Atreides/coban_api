package common

import (
	"coban/api/src/databases"
	"errors"
)

func GetCompanies() []databases.Company {
	var companies []databases.Company

	databases.DB.Find(&companies)
	for i, _ := range companies {
		companies[i].LoadRelated()
	}

	return companies
}

func GetCompanyByID(id uint) (databases.Company, error) {
	var company databases.Company

	databases.DB.First(&company, id)
	if company.ID == 0 {
		return company, errors.New("This company doesn't exist.")
	}
	company.LoadRelated()

	return company, nil
}

func CreateCompany(name string) (databases.Company, error) {
	company := databases.Company{Name:name}

	if err := company.IsValid(); err != nil {
		return company, err
	}
	databases.DB.Save(&company)

	return company, databases.DB.Error
}

func UpdateCompany(name string, id uint) (databases.Company, error) {
	company := databases.Company{Name:name, ID:id}

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

func DeleteCompany(id uint) error {
	var company databases.Company

	databases.DB.First(&company, id)
	if company.ID == 0 {
		return errors.New("This company doesn't exist.")
	}
	databases.DB.Delete(&company)

	return databases.DB.Error
}