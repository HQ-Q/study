package main

import (
	_ "database/sql"
	"demo/Chapter01/task03/gormStudy"
	"fmt"

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
	db.AutoMigrate(&gormStudy.User{}, &gormStudy.Post{}, &gormStudy.Comment{})
}

func query() {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接数据库失败")
		panic(err)
	}
	defer db.Close()

	var posts []gormStudy.Post
	// 预加载 Author 和 Comments 关联数据
	db.Where("user_id = ?", 2).Preload("Author").Preload("Comments").Preload("Comments.FromUser").Find(&posts)

	for _, post := range posts {
		fmt.Println("文章标题：", post.Title)
		fmt.Println("文章内容：", post.Content)
		fmt.Println("文章作者Id：", post.UserId)
		fmt.Println("文章作者：", post.Author.UserName)
		fmt.Println("文章发布时间：", post.CreatedTime.Format("2006-01-02 15:04:05"))
		fmt.Println("文章更新时间：", post.UpdatedTime.Format("2006-01-02 15:04:05"))
		fmt.Println("文章评论：")
		fmt.Println("**********************************")
		for _, comment := range post.Comments {
			fmt.Println("文章Id：", comment.PostId)
			fmt.Println("文章作者Id：", comment.ToUserId)
			fmt.Println("评论作者Id：", comment.FromUserId)
			fmt.Println("评论作者：", comment.FromUser.UserName)
			fmt.Println("评论内容：", comment.Content)
			fmt.Println("评论时间：", comment.CreatedTime.Format("2006-01-02 15:04:05"))
			fmt.Println("评论更新时间：", comment.UpdatedTime.Format("2006-01-02 15:04:05"))
			fmt.Println("--------------------------------")
		}
	}

	//var posts = make([]gormStudy.Post, 10)
	//db.Where("user_id = ?", 1).Find(&posts)
	//for _, post := range posts {
	//	var comments = make([]gormStudy.Comment, 10)
	//	db.Where("post_id = ?", post.Id).Find(&comments)
	//	fmt.Println("文章标题：", post.Title)
	//	fmt.Println("文章内容：", post.Content)
	//	fmt.Println("文章作者Id：", post.UserId)
	//
	//	fmt.Println("文章作者：", post.Author.UserName)
	//	fmt.Println("文章发布时间：", post.CreatedTime.Format("2006-01-02 15:04:05"))
	//	fmt.Println("文章更新时间：", post.UpdatedTime.Format("2006-01-02 15:04:05"))
	//	fmt.Println("文章评论：")
	//	fmt.Println("**********************************")
	//	for _, comment := range comments {
	//		fmt.Println("评论内容：", comment.Content)
	//		fmt.Println("评论作者：", comment.FromUserId)
	//		fmt.Println("评论时间：", comment.CreatedTime.Format("2006-01-02 15:04:05"))
	//		fmt.Println("评论更新时间：", comment.UpdatedTime.Format("2006-01-02 15:04:05"))
	//		fmt.Println("--------------------------------")
	//	}
	//}
}

func main() {
	query()
}
