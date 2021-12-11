package user

import (
	"github.com/BigGroupProgramming/discord-clone/app/services/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		us := user.New(db)

		u, err := us.GetUser()
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error getting user",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": u,
		})
	}
}
