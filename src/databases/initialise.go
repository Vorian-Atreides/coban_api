package databases

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
)

// DB is an instance of the database used for the queries
var DB gorm.DB

// Database is a structure describing the database environment
type Database struct {
	User      string
	Password  string
	Name      string
	Host      string
	Port      int
	Migration string
}

const filePath = "/etc/configurations/environments.yml"

// ParseConfigurationFile get the database environment from the unix environment
func ParseConfigurationFile() map[string]Database {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Read environment's file: ", err)
	}
	env := map[string]Database{}
	err = yaml.Unmarshal(data, &env)
	if err != nil {
		log.Fatal("Unmarshal environment: ", err)
	}
	return env
}

// ConnectionString build a connection string with the environment configuration
func ConnectionString(env Database) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?&parseTime=True", env.User, env.Password, env.Host, env.Port, env.Name)
}

// GetDBEnv parse the unix environment for finding the DB_ENV
func GetDBEnv() string {
	dbEnv := os.Getenv("DB_ENV")
	if dbEnv == "" {
		log.Fatal("Database environment wasn't set.")
	}

	return dbEnv
}

func configuration() string {

	arg := GetDBEnv()
	environments := ParseConfigurationFile()
	if env, ok := environments[arg]; ok {
		return ConnectionString(env)
	}
	log.Fatal("Undefined environment: ", arg)
	return ""
}

func init() {
	var err error

	DB, err = gorm.Open("mysql", configuration())
	if err != nil {
		log.Fatal("Gorm open: ", err)
	}

	DB.LogMode(false)
}
