package main

import (
	endpoints "bande-a-part/endPoints"

	"github.com/gin-gonic/gin"
)

func setBooksEP(router *gin.Engine) {
	// Get All Books (for debug purpose)
	router.GET("/bookAll", endpoints.GetAllBook)
	// Get Book by ID
	router.GET("/book/:id", endpoints.GetBookById)
	// Get Book by filter
	router.GET("/book", endpoints.GetBookByFilter)
	// Post a set of Book (must be admin) (isbn is obtained as a string of size 13)
	router.POST("/book", endpoints.PostBooks)
	// Put a Book (must be admin)
	router.PUT("/book/:id", endpoints.PutBook)
	// Delete a Book (must be admin)
	router.DELETE("/book/:id", endpoints.DeleteBook)
}

func setBookListEP(router *gin.Engine) {
	// Get All BookList
	router.GET("/bookList", endpoints.GetAllBookList)
	// Post a BookList (must be admin)
	router.POST("/bookList", endpoints.PostBookList)
	// Put a BookList (must be admin)
	router.PUT("/bookList/:id", endpoints.PutBookList)
	// Delete a BookList (must be admin)
	router.DELETE("/bookList/:id", endpoints.DeleteBookList)
}

func setCommandEP(router *gin.Engine) {
	// Get All Command of a User (either is admin or ask for a temporary token associated with the user)
	router.GET("/commandById/:userid", endpoints.GetCommandByUser)
	// Get All Command of a certain state (is admin)
	router.GET("/commandByStatus/:status", endpoints.GetCommandByStatus)
	// Post a Command (Put the date to now and compute the total from the list of book) (raise a message for the admin)
	router.POST("/command/:userid", endpoints.PostCommand)
	// Delete a Command (either is admin, or is the user whom own the command and the command has not yet been approuved)
	router.DELETE("/command/:id", endpoints.DeleteCommand)
}

func setContributorEP(router *gin.Engine) {
	// Get All Contributor (authors, illustrators, translators)
	router.GET("/contributor", endpoints.GetAllContributors)
	// Post a set of Contributor (must be admin)
	router.POST("/contributor", endpoints.PostContributor)
	// Put a Contributor (must be admin)
	router.PUT("/contributor/:id", endpoints.PutContributor)
	// Delete a Contributor (must be admin)
	router.DELETE("/contributor/:id", endpoints.DeleteContributor)
}

func setEditorEP(router *gin.Engine) {
	// Get All Editor
	router.GET("/editor", endpoints.GetAllEditors)
	// Post a set of Editor (must be admin)
	router.POST("/editor", endpoints.PostEditors)
	// Put an Editor (must be admin)
	router.PUT("/editor/:id", endpoints.PutEditor)
	// Delete an Editor (must be admin)
	router.DELETE("/editor/:id", endpoints.DeleteEditor)
}

func setGenreEP(router *gin.Engine) {
	// Get All Genre
	router.GET("/genre", endpoints.GetAllGenre)
	// Post a set of Genre (must be admin) (cannot post a genre already pressent in the database)
	router.POST("/genre", endpoints.PostGenre)
	// Put a Genre (must be admin) (cannot put a genre already existing in the database)
	router.PUT("/genre/:id", endpoints.PutGenre)
	// Delete a Genre (must be admin)
	router.DELETE("/genre/:id", endpoints.DeleteGenre)
}

func setLibraryEP(router *gin.Engine) {
	// Get All Library (should be unique)
	router.GET("/library", endpoints.GetAllLibraries)
	// Post a Library (must be admin) (check if address, email, phone exists or null) or if account doesn t already exist vias email
	router.POST("/library", endpoints.PostLibrary)
	// Put a Library (must be admin) (check if address, email, phone exists or null)
	router.PUT("/library/:id", endpoints.PutLibrary)
	// Delete a Library (must be admin)
	router.DELETE("/library/:id", endpoints.DeleteLibrary)
}

func setUserEP(router *gin.Engine) {
	// Get All User (only if admin)
	router.GET("/user", endpoints.GetAllUser)
	// Get User by ID (must be identificated, will give a temporary token)
	router.GET("/user/:id", endpoints.GetUserById)
	// Post a User (check if email exists and username not takken)
	router.POST("/user", endpoints.PostUser)
	// Put a User (must be admin, or said user)
	router.PUT("/user/:id", endpoints.PutUser)
	// Delete a User (either is admin or is the user in question, also delete the commands associated with the user and any informations about them)
	router.DELETE("/user/:id", endpoints.DeleteUser)
}

func SetEndPoints(router *gin.Engine) {
	setBooksEP(router)
	setBookListEP(router)
	setCommandEP(router)
	setContributorEP(router)
	setEditorEP(router)
	setGenreEP(router)
	setLibraryEP(router)
	setUserEP(router)
}
