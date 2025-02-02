package database

import (
	"bande-a-part/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetGenre() ([]models.Genre, error) {
	var res []models.Genre
	curr, err := DB_MANAGER.GetCollection("genre").Find(
		DB_MANAGER.GetContext(),
		bson.M{},
	)
	if err != nil {
		return []models.Genre{}, err
	}

	err = curr.All(
		DB_MANAGER.GetContext(),
		&res,
	)
	if err != nil {
		return []models.Genre{}, err
	}

	if len(res) == 0 {
		return []models.Genre{}, nil
	}
	return res, err
}

func CreateMultGenre(editors []models.Genre) ([]models.Genre, error) {
	mGenre := []any{}
	corrGenre := []models.Genre{}
	for _, el := range editors {
		ed, err := bson.Marshal(el)
		if err != nil {
			log.Println("error marshalling genre: ", err)
		} else {
			mGenre = append(mGenre, ed)
			corrGenre = append(corrGenre, el)
		}
	}
	result, err := DB_MANAGER.
		GetCollection("genre").
		InsertMany(
			DB_MANAGER.GetContext(),
			mGenre,
		)
	if err != nil {
		return []models.Genre{}, err
	}

	for j, id := range result.InsertedIDs {
		corrGenre[j].ID = id.(primitive.ObjectID)
	}

	return corrGenre, nil
}

/*
Modify the Editor corresponding to the ID
filter is of the shape {<field to modify>: <update>}
*/
func UpdateGenre(id primitive.ObjectID, update map[string]any) error {
	updates := bson.D{}

	for k, v := range update {
		updates = append(
			updates,
			bson.E{Key: k, Value: v},
		)
	}

	updates = bson.D{{"$set", updates}}

	filter := bson.D{{"_id", id}}

	_, err := DB_MANAGER.GetCollection("genre").UpdateOne(
		DB_MANAGER.GetContext(),
		filter,
		updates,
	)
	return err
}

func DeleteGenre(id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}
	_, err := DB_MANAGER.GetCollection("genre").DeleteOne(
		DB_MANAGER.GetContext(),
		filter,
	)
	return err
}
