package models

/*
The database model of an Editor

Attributes:
----------
name : str

	The name of the editor
*/
type Editor struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
