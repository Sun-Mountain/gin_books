package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Sun-Mountain/gin_books/controllers"
	"github.com/Sun-Mountain/gin_books/models"
)

func main() {
  r := gin.Default()

  models.ConnectDatabase()

  r.GET("/books", controllers.FindBooks)

  r.Run()
}

