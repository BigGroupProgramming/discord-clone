package app

import (
	"github.com/BigGroupProgramming/discord-clone/app/handlers/user"
	"github.com/gin-gonic/gin"
)

func Router(s Server) (*gin.Engine, error){
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//main -> init -> router -> handlers -> service -> repos

	r.GET("/users", user.GetUser(s.DB))
	return r, nil
}