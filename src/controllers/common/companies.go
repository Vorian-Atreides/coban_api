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

func GetCompanyByID(id uint) databases.Company {
	var company databases.Company

	databases.DB.First(&company, id)
	company.LoadRelated()

	return company
}

func CreateCompany(name string) (databases.Company, error) {
	company := databases.Company{Name:name}

	if err := company.IsValid(true); err != nil {
		return company, err
	}
	databases.DB.Save(&company)

	return company, nil
}

func UpdateCompany(name string, id uint) (databases.Company, error) {
	company := databases.Company{Name:name, ID:id}

	if err := company.IsValid(false); err != nil {
		return company, err
	}
	databases.DB.Save(&company)

	return company, nil
}

func DeleteCompany(id uint) error {
	var company databases.Company

	databases.DB.Where(databases.Address{CompanyID:id}).Delete(databases.Address{})
	databases.DB.First(&company, id)
	databases.DB.Delete(&company)

	if company.ID != 0 {
		return errors.New("This company can't be deleted.")
	}

	return nil
}