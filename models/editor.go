package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
The database model of an Editor

Attributes:
----------
name : str

	The name of the editor
*/
type Editor struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
}
