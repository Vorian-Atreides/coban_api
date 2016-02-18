package common_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
)

type usersTestSuite struct {
	suite.Suite
}

func TestUsers(t *testing.T) {
	suite.Run(t, new(usersTestSuite))
}

func (s *usersTestSuite) Test01Get_Users() {
	expectedUsers := []databases.User{
		databases.User{ID: 1, FirstName: "青木", LastName: "真琳", AccountID: 1,
			CompanyID: 1,
			Account: &databases.Account{ID: 1, Email: "user@coban.jp",
				Password: "b14361404c078ffd549c03db443c3fede2f3e534d" +
					"73f78f77301ed97d4a436a9fd9db05ee8b325c0ad36438b" +
					"43fec8510c204fc1c1edb21d0941c00e9e2c1ce2", Scope: 1},
			Company: &databases.Company{ID: 1, Name: "アコム株式会社"},
			Device:  &databases.Device{ID: 1, IsPaired: false, UserID: 1},
		},
		databases.User{ID: 2, FirstName: "織田", LastName: "信長", AccountID: 2,
			CompanyID: 2,
			Account: &databases.Account{ID: 2, Email: "office@coban.jp",
				Password: "f358a8caf95e1889d88444d054c847506d1448fcf03" +
					"336a621bb9d62ad228d47ca467c0a61d56933f59cc59edb06" +
					"88270549b4c5d17a6f4937b077d643b868ce", Scope: 2},
			Company: &databases.Company{ID: 2, Name: "株式会社愛知銀行"},
			Device:  &databases.Device{},
		},
		databases.User{ID: 3, FirstName: "豊臣", LastName: "秀吉", AccountID: 3,
			CompanyID: 3,
			Account: &databases.Account{ID: 3, Email: "admin@coban.jp",
				Password: "c7ad44cbad762a5da0a452f9e854fdc1e0e7a52a38" +
					"015f23f3eab1d80b931dd472634dfac71cd34ebc35d16ab7" +
					"fb8a90c81f975113d6c7538dc69dd8de9077ec", Scope: 4},
			Company: &databases.Company{ID: 3, Name: "AOCホールディングス株式会社"},
			Device:  &databases.Device{},
		},
		databases.User{ID: 4, FirstName: "徳川", LastName: "家康", AccountID: 4,
			CompanyID: 4,
			Account: &databases.Account{ID: 4, Email: "root@coban.jp",
				Password: "99adc231b045331e514a516b4b7680f588e382321" +
					"3abe901738bc3ad67b2f6fcb3c64efb93d18002588d3ccc" +
					"1a49efbae1ce20cb43df36b38651f11fa75678e8", Scope: 7},
			Company: &databases.Company{ID: 4, Name: "旭化成株式会社"},
			Device:  &databases.Device{ID: 2, IsPaired: true, UserID: 4},
		},
	}

	users := common.GetUsers(0)
	s.Equal(expectedUsers, users)
}

func (s *usersTestSuite) Test10Get_Users_Paginated() {
	expected := []databases.User{}

	users := common.GetUsers(50)
	s.Equal(expected, users)
}

func (s *usersTestSuite) Test02Get_User_ByValidID() {
	expectedUser := databases.User{ID: 1, FirstName: "青木", LastName: "真琳",
		AccountID: 1, CompanyID: 1,
		Account: &databases.Account{ID: 1, Email: "user@coban.jp",
			Password: "b14361404c078ffd549c03db443c3fede2f3e534d" +
				"73f78f77301ed97d4a436a9fd9db05ee8b325c0ad36438b" +
				"43fec8510c204fc1c1edb21d0941c00e9e2c1ce2", Scope: 1},
		Company: &databases.Company{ID: 1, Name: "アコム株式会社"},
		Device:  &databases.Device{ID: 1, IsPaired: false, UserID: 1},
	}

	user, err := common.GetUserByID(expectedUser.ID)
	s.NoError(err)
	s.Equal(expectedUser, user)
}

func (s *usersTestSuite) Test03Get_User_ByInvalidID() {
	user, err := common.GetUserByID(0)
	s.Error(err, "This user doesn't exist.")
	s.Equal(uint(0), user.ID)

	user, err = common.GetUserByID(10)
	s.Error(err, "This user doesn't exist.")
	s.Equal(uint(0), user.ID)
}

func (s *usersTestSuite) Test04Create_User() {
	expectedUser := databases.User{FirstName: "Gaston",
		LastName: "Siffert", CompanyID: 1}

	account, _ := common.CreateAccount("test04_create_user@coban.jp", 1,
		"a323236d73d145dd4d18cf1a212ba7a99e20564323fb8d219a2232cca6"+
			"6fb6fbb46301233e4b7b181ebad3e045247113bf32d0dad1ab912f"+
			"b80461717857ccad")
	user, err := common.CreateUser(expectedUser.FirstName,
		expectedUser.LastName, account.ID, expectedUser.CompanyID)
	s.NoError(err)
	s.NotEqual(uint(0), user.ID)
	s.Equal(account.Email, user.Account.Email)
	s.Equal(account.Password, user.Account.Password)
	s.Equal(account.Scope, user.Account.Scope)
	s.Equal(expectedUser.FirstName, user.FirstName)
	s.Equal(expectedUser.LastName, user.LastName)
	s.Equal(expectedUser.CompanyID, user.CompanyID)
}

func (s *usersTestSuite) Test05CreateInvalid_User() {

	user, err := common.CreateUser("Gaston", "Siffert", 1, 1)
	s.Error(err, "ACCOUNT: This account is already used.")
	s.Equal(uint(0), user.ID)

	user, err = common.CreateUser("Gaston", "Siffert", 5, 0)
	s.Error(err, "USER: The company is mandatory.")
	s.Equal(uint(0), user.ID)

	user, err = common.CreateUser("Gaston", "Siffert", 0, 1)
	s.Error(err, "USER: The account is mandatory.")
	s.Equal(uint(0), user.ID)

	user, err = common.CreateUser("Gaston", "", 5, 1)
	s.Error(err, "USER: The last name is mandatory.")
	s.Equal(uint(0), user.ID)

	user, err = common.CreateUser("", "Siffert", 5, 1)
	s.Error(err, "USER: The first name is mandatory.")
	s.Equal(uint(0), user.ID)
}

func (s *usersTestSuite) Test06UpdateValid_User_ByValidID() {
	var target databases.User
	databases.DB.Where(databases.User{FirstName: "Gaston",
		LastName: "Siffert", CompanyID: 1}).First(&target)

	expectedUser := databases.User{ID: target.ID, FirstName: "Tastsuya",
		LastName: "Zembutsu", CompanyID: 2, AccountID: target.AccountID}
	user, err := common.UpdateUserByID(expectedUser.FirstName,
		expectedUser.LastName, expectedUser.CompanyID,
		expectedUser.ID)
	s.NoError(err)
	s.Equal(expectedUser.ID, user.ID)
	s.Equal(expectedUser.FirstName, user.FirstName)
	s.Equal(expectedUser.LastName, user.LastName)
	s.Equal(expectedUser.AccountID, user.AccountID)
	s.Equal(expectedUser.CompanyID, user.CompanyID)
}

func (s *usersTestSuite) Test07UpdateValid_User_ByInvalidID() {
	_, err := common.UpdateUserByID("Bernard", "Siffert", 2, 0)
	s.Error(err, "This user doesn't exist.")

	_, err = common.UpdateUserByID("Bernard", "Siffert", 2, 10)
	s.Error(err, "This user doesn't exist.")
}

func (s *usersTestSuite) Test08UpdateInvalid_User_ByValidID() {
	var target databases.User
	databases.DB.Where(databases.User{FirstName: "Tastsuya",
		LastName: "Zembutsu", CompanyID: 2}).First(&target)

	_, err := common.UpdateUserByID("Bernard", "Siffert", 0, target.ID)
	s.Error(err, "USER: The company is mandatory.")

	_, err = common.UpdateUserByID("Bernard", "Siffert", 10, target.ID)
	s.Error(err, "USER: This company doesn't exist.")

	_, err = common.UpdateUserByID("Bernard", "", 1, target.ID)
	s.Error(err, "USER: The last name is mandatory.")

	_, err = common.UpdateUserByID("", "Siffert", 1, target.ID)
	s.Error(err, "USER: The first name is mandatory.")
}

func (s *usersTestSuite) Test09Delete_User_ByValidID() {
	var target databases.User
	databases.DB.Where(databases.User{FirstName: "Tastsuya",
		LastName: "Zembutsu", CompanyID: 2}).First(&target)

	err := common.DeleteUser(target.ID)
	s.NoError(err)

	target = databases.User{}
	databases.DB.Where(databases.User{FirstName: "Tastsuya",
		LastName: "Zembutsu", CompanyID: 2}).First(&target)
	s.Equal(uint(0), target.ID)
}

func (s *usersTestSuite) Test09Delete_User_ByInvalidID() {
	err := common.DeleteUser(0)
	s.Error(err, "This user doesn't exist.")

	err = common.DeleteUser(10)
	s.Error(err, "This user doesn't exist.")
}
