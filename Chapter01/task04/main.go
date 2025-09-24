package main

import (
	_ "database/sql"
	"demo/Chapter01/task04/controller"
	"demo/Chapter01/task04/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func createTable() {
	// 创建数据库连接
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 创建表
	db.AutoMigrate(&models.Users{}, &models.Posts{}, &models.Comments{})
}

func main() {
	//createTable()
	r := gin.Default()
	controller.UserControllerInit(r)
	controller.PostControllerInit(r)
	controller.CommentControllerInit(r)
	r.Run(":8080")
}
