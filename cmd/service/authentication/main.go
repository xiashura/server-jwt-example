package main

import (
	"github.com/server-jwt-exmple/middleware"
	"github.com/server-jwt-exmple/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/Registration", func(c *gin.Context) {
		var client middleware.Client
		c.BindJSON(&client)
		token := middleware.Registration(client)
		c.JSON(200, gin.H{
			"Token": model.Token{
				Key:        token.Key,
				Time:       token.Time,
				Authorized: token.Authorized,
			},
		})
	})
	r.GET("/Authentication", func(c *gin.Context) {
		var client middleware.Client
		c.BindJSON(&client)
		token := middleware.Authentication(client)
		c.JSON(200, gin.H{
			"client": token,
		})
	})
	r.GET("/Unauthenticated", func(c *gin.Context) {
		var client middleware.Client
		c.BindJSON(&client)
		token := middleware.Unauthenticated(client)
		c.JSON(200, gin.H{
			"client": token,
		})
	})
	r.GET("/Expired", func(c *gin.Context) {

		var client middleware.Client
		c.BindJSON(&client)
		token := middleware.Expired(client)
		c.JSON(200, gin.H{
			"Token": model.Token{
				Key:        token.Key,
				Time:       token.Time,
				Authorized: token.Authorized,
			},
		})
	})
	r.Run()
}
