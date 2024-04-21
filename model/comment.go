package model

type Comment struct {
	ID       int    `json:"id" gorm:"primary_key"`
	PostID   int    `json:"postid"`   //留言id
	Username string `json:"username"` //当前用户
	Content  string `json:"content"`
	ParentID int    `gorm:"default:-1" json:"parentid"` //上一级评论id

	Replies []Comment `gorm:"foreignkey:ParentID"`
}
