package common

import (
	"errors"
	"time"

	"coban/api/src/databases"
)

// GetTransportHistories get every transport histories from the database
func GetTransportHistories() []databases.TransportHistory {
	var transportHistories []databases.TransportHistory

	databases.DB.Find(&transportHistories)
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
