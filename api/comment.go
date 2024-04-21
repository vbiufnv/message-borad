package api

import (
	"github.com/gin-gonic/gin"
	"message-borad/resp"
	"message-borad/service"
	"strconv"
)

func CreateComment(c *gin.Context) {
	val, _ := c.Get("username")
	username := val.(string)
	//postid&content&parentid
	content := c.PostForm("content")
	postId, err := strconv.Atoi(c.PostForm("postid"))
	if err != nil {
		resp.ParamError(c)
		return
	}
	parentId, err := strconv.Atoi(c.PostForm("parentid"))
	if err != nil {
		resp.ParamError(c)
		return
	}

	if content == "" || postId < 1 {
		resp.OKWithData(c, "内容为空或留言不存在")
	} else {
		service.CreateComment(c, username, content, postId, parentId)
	}
}

func UpdateComment(c *gin.Context) {
	/*val, _ := c.Get("username")
	username := val.(string)*/
	//id&content
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		resp.ParamError(c)
	}
	//newContent
	content := c.PostForm("content")

	service.UpdateComment(c, content, id)
}

func DeleteComment(c *gin.Context) {
	//val, _ := c.Get("username")
	//username := val.(string)

	//id
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		resp.ParamError(c)
	}

	service.DeleteComment(c, id)
}
