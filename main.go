package main

import "github.com/gin-gonic/gin"

func setEndPoints(router *gin.Engine) {
	router.GET("/contributor", getContributor)
	router.POST("/contributor", createContributor)
	router.GET("/contributor/:id", getContributorId)

	// Get All BookList
	// Get All Editor
	// Get All Contributors (authors, illustrators, translators)
	// Get All Genre
	// Get All Users (only if admin)
	// Get All Commands of a User (either is admin or ask for a temporary token associated with the user)
	// Get All Library (should be unique)

	// Get Book by ID
	// Get Book by filter
	// Get Contributor by ID

	// Post a BookList (must be admin)
	// Post a set of Books (must be admin) (isbn is obtained as a string of size 13)
	// Post a set of Editor (must be admin)
	// Post a set of Genre (must be admin) (cannot post a genre already pressent in the database)
	// Post a Command (Put the date to now and compute the total from the list of book) (raise a message for the admin)
	// Post a User (check if email exists)
	// Post a Library (must be admin) (check if address, email, phone exists or null)

	// Put a BookList (must be admin)
	// Put a Book (must be admin)
	// Put a User (must be admin)
	// Put a Genre (must be admin) (cannot put a genre already existing in the database)
	// Put a Contributor (must be admin)
	// Put an Editor (must be admin)
	// Put a Library (must be admin) (check if address, email, phone exists or null)

	// Delete a BookList (must be admin)
	// Delete a Genre (must be admin)
	// Delete a User (either is admin or is the user in question, also delete the commands associated with the user and any informations about them)
	// Delete a Command (either is admin, or is the user whom own th ecommand and the command has not yet been approuved)
	// Delete an Editor (must be admin)
	// Delete a Contributor (must be admin)
	// Delete a Library (must be admin)
}

func main() {
	router := gin.Default()
	setEndPoints(router)
	router.Run("localhost:8080")
}
