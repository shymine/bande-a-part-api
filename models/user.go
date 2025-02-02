package models

/*
User contains the informations about a specific user

Attributes:

	id : str

		The ID associated with the specific user

	email : str

		The contact email of the user

	username: str

		The name used by the user to identify itself

	commands : []Command

		The commands made by the user

	bookmarks : []Book

		The Book marked by the user for whatever they use it for
*/
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Commands  []Command `json:"commands"`
	Bookmarks []Book    `json:"bookmarks"`
}

// TODO: Handle password
