package handlers

import "github.com/gin-gonic/gin"

func GetUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{"id": id})
}

func CreateUser(c *gin.Context) {
	c.JSON(201, gin.H{"created": true})
}
