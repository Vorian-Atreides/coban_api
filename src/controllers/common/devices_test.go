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

func (s *devicesTestSuite) Test01Get_Devices() {
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

func (s *devicesTestSuite) Test02Get_Device_ByValidID() {
	expectedUser := &databases.User{ID:1, FirstName:"青木", LastName:"真琳", AccountID:1, CompanyID:1}
	expectedDevice := databases.Device{ID:1, IsPaired:false, UserID:1, User:expectedUser}

	device, err := common.GetDeviceByID(expectedDevice.ID)
	s.NoError(err)
	s.Equal(expectedDevice, device)
}

func (s *devicesTestSuite) Test03Get_Device_ByInvalidID() {
	device, err := common.GetDeviceByID(0)
	s.Error(err, "This device doesn't exist.")
	s.Equal(uint(0), device.ID)

	device, err = common.GetDeviceByID(10)
	s.Error(err, "This device doesn't exist.")
	s.Equal(uint(0), device.ID)
}

func (s *devicesTestSuite) Test04CreateValid_Device() {
	expectedDevice := databases.Device{IsPaired:false, UserID:3}

	device, err := common.CreateDevice(expectedDevice.UserID)
	s.NoError(err)
	s.NotEqual(uint(0), device.ID)
	s.Equal(expectedDevice.UserID, device.UserID)
	s.Equal(expectedDevice.IsPaired, device.IsPaired)
}

func (s *devicesTestSuite) Test05CreateInvalid_Device() {
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

func (s *devicesTestSuite) Test06UpdateValid_Device_ByValidID() {
	var expectedDevice databases.Device

	databases.DB.Where(databases.Device{UserID:3}).First(&expectedDevice)
	expectedDevice.IsPaired = true
	device, err := common.UpdateDevice(expectedDevice.IsPaired, expectedDevice.UserID, expectedDevice.ID)

	s.NoError(err)
	s.Equal(expectedDevice, device)
}

func (s *devicesTestSuite) Test07UpdateValid_Device_ByInvalidID() {
	_, err := common.UpdateDevice(false, 2, 0)
	s.Error(err, "This device doesn't exist.")

	_, err = common.UpdateDevice(false, 2, 10)
	s.Error(err, "This device doesn't exist.")
}

func (s *devicesTestSuite) Test08UpdateInvalid_Device_ByValidID() {
	_, err := common.UpdateDevice(false, 0, 1)
	s.Error(err, "DEVICE: The user is mandatory.")

	_, err = common.UpdateDevice(false, 10, 1)
	s.Error(err, "DEVICE: This user doesn't exist.")
}

func (s *devicesTestSuite) Test09Delete_Device_ByValidID() {
	var target databases.Device

	databases.DB.Where(databases.Device{UserID:3}).First(&target)
	err := common.DeleteDevice(target.ID)
	s.NoError(err)

	target = databases.Device{}
	databases.DB.Where(databases.Device{UserID:3}).First(&target)
	s.Equal(uint(0), target.ID)
}

func (s *devicesTestSuite) Test10Delete_Device_ByInvalidID() {
	err := common.DeleteDevice(0)
	s.Error(err, "This device doesn't exist.")

	err = common.DeleteDevice(10)
	s.Error(err, "This device doesn't exist.")
}