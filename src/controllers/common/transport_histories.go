package common

import (
	"time"

	"coban/api/src/databases"
)

func GetTransportHistories() []databases.TransportHistory {
	var transportHistories []databases.TransportHistory

	databases.DB.Find(&transportHistories)
	for i, _ := range transportHistories {
		transportHistories[i].LoadRelated()
	}

	return transportHistories
}

func GetTransportHistoryByID(id uint) databases.TransportHistory {
	var transportHistory databases.TransportHistory

	databases.DB.First(&transportHistory, id)
	transportHistory.LoadRelated()

	return transportHistory
}

func CreateTransportHistory(date time.Time, stock uint, expense uint,
	entranceID uint, exitID uint) (databases.TransportHistory, error) {

	transportHistory := databases.TransportHistory{Date:date, Stock:stock, Expense:expense,
		EntranceID:entranceID, ExitID:exitID}

	if err := transportHistory.IsValid(); err != nil {
		return transportHistory, err
	}
	databases.DB.Save(&transportHistory)

	return transportHistory, nil
}