package model

type Post struct {
	ID       int `json:"id" gorm:"primary_key" `
	Title    string
	Content  string
	Username string
	Status   int `gorm:"default:1"` //默认公开
	Star     int `gorm:"default:0"`
}
