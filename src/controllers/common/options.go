package common

import (
	"net/http"

    "coban/api/src/databases"
	"coban/api/src/utils"
)

func AllowAccessControl(w http.ResponseWriter, r *http.Request,
    scopeChecker databases.IsScope) (int, error) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    _, status, err := utils.CheckTokenAndScope(r, scopeChecker)
    if err != nil {
        return status, err
    }
    w.WriteHeader(http.StatusNoContent)
    return 0, nil
}

func AllowAccessControlForClients(w http.ResponseWriter, r *http.Request) {
    status, err := AllowAccessControl(w, r, databases.IsClient)
    if err != nil {
        utils.Error(w, err, status)
        return
    }
}

func AllowAccessControlForOffices(w http.ResponseWriter, r *http.Request) {
    status, err := AllowAccessControl(w, r, databases.IsOffice)
    if err != nil {
        utils.Error(w, err, status)
        return
    }
}

func AllowAccessControlForAdmin(w http.ResponseWriter, r *http.Request) {
    status, err := AllowAccessControl(w, r, databases.IsAdmin)
    if err != nil {
        utils.Error(w, err, status)
        return
    }
}
