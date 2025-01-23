package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*
The database model of a Genre

Attributes:
----------
name : str

	The name of the Genre (sci-fi, horror, etc.)
*/
type Genre struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
}
