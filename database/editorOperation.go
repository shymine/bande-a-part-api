package database

import (
	"bande-a-part/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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

func CreateEditor(editor models.Editor) (models.Editor, error) {
	mEditor, err := bson.Marshal(editor)
	if err != nil {
		return models.Editor{}, err
	}

	_, err = DB_MANAGER.
		GetCollection("editor").
		InsertOne(
			DB_MANAGER.GetContext(),
			mEditor,
	)
	if err != nil {
		return models.Editor{}, err
	}

	var newEditor models.Editor
	err = DB_MANAGER.GetCollection("editor").FindOne(
		DB_MANAGER.GetContext(),
		mEditor,
	).Decode(&newEditor)
	if err != nil {
		return models.Editor{}, err
	}

	log.Println("Create editor", newEditor)

	return newEditor, nil
}