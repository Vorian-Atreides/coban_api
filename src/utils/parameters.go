package utils

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetUINT64Parameter read and parse the uint
// GET's parameter related to the name's function parameter
func GetUINT64Parameter(r *http.Request, name string) (uint64, error) {
	value, err := strconv.ParseUint(mux.Vars(r)[name], 10, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// GetStringParameter read and parse the string
// GET's parameter related to the name's function parameter
func GetStringParameter(r *http.Request, name string) string {
	return mux.Vars(r)[name]
}
