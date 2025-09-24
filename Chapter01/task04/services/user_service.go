package services

import (
	"demo/Chapter01/task04/models"
	"demo/Chapter01/task04/utils"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var TokenMap = make(map[string]string)

type UserService struct {
}

// Register 注册
func (s UserService) Register(user models.Users) (bool, string) {
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return false, "System error"
	}
	user.Password = string(hashedPassword)
	db := utils.DBUtil{}.Connect()
	var existingUser models.Users
	db.Where("username = ?", user.Username).Find(&existingUser)
	if existingUser.Id != 0 {
		return false, "Username already exists"
	}
	db.Where("email = ?", user.Email).Find(&existingUser)
	if existingUser.Id != 0 {
		return false, "Email already exists"
	}
	if err := db.Create(&user).Error; err != nil {
		return false, "Failed to create user"
	}
	utils.DBUtil{}.Close(db)
	return true, "Success"
}

// Login 登录
func (s UserService) Login(username, password string) (bool, string) {
	db := utils.DBUtil{}.Connect()
	var user models.Users
	db.Where("username = ?", username).Find(&user)
	var storedUser models.Users
	if err := db.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		return false, "Invalid username or password"
	}
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(password)); err != nil {
		return false, "Invalid username or password"
	}
	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.Id,
		"username": storedUser.Username,
		"email":    storedUser.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte("abcdefg"))
	if err != nil {
		return false, "Failed to generate token"
	}
	// 保存token 到 map
	TokenMap[storedUser.Username] = tokenString
	return true, tokenString
}

// ParseToken 解析JWT token
func (s UserService) ParseToken(tokenString string) (bool, models.Users) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("abcdefg"), nil
	})
	var user models.Users
	if err != nil {
		return false, user
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := int(claims["id"].(float64))
		username := claims["username"].(string)
		email := claims["email"].(string)

		// 判断token 是否一致
		existToken := TokenMap[username]
		if tokenString != existToken {
			return false, user
		}
		return true, models.Users{
			Id:       id,
			Username: username,
			Email:    email,
		}
	}
	return false, user
}

// Logout 注销登录 删除token
func (s UserService) Logout(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("abcdefg"), nil
	})
	if err != nil {
		return false, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		// 判断token 是否一致
		existToken := TokenMap[username]
		if tokenString == existToken {
			delete(TokenMap, username)
			return ok, nil
		}
	}
	return false, nil
}
