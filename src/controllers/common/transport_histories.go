package common

import (
	"errors"
	"time"

	"coban/api/src/databases"
	"coban/api/src/utils"
)

// GetTransportHistories get every transport histories from the database
func GetTransportHistories(offset uint) []databases.TransportHistory {
	var transportHistories []databases.TransportHistory

	databases.DB.Offset(offset).Find(&transportHistories).Limit(utils.PageSize)
	for i := range transportHistories {
		transportHistories[i].LoadRelated()
	}

	return transportHistories
}

// GetTransportHistoriesBetweenDates get the transport histories between two
// range of date
func GetTransportHistoriesBetweenDates(offset uint, begin time.Time,
	end time.Time) []databases.TransportHistory {
	var transportHistories []databases.TransportHistory

	databases.DB.Offset(offset).
		Where("date BETWEEN ? and ?", begin, end).
		Limit(utils.PageSize).
		Find(&transportHistories)
	for i := range transportHistories {
		transportHistories[i].LoadRelated()
	}

	return transportHistories
}

// GetTransportHistoryByID get a station by its ID
func GetTransportHistoryByID(id uint) (databases.TransportHistory, error) {
	var transportHistory databases.TransportHistory

	databases.DB.First(&transportHistory, id)
	if transportHistory.ID == 0 {
		return transportHistory, errors.New("This history doesn't exist.")
	}
	transportHistory.LoadRelated()

	return transportHistory, databases.DB.Error
}

// CreateTransportHistory try to create a new transport history
func CreateTransportHistory(date time.Time, stock uint, entranceID uint,
	exitID uint, userID uint) (databases.TransportHistory, error) {

	transportHistory := databases.TransportHistory{Date: date, Stock: stock,
		EntranceID: entranceID, ExitID: exitID, UserID: userID}

	if err := transportHistory.IsValid(); err != nil {
		return transportHistory, err
	}
	databases.DB.Save(&transportHistory)

	return transportHistory, databases.DB.Error
}
