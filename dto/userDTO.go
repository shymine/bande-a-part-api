package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
UserDTO is a Data Transfer Object where the Commands and the Bookmarks are referenced by their ID
*/
type UserDTO struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Email     string               `json:"email" bson:"email"`
	Username  string               `json:"username" bson:"username"`
	Commands  []primitive.ObjectID `json:"commands" bson:"commands"`
	Bookmarks []primitive.ObjectID `json:"bookmarks" bson:"bookmarks"`
}

// func UserToDTO(user models.User) UserDTO {
// 	commmands := []string{}
// 	for _, c := range user.Commands {
// 		commmands = append(commmands, c.ID)
// 	}

// 	bookmarks := []string{}
// 	for _, c := range user.Bookmarks {
// 		bookmarks = append(bookmarks, c.ID)
// 	}

// 	return UserDTO{
// 		ID:        user.ID,
// 		Email:     user.Email,
// 		Username:  user.Username,
// 		Commands:  commmands,
// 		Bookmarks: bookmarks,
// 	}
// }

// func DTOToUser(userDTO UserDTO) (models.User, error) {
// 	commands := []models.Command{}
// 	for _, c := range userDTO.Commands {
// 		command, err := database.FindCommandById(c)
// 		if err != nil {
// 			return models.User{}, err
// 		}
// 		commands = append(commands, command)
// 	}

// 	bookmarks := []models.Book{}
// 	for _, c := range userDTO.Bookmarks {
// 		book, err := database.FindBookById(c)
// 		if err != nil {
// 			return models.User{}, err
// 		}
// 		bookmarks = append(bookmarks, book)
// 	}

// 	return models.User{
// 		ID:        userDTO.ID,
// 		Email:     userDTO.Email,
// 		Username:  userDTO.Username,
// 		Commands:  commands,
// 		Bookmarks: bookmarks,
// 	}, nil
// }
