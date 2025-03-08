package main

import (
	"fmt"
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Quantity: 3},
	{ID: "2", Title: "The Hobbit", Author: "J.R.R. Tolkien", Quantity: 5},
	{ID: "3", Title: "Harry Potter and the Philosopher's Stone", Author: "J.K. Rowling", Quantity: 7},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func checkoutBook(c *gin.Context) {
	id := c.Param("id")

	book, err := getBookByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "no books available"})
		return
	}

	book.Quantity--
	c.IndentedJSON(http.StatusOK, book)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	b, err := getBookByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, b)
}

func returnBook(c *gin.Context) {
	id := c.Param("id")
	b, err := getBookByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	b.Quantity++
	c.IndentedJSON(http.StatusOK, b)

}

func getBookByID(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		// c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("incorrect")
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.GET("/checkout/:id", checkoutBook)
	router.GET("/return/:id", returnBook)
	router.Run("localhost:8080")
}
