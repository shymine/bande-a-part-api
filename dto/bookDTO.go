package dto

import (
	"bande-a-part/database"
	"bande-a-part/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
BookDTO is a Data Transfer object so that the contributor is referenced by a BookContributorDTO, editor values by their ID
and the Genres value by the name of the genre
*/
type BookDTO struct {
	ID           string               `json:"id"`
	Title        string               `json:"title"`
	Parution     time.Time            `json:"parution"`
	Price        float32              `json:"price"`
	Synopsis     string               `json:"synopsis"`
	ISBN         string               `json:"isbn"`
	Stock        uint                 `json:"stock"`
	Note         string               `json:"note"`
	Contributors []BookContributorDTO `json:"contributors"`
	Editor       primitive.ObjectID               `json:"editor"`
	Genres       []string             `json:"genres"`
}

type BookContributorDTO struct {
	ID   string                 `json:"id"`
	Type models.ContributorType `json:"type"`
}

func BookToDTO(book models.Book) BookDTO {
	contributors := []BookContributorDTO{}
	for _, c := range book.Contributors {
		contributors = append(contributors, BookContributorDTO{c.Contributor.ID, c.Type})
	}

	genres := []string{}
	for _, c := range book.Genres {
		genres = append(genres, c.Name)
	}

	return BookDTO{
		ID:           book.ID,
		Title:        book.Title,
		Parution:     book.Parution,
		Price:        book.Price,
		Synopsis:     book.Synopsis,
		ISBN:         book.ISBN,
		Stock:        book.Stock,
		Note:         book.Note,
		Contributors: contributors,
		Editor:       book.Editor.ID,
		Genres:       genres,
	}
}

func DTOToBook(bookDTO BookDTO) (models.Book, error) {
	contributors := []models.BookContributor{}
	for _, c := range bookDTO.Contributors {
		contrib, err := database.FindContributorById(c.ID)
		if err != nil {
			return models.Book{}, err
		}
		contributors = append(contributors, models.BookContributor{
			Contributor: contrib,
			Type:        c.Type,
		})
	}

	editor, err := database.FindEditorById(bookDTO.Editor)
	if err != nil {
		return models.Book{}, err
	}

	genres := []models.Genre{}
	for _, c := range bookDTO.Genres {
		genre, err := database.FindGenreByName(c)
		if err != nil {
			return models.Book{}, err
		}
		genres = append(genres, genre)
	}

	return models.Book{
		ID:           bookDTO.ID,
		Title:        bookDTO.Title,
		Parution:     bookDTO.Parution,
		Price:        bookDTO.Price,
		Synopsis:     bookDTO.Synopsis,
		ISBN:         bookDTO.ISBN,
		Stock:        bookDTO.Stock,
		Note:         bookDTO.Note,
		Contributors: contributors,
		Editor:       editor,
		Genres:       genres,
	}, nil
}
