package endpoints

import (
	"bande-a-part/database"
	"bande-a-part/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all Genre
func GetAllGenre(c *gin.Context) {
	genres := database.Genres

	c.IndentedJSON(http.StatusOK, genres)
}

// Post a set of Genre
// TODO: check if there are close equivalent of the genre is so, do not add them
func PostGenre(c *gin.Context) {
	var genres []models.Genre

	if err := c.BindJSON(&genres); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	database.Genres = append(database.Genres, genres...)
	c.IndentedJSON(http.StatusCreated, genres)
}

// Put a Genre
// TODO: check if there are close equivalents of the new genre if so, do not modify
// TODO: check if the ID exist
func PutGenre(c *gin.Context) {
	var incoming models.Genre

	if err := c.BindJSON(&incoming); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	for i, a := range database.Genres {
		if a.ID == incoming.ID {
			database.Genres[i] = incoming
			break
		}
	}
	c.IndentedJSON(http.StatusOK, incoming)
}

// Delete a Genre
func DeleteGenre(c *gin.Context) {
	id := c.Param("id")

	var index = -1
	var element models.Genre

	for i, a := range database.Genres {
		if a.ID == id {
			index = i
			element = a
			break
		}
	}

	if index == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Genre match the id " + id})
		return
	}
	database.Genres = append(database.Genres[:index], database.Genres[index+1:]...)
	c.IndentedJSON(http.StatusOK, element)
}
