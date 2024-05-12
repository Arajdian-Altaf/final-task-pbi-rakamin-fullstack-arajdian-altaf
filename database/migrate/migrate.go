package main

import (
	"log"

	"github.com/Arajdian-Altaf/final-task-pbi/database"
	"github.com/Arajdian-Altaf/final-task-pbi/helpers"
	"github.com/Arajdian-Altaf/final-task-pbi/models"
)

func init() {
	helpers.LoadEnv()
}

func main() {
	db, err := database.ConnectToDB()

	if err != nil {
		log.Fatal(err)
	}

	gormDB := db.GetDB()
	migrateErr := gormDB.AutoMigrate(&models.User{}, &models.Photo{})
	if migrateErr != nil {
		log.Fatal(migrateErr)
	}

}