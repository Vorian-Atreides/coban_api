package common

import (
	"errors"

	"coban/api/src/databases"
)

func GetDevices() []databases.Device {
	var devices []databases.Device

	databases.DB.Find(&devices)
	for i, _ := range devices {
		devices[i].LoadRelated()
	}

	return devices
}

func GetDeviceByID(id uint) (databases.Device, error) {
	var device databases.Device

	databases.DB.First(&device, id)
	if device.ID == 0 {
		return device, errors.New("This device doesn't exist.")
	}
	device.LoadRelated()

	return device, nil
}

func CreateDevice(userID uint) (databases.Device, error) {
	device := databases.Device{IsPaired:false, UserID:userID}

	if err := device.IsValid(); err != nil {
		return device, err
	}
	databases.DB.Save(&device)

	return device, databases.DB.Error
}

func UpdateDevice(isPaired bool, userID uint, id uint) (databases.Device, error) {
	device := databases.Device{IsPaired:isPaired, UserID:userID, ID:id}

	var existingDevice databases.Device
	databases.DB.First(&existingDevice, id)
	if existingDevice.ID == 0 {
		return device, errors.New("This device doesn't exist.")
	}

	if err := device.IsValid(); err != nil {
		return device, err
	}
	databases.DB.Save(&device)

	return device, databases.DB.Error
}

func DeleteDevice(id uint) error {
	var device databases.Device

	databases.DB.First(&device, id)
	if device.ID == 0 {
		return errors.New("This device doesn't exist.")
	}
	databases.DB.Delete(&device)

	return databases.DB.Error
}