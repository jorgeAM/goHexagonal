package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	e.GET("/ping", func(c *gin.Context) {
		c.JSON(201, "pong")
	})

	if err := e.Run(); err != nil {
		log.Fatal(err)
	}
}
