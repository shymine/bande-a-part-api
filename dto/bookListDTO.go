package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
BookListDTO is a Data Transfer Object representing a BookList where the Books are referenced by their ID
*/
type BookListDTO struct {
	ID          primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name        string               `json:"name" bson:"name"`
	Description string               `json:"description" bson:"description"`
	Priority    uint                 `json:"priority" bson:"priority"`
	Books       []primitive.ObjectID `json:"books" bson:"books"`
}

// func BookListToDTO(bookList models.BookList) BookListDTO {
// 	books := []string{}
// 	for _, c := range bookList.Books {
// 		books = append(books, c.ID)
// 	}
// 	return BookListDTO{
// 		ID:          bookList.ID,
// 		Name:        bookList.Name,
// 		Description: bookList.Description,
// 		Priority:    bookList.Priority,
// 		Books:       books,
// 	}
// }

// func DTOToBookList(bookListDTO BookListDTO) (models.BookList, error) {
// 	books := []models.Book{}
// 	for _, c := range bookListDTO.Books {
// 		book, err := database.FindBookById(c)
// 		if err != nil {
// 			return models.BookList{}, err
// 		}
// 		books = append(books, book)
// 	}

// 	return models.BookList{
// 		ID:          bookListDTO.ID,
// 		Name:        bookListDTO.Name,
// 		Description: bookListDTO.Description,
// 		Priority:    bookListDTO.Priority,
// 		Books:       books,
// 	}, nil
// }
