package endpoints

import (
	"bande-a-part/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all user
func GetAllUser(c *gin.Context) {
	users := Users

	c.IndentedJSON(http.StatusOK, users)
}

// Get User by ID
func GetUserById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range Users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No User with id " + id})
}

// Post a User
func PostUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	Users = append(Users, user)
	c.IndentedJSON(http.StatusOK, user)
}

// Put a User
func PutUser(c *gin.Context) {
	var incoming models.User

	if err := c.BindJSON(&incoming); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	for i, a := range Users {
		if a.ID == incoming.ID {
			Users[i] = incoming
			break
		}
	}
	c.IndentedJSON(http.StatusOK, incoming)
}

// Delete a User
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var index = -1
	var element models.User

	for i, a := range Users {
		if a.ID == id {
			index = i
			element = a
			break
		}
	}

	if index == -1 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No User match the id " + id})
		return
	}
	Users = append(Users[:index], Users[index+1:]...)
	c.IndentedJSON(http.StatusOK, element)
}
