package dao

import (
	"message-borad/model"
)

func CreateComment(username, content string, postId, parentId int) (error, model.Comment) {
	comment := model.Comment{Username: username, PostID: postId, Content: content}

	if parentId > 0 {
		comment.ParentID = parentId
	}

	err := DB.Create(&comment).Error
	if err != nil {
		return err, model.Comment{}
	}
	return nil, comment
}

func UpdateComment(content string, comment model.Comment) (error, model.Comment) {
	err := DB.Model(&comment).Select("content").Updates(&model.Comment{Content: content}).Error
	if err != nil {
		return err, model.Comment{}
	}
	return nil, comment
}

func SearchComment(id int) (int, model.Comment) {
	var comment model.Comment
	result := DB.Where("id=?", id).First(&comment)
	if result.RecordNotFound() {
		return 0, model.Comment{}
	}
	return 1, comment
}

func DeleteComment(comment model.Comment) error {
	return DB.Delete(&comment).Error
}

func SearchComments(id int) (int, []model.Comment) {
	var comments []model.Comment
	result := DB.Where("post_id=?", id).Find(&comments)
	if result.RowsAffected == 0 {
		return 0, []model.Comment{}
	}
	return 1, comments
}

func GetReplyComments(comment model.Comment) []model.Comment {

	DB.Preload("Replies").Find(&comment, comment.ID)
	return comment.Replies

}
