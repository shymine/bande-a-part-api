package database

import (
	"bande-a-part/dto"
	"bande-a-part/models"
	"errors"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCommandByUser(userId primitive.ObjectID) ([]dto.CommandDTO, error) {
	var user dto.UserDTO
	var commands []dto.CommandDTO

	filter := bson.D{{"_id", userId}}
	if err := DB_MANAGER.GetCollection("user").FindOne(
		DB_MANAGER.GetContext(),
		filter,
	).Decode(&user); err != nil {
		var s strings.Builder
		s.WriteString("GetCommandByUser decode user: ")
		s.WriteString(" ")
		s.WriteString(err.Error())
		return []dto.CommandDTO{}, errors.New(s.String())
	}

	for _, c := range user.Commands {
		var command dto.CommandDTO
		filter = bson.D{{"_id", c}}
		if err := DB_MANAGER.GetCollection("command").FindOne(
			DB_MANAGER.GetContext(),
			filter,
		).Decode(&command); err != nil {
			return []dto.CommandDTO{}, errors.New("a command in user, doesn t exist")
		}

		commands = append(commands, command)
	}

	if len(commands) == 0 {
		return commands, errors.New("commands are empty but " + fmt.Sprint(len(user.Commands)))
	}

	return commands, nil
}

func GetCommandByStatus(status models.CommandStatus) ([]dto.CommandDTO, error) {
	var commands []dto.CommandDTO

	filter := bson.D{{"status", status}}
	curr, err := DB_MANAGER.GetCollection("command").Find(
		DB_MANAGER.GetContext(),
		filter,
	)
	if err != nil {
		return []dto.CommandDTO{}, err
	}
	err = curr.All(
		DB_MANAGER.GetContext(),
		&commands,
	)
	if err != nil {
		return []dto.CommandDTO{}, err
	}
	return commands, nil
}

func GetCommandById(commandId primitive.ObjectID) (dto.CommandDTO, error) {
	filter := bson.D{{"_id", commandId}}

	var command dto.CommandDTO
	if err := DB_MANAGER.GetCollection("command").FindOne(
		DB_MANAGER.GetContext(),
		filter,
	).Decode(&command); err != nil {
		return dto.CommandDTO{}, err
	}

	return command, nil

}

func CreateCommand(command dto.CommandDTO) (dto.CommandDTO, error) {
	corrCommand := command

	mLibrary, err := bson.Marshal(command)
	if err != nil {
		return dto.CommandDTO{}, err
	}

	result, err := DB_MANAGER.
		GetCollection("command").
		InsertOne(
			DB_MANAGER.GetContext(),
			mLibrary,
		)
	if err != nil {
		return dto.CommandDTO{}, err
	}

	corrCommand.ID = result.InsertedID.(primitive.ObjectID)

	return corrCommand, nil
}

func SetCommandStatus(id primitive.ObjectID, statusUpdate models.CommandStatus) error {
	updates := bson.D{{"status", statusUpdate}}

	updates = bson.D{{"$set", updates}}

	filter := bson.D{{"_id", id}}
	_, err := DB_MANAGER.GetCollection("command").UpdateOne(
		DB_MANAGER.GetContext(),
		filter,
		updates,
	)
	return err
}

// TODO: when delete, also remove it from the user
func DeleteCommand(id primitive.ObjectID) error {
	var user dto.UserDTO

	userFilter := bson.D{{"commands", id}}
	if err := DB_MANAGER.GetCollection("user").FindOne(
		DB_MANAGER.GetContext(),
		userFilter,
	).Decode(&user); err != nil {
		return errors.New("not finding user for command " + id.String())
	}

	commandIndex := -1
	for i, el := range user.Commands {
		if el == id {
			commandIndex = i
		}
	}
	if commandIndex == -1 {
		return errors.New("not finding command in user commands " + id.String())
	}

	update := bson.D{{"commands", append(user.Commands[:commandIndex], user.Commands[commandIndex+1:]...)}}
	update = bson.D{{"$set", update}}
	_, err := DB_MANAGER.GetCollection("user").UpdateOne(
		DB_MANAGER.GetContext(),
		userFilter,
		update,
	)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", id}}
	_, err = DB_MANAGER.GetCollection("command").DeleteOne(
		DB_MANAGER.GetContext(),
		filter,
	)
	return err
}
