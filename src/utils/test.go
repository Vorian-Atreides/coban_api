package utils

import (
	"os"
	"log"
	"fmt"
)

var IP string
var Port string
var Address string

func InitTest() {
	IP = os.Getenv("API_IP")
	Port = os.Getenv("API_PORT")
	Address = fmt.Sprintf("http://%s:%s", IP, Port)

	if IP == "" || Port == "" {
		log.Fatal("The api's ip and port aren't set.")
	}
}