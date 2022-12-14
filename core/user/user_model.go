package user

import "gorm.io/gorm"

type User struct {
	gorm.Model `swaggerignore:"true"`
	Username   string `json:"username" gorm:"uniqueIndex"`
	Password   string `json:"-"`
	Email      string `json:"email" gorm:"uniqueIndex"`
}
