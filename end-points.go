package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// getContributor respond with th elist of all contributor as JSON
func getContributor(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Authors)
}
