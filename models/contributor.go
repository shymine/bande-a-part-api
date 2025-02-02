package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Contributor struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name"`
	Surname string             `json:"surname" bson:"surname"`
	ISNI    string             `json:"isni" bson:"isni"`
}
