package endpoints

import (
	"bande-a-part/database"
	"bande-a-part/models"
	"errors"
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

	res := []models.Editor{}
	for _, ed := range editors {
		newEd, newErr := database.CreateEditor(ed)
		if newErr != nil {
			log.Println(newErr, " for ", ed)
		} else {
			res = append(res, newEd)
		}
	}
	
	c.IndentedJSON(http.StatusCreated, res)
}

// Put and Editor
func PutEditor(c *gin.Context) {
	var incoming models.Editor

	// var editor models.Editor
	var index int
	var err error

	if err := c.BindJSON(&incoming); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	if _, index, err = getEditorById(incoming.ID); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	database.Editors[index] = incoming
	c.IndentedJSON(http.StatusOK, incoming)
}

// Get Editor by their ID
func getEditorById(id primitive.ObjectID) (models.Editor, int, error) {
	for i, a := range database.Editors {
		if a.ID == id {
			return a, i, nil
		}
	}
	return models.Editor{}, 0, errors.New("no Editor found with id: " + id.String())
}


// Delete an Editor
func DeleteEditor(c *gin.Context) {
	id := c.Param("id")

	var index int
	var err error
	var element models.Editor
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	if element, index, err = getEditorById(objId); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	database.Editors = append(database.Editors[:index], database.Editors[index+1:]...)
	c.IndentedJSON(http.StatusOK, element)
}
