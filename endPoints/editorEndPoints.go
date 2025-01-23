package endpoints

import (
	"bande-a-part/database"
	"bande-a-part/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get All Editor
func GetAllEditors(c *gin.Context) {
	log.Println("new get all editors")
	editors, err := database.GetEditor()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error getting the Editors " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, editors)
}

// Post a set of Editor
func PostEditors(c *gin.Context) {
	var editors []models.Editor

	if err := c.BindJSON(&editors); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	newEd, newErr := database.CreateMultEditor(editors)
	if newErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error creating the Editors " + newErr.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newEd)
}

/*
Put and Editor
Modify the Editor corresponding to the ID
filter is of the shape {<field to modify>: <update>}
*/
func PutEditor(c *gin.Context) {
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

	err = database.UpdateEditor(objId, update)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "pb while updating " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

// Delete an Editor
func DeleteEditor(c *gin.Context) {
	id := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed ID " + err.Error()})
		return
	}
	err = database.DeleteEditor(objId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "pb while deleting " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}
