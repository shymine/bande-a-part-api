package models

/*
The database model of a BookList

This list can be for promotionnal events, for themed package of books, etc.

Attributes:
----------
name        : str

	The name of the book list

description : str

	A description of this list

priority    : positive int

	The priority of this list against the other lists, determine the order they will be displayed

books       : []Book

	The books composing the list
*/
type BookList struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    uint   `json:"priority"`
	Books       []Book `json:"books"`
}
