package utils

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// PageSize is the size used for the pagination of the long list
const PageSize = 50

// GetUINT64Parameter read and parse the uint
// GET's parameter related to the name's function parameter
func GetUINT64Parameter(r *http.Request, name string) (uint, error) {
	str, ok := mux.Vars(r)[name]
	if !ok {
		return 0, errors.New("Parameter not found")
	}
	value, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(value), nil
}

// GetStringParameter read and parse the string
// GET's parameter related to the name's function parameter
func GetStringParameter(r *http.Request, name string) (string, error) {
	value, ok := mux.Vars(r)[name]
	if !ok {
		return "", errors.New("Parameter not found")
	}
	return value, nil
}

// GetDateParameter read and parse the string as a GMT date
func GetDateParameter(r *http.Request, name string) (time.Time, error) {
	value := r.FormValue(name)
	log.Println(value)
	return time.Parse("2006-01-02", value)
}

// GetPageOffset get the optional page argument and retrieve its offset
func GetPageOffset(r *http.Request) (int, error) {
	str := r.FormValue("page")
	value, err := strconv.ParseInt(str, 10, 64)
	if err != nil || value <= 0 {
		return 0, err
	}
	return int((value - 1) * PageSize), err
}
