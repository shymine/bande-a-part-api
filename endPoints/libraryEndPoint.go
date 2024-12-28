package endpoints

import (
	"bande-a-part/database"
	"bande-a-part/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get All Library
func GetAllLibraries(c *gin.Context) {
	libraries := database.Libraries

	c.IndentedJSON(http.StatusOK, libraries)
}

// Post a Library
func PostLibrary(c *gin.Context) {
	var library []models.Library

	if err := c.BindJSON(&library); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	database.Libraries = append(database.Libraries, library...)
	c.IndentedJSON(http.StatusOK, library)
}

// Put a Library
func PutLibrary(c *gin.Context) {
	var incoming models.Library

	if err := c.BindJSON(&incoming); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	for i, a := range database.Libraries {
		if a.ID == incoming.ID {
			database.Libraries[i] = incoming
			break
		}
	}
	c.IndentedJSON(http.StatusOK, incoming)
}

// Delete a Library
func DeleteLibrary(c *gin.Context) {
	id := c.Param("id")

	var index = -1
	var element models.Library

	for i, a := range database.Libraries {
		if a.ID == id {
			index = i
			element = a
			break
		}
	}

	if index == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Library match the id " + id})
		return
	}
	database.Libraries = append(database.Libraries[:index], database.Libraries[index+1:]...)
	c.IndentedJSON(http.StatusOK, element)
}
