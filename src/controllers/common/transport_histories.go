package common

import (
	"errors"
	"log"
	"time"

	"github.com/jinzhu/gorm"

	"coban/api/src/databases"
	"coban/api/src/utils"
)

func generalGetQuery(offset uint, begin time.Time, end time.Time) *gorm.DB {
	query := databases.DB.Offset(offset).Limit(utils.PageSize)
	log.Println(begin, end)
	if !begin.IsZero() {
		query = query.Where("date > ?", begin)
	}
	if !end.IsZero() {
		query = query.Where("date < ?", end)
	}

	return query
}

// GetTransportHistories get every transport histories from the database
func GetTransportHistories(offset uint, begin time.Time,
	end time.Time) []databases.TransportHistory {
	var transportHistories []databases.TransportHistory

	query := generalGetQuery(offset, begin, end)
	query.Find(&transportHistories)
	for i := range transportHistories {
		transportHistories[i].LoadRelated()
	}

	return transportHistories
}

// GetTransportHistoriesByUserID get every transport histories
// related to an user
func GetTransportHistoriesByUserID(offset uint, begin time.Time,
	end time.Time, id uint) []databases.TransportHistory {
	var transportHistories []databases.TransportHistory

	query := generalGetQuery(offset, begin, end)
	query.Where(&databases.TransportHistory{UserID: id}).
		Find(&transportHistories)
	for i := range transportHistories {
		transportHistories[i].LoadRelated()
	}

	return transportHistories
}

// GetTransportHistoriesByCompanyID get every transport histories
// related to the company's employees
func GetTransportHistoriesByCompanyID(offset uint, begin time.Time,
	end time.Time, id uint) []databases.TransportHistory {
	var transportHistories []databases.TransportHistory

	query := generalGetQuery(offset, begin, end)
	query.Joins("left join users on users.id = transport_histories.user_id").
		Where(&databases.User{CompanyID: id}).Find(&transportHistories)
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
