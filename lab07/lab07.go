package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"strings"
)
var nextID = 1
type Book struct {
	// TODO: Finish struct
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Pages int    `json:"pages"`
}
var bookshelf = []Book{
	// TODO: Init bookshelf
	{ID: 1, Name: "Blue Bird", Pages: 500},
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, bookshelf)
}
func getBook(c *gin.Context) {
	id := c.Param("id")

	for _, book := range bookshelf {
		if strconv.Itoa(book.ID) == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})	
}
func addBook(c *gin.Context) {
	var newBook Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if isDuplicateName(newBook.Name) {
		// Return 409 Conflict with an error message
		c.JSON(http.StatusConflict, gin.H{"message": "duplicate book name"})
		return
	}
	newBook.ID = getNextID()
	bookshelf = append(bookshelf, newBook)

	c.JSON(http.StatusCreated, newBook)
}
func isDuplicateName(newName string) bool {
	for _, book := range bookshelf {
		if strings.EqualFold(book.Name, newName) {
			return true
		}
	}
	return false
}
func deleteBook(c *gin.Context) {
	id := c.Param("id")

	for i, book := range bookshelf {
		if strconv.Itoa(book.ID) == id {
			bookshelf = append(bookshelf[:i], bookshelf[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.Status(http.StatusNoContent)
}
func updateBook(c *gin.Context) {
	id := c.Param("id")

	var updatedBook Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	num_0 ,_ := strconv.Atoi(id)
	if num_0 > nextID{
		c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	for _, book := range bookshelf {
		if strings.EqualFold(book.Name, updatedBook.Name) && strconv.Itoa(book.ID) != id{
			c.JSON(http.StatusConflict, gin.H{"message": "duplicate book name"})
			return 
		}
	}
	
	for i, book := range bookshelf {
		if strconv.Itoa(book.ID) == id {
			num ,err := strconv.Atoi(id)
			if err != nil {
				return
			}

			updatedBook.ID = num
		
			bookshelf[i] = updatedBook
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
func getNextID() int {
	// if len(bookshelf) == 0 {
	// 	return 1
	// }
	// return bookshelf[len(bookshelf)-1].ID + 1
	nextID +=1
	return nextID
}
func main() {
	r := gin.Default()
	r.RedirectFixedPath = true

	// TODO: Add routes
	r.GET("/bookshelf", getBooks)
	r.GET("/bookshelf/:id", getBook)
	r.POST("/bookshelf", addBook)
	r.DELETE("/bookshelf/:id", deleteBook)
	r.PUT("/bookshelf/:id", updateBook)
	err := r.Run(":8087")
	if err != nil {
		return
	}
}
