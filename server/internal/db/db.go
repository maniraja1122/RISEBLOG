package db

import (
	"os"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
);

var db *sql.DB

type DBConfig struct {
    User     string
    Password string
    Host     string
    Port     string
    Name     string
}
func LoadDBConfig() *DBConfig {
	return &DBConfig{
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME"),
	}
}
func Connect() error {
	var err error
	config:=LoadDBConfig()
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s", config.User, config.Password,config.Host,config.Port,config.Name))
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	err=CreateTables()
	if err!=nil{
		return err
	}
	return nil
}

// Model Tables SQL Mapping Functions
func CreateUserTable() error{
	_,error:=db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
	username VARCHAR(255) NOT NULL PRIMARY KEY UNIQUE,
	password TEXT NOT NULL
);`)
	return error
}

func CreatePostTable() error {
	_,error:=db.Exec(`
	CREATE TABLE IF NOT EXISTS posts (
	id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	title TEXT NOT NULL,
	content TEXT NOT NULL,
	category TEXT NOT NULL,
	username VARCHAR(255) NOT NULL
);`);
	return error
}

func CreateCommentTable() error{
	_,error:=db.Exec(`
	CREATE TABLE IF NOT EXISTS comments (
	id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	username VARCHAR(255) NOT NULL,
	postid BIGINT NOT NULL,
	content TEXT NOT NULL
);`)
	return error
}

func CreateLikeTable() error{
	_,error:=db.Exec(`
	CREATE TABLE IF NOT EXISTS likes (
	username VARCHAR(255) NOT NULL,
	postid BIGINT NOT NULL,
	CONSTRAINT USER_POST PRIMARY KEY (username,postid)
);
	`)
	return error
}

func CreateTables() error{
	err:=CreateUserTable()
	if err!=nil{
		return err
	}
	err=CreatePostTable()
	if err!=nil{
		return err
	}
	err=CreateLikeTable()
	if err!=nil{
		return err
	}
	err=CreateCommentTable()
	if err!=nil{
		return err
	}
	return nil
}

// Queries