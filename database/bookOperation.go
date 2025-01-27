package database

import (
	"bande-a-part/dto"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllBooks() ([]dto.BookDTO, error) {
	var res []dto.BookDTO
	curr, err := DB_MANAGER.GetCollection("book").Find(
		DB_MANAGER.GetContext(),
		bson.M{},
	)
	if err != nil {
		return []dto.BookDTO{}, err
	}

	err = curr.All(
		DB_MANAGER.GetContext(),
		&res,
	)
	if err != nil {
		return []dto.BookDTO{}, err
	}
	return res, err
}

func GetBookById(id primitive.ObjectID) (dto.BookDTO, error) {
	var res dto.BookDTO
	filter := bson.D{{"_id", id}}
	if err := DB_MANAGER.GetCollection("book").FindOne(
		DB_MANAGER.GetContext(),
		filter,
	).Decode(&res); err != nil {
		return dto.BookDTO{}, err
	}
	return res, nil
}

func CreateBook(book dto.BookDTO) (dto.BookDTO, error) {
	corrBook := book

	mLibrary, err := bson.Marshal(book)
	if err != nil {
		log.Println("error marshalling book: ", err)
	}

	result, err := DB_MANAGER.
		GetCollection("book").
		InsertOne(
			DB_MANAGER.GetContext(),
			mLibrary,
		)
	if err != nil {
		return dto.BookDTO{}, err
	}

	corrBook.ID = result.InsertedID.(primitive.ObjectID)

	return corrBook, nil
}

// TODO: check if the update concern a referenced field that the references are correct
func UpdateBook(id primitive.ObjectID, update map[string]any) error {
	updates := bson.D{}

	for k, v := range update {
		updates = append(
			updates,
			bson.E{Key: k, Value: v},
		)
	}

	updates = bson.D{{"$set", updates}}

	filter := bson.D{{"_id", id}}

	_, err := DB_MANAGER.GetCollection("book").UpdateOne(
		DB_MANAGER.GetContext(),
		filter,
		updates,
	)
	return err
}

func DeleteBook(id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}
	_, err := DB_MANAGER.GetCollection("book").DeleteOne(
		DB_MANAGER.GetContext(),
		filter,
	)
	return err
}
