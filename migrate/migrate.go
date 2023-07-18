package main

import (
	"github.com/19abhishek/todo/initializers"
	"github.com/19abhishek/todo/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}