package main

// @APIVersion 0.1.0
// @APITitle ATES
// @APIDescription API for handling the ATES resources
// @Contact support@coban.jp
// @TermsOfServiceUrl https://www.coban.co.jp/terms-of-service
// @License Copyright Coban
// @LicenseUrl http://coban.co.jp/license

import (
	"log"
	"net/http"
	"os"
)

func main() {
	router := NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment wasn't set.")
	}

	log.Fatal(http.ListenAndServe(":" + port, router))
}
