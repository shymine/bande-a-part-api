package endpoints

import (
	"bande-a-part/database"
	"bande-a-part/dto"
	"bande-a-part/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get all Command of a User
func GetCommandByUser(c *gin.Context) {
	id := c.Param("userid")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed ID " + err.Error()})
		return
	}

	commands, err := database.GetCommandByUser(objId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error getting the Commands " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, commands)
}

// Get all Command of a certain state
func GetCommandByStatus(c *gin.Context) {
	status := c.Param("status")
	cStatus, err := models.StringToCommandStatus(status)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed status " + err.Error()})
		return
	}

	commands, err := database.GetCommandByStatus(cStatus)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error getting the Commands " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, commands)
}

// Post a Command
func PostCommand(c *gin.Context) {
	var command dto.CommandDTO

	if err := c.BindJSON(&command); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	command.Date = time.Now()
	command.Status = models.TOAPPROUVE

	var bookSum float32 = 0.
	for _, bookId := range command.Books {
		book, err := database.GetBookById(bookId)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error getting books from command " + err.Error()})
			return
		}
		bookSum += book.Price
	}

	command.Total = bookSum

	newC, newErr := database.CreateCommand(command)
	if newErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error creating the Command " + newErr.Error()})
		return
	}

	userId := c.Param("userid")
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed ID " + err.Error()})
		return
	}

	err = database.AddCommand(objId, newC.ID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error adding command to user " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newC)
}

func CommandNextStatus(c *gin.Context) {
	id := c.Param(("id"))

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed ID " + err.Error()})
		return
	}
	command, err := database.GetCommandById(objId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No Command with this ID " + err.Error()})
		return
	}

	cStatus, err := models.NextCommandStatus(command.Status)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error getting next command status " + err.Error()})
		return
	}

	err = database.SetCommandStatus(objId, cStatus)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error with updating command status " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func CommandReject(c *gin.Context) {
	id := c.Param(("id"))

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed ID " + err.Error()})
		return
	}

	err = database.SetCommandStatus(objId, models.REJECTED)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error with updating command status " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

// Delete a Command
func DeleteCommand(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed ID " + err.Error()})
		return
	}
	err = database.DeleteCommand(objId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "pb while deleting " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}
