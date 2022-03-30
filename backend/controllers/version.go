package controllers

import (
	"github.com/gin-gonic/gin"
	"os"
)

func Version(c *gin.Context) {
	c.JSON(200, gin.H{
		"version": os.Getenv("VERSION"),
	})
}
