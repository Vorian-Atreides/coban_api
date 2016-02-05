package databases_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"coban/api/src/databases"
)

type accountTestSuite struct {
	suite.Suite
}

func TestAccount(t *testing.T) {
	suite.Run(t, new(accountTestSuite))
}

func (s *accountTestSuite) Test01_Account_BuildScopes() {
	s.Equal(byte(0x03), databases.BuildScope(databases.ClientScope, databases.OfficeScope))
	s.Equal(byte(0x05), databases.BuildScope(databases.ClientScope, databases.AdminScope))
	s.Equal(byte(0x06), databases.BuildScope(databases.OfficeScope, databases.AdminScope))
	s.Equal(byte(0x01), databases.BuildScope(databases.ClientScope, databases.ClientScope))
}

func (s *accountTestSuite) Test02_Account_IsValidClient() {
	s.True(databases.IsClient(databases.ClientScope))
	s.True(databases.IsClient(databases.BuildScope(databases.ClientScope, databases.OfficeScope)))
	s.True(databases.IsClient(databases.BuildScope(databases.ClientScope, databases.AdminScope)))
	s.True(databases.IsClient(databases.BuildScope(databases.ClientScope, databases.OfficeScope, databases.AdminScope)))
}

func (s *accountTestSuite) Test03_Account_IsInvalidClient() {
	s.False(databases.IsClient(databases.OfficeScope))
	s.False(databases.IsClient(databases.AdminScope))
	s.False(databases.IsClient(databases.BuildScope(databases.AdminScope, databases.OfficeScope)))
}

func (s *accountTestSuite) Test04_Account_IsValidOffice() {
	s.True(databases.IsOffice(databases.OfficeScope))
	s.True(databases.IsOffice(databases.BuildScope(databases.OfficeScope, databases.ClientScope)))
	s.True(databases.IsOffice(databases.BuildScope(databases.OfficeScope, databases.AdminScope)))
	s.True(databases.IsOffice(databases.BuildScope(databases.OfficeScope, databases.ClientScope, databases.AdminScope)))
}

func (s *accountTestSuite) Test05_Account_IsInvalidOffice() {
	s.False(databases.IsOffice(databases.ClientScope))
	s.False(databases.IsOffice(databases.AdminScope))
	s.False(databases.IsOffice(databases.BuildScope(databases.AdminScope, databases.ClientScope)))
}

func (s *accountTestSuite) Test06_Account_IsValidAdmin() {
	s.True(databases.IsAdmin(databases.AdminScope))
	s.True(databases.IsAdmin(databases.BuildScope(databases.AdminScope, databases.ClientScope)))
	s.True(databases.IsAdmin(databases.BuildScope(databases.AdminScope, databases.OfficeScope)))
	s.True(databases.IsAdmin(databases.BuildScope(databases.AdminScope, databases.ClientScope, databases.OfficeScope)))
}

func (s *accountTestSuite) Test07_Account_IsInvalidAdmin() {
	s.False(databases.IsAdmin(databases.ClientScope))
	s.False(databases.IsAdmin(databases.OfficeScope))
	s.False(databases.IsAdmin(databases.BuildScope(databases.ClientScope, databases.OfficeScope)))
}