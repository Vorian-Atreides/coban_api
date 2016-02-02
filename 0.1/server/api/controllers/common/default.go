package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"

	"coban/api/0.1/server/api/databases"
	"coban/api/0.1/server/api/utils"
)

func WriteBody(w http.ResponseWriter, content interface {}) error {
	data, err := json.Marshal(content)
	if err != nil {
		return err
	}

	if _, err = fmt.Fprint(w, string(data)); err != nil {
		return err
	}
	return nil
}

func UpdateOrInsertInitialisation(w http.ResponseWriter, r *http.Request, model databases.Model) error {
	if err := model.FromBody(r); err != nil {
		return err
	}
	model.LoadRelated()

	log.Println(model)
	if err := model.IsValid(true); err != nil {
		log.Println("Test")

		return err
	}
	return nil
}

func CheckTokenAndScope(r *http.Request, scopeChecker IsScope) (databases.User, error) {
	var user databases.User

	token, err := utils.ParseTokenFromRequest(r)
	if err != nil {
		return user, err
	}
	if !token.Valid {
		return user, errors.New("This token isn't valid.")
	}
	scope, found := token.Claims["scope"].(float64)
	if !found {
		return user, errors.New("There aren't any scope in the token.")
	}
	if !scopeChecker(byte(scope)) {
		return user, errors.New("Unauthorised user.")
	}
	user.LoadRelated()

	return user, nil
}

func Error(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, err)
}

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