package endpoints

import (
	"bande-a-part/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get All Editor
func GetAllEditors(c *gin.Context) {
	editors := Editors

	c.IndentedJSON(http.StatusOK, editors)
}

// Post a set of Editor
func PostEditors(c *gin.Context) {
	var editors []models.Editor

	if err := c.BindJSON(&editors); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	Editors = append(Editors, editors...)
	c.IndentedJSON(http.StatusCreated, editors)
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

	Editors[index] = incoming
	c.IndentedJSON(http.StatusOK, incoming)
}

func getEditorById(id string) (models.Editor, int, error) {
	for i, a := range Editors {
		if a.ID == id {
			return a, i, nil
		}
	}
	return models.Editor{}, 0, errors.New("no Editor found with id: " + id)
}

// Delete an Editor
func DeleteEditor(c *gin.Context) {
	id := c.Param("id")

	var index int
	var err error
	var element models.Editor

	if element, index, err = getEditorById(id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	Editors = append(Editors[:index], Editors[index+1:]...)
	c.IndentedJSON(http.StatusOK, element)
}
