package category

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Id   int    `json:"id" gorm:"primaryKey,autoIncrement"`
	Name string `json:"name"`
}
