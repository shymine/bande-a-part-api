package endpoints

import (
	"bande-a-part/database"
	"bande-a-part/dto"
	"bande-a-part/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all Command of a User
func GetCommandByUser(c *gin.Context) {
	userId := c.Param("userid")

	user, err := database.FindUserById(userId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	commandsDTO := []dto.CommandDTO{}
	for _, e := range user.Commands {
		commandsDTO = append(commandsDTO, dto.CommandToDTO(e))
	}

	c.IndentedJSON(http.StatusOK, commandsDTO)
}

// Get all Command of a certain state
func GetCommandByStatus(c *gin.Context) {
	stt, err := models.StringToCommandStatus(c.Param("status"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var result = database.FindCommandByStatus(stt)

	commandsDTO := []dto.CommandDTO{}
	for _, e := range result {
		commandsDTO = append(commandsDTO, dto.CommandToDTO(e))
	}

	c.IndentedJSON(http.StatusOK, commandsDTO)
}

// Post a Command
func PostCommand(c *gin.Context) {
	userId := c.Param(("userid"))
	var commandDTO dto.CommandDTOCreated

	if err := c.BindJSON(&commandDTO); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	user, err := database.FindUserById(userId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	command, err := dto.CreatedToCommand(commandDTO)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	user.Commands = append(user.Commands, command)

	database.Commands = append(database.Commands, command)
	c.IndentedJSON(http.StatusOK, command)
}

// Delete a Command
func DeleteCommand(c *gin.Context) {
	id := c.Param("id")

	var index = -1
	var element models.Command

	for i, a := range database.Commands {
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
	database.Commands = append(database.Commands[:index], database.Commands[index+1:]...)

	for j, user := range database.Users {
		for i, a := range user.Commands {
			if a.ID == element.ID {
				database.Users[j].Commands = append(database.Users[j].Commands[:i], database.Users[j].Commands[i+1:]...)
				break
			}
		}
	}
	c.IndentedJSON(http.StatusOK, dto.CommandToDTO(element))
}
