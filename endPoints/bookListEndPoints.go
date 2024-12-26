package endpoints

import (
	"bande-a-part/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get All BookList
func GetAllBookList(c *gin.Context) {
	bookLists := BookLists

	c.IndentedJSON(http.StatusOK, bookLists)
}

// Post a BookList
// TODO: the Books are coming as the id, create an intermediate representation
func PostBookList(c *gin.Context) {
	var bookList models.BookList

	if err := c.BindJSON(&bookList); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	BookLists = append(BookLists, bookList)
	c.IndentedJSON(http.StatusOK, bookList)
}

// Put a BookList
func PutBookList(c *gin.Context) {
	var incoming models.BookList

	if err := c.BindJSON(&incoming); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	for i, a := range BookLists {
		if a.ID == incoming.ID {
			BookLists[i] = incoming
			break
		}
	}
	c.IndentedJSON(http.StatusOK, incoming)
	// TODO: check that the Books are valid
}

// Delete a BookList
func DeleteBookList(c *gin.Context) {
	id := c.Param("id")

	var index = -1
	var element models.BookList

	for i, a := range BookLists {
		if a.ID == id {
			index = i
			element = a
			break
		}
	}

	if index == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No BookList match the id " + id})
		return
	}
	BookLists = append(BookLists[:index], BookLists[index+1:]...)
	c.IndentedJSON(http.StatusOK, element)
}
