package routes

import (
	"log"
	db "server/internal/db"
	model "server/internal/models"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error{
	username:=c.Params("username")
	user:=model.User{}
	rows,err:=db.DB.Query("SELECT username,password FROM users where username=?",username)
	if err!=nil{
		return err
	}
	defer rows.Close()
	if rows.Next(){
		if err=rows.Scan(&user.UserName,&user.Password); err!=nil{
			return err
		}
	}
	return c.JSON(user)
}

func PostUser(c *fiber.Ctx) error{
	user:=new(model.User)
	if err:=c.BodyParser(user); err!=nil{
		return err;
	}
	res,err:= db.DB.Query("INSERT INTO users (username,password) VALUES (?,?)",user.UserName,user.Password)
	if err!=nil{
		return err
	}
	log.Println(res)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error{
	username:=c.Params("username")
	res,err:=db.DB.Query("DELETE FROM users WHERE username=?",username)
	if err!=nil{
		return err
	}
	log.Println(res)
	return c.JSON("Deleted")
}