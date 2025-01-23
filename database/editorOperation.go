package database

import (
	"bande-a-part/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetEditor() ([]models.Editor, error) {
	var res []models.Editor
	curr, err := DB_MANAGER.GetCollection("editor").Find(
		DB_MANAGER.GetContext(),
		bson.M{},
	)
	if err != nil {
		return []models.Editor{}, err
	}

	err = curr.All(
		DB_MANAGER.GetContext(),
		&res,
	)
	if err != nil {
		return []models.Editor{}, err
	}
	return res, err
}

func CreateMultEditor(editors []models.Editor) ([]models.Editor, error) {
	mEditors := []any{}
	corrEditors := []models.Editor{}
	for _, el := range editors {
		ed, err := bson.Marshal(el)
		if err != nil {
			log.Println("error marshalling editor: ", err)
		} else {
			mEditors = append(mEditors, ed)
			corrEditors = append(corrEditors, el)
		}
	}
	result, err := DB_MANAGER.
		GetCollection("editor").
		InsertMany(
			DB_MANAGER.GetContext(),
			mEditors,
		)
	if err != nil {
		return []models.Editor{}, err
	}

	for j, id := range result.InsertedIDs {
		corrEditors[j].ID = id.(primitive.ObjectID)
	}

	return corrEditors, nil
}

/*
Modify the Editor corresponding to the ID
filter is of the shape {<field to modify>: <update>}
*/
func UpdateEditor(id primitive.ObjectID, update map[string]any) error {
	updates := bson.D{}

	for k, v := range update {
		updates = append(
			updates,
			bson.E{Key: k, Value: v},
		)
	}

	updates = bson.D{{"$set", updates}}

	filter := bson.D{{"_id", id}}

	_, err := DB_MANAGER.GetCollection("editor").UpdateOne(
		DB_MANAGER.GetContext(),
		filter,
		updates,
	)
	return err
}

func DeleteEditor(id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}
	_, err := DB_MANAGER.GetCollection("editor").DeleteOne(
		DB_MANAGER.GetContext(),
		filter,
	)
	return err
}
