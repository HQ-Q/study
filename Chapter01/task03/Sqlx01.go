package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("连接数据库失败")
		panic(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	fmt.Println("连接数据库成功")
	return nil
}

// id INT AUTO_INCREMENT PRIMARY KEY,  -- 主键，自增，唯一标识员工
//
//	name VARCHAR(100) NOT NULL,         -- 员工姓名，非空，最长100字符
//	department VARCHAR(50),             -- 部门名称，最长50字符，允许为空
//	salary DECIMAL(10, 2) NOT NULL
type Employee struct {
	ID         int             `db:"id" json:"id"`
	Name       string          `db:"name" json:"name"`
	Department string          `db:"department" json:"department"`
	Salary     decimal.Decimal `db:"salary" json:"salary"`
}

// sqlx 入门
func main() {
	initDB()
	var emps []Employee = make([]Employee, 100)
	//查询所有员工数据
	err := db.Select(&emps, "select * from employees")
	if err != nil {
		fmt.Println("查询失败")
		panic(err)
		return
	}
	for _, emp := range emps {
		fmt.Printf("id:%d name:%s department:%s salary:%s\n", emp.ID, emp.Name, emp.Department, emp.Salary.String())
	}
	//查询技术部的所有员工数据
	var emps2 []Employee = make([]Employee, 100)
	err = db.Select(&emps2, "select * from employees where department = ?", "技术部")
	if err != nil {
		fmt.Println("查询失败")
		panic(err)
		return
	}
	fmt.Println("----------------技术部员工数据：---------------")
	for _, emp := range emps2 {
		fmt.Printf("id:%d name:%s department:%s salary:%s\n", emp.ID, emp.Name, emp.Department, emp.Salary.String())
	}

	//工资最高的员工信息，并将结果映射到一个 Employee 结构体中
	var emp Employee
	err = db.Get(&emp, "select * from employees where salary = (select max(salary) from employees)")
	if err != nil {
		fmt.Println("查询失败")
		panic(err)
		return
	}
	fmt.Println("----------------工资最高的员工信息：---------------")
	fmt.Printf("id:%d name:%s department:%s salary:%s\n", emp.ID, emp.Name, emp.Department, emp.Salary.String())

}
