package main

import (
	"example/Movie-Api/movies"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/books", movies.GetBooks)
	router.GET("/books/:id", movies.BookById)
	router.POST("/books", movies.CreatBooks)
	router.PATCH("/checkout", movies.CheckOutBook)
	router.Run("localhost:8080")
}
