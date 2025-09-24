package controller

import (
	"demo/Chapter01/task04/req"
	"demo/Chapter01/task04/services"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
}

func CommentControllerInit(r *gin.Engine) {
	group := r.Group("/comment")
	{
		group.POST("/create", AuthMiddleware(), CreateComment)
		group.POST("/getByPostId", AuthMiddleware(), GetCommentByPostId)
	}
}

func CreateComment(c *gin.Context) {
	req := req.CreateCommentRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	req.UserId = c.GetInt("userId")

	ok := services.CommentService{}.CreateComment(req)
	if ok {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "评论成功",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "评论失败",
		})
	}
}

func GetCommentByPostId(c *gin.Context) {
	req := req.GetCommentByPostIdRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	comments := services.CommentService{}.GetCommentByPostId(req)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": comments,
	})
}
