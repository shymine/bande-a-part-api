package database

import (
	"bande-a-part/models"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Contributors = []models.Contributor{
	{ID: "1", Name: "René", Surname: "Gérard", ISNI: "534"},
	{ID: "2", Name: "Yin", Surname: "Lao", ISNI: "486"},
	{ID: "3", Name: "Ketty", Surname: "Steward", ISNI: "788"},
	{ID: "4", Name: "Alifa", Surname: "Mawi", ISNI: "186"},
}

var Editors = []models.Editor{
	{ID: primitive.NewObjectID(), Name: "Argyll"},
	{ID: primitive.NewObjectID(), Name: "1115"},
}

var Libraries = []models.Library{
	{ID: "111", Name: "Bande a part"},
}

var Genres = []models.Genre{
	{ID: "745", Name: "sci-fi"},
	{ID: "258", Name: "horror"},
}

var Users = []models.User{
	{ID: "753", Email: "eruignp@vrivo", Username: "rémi"},
}

var Books = []models.Book{
	{
		ID: "485", Title: "coucou toi",
		Parution: time.Date(1987, time.April, 31, 0, 0, 0, 0, time.UTC),
		Price:    5.38, Synopsis: "ceci est un synopsis",
		ISBN: "123567894123", Stock: 5, Note: "ceci est une note",
		Contributors: []models.BookContributor{
			{Contributor: Contributors[1], Type: models.AUTHOR},
			{Contributor: Contributors[3], Type: models.ILLUSTRATOR},
		},
		Editor: Editors[0],
		Genres: []models.Genre{Genres[1]},
	}, {
		ID: "845", Title: "salut toi",
		Parution: time.Date(2019, time.April, 31, 0, 0, 0, 0, time.UTC),
		Price:    7.38, Synopsis: "ceci est un synopsis 2",
		ISBN: "123567894147", Stock: 5, Note: "ceci est une note 2",
		Contributors: []models.BookContributor{
			{Contributor: Contributors[0], Type: models.AUTHOR},
		},
		Editor: Editors[0],
		Genres: []models.Genre{Genres[0]},
	}, {
		ID: "256", Title: "yo toi",
		Parution: time.Date(1987, time.August, 31, 0, 0, 0, 0, time.UTC),
		Price:    5.99, Synopsis: "ceci est un synopsis 3",
		ISBN: "177467894123", Stock: 5, Note: "ceci est une note 3",
		Contributors: []models.BookContributor{
			{Contributor: Contributors[1], Type: models.TRANSLATOR},
			{Contributor: Contributors[3], Type: models.AUTHOR},
		},
		Editor: Editors[1],
		Genres: []models.Genre{Genres[0], Genres[1]},
	},
}

var BookLists = []models.BookList{
	{
		ID: "369", Name: "rentrée littéraire",
		Description: "des bouquins en veux tu en voila",
		Priority:    1, Books: []models.Book{Books[0], Books[1]},
	}, {
		ID: "719", Name: "liste de noel",
		Description: "des bouquins pour faire des cadeaux",
		Priority:    4, Books: []models.Book{Books[1], Books[2]},
	},
}

var Commands = []models.Command{
	{
		ID: "203", Date: time.Now(), Total: 55, Books: []models.Book{Books[2]},
		Status: models.TOAPPROUVE,
	},
}

func UpdateFill() {
	Users[0].Commands = append(Users[0].Commands, Commands[0])
}

func FindContributorById(id string) (models.Contributor, error) {
	for _, c := range Contributors {
		if c.ID == id {
			return c, nil
		}
	}
	return models.Contributor{}, errors.New("Contributor not found with id: " + id)
}

func FindEditorById(id primitive.ObjectID) (models.Editor, error) {
	for _, c := range Editors {
		if c.ID == id {
			return c, nil
		}
	}
	return models.Editor{}, errors.New("Editor not found with id: " + id.String())
}

func FindGenreByName(name string) (models.Genre, error) {
	for _, c := range Genres {
		if c.Name == name {
			return c, nil
		}
	}
	return models.Genre{}, errors.New("Genre not found with name: " + name)
}

func FindBookById(id string) (models.Book, error) {
	for _, c := range Books {
		if c.ID == id {
			return c, nil
		}
	}
	return models.Book{}, errors.New("Book not found with id: " + id)
}

func FindCommandById(id string) (models.Command, error) {
	for _, c := range Commands {
		if c.ID == id {
			return c, nil
		}
	}
	return models.Command{}, errors.New("Command not found with id: " + id)
}

func FindBookListById(id string) (models.BookList, error) {
	for _, c := range BookLists {
		if c.ID == id {
			return c, nil
		}
	}
	return models.BookList{}, errors.New("BookList not found with id: " + id)
}

func FindUserById(id string) (models.User, error) {
	for _, c := range Users {
		if c.ID == id {
			return c, nil
		}
	}
	return models.User{}, errors.New("User not found with id: " + id)
}

func FindCommandByStatus(status models.CommandStatus) []models.Command {
	commands := []models.Command{}
	for _, c := range Commands {
		if c.Status == status {
			commands = append(commands, c)
		}
	}
	return commands
}
