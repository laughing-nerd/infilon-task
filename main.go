package main

import (
	"infilon-task/controllers"
	"infilon-task/internal"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {

	// Connect to DB
	internal.InitDb()

	r := gin.Default()

	// Since this task has only 2 endpoints, I am not creating a separate routes directory for this
	r.GET("/person/:id/info", controllers.GetPerson)
	r.POST("/person/create", controllers.CreatePerson)

  port := os.Getenv("PORT")
  if port == "" {
    port = ":8080"
  }
	r.Run(port)
}
