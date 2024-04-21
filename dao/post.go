package dao

import (
	"message-borad/model"
)

func CreatePost(post model.Post) (error, model.Post) {
	err := DB.Create(&post).Error
	var newPost model.Post
	newPost = post
	return err, newPost
}

func SearchPost(postID int) (error, model.Post) {
	var post model.Post
	err := DB.Where("id=?", postID).First(&post).Error
	return err, post
}

func UpdatePost(title, content string, post model.Post) (err error, newPost model.Post) {
	err = DB.Model(&post).Select("title", "content").Updates(model.Post{Title: title, Content: content}).Error
	newPost = model.Post{Title: title, Content: content, Username: post.Username, ID: post.ID}
	return err, newPost
}

func DeletePost(post model.Post) error {
	err := DB.Delete(&post).Error
	return err
}

func ListPosts(username string) (int, []model.Post) {
	var posts []model.Post
	result := DB.Where("username=?", username).Find(&posts)
	if result.RowsAffected == 0 {
		return 0, nil
	}
	return 1, posts
}

func UpdateStar(postid int, post model.Post) error {
	return DB.Model(&post).Select("star").Updates(model.Post{Star: post.Star + 1}).Error
}
