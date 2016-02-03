package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"coban/api/src/databases"
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

func Error(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, err)
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

func CheckTokenAndScope(r *http.Request, scopeChecker databases.IsScope) (databases.User, error) {
	var user databases.User

	token, err := ParseTokenFromRequest(r)
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