package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "please register first!"})
	})

	r.Run(":5500")

}
