package common_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
)

type devicesTestSuite struct {
	suite.Suite
}

func TestDevices(t *testing.T) {
	suite.Run(t, new(devicesTestSuite))
}

func (s *devicesTestSuite) Test01GetDevices() {
	expectedUsers := []*databases.User{
		&databases.User{ID:1, FirstName:"青木", LastName:"真琳", AccountID:1, CompanyID:1},
		&databases.User{ID:4, FirstName:"徳川", LastName:"家康", AccountID:4, CompanyID:4},
	}
	expectedDevices := []databases.Device{
		databases.Device{ID:1, IsPaired:false, UserID:1, User:expectedUsers[0]},
		databases.Device{ID:2, IsPaired:true, UserID:4, User:expectedUsers[1]},
	}

	devices := common.GetDevices()
	s.Equal(expectedDevices, devices)
}

func (s *devicesTestSuite) Test02GetDeviceByValidID() {
	expectedUser := &databases.User{ID:1, FirstName:"青木", LastName:"真琳", AccountID:1, CompanyID:1}
	expectedDevice := databases.Device{ID:1, IsPaired:false, UserID:1, User:expectedUser}

	device, err := common.GetDeviceByID(expectedDevice.ID)
	s.NoError(err)
	s.Equal(expectedDevice, device)
}

func (s *devicesTestSuite) Test03GetDeviceByInvalidID() {
	device, err := common.GetDeviceByID(0)
	s.Error(err, "This device doesn't exist.")
	s.Equal(uint(0), device.ID)

	device, err = common.GetDeviceByID(10)
	s.Error(err, "This device doesn't exist.")
	s.Equal(uint(0), device.ID)
}

func (s *devicesTestSuite) Test04CreateValidDevice() {
	expectedDevice := databases.Device{IsPaired:false, UserID:3}

	device, err := common.CreateDevice(expectedDevice.UserID)
	s.NoError(err)
	s.NotEqual(uint(0), device.ID)
	s.Equal(expectedDevice.UserID, device.UserID)
	s.Equal(expectedDevice.IsPaired, device.IsPaired)
}

func (s *devicesTestSuite) Test05CreateInvalidDevice() {
	device, err := common.CreateDevice(1)
	s.Error(err, "DEVICE: This device already exist.")
	s.Equal(uint(0), device.ID)

	device, err = common.CreateDevice(0)
	s.Error(err, "DEVICE: The user is mandatory.")
	s.Equal(uint(0), device.ID)

	device, err = common.CreateDevice(10)
	s.Error(err, "DEVICE: This user doesn't exist.")
	s.Equal(uint(0), device.ID)
}

func (s *devicesTestSuite) Test06UpdateValidDeviceByValidID() {
	var expectedDevice databases.Device

	databases.DB.Where(databases.Device{UserID:4}).First(&expectedDevice)
	expectedDevice.IsPaired = false
	device, err := common.UpdateDevice(expectedDevice.IsPaired, expectedDevice.UserID, expectedDevice.ID)

	s.NoError(err)
	s.Equal(expectedDevice, device)
}

func (s *devicesTestSuite) Test07UpdateValidByInvalidID() {
	_, err := common.UpdateDevice(false, 2, 0)
	s.Error(err, "This device doesn't exist.")

	_, err = common.UpdateDevice(false, 2, 10)
	s.Error(err, "This device doesn't exist.")
}

func (s *devicesTestSuite) Test08UpdateInvalidDeviceByValidID() {
	_, err := common.UpdateDevice(false, 0, 1)
	s.Error(err, "DEVICE: The user is mandatory.")

	_, err = common.UpdateDevice(false, 10, 1)
	s.Error(err, "DEVICE: This user doesn't exist.")
}

func (s *devicesTestSuite) Test09DeleteByValidID() {
	var target databases.Device

	databases.DB.Where(databases.Device{UserID:4}).First(&target)
	err := common.DeleteDevice(target.ID)
	s.NoError(err)

	target = databases.Device{}
	databases.DB.Where(databases.Device{UserID:4}).First(&target)
	s.Equal(uint(0), target.ID)
}

func (s *devicesTestSuite) Test10DeleteByInvalidID() {
	err := common.DeleteDevice(0)
	s.Error(err, "This device doesn't exist.")

	err = common.DeleteDevice(10)
	s.Error(err, "This device doesn't exist.")
}