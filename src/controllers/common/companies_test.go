package common_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
)

type companiesTestSuite struct {
	suite.Suite
}

func TestCompanies(t *testing.T) {
	suite.Run(t, new(companiesTestSuite))
}

func (s *companiesTestSuite) Test01Get_Companies() {
	expectedCompanies := []databases.Company{
		databases.Company{ID: 1, Name: "アコム株式会社",
			Addresses: []databases.Address{
				databases.Address{ID: 1, Zip: "100-8307",
					Street: "千代田区丸の内二丁目1番1号明治安田生命ビル",
					City:   "東京都", CompanyID: 1},
			}, Employees: []databases.User{
				databases.User{ID: 1, FirstName: "青木", LastName: "真琳",
					AccountID: 1, CompanyID: 1},
			}},
		databases.Company{ID: 2, Name: "株式会社愛知銀行",
			Addresses: []databases.Address{
				databases.Address{ID: 2, Zip: "23106-1",
					Street: "中区栄3-14-12", City: "名古屋市", CompanyID: 2},
			}, Employees: []databases.User{
				databases.User{ID: 2, FirstName: "織田",
					LastName: "信長", AccountID: 2, CompanyID: 2},
			}},
		databases.Company{ID: 3, Name: "AOCホールディングス株式会社",
			Addresses: []databases.Address{
				databases.Address{ID: 3, Zip: "140-0002",
					Street: "品川区東品川二丁目5番8号", City: "東京都",
					CompanyID: 3},
			}, Employees: []databases.User{
				databases.User{ID: 3, FirstName: "豊臣", LastName: "秀吉",
					AccountID: 3, CompanyID: 3},
			}},
		databases.Company{ID: 4, Name: "旭化成株式会社",
			Addresses: []databases.Address{
				databases.Address{ID: 4, Zip: "101-8101",
					Street: "千代田区神田神保町1丁目105番地",
					City:   "東京都", CompanyID: 4},
			}, Employees: []databases.User{
				databases.User{ID: 4, FirstName: "徳川",
					LastName: "家康", AccountID: 4, CompanyID: 4},
			}},
		databases.Company{ID: 5, Name: "株式会社バンダイ",
			Addresses: []databases.Address{
				databases.Address{ID: 5, Zip: "111-8081",
					Street: "台東区駒形1丁目4-8", City: "東京都", CompanyID: 5},
				databases.Address{ID: 6, Zip: "111-8081",
					Street: "東京都台東区駒形2-5-4", City: "東京都", CompanyID: 5},
			}, Employees: []databases.User{}},
	}

	companies := common.GetCompanies(0)
	s.Equal(expectedCompanies, companies)
}

func (s *companiesTestSuite) Test11Get_Companies_Paginated() {
	var expected []databases.User

	companies := common.GetCompanies(50)
	s.Equal(expected, companies)
}

func (s *companiesTestSuite) Test02Get_Company_ByValidID() {
	expectedCompany := databases.Company{ID: 1, Name: "アコム株式会社",
		Addresses: []databases.Address{
			databases.Address{ID: 1, Zip: "100-8307",
				Street: "千代田区丸の内二丁目1番1号明治安田生命ビル",
				City:   "東京都", CompanyID: 1},
		}, Employees: []databases.User{
			databases.User{ID: 1, FirstName: "青木",
				LastName: "真琳", AccountID: 1, CompanyID: 1},
		}}

	company, err := common.GetCompanyByID(expectedCompany.ID)
	s.NoError(err)
	s.Equal(expectedCompany, company)
}

func (s *companiesTestSuite) Test03Get_Company_ByInvalidID() {
	company, err := common.GetCompanyByID(0)
	s.Error(err, "This company doesn't exist.")
	s.Equal(uint(0), company.ID)

	company, err = common.GetCompanyByID(10)
	s.Error(err, "This company doesn't exist.")
	s.Equal(uint(0), company.ID)
}

func (s *companiesTestSuite) Test04CreateValid_Company() {
	expectedCompany := databases.Company{Name: "Coban"}

	company, err := common.CreateCompany(expectedCompany.Name)
	s.NoError(err)
	s.NotEqual(uint(0), company.ID)
	s.Equal(expectedCompany.Name, company.Name)
}

func (s *companiesTestSuite) Test05CreateInvalid_Company() {
	company, err := common.CreateCompany("Coban")
	s.Error(err, "This company already exist.")
	s.Equal(uint(0), company.ID)

	company, err = common.CreateCompany("")
	s.Error(err, "COMPANY: The name is mandatory.")
	s.Equal(uint(0), company.ID)
}

func (s *companiesTestSuite) Test06UpdateValid_Company_ByValidID() {
	var target databases.Company
	databases.DB.Where(databases.Company{Name: "Coban"}).First(&target)

	target.Name = "Coban, corp."
	company, err := common.UpdateCompany(target.Name, target.ID)
	s.NoError(err)
	s.Equal(target.ID, company.ID)
	s.Equal(target.Name, company.Name)
}

func (s *companiesTestSuite) Test07UpdateByValid_Company_ByInvalidID() {
	_, err := common.UpdateCompany("Coban", 0)
	s.Error(err, "This company doesn't exist.")

	_, err = common.UpdateCompany("Coban", 10)
	s.Error(err, "This company doesn't exist.")
}

func (s *companiesTestSuite) Test08UpdateByInvalid_Company_ByValidID() {
	var target databases.Company
	databases.DB.Where(databases.Company{Name: "Coban, corp."}).First(&target)

	_, err := common.UpdateCompany("アコム株式会社", target.ID)
	s.Error(err, "COMPANY: This company already exist.")

	_, err = common.UpdateCompany("", target.ID)
	s.Error(err, "COMPANY: The name is mandatory.")
}

func (s *companiesTestSuite) Test09Delete_Company_ByValidID() {
	var target databases.Company
	databases.DB.Where(databases.Company{Name: "Coban, corp."}).First(&target)

	err := common.DeleteCompany(target.ID)
	s.NoError(err)

	var company databases.Company
	databases.DB.Where(databases.Company{Name: "Coban, corp."}).First(&company)
	s.Equal(uint(0), company.ID)
}

func (s *companiesTestSuite) Test10Delete_Company_ByInvalidID() {
	err := common.DeleteCompany(0)
	s.Error(err, "This company doesn't exist.")

	err = common.DeleteCompany(10)
	s.Error(err, "This company doesn't exist.")
}
