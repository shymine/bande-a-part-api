package endpoints

import (
	"bande-a-part/database"
	"bande-a-part/dto"
	"bande-a-part/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get All BookList
func GetAllBookList(c *gin.Context) {
	bookLists := database.BookLists

	bookListDTO := []dto.BookListDTO{}
	for _, bl := range bookLists {
		bookListDTO = append(bookListDTO, dto.BookListToDTO(bl))
	}

	c.IndentedJSON(http.StatusOK, bookListDTO)
}

// Post a BookList
func PostBookList(c *gin.Context) {
	var bookListDTO dto.BookListDTO

	if err := c.BindJSON(&bookListDTO); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	bookList, err := dto.DTOToBookList(bookListDTO)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	database.BookLists = append(database.BookLists, bookList)
	c.IndentedJSON(http.StatusOK, bookList)
}

// Put a BookList
func PutBookList(c *gin.Context) {
	var incoming dto.BookListDTO

	if err := c.BindJSON(&incoming); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	bookList, err := dto.DTOToBookList(incoming)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	for i, a := range database.BookLists {
		if a.ID == incoming.ID {
			database.BookLists[i] = bookList
			break
		}
	}
	c.IndentedJSON(http.StatusOK, incoming)
}

// Delete a BookList
func DeleteBookList(c *gin.Context) {
	id := c.Param("id")

	var index = -1
	var element models.BookList

	for i, a := range database.BookLists {
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
	database.BookLists = append(database.BookLists[:index], database.BookLists[index+1:]...)
	c.IndentedJSON(http.StatusOK, element)
}
