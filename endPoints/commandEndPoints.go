package endpoints

import (
	"bande-a-part/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all Command of a User
func GetCommandByUser(c *gin.Context) {
	userId := c.Param("id")

	for _, a := range Users {
		if a.ID == userId {
			c.IndentedJSON(http.StatusOK, a.Commands)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No User with id " + userId})

}

// Get all Command of a certain state
func GetCommandByState(c *gin.Context) {
	stt, err := models.StringToCommandStatus(c.Param("status"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var result []models.Command
	for _, a := range Commands {
		if stt == a.Status {
			result = append(result, a)
		}
	}

	c.IndentedJSON(http.StatusOK, result)
}

// Post a Command
// TODO: Put the date to now and compute the total from the list of book
// TODO: check the books are valid
// TODO: compose an intermediate version to handle the addition of total and date
func PostCommand(c *gin.Context) {
	var command models.Command

	if err := c.BindJSON(&command); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	Commands = append(Commands, command)
	c.IndentedJSON(http.StatusOK, command)
}

// Delete a Command
func DeleteCommand(c *gin.Context) {
	id := c.Param("id")

	var index = -1
	var element models.Command

	for i, a := range Commands {
		if a.ID == id {
			index = i
			element = a
			break
		}
	}

	if index == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Command match the id " + id})
		return
	}
	Commands = append(Commands[:index], Commands[index+1:]...)
	c.IndentedJSON(http.StatusOK, element)
}
