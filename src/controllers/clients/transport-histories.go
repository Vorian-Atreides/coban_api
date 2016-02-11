package clients

import (
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"time"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

type lightTransportHistory struct {
	ID uint `json:"id"`

	Date    time.Time `json:"date"`
	Stock   uint      `json:"stock"`
	Expense uint      `json:"expense"`

	Entrance *databases.Station `json:"entrance; omitempty"`
	Exit     *databases.Station `json:"exit; omitempty"`
}

// GetTransportHistories get the transport histories for the current user
func GetTransportHistories(w http.ResponseWriter, r *http.Request) {
	user, status, err := utils.CheckTokenAndScope(r, databases.IsClient)
	if err != nil {
		utils.Error(w, err, status)
		return
	}

	var lightTransportHistories []lightTransportHistory
	var transportHistories []databases.TransportHistory
	databases.DB.Where(databases.TransportHistory{UserID: user.ID}).
		Find(&transportHistories)
	for i := range transportHistories {
		transportHistories[i].LoadRelated()
		light := lightTransportHistory{
			ID:       transportHistories[i].ID,
			Stock:    transportHistories[i].Stock,
			Expense:  transportHistories[i].Expense,
			Date:     transportHistories[i].Date,
			Entrance: transportHistories[i].Entrance,
			Exit:     transportHistories[i].Exit}
		lightTransportHistories = append(lightTransportHistories, light)
	}

	utils.WriteBody(w, lightTransportHistories, http.StatusOK)
}

type bodyTransportHistory struct {
	Data string `json:"data"`
}

// AddTransportHistory Create a new transport history for the current user
func AddTransportHistory(w http.ResponseWriter, r *http.Request) {
	user, status, err := utils.CheckTokenAndScope(r, databases.IsClient)
	if err != nil {
		utils.Error(w, err, status)
		return
	}

	var transportHistories []databases.TransportHistory
	var transports []string
	utils.ReadBody(r, &transports)
	for i := range transports {
		data, err := base64.StdEncoding.DecodeString(transports[i])
		log.Println(data)
		if err != nil {
			utils.Error(w,
				errors.New("One of the history is badly encoded"),
				http.StatusBadRequest)
			return
		}

		transportHistory, err := common.ParseTransportHistory(data)
		if err != nil {
			utils.Error(w,
				err,
				http.StatusBadRequest)
			return
		}
		transport, err := common.CreateTransportHistory(transportHistory.Date,
			transportHistory.Stock, transportHistory.EntranceID,
			transportHistory.ExitID, user.ID)
		if err != nil {
			utils.Error(w,
				err,
				http.StatusBadRequest)
			return
		}
		transport.LoadRelated()
		transportHistories = append(transportHistories, transport)
	}

	utils.WriteBody(w, transportHistories, http.StatusOK)
}
