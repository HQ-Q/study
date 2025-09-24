package services

import (
	"demo/Chapter01/task04/models"
	"demo/Chapter01/task04/req"
	"demo/Chapter01/task04/utils"
	"time"
)

type PostService struct {
}

func (s PostService) CreatePost(request req.CreatePostRequest) bool {
	var post models.Posts
	post.Title = request.Title
	post.Content = request.Content
	post.UserId = request.UserId
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	db := utils.DBUtil{}.Connect()
	db.Create(&post)
	return true
}

func (s PostService) DetailPost(request req.DetailPostRequest) models.Posts {
	var post models.Posts
	db := utils.DBUtil{}.Connect()
	db.Where("id = ?", request.Id).First(&post)
	return post
}

// DeletePost 删除帖子
func (s PostService) DeletePost(request req.DetailPostRequest) bool {
	db := utils.DBUtil{}.Connect()
	//根据id查询帖子
	var post models.Posts
	db.Where("id = ?", request.Id).First(&post)
	if post.Id == 0 {
		return false
	}
	//判断用户是否是帖子的作者
	if post.UserId != request.UserId {
		return false
	}
	return db.Where("id = ?", request.Id).Where("user_id = ?", request.UserId).Delete(&models.Posts{}).Error == nil
}

// ListPost 获取帖子列表
func (s PostService) ListPost(request req.ListPostRequest) []models.Posts {
	var posts []models.Posts
	db := utils.DBUtil{}.Connect()
	db.Limit(request.PageSize).Offset((request.PageNo - 1) * request.PageSize).Find(&posts)
	return posts
}

// UpdatePost 更新帖子
func (s PostService) UpdatePost(request req.UpdatePostRequest) bool {
	db := utils.DBUtil{}.Connect()
	var post models.Posts
	db.Where("id = ?", request.Id).Find(&post)
	//判断post是否存在
	if post.Id == 0 {
		return false
	}
	//判断用户是否是作者
	if post.UserId != request.UserId {
		return false
	}
	//更新post
	db.Model(&post).Update("title", request.Title).Update("content", request.Content).Update("updated_at", time.Now())
	return true
}
