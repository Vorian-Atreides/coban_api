package common

import (
	"time"

	"coban/api/src/databases"
	"errors"
)

func GetTransportHistories() []databases.TransportHistory {
	var transportHistories []databases.TransportHistory

	databases.DB.Find(&transportHistories)
	for i, _ := range transportHistories {
		transportHistories[i].LoadRelated()
	}

	return transportHistories
}

func GetTransportHistoryByID(id uint) (databases.TransportHistory, error) {
	var transportHistory databases.TransportHistory

	databases.DB.First(&transportHistory, id)
	if transportHistory.ID == 0 {
		return transportHistory, errors.New("This history doesn't exist.")
	}
	transportHistory.LoadRelated()

	return transportHistory, databases.DB.Error
}

func CreateTransportHistory(date time.Time, stock uint, expense uint,
	entranceID uint, exitID uint, userID uint) (databases.TransportHistory, error) {

	transportHistory := databases.TransportHistory{Date:date, Stock:stock, Expense:expense,
		EntranceID:entranceID, ExitID:exitID, UserID:userID}

	if err := transportHistory.IsValid(); err != nil {
		return transportHistory, err
	}
	databases.DB.Save(&transportHistory)

	return transportHistory, databases.DB.Error
}