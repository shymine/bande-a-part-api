package endpoints

import (
	"bande-a-part/database"
	"bande-a-part/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get All BookList
func GetAllBookList(c *gin.Context) {
	booksList, err := database.GetBookList()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error getting the BookLists " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, booksList)
}

// Post a BookList
func PostBookList(c *gin.Context) {
	var bookList dto.BookListDTO

	if err := c.BindJSON(&bookList); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	newB, newErr := database.CreateBookList(bookList)
	if newErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error creating the Books " + newErr.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newB)
}

// Put a BookList
func PutBookList(c *gin.Context) {
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

	err = database.UpdateBookList(objId, update)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "pb while updating " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

// Delete a BookList
func DeleteBookList(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed ID " + err.Error()})
		return
	}
	err = database.DeleteBookList(objId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "pb while deleting " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}
