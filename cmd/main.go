package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []Book{
	{ID: "1", Title: "Golang pointers", Author: "Mr. Golang", Quantity: 2},
	{ID: "2", Title: "Goroutines", Author: "Mr. Goroutine", Quantity: 4},
	{ID: "3", Title: "Golang routers", Author: "Mr. Router", Quantity: 6},
	{ID: "4", Title: "Golang concurrency", Author: "Mr. Currency", Quantity: 8},
	{ID: "5", Title: "Golang good parts", Author: "Mr. Good", Quantity: 10},
}

func main() {
	router := gin.Default()
	router.GET("/books", getAlbums)
	router.POST("/books", postBook)

	router.Run("0.0.0.0:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func postBook(c *gin.Context) {
	var newBook Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
