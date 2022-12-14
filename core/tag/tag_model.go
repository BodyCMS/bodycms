package tag

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model `swaggerignore:"true"`
	Id         int    `json:"id" gorm:"primaryKey,autoIncrement"`
	Name       string `json:"name" gorm:"type:varchar(255);not null"`
}
