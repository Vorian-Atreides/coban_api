package common

import (
	"errors"

	"coban/api/src/databases"
	"coban/api/src/utils"
)

// GetStations get every stations in the databases
func GetStations(offset int) []databases.Station {
	var stations []databases.Station

	databases.DB.Offset(offset).Limit(utils.PageSize).Find(&stations)
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
