package models

type Like struct{
	UserName string `json:"username"`
	PostID int `json:"postid"`
}