package database

import (
	"bande-a-part/dto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: Complete before testing commands

func GetAllUser() ([]dto.UserDTO, error) {
	var res []dto.UserDTO
	curr, err := DB_MANAGER.GetCollection("user").Find(
		DB_MANAGER.GetContext(),
		bson.M{},
	)
	if err != nil {
		return []dto.UserDTO{}, err
	}

	err = curr.All(
		DB_MANAGER.GetContext(),
		&res,
	)
	if err != nil {
		return []dto.UserDTO{}, err
	}
	return res, err
}

func GetUserById(id primitive.ObjectID) (dto.UserDTO, error) {
	var res dto.UserDTO
	filter := bson.D{{"_id", id}}
	if err := DB_MANAGER.GetCollection("user").FindOne(
		DB_MANAGER.GetContext(),
		filter,
	).Decode(&res); err != nil {
		return dto.UserDTO{}, err
	}
	return res, nil
}

func CreateUser(user dto.UserDTO) (dto.UserDTO, error) {
	corrUser := user

	mLibrary, err := bson.Marshal(user)
	if err != nil {
		return dto.UserDTO{}, err
	}

	result, err := DB_MANAGER.
		GetCollection("user").
		InsertOne(
			DB_MANAGER.GetContext(),
			mLibrary,
		)
	if err != nil {
		return dto.UserDTO{}, err
	}

	corrUser.ID = result.InsertedID.(primitive.ObjectID)

	return corrUser, nil
}

func UpdateUser(id primitive.ObjectID, update map[string]any) error {
	updates := bson.D{}

	for k, v := range update {
		updates = append(
			updates,
			bson.E{Key: k, Value: v},
		)
	}

	updates = bson.D{{"$set", updates}}

	filter := bson.D{{"_id", id}}

	_, err := DB_MANAGER.GetCollection("user").UpdateOne(
		DB_MANAGER.GetContext(),
		filter,
		updates,
	)
	return err
}

func AddCommand(userId primitive.ObjectID, commandId primitive.ObjectID) error {
	filter := bson.D{{"_id", userId}}
	update := bson.D{{"$push", bson.D{{"commands", commandId}}}}

	_, err := DB_MANAGER.GetCollection("user").UpdateOne(
		DB_MANAGER.GetContext(),
		filter,
		update,
	)
	return err
}

func AddBookmark(userId primitive.ObjectID, bookId primitive.ObjectID) error {
	filter := bson.D{{"_id", userId}}
	update := bson.D{{"$push", bson.D{{"bookmarks", bookId}}}}

	_, err := DB_MANAGER.GetCollection("user").UpdateOne(
		DB_MANAGER.GetContext(),
		filter,
		update,
	)
	return err
}

func DeleteUser(id primitive.ObjectID) error {
	filter := bson.D{{"_id", id}}
	_, err := DB_MANAGER.GetCollection("user").DeleteOne(
		DB_MANAGER.GetContext(),
		filter,
	)
	return err
}
