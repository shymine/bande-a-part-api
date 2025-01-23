package endpoints

import (
	"bande-a-part/database"
	"bande-a-part/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get all Genre
func GetAllGenre(c *gin.Context) {
	genres, err := database.GetGenre()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error getting the Genre " + err.Error()})
		return
	}

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

	newG, newErr := database.CreateMultGenre(genres)
	if newErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error creating the Editors " + newErr.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newG)
}

// Put a Genre
// TODO: check if there are close equivalents of the new genre if so, do not modify
// TODO: check if the ID exist
func PutGenre(c *gin.Context) {
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

	err = database.UpdateGenre(objId, update)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "pb while updating " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

// Delete a Genre
func DeleteGenre(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed ID " + err.Error()})
		return
	}
	err = database.DeleteGenre(objId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "pb while deleting " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}
