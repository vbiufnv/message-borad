package api

import (
	"github.com/gin-gonic/gin"
	"message-borad/resp"
	"message-borad/service"
	"strings"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if strings.ContainsAny(username, " \t\n") || strings.ContainsAny(password, " \t\n") || len(username) < 1 || len(password) < 6 {
		resp.OKWithData(c, resp.UsernameOfPasswordFormatError)
	} else {
		service.Login(c, username, password)
	}
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	//格式&长度
	if strings.ContainsAny(username, " \t\n") || strings.ContainsAny(password, " \t\n") || len(username) < 1 || len(password) < 6 {
		resp.UsernameOrPasswordError(c)
	} else {
		service.Register(c, username, password)
	}
}

func Update(c *gin.Context) {
	val, _ := c.Get("username")
	username := val.(string)

	oldPassword := c.PostForm("old_password")
	newPassword := c.PostForm("new_password")

	service.Update(c, username, oldPassword, newPassword)
}

func QuitLogin(c *gin.Context) {
	val, _ := c.Get("username")
	username := val.(string)

	service.QuitLogin(c, username)

}
