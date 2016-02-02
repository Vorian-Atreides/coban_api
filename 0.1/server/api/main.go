package main

import (
	"log"
	"net/http"
	"os"
	"coban/0.1/server/api/utils"
	"coban/0.1/server/api/controllers/common"
)

func main() {
	router := NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable was not set")
	}

	str, err := utils.GenerateToken(1, 1)
	token, err := utils.ParseToken(str)
	scope, found := token.Claims["scope"].(int)
	println(scope)
	println(common.IsOffice(3))
	found = found
	err = err

	log.Fatal(http.ListenAndServe(":" + port, router))
}
