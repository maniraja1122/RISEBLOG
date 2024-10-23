package routes

import (
	"log"
	db "server/internal/db"
	model "server/internal/models"

	"github.com/gofiber/fiber/v2"
)

func GetPostsbyTopic(c *fiber.Ctx) error{
	topic:=c.Params("topic")
	posts:=model.Posts{}
	rows,err:=db.DB.Query("SELECT * FROM posts WHERE category=?",topic)
	if err!=nil{
		return err
	}
	for rows.Next(){
		post:=model.Post{}
		if err=rows.Scan(&post.ID,&post.Title,&post.Content,&post.Category,&post.Username); err!=nil{
			return err;
		}
		posts.Posts=append(posts.Posts, post)
	}
	return c.JSON(posts)
}

func GetPostsbyUser(c *fiber.Ctx) error{
	username:=c.Params("username")
	posts:=model.Posts{}
	rows,err:=db.DB.Query("SELECT * FROM posts WHERE username=?",username)
	if err!=nil{
		return err
	}
	for rows.Next(){
		post:=model.Post{}
		if err=rows.Scan(&post.ID,&post.Title,&post.Content,&post.Category,&post.Username); err!=nil{
			return err;
		}
		posts.Posts=append(posts.Posts, post)
	}
	return c.JSON(posts)
}

func GetAllPosts(c *fiber.Ctx) error{
	posts:=model.Posts{}
	rows,err:=db.DB.Query("SELECT * FROM posts")
	if err!=nil{
		return err
	}
	for rows.Next(){
		post:=model.Post{}
		if err=rows.Scan(&post.ID,&post.Title,&post.Content,&post.Category,&post.Username); err!=nil{
			return err;
		}
		posts.Posts=append(posts.Posts, post)
	}
	return c.JSON(posts)
}

func PostNewPost(c *fiber.Ctx) error{
	post:=new(model.Post)
	if err:=c.BodyParser(post); err!=nil{
		return err
	}
	res,err:=db.DB.Query("INSERT INTO posts (title,content,category,username) VALUES (?,?,?,?)",post.Title,post.Content,post.Category,post.Username)
	if err!=nil{
		return err
	}
	log.Println(res)
	return c.JSON(post)
}

func PutPost(c *fiber.Ctx) error{
	return nil
}

func DeletePost(c *fiber.Ctx) error{
	return nil
}
