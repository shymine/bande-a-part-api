package models

import "go.mongodb.org/mongo-driver/bson/primitive"

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
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Town      string             `json:"town" bson:"town"`
	Address_1 string             `json:"address_1" bson:"address_1"`
	Address_2 string             `json:"address_2" bson:"address_2"`
	Phone     string             `json:"phone" bson:"phone"`
	Email     string             `json:"email" bson:"email"`
	About     string             `json:"about" bson:"about"`
}
