package models

/*
The database model of a Genre

Attributes:
----------
name : str

	The name of the Genre (sci-fi, horror, etc.)
*/
type Genre struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
