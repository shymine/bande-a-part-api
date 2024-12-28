package models

/*
The database model of a Library

# Represent the place where library is and contact infos

Attributes:
----------
name      : str

	The name of the library

town      : str

	The town where it is located

address_1 : str

	The first address line

address_2 : str

	The second address line

phone     : str

	The phone number of the library (can be dot separated)

email     : str

	The email address to contact the library

about     : str

	A short description of the library and tell about the people and its purpose / atmosphere
*/
type Library struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Town      string `json:"town"`
	Address_1 string `json:"address_1"`
	Address_2 string `json:"address_2"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	About     string `json:"about"`
}
