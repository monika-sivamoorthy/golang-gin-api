package main

import (
	"fmt"
	"net/http" 
	"github.com/gin-gonic/gin"
	// "errors"
) 

type book struct {
	ID string `json:"id"` 
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Hary potter 1", Author: "J.K R 1", Quantity: 2},
	{ID: "2", Title: "Hary potter 2", Author: "J.K R 2", Quantity: 5},
	{ID: "3", Title: "Hary potter 3", Author: "J.K R 3", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book 
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main(){
	fmt.Println("Executing main.go ...")
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.Run("localhost:8080")
}
