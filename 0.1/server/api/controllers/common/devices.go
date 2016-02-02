package common

import (
	"coban/0.1/server/api/databases"
	"net/http"
)

func GetDevices(w http.ResponseWriter, r *http.Request) {
	var devices []databases.Device

	databases.DB.Find(&devices)
	for i, _ := range devices {
		devices[i].LoadRelated()
	}

	WriteBody(w, devices)
}

func CreateDevice() {

}