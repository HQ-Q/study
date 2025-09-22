package main

import (
	"fmt"
	"sync"
)

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
//启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

var counter int
var sy sync.Mutex
var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				sy.Lock()
				counter++
				fmt.Println(counter)
				sy.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
