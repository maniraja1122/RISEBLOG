package models

type Post struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Category string `json:"category"`
	Username string `json:"username"`
}

type Posts struct{
	Posts []Post `json:"posts"`
}