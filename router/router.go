package router

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine{
	r := gin.Default()

	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "I am root end point",
		})
	})
	return r
}
