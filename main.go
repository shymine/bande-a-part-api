package main

import "github.com/gin-gonic/gin"

func setEndPoints(router *gin.Engine) {
	router.GET("/contributor", getContributor)
}

func main() {
	router := gin.Default()
	setEndPoints(router)
	router.Run("localhost:8080")
}
