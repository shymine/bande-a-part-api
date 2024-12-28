package endpoints

import (
	"bande-a-part/database"
	"bande-a-part/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all contributors
func GetAllContributors(c *gin.Context) {
	contributors := database.Contributors

	c.IndentedJSON(http.StatusOK, contributors)
}

// Get a contributor by ID
func GetContributorById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range database.Contributors {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Contributor with id " + id})
}

// Post a contributor
func PostContributor(c *gin.Context) {
	var contributor models.Contributor

	if err := c.BindJSON(&contributor); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	database.Contributors = append(database.Contributors, contributor)
	c.IndentedJSON(http.StatusOK, contributor)
}

// Put a contributor
func PutContributor(c *gin.Context) {
	var incoming models.Contributor

	if err := c.BindJSON(&incoming); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}
	for i, a := range database.Contributors {
		if a.ID == incoming.ID {
			database.Contributors[i] = incoming
			break
		}
	}
	c.IndentedJSON(http.StatusOK, incoming)
}

// Delete a contributor
func DeleteContributor(c *gin.Context) {
	id := c.Param("id")

	var index = -1
	var element models.Contributor

	for i, a := range database.Contributors {
		if a.ID == id {
			index = i
			element = a
			break
		}
	}

	if index == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Contributo match the id " + id})
		return
	}
	database.Contributors = append(database.Contributors[:index], database.Contributors[index+1:]...)
	c.IndentedJSON(http.StatusOK, element)
}
