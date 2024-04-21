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
	u = model.User{Username: username, Password: password}
	if err := DB.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePassword(pasword string, u model.User) error {
	err := DB.Model(&u).Select("password").Updates(model.User{Password: pasword}).Error
	return err
}
