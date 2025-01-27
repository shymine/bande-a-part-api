package database

import (
	"bande-a-part/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetBookList() ([]dto.BookListDTO, error) {
	var res []dto.BookListDTO
	curr, err := DB_MANAGER.GetCollection("bookList").Find(
		DB_MANAGER.GetContext(),
		bson.M{},
	)
	if err != nil {
		return []dto.BookListDTO{}, err
	}

	err = curr.All(
		DB_MANAGER.GetContext(),
		&res,
	)
	if err != nil {
		return []dto.BookListDTO{}, err
	}
	return res, err
}

func CreateBookList(bookList dto.BookListDTO) (dto.BookListDTO, error) {
	corrBookList := bookList

	mLibrary, err := bson.Marshal(bookList)
	if err != nil {
		return dto.BookListDTO{}, err
	}

	result, err := DB_MANAGER.
		GetCollection("bookList").
		InsertOne(
			DB_MANAGER.GetContext(),
			mLibrary,
		)
	if err != nil {
		return dto.BookListDTO{}, err
	}

	corrBookList.ID = result.InsertedID.(primitive.ObjectID)

	return corrBookList, nil
}

// TODO: check if the update concern a referenced field that the references are correct
func UpdateBookList(id primitive.ObjectID, update map[string]any) error {
	updates := bson.D{}

	for k, v := range update {
		updates = append(
			updates,
			bson.E{Key: k, Value: v},
		)
	}

	updates = bson.D{{"$set", updates}}

	filter := bson.D{{"_id", id}}

	_, err := DB_MANAGER.GetCollection("bookList").UpdateOne(
		DB_MANAGER.GetContext(),
		filter,
		updates,
	)
	return err
}

func DeleteBookList(id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}
	_, err := DB_MANAGER.GetCollection("bookList").DeleteOne(
		DB_MANAGER.GetContext(),
		filter,
	)
	return err
}
