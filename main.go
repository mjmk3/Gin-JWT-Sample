package main

import (
	"Gin-JWT-Sample/api"
	middlewares "Gin-JWT-Sample/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	public := r.Group("/api")
	public.POST("/register", api.Register)
	public.POST("/login", api.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", api.CurrentUser)

	r.Run(":5500")

}
