package database

import (
	"bande-a-part/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetLibrary() ([]models.Library, error) {
	var res []models.Library
	curr, err := DB_MANAGER.GetCollection("library").Find(
		DB_MANAGER.GetContext(),
		bson.M{},
	)
	if err != nil {
		return []models.Library{}, err
	}

	err = curr.All(
		DB_MANAGER.GetContext(),
		&res,
	)
	if err != nil {
		return []models.Library{}, err
	}
	if len(res) == 0 {
		return []models.Library{}, nil
	}
	return res, err
}

func CreateLibrary(library models.Library) (models.Library, error) {
	corrlibrary := library

	mLibrary, err := bson.Marshal(library)
	if err != nil {
		log.Println("error marshalling library: ", err)
	}

	result, err := DB_MANAGER.
		GetCollection("library").
		InsertOne(
			DB_MANAGER.GetContext(),
			mLibrary,
		)
	if err != nil {
		return models.Library{}, err
	}

	corrlibrary.ID = result.InsertedID.(primitive.ObjectID)

	return corrlibrary, nil
}

/*
Modify the library corresponding to the ID
filter is of the shape {<field to modify>: <update>}
*/
func UpdateLibrary(id primitive.ObjectID, update map[string]any) error {
	updates := bson.D{}

	for k, v := range update {
		updates = append(
			updates,
			bson.E{Key: k, Value: v},
		)
	}

	updates = bson.D{{"$set", updates}}

	filter := bson.D{{"_id", id}}

	_, err := DB_MANAGER.GetCollection("library").UpdateOne(
		DB_MANAGER.GetContext(),
		filter,
		updates,
	)
	return err
}

func DeleteLibrary(id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}
	_, err := DB_MANAGER.GetCollection("library").DeleteOne(
		DB_MANAGER.GetContext(),
		filter,
	)
	return err
}
