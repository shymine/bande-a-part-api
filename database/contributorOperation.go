package database

import (
	"bande-a-part/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetContributor() ([]models.Contributor, error) {
	var res []models.Contributor
	curr, err := DB_MANAGER.GetCollection("contributor").Find(
		DB_MANAGER.GetContext(),
		bson.M{},
	)
	if err != nil {
		return []models.Contributor{}, err
	}

	err = curr.All(
		DB_MANAGER.GetContext(),
		&res,
	)
	if err != nil {
		return []models.Contributor{}, err
	}
	return res, err
}

func CreateMultContributor(contributors []models.Contributor) ([]models.Contributor, error) {
	mContributors := []any{}
	corrContributors := []models.Contributor{}
	for _, el := range contributors {
		ed, err := bson.Marshal(el)
		if err != nil {
			log.Println("error marshalling contributor: ", err)
		} else {
			mContributors = append(mContributors, ed)
			corrContributors = append(corrContributors, el)
		}
	}
	result, err := DB_MANAGER.
		GetCollection("contributor").
		InsertMany(
			DB_MANAGER.GetContext(),
			mContributors,
		)
	if err != nil {
		return []models.Contributor{}, err
	}

	for j, id := range result.InsertedIDs {
		corrContributors[j].ID = id.(primitive.ObjectID)
	}

	return corrContributors, nil
}

/*
Modify the Contributor corresponding to the ID
filter is of the shape {<field to modify>: <update>}
*/
func UpdateContributor(id primitive.ObjectID, update map[string]any) error {
	updates := bson.D{}

	for k, v := range update {
		updates = append(
			updates,
			bson.E{Key: k, Value: v},
		)
	}

	updates = bson.D{{"$set", updates}}

	filter := bson.D{{"_id", id}}

	_, err := DB_MANAGER.GetCollection("contributor").UpdateOne(
		DB_MANAGER.GetContext(),
		filter,
		updates,
	)
	return err
}

func DeleteContributor(id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}
	_, err := DB_MANAGER.GetCollection("contributor").DeleteOne(
		DB_MANAGER.GetContext(),
		filter,
	)
	return err
}
