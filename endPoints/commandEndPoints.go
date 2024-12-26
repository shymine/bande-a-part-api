package endpoints

import (
	"bande-a-part/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all Command of a User
func GetCommandByUser(c *gin.Context) {
	userId := c.Param("userid")

	for _, a := range Users {
		if a.ID == userId {
			c.IndentedJSON(http.StatusOK, a.Commands)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No User with id " + userId})

}

// Get all Command of a certain state
func GetCommandByStatus(c *gin.Context) {
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
// TODO: the books of the command use the IDs
// TODO: the status is not communicated by the client, it is set here as TOAPPROUVE
func PostCommand(c *gin.Context) {
	userId := c.Param(("userid"))
	var command models.Command

	if err := c.BindJSON(&command); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	var user models.User
	for _, a := range Users {
		if a.ID == userId {
			user = a
			break
		}
	}

	user.Commands = append(user.Commands, command)

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

	for j, user := range Users {
		for i, a := range user.Commands {
			if a.ID == element.ID {
				Users[j].Commands = append(Users[j].Commands[:i], Users[j].Commands[i+1:]...)
				break
			}
		}
	}
	c.IndentedJSON(http.StatusOK, element)
}
