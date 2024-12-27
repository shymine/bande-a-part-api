package dto

import (
	"bande-a-part/database"
	"bande-a-part/models"
	"time"
)

/*
CommandDTOCreated represent a Command that had just been created where it is necessary to compute

	the current Date,
	the Total price,

and set the status tu TOAPPROUVE
Moreover the Books are referenced by their ID
*/
type CommandDTOCreated struct {
	ID    string   `json:"id"`
	Books []string `json:"books"`
}

/*
CommandDTO represent a Command where the Books are referenced by their ID
*/
type CommandDTO struct {
	ID     string               `json:"id"`
	Date   time.Time            `json:"date"`
	Total  float32              `json:"total"`
	Books  []string             `json:"books"`
	Status models.CommandStatus `json:"status"`
}

func CreatedToCommand(commandC CommandDTOCreated) (models.Command, error) {
	books := []models.Book{}
	var total float32 = 0.

	for _, c := range commandC.Books {
		book, err := database.FindBookById(c)
		if err != nil {
			return models.Command{}, err
		}
		books = append(books, book)
		total += book.Price
	}

	return models.Command{
		ID:     commandC.ID,
		Date:   time.Now(),
		Status: models.TOAPPROUVE,
		Total:  total,
		Books:  books,
	}, nil
}

func CommandToDTO(command models.Command) CommandDTO {
	books := []string{}
	for _, c := range command.Books {
		books = append(books, c.ID)
	}

	return CommandDTO{
		ID:     command.ID,
		Date:   command.Date,
		Total:  command.Total,
		Books:  books,
		Status: command.Status,
	}
}

func DTOToCommand(commDTO CommandDTO) (models.Command, error) {
	books := []models.Book{}

	for _, c := range commDTO.Books {
		book, err := database.FindBookById(c)
		if err != nil {
			return models.Command{}, err
		}
		books = append(books, book)
	}
	return models.Command{
		ID:     commDTO.ID,
		Date:   commDTO.Date,
		Total:  commDTO.Total,
		Status: commDTO.Status,
		Books:  books,
	}, nil
}
