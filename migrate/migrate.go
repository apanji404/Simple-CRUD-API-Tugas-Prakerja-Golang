package main

import (
	"github.com/apanji404/Prakerja_Golang/initializers"
	"github.com/apanji404/Prakerja_Golang/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.DBConnection()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
