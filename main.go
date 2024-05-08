package main

import (
	"github.com/Arajdian-Altaf/final-task-pbi/database"
	"github.com/Arajdian-Altaf/final-task-pbi/helpers"
	"github.com/Arajdian-Altaf/final-task-pbi/router"
	"github.com/gin-gonic/gin"
)

func init() {
	helpers.LoadEnv()
	database.ConnectToDB()
}

func main() {
	r := gin.Default()

	router.PhotoRoutes(r)
	router.UserRoutes(r)

	r.Run(":8080")
}
