package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"coban/api/src/databases"
)

// WriteBody serialise an object in JSON and write it
// into the http.ResponseWriter's body
func WriteBody(w http.ResponseWriter, content interface{},
	statusHTTP int) error {
	data, err := json.Marshal(content)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusHTTP)
	if _, err = fmt.Fprint(w, string(data)); err != nil {
		return err
	}
	return nil
}

// ReadBody unserialise an object fromt the http.Request's body
func ReadBody(r *http.Request, model interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(model)

	return err
}

// Error Write the go's error into the http.ResponseWriter's body
func Error(w http.ResponseWriter, err error, statusHTTP int) {
	w.WriteHeader(statusHTTP)
	fmt.Fprint(w, err)
}

// CheckTokenAndScope Ensure that the current token is:
// 		- signed with the server's private key
//		- isn't expired
//		- has a scope
//		- the scope is valid for the user
//		- the user exist and is valid
func CheckTokenAndScope(r *http.Request,
	scopeChecker databases.IsScope) (databases.User, int, error) {
	var user databases.User

	token, err := ParseTokenFromRequest(r)
	if err != nil {
		return user, http.StatusBadRequest, err
	}
	if !token.Valid {
		return user,
			http.StatusUnauthorized,
			errors.New("This token isn't valid.")
	}
	scope, found := token.Claims["scope"].(float64)
	if !found {
		return user,
			http.StatusBadRequest,
			errors.New("There aren't any scope in the token.")
	}
	if !scopeChecker(byte(scope)) {
		return user,
			http.StatusUnauthorized,
			errors.New("Unauthorised user.")
	}
	id, found := token.Claims["user"].(float64)
	if !found {
		return user,
			http.StatusBadRequest,
			errors.New("There aren't any user in the token.")
	}
	databases.DB.First(&user, uint(id))
	if user.ID == 0 {
		return user,
			http.StatusBadRequest,
			errors.New("User not found.")
	}
	user.LoadRelated()

	return user, 0, nil
}
