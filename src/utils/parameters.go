package utils

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetUINT64Parameter(w http.ResponseWriter, r *http.Request, name string) (uint64, error) {
	value, err := strconv.ParseUint(mux.Vars(r)[name], 10, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func GetStringParameter(r *http.Request, name string) string {
	return mux.Vars(r)[name]
}