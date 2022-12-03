package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("", Hand)
	router.Run(":8080")
}

func Hand(context *gin.Context) {
	fmt.Println("Ok")
}
