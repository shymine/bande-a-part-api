package main

import (
	"bande-a-part/database"
	"log"

	"github.com/gin-gonic/gin"
)

type Params struct {
	DbName string `json:"dbName"`
	Collections []string `json:"collections"`
}



func main() {
	database.UpdateFill()
	
	router := gin.Default()
	SetEndPoints(router)
	database.SetDBManager("bande-a-part", []string{
		"book",
		"bookList",
		"command",
		"contributor",
		"editor",
		"genre",
		"library",
		"user",
	})
	defer database.DB_MANAGER.Disconnect()

	// 0.0.0.0 => listen to all
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal(err)
	}
}
