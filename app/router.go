package app

import (
	"github.com/gin-gonic/gin"
)

func Router(s Server) (*gin.Engine, error){
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})


	r.GET("/users", func(c *gin.Context) {
		var u User
		s.DB.First(&u, 1)
		//tx := db.First(&u, 1)
		//if tx.Error != nil {
		//	return Server{}, fmt.Errorf("error finding user | %v", err)
		//}
		c.JSON(200, gin.H{"message": u})
	})
	return r, nil
}