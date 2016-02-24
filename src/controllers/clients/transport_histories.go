package clients

import (
	"encoding/base64"
	"net/http"

	"coban/api/src/controllers/common"
	"coban/api/src/databases"
	"coban/api/src/utils"
)

// GetTransportHistories get the transport histories for the current user
func GetTransportHistories(w http.ResponseWriter, r *http.Request) {
	user, status, err := utils.CheckTokenAndScope(r, databases.IsClient)
	if err != nil {
		utils.Error(w, err, status)
		return
	}

	offset, err := utils.GetPageOffset(r)
	begin, _ := utils.GetDateParameter(r, "begin")
	end, _ := utils.GetDateParameter(r, "end")

	transportHistories := common.GetTransportHistoriesByUserID(offset,
		begin, end, user.ID)
	for i := range transportHistories {
		transportHistories[i].LoadRelated()
	}

	utils.WriteBody(w, transportHistories, http.StatusOK)
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
	status = http.StatusOK

	utils.ReadBody(r, &transports)
	for i := range transports {
		data, err := base64.StdEncoding.DecodeString(transports[i])
		if err != nil {
			status = http.StatusBadRequest
			continue
		}

		transportHistory, err := common.ParseTransportHistory(data)
		if err != nil {
			status = http.StatusBadRequest
			continue
		}
		transport, err := common.CreateTransportHistory(transportHistory.Date,
			transportHistory.Stock, transportHistory.EntranceID,
			transportHistory.ExitID, user.ID)
		if err != nil {
			status = http.StatusBadRequest
			continue
		}
		transport.LoadRelated()
		transportHistories = append(transportHistories, transport)
	}

	utils.WriteBody(w, transportHistories, status)
}
