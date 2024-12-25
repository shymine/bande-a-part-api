package main

import (
	"bande-a-part/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getContributor respond with the list of all contributor as JSON
func getContributor(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Authors)
}

// getContributorId respond with the contributor associated to the ID
func getContributorId(c *gin.Context) {
	id := c.Param("id")

	for _, a := range Authors {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "contributor not found"})
}

// createContributor add a contributor to th elsit of all contributors
func createContributor(c *gin.Context) {
	var newContrib models.Contributor

	if err := c.BindJSON(&newContrib); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "json badly formated"})
		return
	}

	Authors = append(Authors, newContrib)
	c.IndentedJSON(http.StatusCreated, newContrib)
}
