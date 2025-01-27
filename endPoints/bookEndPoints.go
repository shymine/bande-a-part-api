package endpoints

import (
	"bande-a-part/database"
	"bande-a-part/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllBook(c *gin.Context) {
	books, err := database.GetAllBooks()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error getting the Books " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, books)
}

// Get a Book by ID
func GetBookById(c *gin.Context) {
	id := c.Param("id")

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed ID " + err.Error()})
		return
	}

	book, err := database.GetBookById(objId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

// Get a Book by filter
func GetBookByFilter(c *gin.Context) {
	c.IndentedJSON(http.StatusNotImplemented, nil)
}

// Post a set of Book
func PostBooks(c *gin.Context) {
	var book dto.BookDTO

	if err := c.BindJSON(&book); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	newB, newErr := database.CreateBook(book)
	if newErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error creating the Books " + newErr.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newB)
}

// Put a Book
func PutBook(c *gin.Context) {
	id := c.Param(("id"))

	var update map[string]any
	if err := c.BindJSON(&update); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed ID " + err.Error()})
		return
	}

	err = database.UpdateBook(objId, update)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "pb while updating " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

// Delete a Book
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed ID " + err.Error()})
		return
	}
	err = database.DeleteBook(objId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "pb while deleting " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}
