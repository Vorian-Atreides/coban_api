package common_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
)

type addressesTestSuite struct {
	suite.Suite
}

func TestAddresses(t *testing.T) {
	suite.Run(t, new(addressesTestSuite))
}

func (s *addressesTestSuite) Test01Get_Addresses() {
	companies := []*databases.Company{
		&databases.Company{ID: 1, Name: "アコム株式会社"},
		&databases.Company{ID: 2, Name: "株式会社愛知銀行"},
		&databases.Company{ID: 3, Name: "AOCホールディングス株式会社"},
		&databases.Company{ID: 4, Name: "旭化成株式会社"},
		&databases.Company{ID: 5, Name: "株式会社バンダイ"},
	}

	expectedAddresses := []databases.Address{
		databases.Address{ID: 1, Zip: "100-8307",
			Street: "千代田区丸の内二丁目1番1号明治安田生命ビル",
			City:   "東京都", CompanyID: 1, Company: companies[0]},
		databases.Address{ID: 2, Zip: "23106-1",
			Street: "中区栄3-14-12",
			City:   "名古屋市", CompanyID: 2, Company: companies[1]},
		databases.Address{ID: 3, Zip: "140-0002",
			Street: "品川区東品川二丁目5番8号",
			City:   "東京都", CompanyID: 3, Company: companies[2]},
		databases.Address{ID: 4, Zip: "101-8101",
			Street: "千代田区神田神保町1丁目105番地",
			City:   "東京都", CompanyID: 4, Company: companies[3]},
		databases.Address{ID: 5, Zip: "111-8081",
			Street: "台東区駒形1丁目4-8",
			City:   "東京都", CompanyID: 5, Company: companies[4]},
		databases.Address{ID: 6, Zip: "111-8081",
			Street: "東京都台東区駒形2-5-4",
			City:   "東京都", CompanyID: 5, Company: companies[4]},
	}

	addresses := common.GetAddresses(0)
	s.Equal(expectedAddresses, addresses)
}

func (s *addressesTestSuite) Test10Get_Addresses_Paginated() {
	var expected []databases.Address

	addresses := common.GetAddresses(50)
	s.Equal(expected, addresses)
}

func (s *addressesTestSuite) Test02Get_Address_ByValidID() {
	expectedCompany := new(databases.Company)
	expectedCompany.ID = 4
	expectedCompany.Name = "旭化成株式会社"
	expectedAddress := databases.Address{ID: 4, Zip: "101-8101",
		Street: "千代田区神田神保町1丁目105番地",
		City:   "東京都", CompanyID: 4, Company: expectedCompany}

	address, err := common.GetAddressByID(expectedAddress.ID)
	s.NoError(err)
	s.Equal(expectedAddress, address)
}

func (s *addressesTestSuite) Test03Get_Address_ByInvalidID() {
	address, err := common.GetAddressByID(10)
	s.Error(err, "This address doesn't exist")
	s.Equal(uint(0), address.ID)

	address, err = common.GetAddressByID(0)
	s.Error(err, "This address doesn't exist")
	s.Equal(uint(0), address.ID)
}

func (s *addressesTestSuite) Test04Create_Address() {
	expectedAddress := databases.Address{Street: "123 Ginza",
		Zip: "109-8320", City: "Tokyo", CompanyID: 1}

	address, err := common.CreateAddress(expectedAddress.Street,
		expectedAddress.Zip,
		expectedAddress.City, expectedAddress.CompanyID)
	s.NoError(err)
	s.NotEqual(uint(0), address.ID)
	s.Equal(expectedAddress.City, address.City)
	s.Equal(expectedAddress.Street, address.Street)
	s.Equal(expectedAddress.Zip, address.Zip)
	s.Equal(expectedAddress.CompanyID, address.CompanyID)
}

func (s *addressesTestSuite) Test05CreateInvalid_Address() {
	address, err := common.CreateAddress("123 Ginza", "109-8320", "Tokyo", 1)
	s.Error(err, "ADDRESS: This address already exist.")
	s.Equal(uint(0), address.ID)

	address, err = common.CreateAddress("", "110-8320", "Tokyo", 1)
	s.Error(err, "ADDRESS: The street is mandatory.")
	s.Equal(uint(0), address.ID)

	address, err = common.CreateAddress("145 Ginza", "", "Tokyo", 1)
	s.Error(err, "ADDRESS: The zip is mandatory.")
	s.Equal(uint(0), address.ID)

	address, err = common.CreateAddress("145 Ginza", "110-8320", "", 1)
	s.Error(err, "ADDRESS: The city is mandatory.")
	s.Equal(uint(0), address.ID)

	address, err = common.CreateAddress("145 Ginza", "110-8320", "Tokyo", 0)
	s.Error(err, "ADDRESS: The company is mandatory.")
	s.Equal(uint(0), address.ID)

	address, err = common.CreateAddress("145 Ginza", "110-8320", "Tokyo", 10)
	s.Error(err, "ADDRESS: This company doesn't exist.")
	s.Equal(uint(0), address.ID)
}

func (s *addressesTestSuite) Test06UpdateValid_Address_ByValidID() {
	var target databases.Address
	databases.DB.Where(databases.Address{Street: "123 Ginza"}).First(&target)

	expectedAddress := databases.Address{Street: "155 Ginza", Zip: "109-8320",
		City: "Nagoya", CompanyID: 2, ID: target.ID}
	address, err := common.UpdateAddress(expectedAddress.Street,
		expectedAddress.Zip,
		expectedAddress.City, expectedAddress.CompanyID, expectedAddress.ID)
	s.NoError(err)
	s.Equal(expectedAddress, address)
}

func (s *addressesTestSuite) Test07UpdateValid_Address_ByInvalidID() {
	expectedAddress := databases.Address{Street: "188 Ginza", Zip: "110-8320",
		City: "Nagazaki", CompanyID: 2, ID: 0}
	_, err := common.UpdateAddress(expectedAddress.Street, expectedAddress.Zip,
		expectedAddress.City, expectedAddress.CompanyID, expectedAddress.ID)
	s.Error(err, "This address doesn't exist.")

	expectedAddress = databases.Address{Street: "188 Ginza", Zip: "110-8320",
		City: "Nagazaki", CompanyID: 2, ID: 10}
	_, err = common.UpdateAddress(expectedAddress.Street, expectedAddress.Zip,
		expectedAddress.City, expectedAddress.CompanyID, expectedAddress.ID)
	s.Error(err, "This address doesn't exist.")
}

func (s *addressesTestSuite) Test08UpdateInvalid_Address_ByValidID() {
	var target databases.Address
	databases.DB.Where(databases.Address{Street: "155 Ginza"}).First(&target)

	expectedAddress := databases.Address{Street: "品川区東品川二丁目5番8号",
		Zip: "140-0002", City: "東京都", CompanyID: 3, ID: target.ID}
	_, err := common.UpdateAddress(expectedAddress.Street, expectedAddress.Zip,
		expectedAddress.City, expectedAddress.CompanyID, expectedAddress.ID)
	s.Error(err, "ADDRESS: This address already exist.")

	expectedAddress = databases.Address{Street: "", Zip: "101-8101",
		City: "東京都", CompanyID: 4, ID: target.ID}
	_, err = common.UpdateAddress(expectedAddress.Street, expectedAddress.Zip,
		expectedAddress.City, expectedAddress.CompanyID, expectedAddress.ID)
	s.Error(err, "ADDRESS: The street is mandatory.")

	expectedAddress = databases.Address{Street: "千代田区神田神保町1丁目105番地",
		Zip: "", City: "東京都", CompanyID: 4, ID: target.ID}
	_, err = common.UpdateAddress(expectedAddress.Street, expectedAddress.Zip,
		expectedAddress.City, expectedAddress.CompanyID, expectedAddress.ID)
	s.Error(err, "ADDRESS: The zip is mandatory.")

	expectedAddress = databases.Address{Street: "千代田区神田神保町1丁目105番地",
		Zip: "101-8101", City: "", CompanyID: 4, ID: target.ID}
	_, err = common.UpdateAddress(expectedAddress.Street, expectedAddress.Zip,
		expectedAddress.City, expectedAddress.CompanyID, expectedAddress.ID)
	s.Error(err, "ADDRESS: The city is mandatory.")

	expectedAddress = databases.Address{Street: "千代田区神田神保町1丁目105番地",
		Zip: "101-8101", City: "東京都", CompanyID: 0, ID: target.ID}
	_, err = common.UpdateAddress(expectedAddress.Street, expectedAddress.Zip,
		expectedAddress.City, expectedAddress.CompanyID, expectedAddress.ID)
	s.Error(err, "ADDRESS: The company is mandatory.")

	expectedAddress = databases.Address{Street: "千代田区神田神保町1丁目105番地",
		Zip: "101-8101", City: "東京都", CompanyID: 10, ID: target.ID}
	_, err = common.UpdateAddress(expectedAddress.Street, expectedAddress.Zip,
		expectedAddress.City, expectedAddress.CompanyID, expectedAddress.ID)
	s.Error(err, "ADDRESS: This company doesn't exist.")
}

func (s *addressesTestSuite) Test09Delete_Address_ByValidID() {
	var target databases.Address
	databases.DB.Where(databases.Address{Street: "155 Ginza"}).First(&target)

	err := common.DeleteAddressByID(target.ID)
	s.NoError(err)

	target = databases.Address{}
	databases.DB.Where(databases.Address{Street: "155 Ginza"}).First(&target)
	s.Equal(uint(0), target.ID)
}

func (s *addressesTestSuite) Test09Delete_Address_ByInvalidID() {
	err := common.DeleteAddressByID(0)
	s.Error(err, "This address doesn't exist")
	err = common.DeleteAddressByID(10)
	s.Error(err, "This address doesn't exist")
}
