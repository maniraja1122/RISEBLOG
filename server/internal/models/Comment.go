package models

type Comment struct{
	ID string `json:"id"`
	UserName string `json:"username"`
	PostID int `json:"postid"`
	Content string `json:"content"`
}