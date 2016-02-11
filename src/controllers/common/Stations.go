package common

import (
	"errors"

	"coban/api/src/databases"
)

// GetStations get every stations in the databases
func GetStations() []databases.Station {
	var stations []databases.Station

	databases.DB.Find(&stations)
	for i := range stations {
		stations[i].LoadRelated()
	}

	return stations
}

// GetStationByID get a station by its ID
func GetStationByID(id uint) (databases.Station, error) {
	var station databases.Station

	databases.DB.First(&station, id)
	if station.ID == 0 {
		return station, errors.New("This station doesn't exist.")
	}
	station.LoadRelated()

	return station, nil
}

// CreateStation try to create a new station
func CreateStation(name string, transportType string) (databases.Station, error) {
	station := databases.Station{Name: name, Type: transportType}

	if err := station.IsValid(); err != nil {
		return station, err
	}
	databases.DB.Save(&station)

	return station, databases.DB.Error
}

// UpdateStation try to update a station
func UpdateStation(name string, transporType string, id uint) (databases.Station, error) {
	station := databases.Station{Name: name, Type: transporType, ID: id}

	var target databases.Station
	databases.DB.First(&target, id)
	if target.ID == 0 {
		return station, errors.New("This station doesn't exist.")
	}
	if err := station.IsValid(); err != nil {
		return station, err
	}
	databases.DB.Save(&station)

	return station, databases.DB.Error
}
