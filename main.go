package main

import (
	"log"

	"github.com/Arajdian-Altaf/final-task-pbi/database"
	"github.com/Arajdian-Altaf/final-task-pbi/helpers"
	"github.com/Arajdian-Altaf/final-task-pbi/middlewares"
	"github.com/Arajdian-Altaf/final-task-pbi/router"
	"github.com/gin-gonic/gin"
)

func init() {
	helpers.LoadEnv()
}

func main() {
	r := gin.Default()

	db, err := database.ConnectToDB()

	if err != nil {
		log.Fatal(err)
		return
	}

	r.Use(middlewares.DBMiddleware(db.GetDB()))
	router.PhotoRoutes(r)
	router.UserRoutes(r)

	r.Run(":8080")
}
