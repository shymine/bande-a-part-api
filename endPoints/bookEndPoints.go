package endpoints

import (
	"bande-a-part/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBook(c *gin.Context) {
	books := Books

	c.IndentedJSON(http.StatusOK, books)
}

// Get a Book by ID
func GetBookById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range Books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Book with id " + id})
}

// Get a Book by filter
func GetBookByFilter(c *gin.Context) {
	c.IndentedJSON(http.StatusNotImplemented, nil)
}

// Post a set of Book
// TODO: The Contributor, Editor and Genre are coming as their id, create an intermediate representations
func PostBooks(c *gin.Context) {
	var books []models.Book

	if err := c.BindJSON(&books); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	Books = append(Books, books...)
	c.IndentedJSON(http.StatusOK, books)
}

// Put a Book
func PutBook(c *gin.Context) {
	var incoming models.Book

	if err := c.BindJSON(&incoming); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	for i, a := range Books {
		if a.ID == incoming.ID {
			Books[i] = incoming
			break
		}
	}
	c.IndentedJSON(http.StatusOK, incoming)
	// TODO: make a check that all Contributor exists
	// TODO: make a check that Editor exists
	// TODO: make a check that all Genre exists
}

// Delete a Book
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	var index = -1
	var element models.Book

	for i, a := range Books {
		if a.ID == id {
			index = i
			element = a
			break
		}
	}

	if index == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Book match the id " + id})
		return
	}
	Books = append(Books[:index], Books[index+1:]...)
	c.IndentedJSON(http.StatusOK, element)
}
