package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

func main() {
	var wg sync.WaitGroup
	var atomicCounter int32 = 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&atomicCounter, 1)
				//fmt.Println(atomicCounter)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(atomicCounter)
}
