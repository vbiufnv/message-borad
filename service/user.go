package service

import (
	"github.com/gin-gonic/gin"
	"message-borad/dao"
	"message-borad/model"
	"message-borad/resp"
)

func Login(c *gin.Context, username, password string) {
	var u model.User
	var result bool
	result, u = dao.SearchUser(username)

	if !result || u.Password != password {
		resp.UsernameOrPasswordError(c)
	} else {
		//token
		tokenString := CreateToken(username)
		resp.OKWithData(c, map[string]string{"token": tokenString})
	}
}

func Register(c *gin.Context, username, password string) {
	result, _ := dao.SearchUser(username)
	if result {
		resp.OKWithData(c, "用户名已存在")
	} else {
		var err error
		err = dao.CreateUser(username, password)
		if err != nil {
			resp.FindError(c, "注册失败", err)
		} else {
			resp.OKWithData(c, "注册成功")
		}
	}
}

func Update(c *gin.Context, username, oldPassword, newPassword string) {
	var u model.User
	var result bool
	result, u = dao.SearchUser(username)
	if !result {
		resp.OKWithData(c, "用户名不存在")

	} else {
		if u.Password != oldPassword {
			resp.OKWithData(c, "原密码错误")
		}
		err := dao.UpdatePassword(newPassword, u)
		if err != nil {
			resp.FindError(c, "修改失败", err)
		} else {
			resp.OKWithData(c, "修改成功")
		}
	}
}
