package gormStudy

import (
	"time"

	"github.com/jinzhu/gorm"
)

// 假设你要开发一个博客系统，有以下几个实体： Author （用户）、 Post （文章）、 Comment （评论）
// 要求 ：
// 使用Gorm定义 Author 、 Post 和 Comment 模型，其中 Author 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
// 编写Go代码，使用Gorm创建这些模型对应的数据库表。

// User 用户
type User struct {
	Id          int       `db:"id" json:"id"`
	UserName    string    `db:"user_name" json:"username" gorm:"not null;index"`
	Password    string    `db:"password" json:"password" gorm:"not null"`
	NickName    string    `db:"nick_name" json:"nickname" gorm:"not null;index"`
	PostCount   int       `db:"post_count" json:"postCount" gorm:"not null;default:0"` // 文章数量统计
	CreatedTime time.Time `db:"created_time" json:"createdTime" gorm:"not null"`
	UpdatedTime time.Time `db:"updated_time" json:"updatedTime" gorm:"not null"`
}

func (u *User) TableName() string {
	return "user"
}

// Post 文章 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
type Post struct {
	Id          int       `db:"id" json:"id"`
	Title       string    `db:"title" json:"title" gorm:"not null"`
	Content     string    `db:"content" json:"content" gorm:"not null"`
	UserId      int       `db:"user_id" json:"user_id" gorm:"not null;index"`
	Author      User      `gorm:"foreignkey:UserId" json:"author"`   // 添加用户关联
	Comments    []Comment `gorm:"foreignkey:PostId" json:"comments"` // 添加评论关联
	CreatedTime time.Time `db:"created_time" json:"createdTime" gorm:"not null"`
	UpdatedTime time.Time `db:"updated_time" json:"updatedTime" gorm:"not null"`
}

// AfterCreate 文章创建后的钩子函数，用于更新用户的文章数量
func (p *Post) AfterCreate(tx *gorm.DB) error {
	// 更新用户的文章数量，自增1
	return tx.Model(&User{}).Where("id = ?", p.UserId).Update("post_count", gorm.Expr("post_count + ?", 1)).Error
}

func (p *Post) TableName() string {
	return "post"
}

// Comment 评论
type Comment struct {
	Id          int       `db:"id" json:"id" gorm:"primary_key"`
	PostId      int       `db:"post_id" json:"postId" gorm:"not null;index"`
	FromUserId  int       `db:"from_user_id" json:"fromUserId" gorm:"not null;index"`
	FromUser    User      `gorm:"foreignkey:FromUserId" json:"fromUser"`
	ToUserId    int       `db:"to_user_id" json:"toUserId" gorm:"not null;index"`
	Content     string    `db:"content" json:"content" gorm:"not null"`
	CreatedTime time.Time `db:"created_time" json:"createdTime" gorm:"not null"`
	UpdatedTime time.Time `db:"updated_time" json:"updatedTime" gorm:"not null"`
}

func (c *Comment) TableName() string {
	return "comment"
}
