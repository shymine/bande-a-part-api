package endpoints

import (
	"bande-a-part/database"
	"bande-a-part/dto"
	"bande-a-part/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all user
func GetAllUser(c *gin.Context) {
	users := database.Users

	usersDTO := []dto.UserDTO{}
	for _, u := range users {
		usersDTO = append(usersDTO, dto.UserToDTO(u))
	}
	c.IndentedJSON(http.StatusOK, usersDTO)
}

// Get User by ID
func GetUserById(c *gin.Context) {
	id := c.Param("id")

	user, err := database.FindUserById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, dto.UserToDTO(user))
}

// Post a User
func PostUser(c *gin.Context) {
	var userDTO dto.UserDTO

	if err := c.BindJSON(&userDTO); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	user, err := dto.DTOToUser(userDTO)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	database.Users = append(database.Users, user)
	c.IndentedJSON(http.StatusOK, userDTO)
}

// Put a User
func PutUser(c *gin.Context) {
	var incoming dto.UserDTO

	if err := c.BindJSON(&incoming); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "badly formed JSON " + err.Error()})
		return
	}

	user, err := dto.DTOToUser(incoming)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	for i, a := range database.Users {
		if a.ID == incoming.ID {
			database.Users[i] = user
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

	for i, a := range database.Users {
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
	database.Users = append(database.Users[:index], database.Users[index+1:]...)
	c.IndentedJSON(http.StatusOK, dto.UserToDTO(element))
}
