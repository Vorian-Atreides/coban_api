package common

import (
	"coban/api/src/databases"
)

func GetStations() []databases.Station {
	var stations []databases.Station

	databases.DB.Find(&stations)
	for i, _ := range stations {
		stations[i].LoadRelated()
	}

	return stations
}

func GetstationByID(id uint) databases.Station {
	var station databases.Station

	databases.DB.First(&station, id)
	station.LoadRelated()

	return station
}

func CreateStation(name string, transportTypeID uint) (databases.Station, error) {
	station := databases.Station{Name:name, TypeID:transportTypeID}

	if err := station.IsValid(true); err != nil {
		return station, err
	}
	databases.DB.Save(&station)

	return station, nil
}

func UpdateStation(name string, transporType uint, id uint) (databases.Station, error) {
	station := databases.Station{Name:name, TypeID:transporType, ID:id}

	if err := station.IsValid(false); err != nil {
		return station, err
	}
	databases.DB.Update(&station)

	return station, nil
}