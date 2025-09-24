package controller

import (
	"demo/Chapter01/task04/models"
	"demo/Chapter01/task04/services"
	"fmt"
	"net/http"

	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func UserControllerInit(r *gin.Engine) {
	v1 := r.Group("user")
	{
		v1.POST("/register", Register)
		v1.POST("/login", Login)
		v1.GET("/logout", AuthMiddleware(), Logout)
		v1.GET("/getUserInfo", AuthMiddleware(), GetUserInfo)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		// 解析token
		ok, user := services.UserService{}.ParseToken(token)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
		}
		// 将解析出的用户信息存入上下文，供后续处理使用
		c.Set("userId", user.Id)
		c.Set("username", user.Username)
		c.Set("email", user.Email)
		c.Set("token", token)
		c.Next()
	}
}

// Register 注册
func Register(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ok, msg := services.UserService{}.Register(user)
	if ok {
		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": msg})
	}
}

// Login 登录
func Login(c *gin.Context) {
	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ok, msg := services.UserService{}.Login(user.Username, user.Password)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successfully",
			"token":   msg,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}
}

// Logout 登出
func Logout(c *gin.Context) {
	//删除 token
	token, _ := c.Get("token")
	//token转为字符串
	tokenStrings := fmt.Sprint(token)
	ok, err := services.UserService{}.Logout(tokenStrings)
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"message": "Logout successfully",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {

	userId, _ := c.Get("userId")
	username, _ := c.Get("username")
	email, _ := c.Get("email")

	c.JSON(http.StatusOK, gin.H{
		"message":  "Get user info successfully",
		"userId":   userId,
		"username": username,
		"email":    email,
	})
}
