package service

import (
	"github.com/gin-gonic/gin"
	"message-borad/dao"
	"message-borad/model"
	"message-borad/resp"
)

type DataComment struct {
	Message string
	Comm    model.Comment
}

func CreateComment(c *gin.Context, username, content string, postId, parentId int) {
	var comment model.Comment
	var err error
	err, comment = dao.CreateComment(username, content, postId, parentId)
	if err != nil {
		resp.FindError(c, "评论失败", err)
	} else {
		data := DataComment{Message: "评论成功", Comm: comment}
		resp.OKWithData(c, data)
	}
}

func UpdateComment(c *gin.Context, content string, id int) {
	var comment model.Comment
	var result int
	result, comment = dao.SearchComment(id)
	if result == 0 {
		resp.OKWithData(c, "评论id有误")
	} else {
		var newComment model.Comment
		var err error
		err, newComment = dao.UpdateComment(content, comment)
		if err != nil {
			resp.FindError(c, "评论更新失败", err)
		} else {
			data := DataComment{Message: "评论更新成功", Comm: newComment}
			resp.OKWithData(c, data)
		}
	}
}

func DeleteComment(c *gin.Context, id int) {
	var comment model.Comment
	var result int
	result, comment = dao.SearchComment(id)

	if result == 0 {
		resp.OKWithData(c, "id有误，评论不存在")
	} else {
		err := dao.DeleteComment(comment)
		if err != nil {
			resp.FindError(c, "删除出错", err)
		} else {
			resp.OKWithData(c, "删除成功")
		}
	}
}
