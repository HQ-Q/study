package utils

import "github.com/jinzhu/gorm"

type DBUtil struct {
}

// Connect 连接数据库
func (u DBUtil) Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
}

// Close 关闭数据库
func (u DBUtil) Close(db *gorm.DB) {
	db.Close()
}
