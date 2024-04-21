package resp

import (
	"github.com/gin-gonic/gin"
	"message-borad/model"
	"net/http"
)

type ResponseForm struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
}

type DataComments struct {
	Message string
	Comms   []model.Comment
}

type Data struct {
	Posts   []model.Post
	Message string
}

var (
	// Success 请求成功
	Success = ResponseForm{
		Status: 10000,
		Info:   "success",
	}

	// Param 请求参数错误
	Param = ResponseForm{
		Status: 20001,
		Info:   "param error",
	}

	//token验证错误
	VerifyFailed = ResponseForm{
		Status: 20002,
		Info:   "verify failed",
	}

	// TokenExpired token过期
	TokenExpired = ResponseForm{
		Status: 20003,
		Info:   "token expired",
	}

	// RefreshTokenExpired refresh token过期
	RefreshTokenExpired = ResponseForm{
		Status: 20004,
		Info:   "登录过期，请重新登录",
	}

	//未找到留言
	NotFindPostForm = ResponseForm{
		Status: 20005,
		Info:   "未找到留言",
	}

	// UsernameOfPasswordError 登录时账号密码错误
	UsernameOfPasswordError = ResponseForm{
		Status: 20006,
		Info:   "账号或密码错误",
	}

	// 账号密码格式错误
	UsernameOfPasswordFormatError = ResponseForm{
		Status: 20007,
		Info:   "账号或密码格式错误",
	}

	// 注册失败
	RegisterErrorForm = ResponseForm{
		Status: 20008,
		Info:   "注册失败",
	}

	//原密码错误
	OldPasswordError = ResponseForm{
		Status: 20009,
		Info:   "原密码错误",
	}

	// DatabaseUnavailable 数据库不可用
	DatabaseUnavailable = ResponseForm{
		Status: 40001,
		Info:   "database unavailable",
	}
)

func OK(c *gin.Context) {
	c.JSON(http.StatusOK, Success)
}

func OKWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": 10000,
		"info":   "success",
		"data":   data,
	})
}

func VerifyError(c *gin.Context) {
	c.JSON(http.StatusForbidden, VerifyFailed)
}

func ParamError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, Param)
}

func UsernameOrPasswordError(c *gin.Context) {
	c.JSON(http.StatusOK, UsernameOfPasswordError)
}

func FindError(c *gin.Context, data string, err error) {
	c.JSON(http.StatusOK, ResponseForm{
		Status: 50099,
		Info:   data + err.Error(),
	})
}

func NotFindPost(c *gin.Context) {
	c.JSON(http.StatusOK, NotFindPostForm)
}
