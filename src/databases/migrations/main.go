package main

import (
	"database/sql"
	"log"

	"github.com/DavidHuie/gomigrate"
	_ "github.com/go-sql-driver/mysql"

	"coban/api/src/databases"
)

const path = "/go/src/coban/migrations"

func main() {

	dbEnv := databases.GetDBEnv()
	env := databases.ParseConfigurationFile()
	db, err := sql.Open("mysql", databases.ConnectionString(env[dbEnv]))
	if err != nil {
		log.Fatal(err)
	}
	migrator, err := gomigrate.NewMigrator(db, gomigrate.Mysql{}, path+env[dbEnv].Migration)
	if err != nil {
		log.Fatal(err)
	}
	if dbEnv == "test_unit" {
		if err := migrator.RollbackAll(); err != nil {
			log.Fatal(err)
		}
	}

	migrator.Migrate()
}

//
//Migration files need to follow a standard format and must be present in the same directory.
//Given "up" and "down" steps for a migration, create a file for each by following this template:
//
//{{ id }}_{{ name }}_{{ "up" or "down" }}.sql
//For a given migration, the id and name fields must be the same.
//The id field is an integer that corresponds to the order in which the migration should run relative to the other migrations.
//
//id should not be 0 as that value is used for internal validations.
//
