package utils

import (
	"encoding/json"
	"errors"
	"fmt"
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
	id, found := token.Claims["user"].(float64)
	if !found {
		return user, errors.New("There aren't any user in the token.")
	}
	databases.DB.First(&user, uint(id))
	if user.ID == 0 {
		return user, errors.New("User not found.")
	}
	user.LoadRelated()

	return user, nil
}