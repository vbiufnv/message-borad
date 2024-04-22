package dao

import (
	"message-borad/model"
)

func SearchUser(username string) (bool, model.User) {
	var u model.User
	result := DB.Where("username=?", username).First(&u)
	if result.RecordNotFound() {
		return false, u
	}
	return true, u
}

func CreateUser(username, password string) error {
	var u model.User
	u = model.User{Username: username, Password: password, Status: 0}
	if err := DB.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePassword(password string, u model.User) error {
	err := DB.Model(&u).Select("password").Updates(model.User{Password: password}).Error
	return err
}

func UpdateLoginStatus(username string, status int) error {
	return DB.Model(&model.User{}).Where("username = ?", username).Update("status", status).Error
}
