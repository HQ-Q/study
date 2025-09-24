package controller

import (
	"demo/Chapter01/task04/req"
	"demo/Chapter01/task04/services"

	"github.com/gin-gonic/gin"
)

type PostController struct {
}

func PostControllerInit(r *gin.Engine) {
	post := r.Group("/post")
	{
		post.POST("/create", AuthMiddleware(), CreatePost)
		post.POST("/list", AuthMiddleware(), ListPost)
		post.POST("/detail", AuthMiddleware(), DetailPost)
		post.DELETE("/delete", AuthMiddleware(), DeletePost)
		post.PUT("/update", AuthMiddleware(), UpdatePost)
	}
}

// CreatePost 创建帖子
func CreatePost(c *gin.Context) {
	var request req.CreatePostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	request.UserId = c.GetInt("userId")
	ok := services.PostService{}.CreatePost(request)
	if ok {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "创建成功",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "创建失败",
		})
	}
}

// DetailPost 获取帖子详情
func DetailPost(c *gin.Context) {
	var request req.DetailPostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	post := services.PostService{}.DetailPost(request)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": post,
	})

}

// DeletePost 删除帖子
func DeletePost(c *gin.Context) {
	var request req.DetailPostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	request.UserId = c.GetInt("userId")
	ok := services.PostService{}.DeletePost(request)
	if ok {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "删除成功",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "删除失败",
		})
	}
}

// ListPost 获取帖子列表
func ListPost(c *gin.Context) {
	var request req.ListPostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	posts := services.PostService{}.ListPost(request)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": posts,
	})
}

// UpdatePost 修改帖子
func UpdatePost(c *gin.Context) {
	var request req.UpdatePostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	request.UserId = c.GetInt("userId")
	ok := services.PostService{}.UpdatePost(request)
	if ok {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "修改成功",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "修改失败",
		})

	}

}
