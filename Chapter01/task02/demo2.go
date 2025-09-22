package main

import (
	"fmt"
	"sync"
)

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func printJiShu() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("奇数：", i)
	}
	wait.Done()
}

func printOuShu() {
	for i := 2; i <= 10; i += 2 {
		fmt.Println("偶数：", i)
	}
	wait.Done()
}

var wait sync.WaitGroup

func main() {

	wait.Add(1)
	go printJiShu()
	wait.Add(1)
	go printOuShu()
	wait.Wait()

}
