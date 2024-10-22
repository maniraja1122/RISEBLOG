package main

import (
	"fmt"
	"log"
	"math/rand"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	db "server/internal/db"
	config "server/config"
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
	log.Fatal(app.Listen(fmt.Sprintf(":%d",port)))
}
