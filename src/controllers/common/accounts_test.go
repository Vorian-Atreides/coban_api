package common_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
)

type accountsTestSuite struct {
	suite.Suite
}

func TestAccounts(t *testing.T) {
	suite.Run(t, new(accountsTestSuite))
}

func (s *accountsTestSuite) Test01Get_Accounts() {
	expectedAccounts := []databases.Account{
		databases.Account{ID: 1, Email: "user@coban.jp",
			Password: "b14361404c078ffd549c03db443c3fede2f3e534d73f78f7730" +
				"1ed97d4a436a9fd9db05ee8b325c0ad36438b43fec8510c204fc1c1ed" +
				"b21d0941c00e9e2c1ce2", Scope: 1,
			User: &databases.User{ID: 1, FirstName: "青木", LastName: "真琳",
				AccountID: 1, CompanyID: 1},
		},
		databases.Account{ID: 2, Email: "office@coban.jp",
			Password: "f358a8caf95e1889d88444d054c847506d1448fcf0" +
				"3336a621bb9d62ad228d47ca467c0a61d56933f59cc59edb" +
				"0688270549b4c5d17a6f4937b077d643b868ce", Scope: 2,
			User: &databases.User{ID: 2, FirstName: "織田",
				LastName: "信長", AccountID: 2, CompanyID: 2},
		},
		databases.Account{ID: 3, Email: "admin@coban.jp",
			Password: "c7ad44cbad762a5da0a452f9e854fdc1e0e7a52a380" +
				"15f23f3eab1d80b931dd472634dfac71cd34ebc35d16ab7fb" +
				"8a90c81f975113d6c7538dc69dd8de9077ec", Scope: 4,
			User: &databases.User{ID: 3, FirstName: "豊臣",
				LastName: "秀吉", AccountID: 3, CompanyID: 3},
		},
		databases.Account{ID: 4, Email: "root@coban.jp",
			Password: "99adc231b045331e514a516b4b7680f588e3823213ab" +
				"e901738bc3ad67b2f6fcb3c64efb93d18002588d3ccc1a49ef" +
				"bae1ce20cb43df36b38651f11fa75678e8", Scope: 7,
			User: &databases.User{ID: 4, FirstName: "徳川",
				LastName: "家康", AccountID: 4, CompanyID: 4},
		},
		databases.Account{ID: 5, Email: "other@coban.jp",
			Password: "e25ac3845f8cbe12801a2dfa5a89d4c55dc47900f3b6" +
				"edc9a9ee590f3c2b9312f665d0039c93828b7b58f33950bc81" +
				"7a0955a9c5000a8d3e280569f08745ca68",
			Scope: 1, User: &databases.User{}},
		databases.Account{ID: 6, Email: "other2@coban.jp",
			Password: "acd97a6214b6648dd2859dcc48afe1f2ad0603a634d60" +
				"652f00028f17df150d4298337101cee4456ffa326f18077f2cc" +
				"1f6ba9f52d13020816127fa0e4378fdf",
			Scope: 1, User: &databases.User{}},
	}

	accounts := common.GetAccounts(0)
	s.Equal(expectedAccounts, accounts)
}

func (s *accountsTestSuite) Test16Get_Accounts_Paginated() {
	expected := []databases.Account{}

	accounts := common.GetAccounts(50)
	s.Equal(expected, accounts)
}

func (s *accountsTestSuite) Test02Get_Account_ByValidID() {
	expectedAccount := databases.Account{ID: 1, Email: "user@coban.jp",
		Password: "b14361404c078ffd549c03db443c3fede2f3e534d73f78f773" +
			"01ed97d4a436a9fd9db05ee8b325c0ad36438b43fec8510c204fc1c1" +
			"edb21d0941c00e9e2c1ce2", Scope: 1,
		User: &databases.User{ID: 1, FirstName: "青木",
			LastName: "真琳", AccountID: 1, CompanyID: 1},
	}

	account, err := common.GetAccountByID(expectedAccount.ID)
	s.NoError(err)
	s.Equal(expectedAccount, account)
}

func (s *accountsTestSuite) Test03Get_Account_ByInvalidID() {
	account, err := common.GetAccountByID(0)
	s.Error(err, "This account doesn't exist")
	s.Equal(uint(0), account.ID)

	account, err = common.GetAccountByID(10)
	s.Error(err, "This account doesn't exist")
	s.Equal(uint(0), account.ID)
}

func (s *accountsTestSuite) Test04Create_Account() {
	expectedAccount := databases.Account{Email: "gaston.siffert@coban.jp", Password: "1234567890", Scope: 1}

	account, err := common.CreateAccount(expectedAccount.Email, expectedAccount.Scope, expectedAccount.Password)
	s.NoError(err)
	s.NotEqual(uint(0), account.ID)
	s.Equal(expectedAccount.Email, account.Email)
	s.Equal("12b03226a6d8be9c6e8cd5e55dc6c7920caaa39df14aab92d5e3ea9340d1c8a4d3d0b8e4314f1f6ef131ba4bf1ceb9186ab87c801af0d5c95b1befb8cedae2b9", account.Password)
	s.Equal(expectedAccount.Scope, account.Scope)
}

func (s *accountsTestSuite) Test05CreateInvalid_Account() {
	account, err := common.CreateAccount("gaston.siffert@coban.jp", 1, "1234567890")
	s.Error(err, "ACCOUNT: This account already exist.")
	s.Equal(uint(0), account.ID)

	account, err = common.CreateAccount("gaston.siffert@coban.jp", 1, "")
	s.Error(err, "ACCOUNT: The password is mandatory.")
	s.Equal(uint(0), account.ID)

	account, err = common.CreateAccount("gaston.siffert@coban.jp", 0, "1234567890")
	s.Error(err, "ACCOUNT: The scope is mandatory.")
	s.Equal(uint(0), account.ID)

	account, err = common.CreateAccount("", 1, "1234567890")
	s.Error(err, "ACCOUNT: The email is mandatory..")
	s.Equal(uint(0), account.ID)
}

func (s *accountsTestSuite) Test06UpdateValid_Account_ByValidID() {
	var expected databases.Account
	databases.DB.Where(databases.Account{Email: "gaston.siffert@coban.jp"}).First(&expected)
	expected.Email = "gaston@coban.jp"
	expected.Scope = 4

	account, err := common.UpdateAccount(expected.Email, expected.Scope, expected.ID)
	s.NoError(err)
	s.Equal(expected, account)
}

func (s *accountsTestSuite) Test07UpdatePasswordValid_Account_ByValidID() {
	var expected databases.Account
	databases.DB.Where(databases.Account{Email: "gaston@coban.jp"}).First(&expected)
	expected.Password = "9c7ce0e5f382bcb4a5032ea44d1a63d0438322532f431610ed482696100c4d2ae1f048f70cda120fcb799f632231fd090d7779bf35e15e796f4e7212df74ff11"

	account, err := common.UpdateAccountPassword("cobanpassword", expected.ID)
	s.NoError(err)
	s.Equal(expected, account)
}

func (s *accountsTestSuite) Test08UpdateValid_Account_ByInvalidID() {
	_, err := common.UpdateAccount("siffert@coban.jp", 1, 0)
	s.Error(err, "This account doesn't exist.")

	_, err = common.UpdateAccount("siffert@coban.jp", 1, 10)
	s.Error(err, "This account doesn't exist.")
}

func (s *accountsTestSuite) Test09UpdateInvalid_Account_ByValidID() {
	var target databases.Account
	databases.DB.Where(databases.Account{Email: "gaston@coban.jp"}).First(&target)

	_, err := common.UpdateAccount("admin@coban.jp", 1, target.ID)
	s.Error(err, "ACCOUNT: This account already exist.")

	_, err = common.UpdateAccount("siffert@coban.jp", 0, target.ID)
	s.Error(err, "ACCOUNT: The scope is mandatory.")

	_, err = common.UpdateAccount("", 1, target.ID)
	s.Error(err, "ACCOUNT: The email is mandatory.")
}

func (s *accountsTestSuite) Test10UpdateValidPassword_Account_ByInvalidID() {
	_, err := common.UpdateAccountPassword("PasswordTest", 0)
	s.Error(err, "This account doesn't exist.")

	_, err = common.UpdateAccountPassword("PasswordTest", 10)
	s.Error(err, "This account doesn't exist.")
}

func (s *accountsTestSuite) Test11UpdateInvalidPassword_Account_ByValidID() {
	var target databases.Account
	databases.DB.Where(databases.Account{Email: "gaston@coban.jp"}).First(&target)

	_, err := common.UpdateAccountPassword("", target.ID)
	s.Error(err, "ACCOUNT:  The password is mandatory.")
}

func (s *accountsTestSuite) Test12Delete_Account_ByValidID() {
	var target databases.Account
	databases.DB.Where(databases.Account{Email: "gaston@coban.jp"}).First(&target)

	err := common.DeleteAccount(target.ID)
	s.NoError(err)

	target = databases.Account{}
	databases.DB.Where(databases.Account{Email: "gaston@coban.jp"}).First(&target)
	s.Equal(uint(0), target.ID)
}

func (s *accountsTestSuite) Test13Delete_Account_ByInvalidID() {
	err := common.DeleteAccount(0)
	s.Error(err, "This account doesn't exist.")

	err = common.DeleteAccount(10)
	s.Error(err, "This account doesn't exist.")
}

func (s *accountsTestSuite) Test14AuthenticateValid_Account() {
	_, err := common.Authenticate("user@coban.jp", "user")
	s.NoError(err)
}

func (s *accountsTestSuite) Test15AuthenticateInvalid_Account() {
	_, err := common.Authenticate("user@coban.jp", "wrong")
	s.Error(err, "The credentials are invalid.")

	_, err = common.Authenticate("wrong_user@coban.jp", "wrong")
	s.Error(err, "This account doesn't exist.")
}
