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

func GetDeviceByID(id uint) databases.Device {
	var device databases.Device

	databases.DB.First(&device, id)
	device.LoadRelated()

	return device
}

func CreateDevice(userID uint) (databases.Device, error) {
	device := databases.Device{IsPaired:false, UserID:userID}

	if err := device.IsValid(); err != nil {
		return device, err
	}
	databases.DB.Save(&device)

	return device, nil
}

func UpdateDevice(isPaired bool, userID uint, id uint) (databases.Device, error) {
	device := databases.Device{IsPaired:isPaired, UserID:userID, ID:id}

	if err := device.IsValid(); err != nil {
		return device, err
	}
	databases.DB.Update(&device)

	return device, nil
}

func DeleteDevice(id uint) error {
	var device databases.Device

	databases.DB.First(&device, id)
	databases.DB.Delete(&device)

	if device.ID != 0 {
		return errors.New("This device can't be deleted.")
	}

	return nil
}