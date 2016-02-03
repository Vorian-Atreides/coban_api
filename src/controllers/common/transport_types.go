package common

import (
	"coban/api/src/databases"
)

func GetTransportTypes() []databases.TransportType {
	var transportTypes []databases.TransportType

	databases.DB.Find(&transportTypes)
	for i, _ := range transportTypes {
		transportTypes[i].LoadRelated()
	}

	return transportTypes
}

func GetTransportTypeByID(id uint) databases.TransportType {
	var transportType databases.TransportType

	databases.DB.First(&transportType)
	transportType.LoadRelated()

	return transportType
}

func CreateTransportType(name string) (databases.TransportType, error) {
	transportType := databases.TransportType{Name:name}

	if err := transportType.IsValid(true); err != nil {
		return transportType, err
	}
	databases.DB.Save(&transportType)

	return transportType, nil
}

func UpdateTranportType(name string, id uint) (databases.TransportType, error) {
	transportType := databases.TransportType{Name:name, ID:id}

	if err := transportType.IsValid(false); err != nil {
		return transportType, err
	}
	databases.DB.Update(&transportType)

	return transportType, nil
}