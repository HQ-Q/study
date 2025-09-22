package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type Books struct {
	Id     int             `db:"id"`
	Title  string          `db:"title"`
	Author string          `db:"author"`
	Price  decimal.Decimal `db:"price"`
}

var DB *sqlx.DB

func InitDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test"
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("连接数据库失败")
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	fmt.Println("连接数据库成功")
	return nil
}

// 使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
func main() {
	err := InitDB()
	if err != nil {
		panic(err)
	}
	var books = make([]Books, 100)
	err = DB.Select(&books, "select * from books where price > ?", decimal.NewFromFloat(50))
	if err != nil {
		fmt.Println("查询失败")
		panic(err)
		return
	}
	for _, book := range books {
		fmt.Printf("id:%d title:%s author:%s price:%s\n", book.Id, book.Title, book.Author, book.Price.String())
	}
}
