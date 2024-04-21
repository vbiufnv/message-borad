package service

import (
	"github.com/gin-gonic/gin"
	"message-borad/dao"
	"message-borad/model"
	"message-borad/resp"
	"net/http"
)

func CreatePost(c *gin.Context, title, content, username string) {
	var post model.Post
	var err error
	post = model.Post{Title: title, Content: content, Username: username}
	err, post = dao.CreatePost(post)
	if err != nil {
		resp.FindError(c, "发布失败", err)
	} else {
		data := resp.Data{
			Message: "留言发布成功",
			Posts:   []model.Post{post},
		}
		resp.OKWithData(c, data)
	}
}

func SearchPost(postID int) (err error, post model.Post) {
	err, post = dao.SearchPost(postID)
	return err, post
}

func UpdatePost(c *gin.Context, title, content string, postID int) {
	var post model.Post
	var err error
	err, post = dao.SearchPost(postID)
	if err != nil {
		resp.NotFindPost(c)
	} else {
		err, post = dao.UpdatePost(title, content, post)
		if err != nil {
			resp.FindError(c, "更新失败", err)
		} else {
			//传回数据
			data := resp.Data{
				Message: "更改成功",
				Posts:   []model.Post{post},
			}
			resp.OKWithData(c, data)
		}
	}
}

func DeletePost(c *gin.Context, postID int) {

	var post model.Post
	var err error
	err, post = dao.SearchPost(postID)

	if err != nil {
		resp.NotFindPost(c)
	} else {
		err = dao.DeletePost(post)
		if err != nil {
			resp.FindError(c, "删除失败", err)
		} else {
			resp.OKWithData(c, "删除成功")
		}
	}
}

func ListPosts(c *gin.Context, username string) {
	var posts []model.Post
	var result int
	result, posts = dao.ListPosts(username)
	if result == 0 {
		resp.OKWithData(c, "暂无留言")
	} else {
		data := resp.Data{
			Message: "查找成功",
			Posts:   posts,
		}
		resp.OKWithData(c, data)
	}
}

func IdToComments(c *gin.Context, postID int) {
	//根评论
	var rootComments []model.Comment
	var result int
	result, rootComments = dao.SearchComments(postID)

	if result == 0 {
		data := resp.DataComments{Message: "暂无评论或id有误"}
		c.JSON(http.StatusOK, data)

	} else {

		var replyComments []model.Comment

		for i := 0; i < len(rootComments); i++ {

			replyComments = dao.GetReplyComments(rootComments[i])
			rootComments[i].Replies = replyComments
		}

		data := resp.DataComments{Message: "评论如下", Comms: rootComments}
		resp.OKWithData(c, data)

	}
}

func UpdateStar(c *gin.Context, postID int) {
	var post model.Post
	var err error
	err, post = dao.SearchPost(postID)
	if err != nil {
		resp.FindError(c, "留言不存在", err)
	} else {
		err = dao.UpdateStar(postID, post)
		if err != nil {
			resp.FindError(c, "点赞失败", err)
		} else {
			resp.OKWithData(c, "点赞成功")
		}
	}
}
