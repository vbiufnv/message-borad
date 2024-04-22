package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Status   int    `json:"status"` //登录状态
}
