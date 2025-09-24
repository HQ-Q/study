package services

import (
	"demo/Chapter01/task04/models"
	"demo/Chapter01/task04/req"
	"demo/Chapter01/task04/utils"
	"time"
)

type CommentService struct {
}

func (s CommentService) CreateComment(request req.CreateCommentRequest) bool {
	db := utils.DBUtil{}.Connect()
	db.Create(&models.Comments{
		Content:   request.Content,
		UserId:    request.UserId,
		PostId:    request.PostId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	return true
}

func (s CommentService) GetCommentByPostId(request req.GetCommentByPostIdRequest) interface{} {
	db := utils.DBUtil{}.Connect()
	var comments []models.Comments
	db.Preload("Author").Preload("Post").Preload("ToUser").Preload("Post.Author").Where("post_id = ?", request.PostId).Find(&comments)
	return comments
}
