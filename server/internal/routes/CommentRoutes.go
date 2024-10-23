package routes

import (
	"log"
	db "server/internal/db"
	model "server/internal/models"

	"github.com/gofiber/fiber/v2"
)

func PostComment(c *fiber.Ctx) error{
	comment:=new(model.Comment)
	if err:=c.BodyParser(comment); err!=nil{
		return nil
	}
	res,err:=db.DB.Query("INSERT INTO comments (username,postid,content) VALUES (?,?,?)",comment.UserName,comment.PostID,comment.Content)
	if err!=nil{
		return err
	}
	log.Println(res)
	return c.JSON(comment)
}

func GetCommentsByPostID(c *fiber.Ctx) error{
	postid:=c.Params("postid")
	comments:=model.Comments{}
	rows,err:=db.DB.Query("SELECT * FROM comments WHERE postid=?",postid)
	if err!=nil{
		return err
	}
	defer rows.Close()
	for rows.Next(){
		comment:=model.Comment{}
		if err:=rows.Scan(&comment.ID,&comment.UserName,&comment.PostID,&comment.Content);err!=nil{
			return err
		}
		comments.Comments=append(comments.Comments, comment)
	}
	return c.JSON(comments)
}