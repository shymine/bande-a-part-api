package models

import (
	"errors"
	"time"
)

/*
The database model for a User's command
The Books contained in this command may not be up to date as they represent the fixed state at the time of the command

Attributes:

	id : str

		The id of the specific command

	date  : Date

		The date of the command

	total : int

		The total price paid

	books : []Book

		The Books that are part of the command

	status : CommandStatus

		The current status of the command
*/
type Command struct {
	ID     string        `json:"id"`
	Date   time.Time     `json:"date"`
	Total  float32       `json:"total"`
	Books  []Book        `json:"books"`
	Status CommandStatus `json:"status"`
}

/*
CommandStatus represent the current state of the Command

TOAPPROUVE :

	the command has been made, the admin has yet to take notice of it and approuve it

APPROUVED :

	the command has been checked by an admin and has been sent

SHIPPED :

	the command has arrived and need to be retrieved

RETRIEVED :

	the command has been retrieved by the user and is considered closed

REJECTED :

	for whatever reason the command has been rejected, please contact the admin of the library
*/
type CommandStatus string

const (
	TOAPPROUVE CommandStatus = "toapprouve"
	APPROUVED  CommandStatus = "approuved"
	SHIPPED    CommandStatus = "shipped"
	RETRIEVED  CommandStatus = "retrieved"
	REJECTED   CommandStatus = "rejected"
)

func StringToCommandStatus(elem string) (CommandStatus, error) {
	switch elem {
	case "toapprouve":
		return TOAPPROUVE, nil
	case "approuved":
		return APPROUVED, nil
	case "shipped":
		return SHIPPED, nil
	case "retrieved":
		return RETRIEVED, nil
	case "rejected":
		return REJECTED, nil
	default:
		return "", errors.New("Wrong string format for Command status: " + elem)
	}
}
