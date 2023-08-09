package main

import (
	"log"

	"github.com/max38/golang-clean-code-architecture/src/config"

	entitymodels "github.com/max38/golang-clean-code-architecture/src/domain/models"
	postgreshandler "github.com/max38/golang-clean-code-architecture/src/infrastructure/database/postgres"
)

func main() {
	config.Load(".env.dev")

	var dbHandler = postgreshandler.PostgresHandler()
	var dbConnector = dbHandler.GetConnector()

	log.Println("Connect success")
	log.Println("Start migration")

	// Migrate the schema
	dbConnector.AutoMigrate(&entitymodels.UserModel{})
	dbConnector.AutoMigrate(&entitymodels.UserTokenModel{})

	dbHandler.Close()

	log.Println("Migration success !!!")
}
