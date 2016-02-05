package common

import (
	"coban/api/src/databases"
	"errors"
)

func GetStations() []databases.Station {
	var stations []databases.Station

	databases.DB.Find(&stations)
	for i, _ := range stations {
		stations[i].LoadRelated()
	}

	return stations
}

func GetStationByID(id uint) (databases.Station, error) {
	var station databases.Station

	databases.DB.First(&station, id)
	if station.ID == 0 {
		return station, errors.New("This station doesn't exist.")
	}
	station.LoadRelated()

	return station, nil
}

func CreateStation(name string, transportType string) (databases.Station, error) {
	station := databases.Station{Name:name, Type:transportType}

	if err := station.IsValid(); err != nil {
		return station, err
	}
	databases.DB.Save(&station)

	return station, databases.DB.Error
}

func UpdateStation(name string, transporType string, id uint) (databases.Station, error) {
	station := databases.Station{Name:name, Type:transporType, ID:id}

	var target databases.Station
	databases.DB.First(&target, id)
	if target.ID == 0 {
		return station, errors.New("This station doesn't exist.")
	}
	if err := station.IsValid(); err != nil {
		return station, err
	}
	databases.DB.Update(&station)

	return station, databases.DB.Error
}