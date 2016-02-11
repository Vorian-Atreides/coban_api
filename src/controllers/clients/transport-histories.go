package clients

import (
	"net/http"
	"time"

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

// AddTransportHistory Create a new transport history for the current user
func AddTransportHistory(w http.ResponseWriter, r *http.Request) {
	_, status, err := utils.CheckTokenAndScope(r, databases.IsClient)
	if err != nil {
		utils.Error(w, err, status)
		return
	}

	utils.WriteBody(w, "Unimplemented exception !", http.StatusBadRequest)
}
