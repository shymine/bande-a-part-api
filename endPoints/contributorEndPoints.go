package endpoints

import (
	"bande-a-part/database"
	"bande-a-part/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get all contributors
func GetAllContributors(c *gin.Context) {
	log.Println("new get all editors")
	contributors, err := database.GetContributor()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error getting the Contributors " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, contributors)
}

// Post a contributor
func PostContributor(c *gin.Context) {
	var contributors []models.Contributor

	if err := c.BindJSON(&contributors); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	newEd, newErr := database.CreateMultContributor(contributors)
	if newErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error creating the Contributors " + newErr.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newEd)
}

// Put a contributor
func PutContributor(c *gin.Context) {
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

	err = database.UpdateContributor(objId, update)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "pb while updating " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

// Delete a contributor
func DeleteContributor(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed ID " + err.Error()})
		return
	}
	err = database.DeleteContributor(objId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "pb while deleting " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}
