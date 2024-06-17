package main

import (
	"github.com/gin-gonic/gin"
)

func listEvents(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
