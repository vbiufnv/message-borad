package model

import "github.com/jinzhu/gorm"

type Like struct {
	gorm.Model
	Username string `json:"username"`
	PostID   int    `json:"post_id"`
}
