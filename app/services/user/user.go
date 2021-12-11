package user

import (
	"fmt"
	"github.com/BigGroupProgramming/discord-clone/models"
	"gorm.io/gorm"
)

type userService struct {
	db *gorm.DB
}

type UserService interface {
	GetUser() (models.User, error)
}

func New(db *gorm.DB) UserService{
	return &userService{db: db}
}

func (us *userService) GetUser() (models.User, error) {
	var u models.User
	tx := us.db.First(&u, 1)
	if tx.Error != nil {
		return u, fmt.Errorf("error getting user | %v", tx.Error)
	}

	return u, nil
}