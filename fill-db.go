package main

import "bande-a-part/models"

var Authors = []models.Contributor{
	{ID: "1", Name: "René", Surname: "Gérard", ISNI: "534"},
	{ID: "2", Name: "Yin", Surname: "Lao", ISNI: "486"},
	{ID: "3", Name: "Ketty", Surname: "Steward", ISNI: "788"},
	{ID: "4", Name: "Alifa", Surname: "Mawi", ISNI: "186"},
}

var Editor = []models.Editor{
	{ID: "1", Name: "Argyll"},
	{ID: "2", Name: "1115"},
}
