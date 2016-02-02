package databases

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var DB gorm.DB

type Database struct {
	User 		string
	Password 	string
	Name 		string
	Host		string
	Port		int
	Migration	string
}

func ParseConfigurationFile() map[string]Database {
	path, err := filepath.Abs("configurations/environments.yml")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	env := map[string]Database{}
	err = yaml.Unmarshal(data, &env)
	if err != nil {
		log.Fatal(err)
	}
	return env
}

func ConnectionString(env Database) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", env.User, env.Password, env.Host, env.Port, env.Name)
}

func configuration() string {
	var arg string

	if len(os.Args) > 1 {
		arg = os.Args[1]
	} else {
		arg = "development"
	}

	environments := ParseConfigurationFile()
	if env, ok := environments[arg]; ok {
		return ConnectionString(env)
	} else {
		log.Fatal("Undefined environment: ", arg)
	}
	return ""
}

func init() {
	var err error

	DB, err = gorm.Open("mysql", configuration())
	if err != nil {
		log.Fatal(err)
	}

	DB.LogMode(false)
}