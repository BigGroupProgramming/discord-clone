package models

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Password string
}

func (User) TableName() string {
	scheme := "public"
	return fmt.Sprintf("%v.user", scheme)
}
