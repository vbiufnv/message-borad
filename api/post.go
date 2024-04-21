package api

import (
	"github.com/gin-gonic/gin"
	"message-borad/resp"
	"message-borad/service"
	"strconv"
)

// 发布
func CreatePost(c *gin.Context) {

	val, _ := c.Get("username")
	username := val.(string)

	//title&content
	title := c.PostForm("title")
	content := c.PostForm("content")
	if title == "" || content == "" {
		resp.OKWithData(c, "标题or内容为空")
	} else {
		service.CreatePost(c, title, content, username)
	}
}

func UpdatePost(c *gin.Context) {

	//postid
	postID, err := strconv.Atoi(c.PostForm("postid"))
	if err != nil {
		resp.ParamError(c)
	}

	//newTitle&newContent
	title := c.PostForm("title")
	content := c.PostForm("content")

	if title == "" || content == "" {
		resp.OKWithData(c, "标题or内容为空")
	} else {
		service.UpdatePost(c, title, content, postID)
	}
}

func DeletePost(c *gin.Context) {

	postID, err := strconv.Atoi(c.PostForm("postid"))
	if err != nil {
		resp.ParamError(c)
	} else {
		service.DeletePost(c, postID)
	}
}

func ListPosts(c *gin.Context) {
	//查看所有留言
	val, _ := c.Get("username")
	username := val.(string)

	service.ListPosts(c, username)
}

func IdToComments(c *gin.Context) {

	idStr := c.Param("id")[1:]
	postID, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ParamError(c)
	} else {
		service.IdToComments(c, postID)
	}
}

// 点赞
func LikePost(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("post_id")[1:])
	if err != nil {
		resp.ParamError(c)
	} else {
		service.UpdateStar(c, postId)
	}
}
