package models

type Contributor struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	ISNI    string `json:"isni"`
}
