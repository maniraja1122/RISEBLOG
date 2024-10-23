package routes

import (
	"log"
	db "server/internal/db"
	"server/internal/models"

	"github.com/gofiber/fiber/v2"
)

func GetPostLikesCount(c *fiber.Ctx) error{
	postid:=c.Params("postid")
	likes:=new(int)
	row,err:=db.DB.Query("SELECT COUNT(username) AS likes FROM likes WHERE postid=?",postid)
	if err!=nil{
		return err
	}
	if row.Next(){
		if err=row.Scan(likes); err!=nil{
			return err
		}
	}
	return c.JSON(*likes)
}
func GetLikeByUser(c *fiber.Ctx) error{
	username:=c.Params("username")
	postid:=c.Params("postid")
	likes:=new(int)
	row,err:=db.DB.Query("SELECT COUNT(username) AS likes FROM likes WHERE postid=? AND username=?",postid,username)
	if err!=nil{
		return err
	}
	if row.Next(){
		if err=row.Scan(likes); err!=nil{
			return err
		}
	}
	if *likes>0{
		return c.JSON(true)
	} else {
		return c.JSON(false)
	}
}
func PostLike(c *fiber.Ctx) error {
	like:=new(models.Like)
	if err:=c.BodyParser(like); err!=nil{
		return err
	}
	res,err:=db.DB.Query("INSERT INTO likes (username,postid) VALUES (?,?)",like.UserName,like.PostID)
	if err!=nil{
		return err
	}
	log.Println(res)
	return nil
}
func DeleteLike(c *fiber.Ctx) error{
	username:=c.Params("username")
	postid:=c.Params("postid")
	res,err:=db.DB.Query("DELETE FROM likes WHERE username=? AND postid=?",username,postid)
	if err!=nil{
		return err
	}
	log.Println(res)
	return nil
}