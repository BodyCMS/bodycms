package post

import "gorm.io/gorm"

type Post struct {
	gorm.Model `swaggerignore:"true"`
	Id         int    `json:"id" gorm:"primary_key"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	AuthorId   int    `json:"author_id"`
	Author     string `json:"author"`
	Tags       string `json:"tags"`
	Categories string `json:"categories"`
}
