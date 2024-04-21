package model

type Post struct {
	ID       int `json:"id" gorm:"primary_key" `
	Title    string
	Content  string
	Username string
	Star     int `gorm:"default:0"`
}
