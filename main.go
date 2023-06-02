package main

import (
	"os"

	"github.com/apanji404/Prakerja_Golang/controllers"
	"github.com/apanji404/Prakerja_Golang/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.DBConnection()
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1234"
	}
	return port
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsGet)
	r.GET("/posts/:id", controllers.PostsGetByID)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.Run(":" + getPort())
}
