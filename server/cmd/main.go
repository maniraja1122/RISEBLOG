package main

import (
	"fmt"
	"log"
	"math/rand"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	db "server/internal/db"
	config "server/config"
	routes "server/internal/routes"
)

func main(){
	err:=config.LoadEnv()
	if err!=nil{
		log.Fatal(err)
	}
	err=db.Connect()
	if err!=nil{
		log.Fatal(err)
	}
	app:=fiber.New()
	port:=rand.Intn(5000)+3000
	// Defining Routes
	// User
	app.Get("/user/:username",routes.GetUser)
	app.Delete("/user/:username",routes.DeleteUser)
	app.Post("/user",routes.PostUser)
	// Comment
	app.Get("/comment/:postid",routes.GetCommentsByPostID)
	app.Post("/comment",routes.PostComment)
	// Likes
	app.Get("/likes/:postid",routes.GetPostLikesCount)
	app.Get("/likes/:postid/:username",routes.GetLikeByUser)
	app.Post("/likes",routes.PostLike)
	app.Delete("/likes/:postid/:username",routes.DeleteLike)
	// Posts
	app.Get("/posts/topic/:topic",routes.GetPostsbyTopic)
	app.Get("/posts/user/:username",routes.GetPostsbyUser)
	app.Get("/posts",routes.GetAllPosts)
	app.Post("/posts",routes.PostNewPost)
	log.Fatal(app.Listen(fmt.Sprintf(":%d",port)))
}
