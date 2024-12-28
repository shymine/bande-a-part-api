package endpoints

import (
	"bande-a-part/dto"

	"bande-a-part/database"
	"bande-a-part/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBook(c *gin.Context) {
	books := database.Books

	booksDTO := []dto.BookDTO{}
	for _, c := range books {
		book := dto.BookToDTO(c)

		booksDTO = append(booksDTO, book)
	}

	c.IndentedJSON(http.StatusOK, booksDTO)
}

// Get a Book by ID
func GetBookById(c *gin.Context) {
	id := c.Param("id")

	book, err := database.FindBookById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	bookDTO := dto.BookToDTO(book)
	c.IndentedJSON(http.StatusOK, bookDTO)
}

// Get a Book by filter
func GetBookByFilter(c *gin.Context) {
	c.IndentedJSON(http.StatusNotImplemented, nil)
}

// Post a set of Book
func PostBooks(c *gin.Context) {
	var booksDTO []dto.BookDTO

	if err := c.BindJSON(&booksDTO); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	books := []models.Book{}
	for _, b := range booksDTO {
		book, err := dto.DTOToBook(b)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		books = append(books, book)
	}

	database.Books = append(database.Books, books...)
	c.IndentedJSON(http.StatusOK, booksDTO)
}

// Put a Book
func PutBook(c *gin.Context) {
	var incoming dto.BookDTO

	if err := c.BindJSON(&incoming); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	book, err := dto.DTOToBook(incoming)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	for i, a := range database.Books {
		if a.ID == incoming.ID {
			database.Books[i] = book
			break
		}
	}
	c.IndentedJSON(http.StatusOK, incoming)
}

// Delete a Book
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	var index = -1
	var element models.Book

	for i, a := range database.Books {
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

	database.Books = append(database.Books[:index], database.Books[index+1:]...)

	elementDTO := dto.BookToDTO(element)
	c.IndentedJSON(http.StatusOK, elementDTO)
}
