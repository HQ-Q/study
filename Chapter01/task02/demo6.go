package main

import (
	"fmt"
	"sync"
)

// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印

var wa sync.WaitGroup

func producer(ch chan int) {
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Println("生产者协程发送数据:", i)
	}
	close(ch)
	wa.Done()
}

func consumer(ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("消费者协程接收数据:", i)
	}
	wa.Done()
}
func main() {
	ch := make(chan int, 100)
	wa.Add(1)
	go producer(ch)
	wa.Add(1)
	go consumer(ch)
	wa.Wait()
}
