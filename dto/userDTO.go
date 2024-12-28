package dto

import (
	"bande-a-part/database"
	"bande-a-part/models"
)

/*
UserDTO is a Data Transfer Object where the Commands and the Bookmarks are referenced by their ID
*/
type UserDTO struct {
	ID        string   `json:"id"`
	Email     string   `json:"email"`
	Username  string   `json:"username"`
	Commands  []string `json:"commands"`
	Bookmarks []string `json:"bookmarks"`
}

func UserToDTO(user models.User) UserDTO {
	commmands := []string{}
	for _, c := range user.Commands {
		commmands = append(commmands, c.ID)
	}

	bookmarks := []string{}
	for _, c := range user.Bookmarks {
		bookmarks = append(bookmarks, c.ID)
	}

	return UserDTO{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Commands:  commmands,
		Bookmarks: bookmarks,
	}
}

func DTOToUser(userDTO UserDTO) (models.User, error) {
	commands := []models.Command{}
	for _, c := range userDTO.Commands {
		command, err := database.FindCommandById(c)
		if err != nil {
			return models.User{}, err
		}
		commands = append(commands, command)
	}

	bookmarks := []models.Book{}
	for _, c := range userDTO.Bookmarks {
		book, err := database.FindBookById(c)
		if err != nil {
			return models.User{}, err
		}
		bookmarks = append(bookmarks, book)
	}

	return models.User{
		ID:        userDTO.ID,
		Email:     userDTO.Email,
		Username:  userDTO.Username,
		Commands:  commands,
		Bookmarks: bookmarks,
	}, nil
}
