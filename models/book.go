package models

import "time"

/*
The database model for a Book

Attributes:
----------
title     : str

	The title of the Book

parution : date

	The date of release of the Book

price    : positive float

	The price to pay for this book

synopsis : str

	The synopsis of the book, usually found on the book backpage

ISBN     : 13 digits int

	The Identification number of the book

stock    : positive int

	The number of remaining books in stock

new      : boolean

	Is this book considered new (if yes, will be put up front in the "New" BookList)
	After a month new will be put to false

note     : str

	The note of the librarian about this book

contributors  : []BookContributor

	The authors of the book

editor   : Editor

	The editor of the book

genres   : []Genre

	The genres of the book (mainly for tagging) such as sci-fi, horror, etc.
*/
type Book struct {
	ID           string            `json:"id"`
	Title        string            `json:"title"`
	Parution     time.Time         `json:"parution"`
	Price        float32           `json:"price"`
	Synopsis     string            `json:"synopsis"`
	ISBN         string            `json:"isbn"`
	Stock        uint              `json:"stock"`
	New          bool              `json:"new"`
	Note         string            `string:"note"`
	Contributors []BookContributor `json:"books"`
	Editor       Editor            `json:"editor"`
	Genres       []Genre           `json:"genres"`
}

/*
BookContributor encompass the contributor and its role in the publication

Attributes:
----------
contributor : Contributor

	The contributor of the publication

Type : ContributorType

	The type of contribution
*/
type BookContributor struct {
	Contributor Contributor     `json:"contributor"`
	Type        ContributorType `json:"type"`
}

/*
ContributorType explicit the contributor's kind of contribution to the publication
*/
type ContributorType string

const (
	AUTHOR      ContributorType = "author"
	ILLUSTRATOR ContributorType = "illustrator"
	TRANSLATOR  ContributorType = "translator"
)
