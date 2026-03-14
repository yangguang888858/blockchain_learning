package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//5 锁机制:实现并发数据安全

var (
	wg sync.WaitGroup
)

func main() {
	//5.2 第2种方式:使用atomic原子方式(无锁)
	var count int64
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 1; i <= 1000; i++ {
				atomic.AddInt64(&count, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("预计结果", 10*1000)
	fmt.Println("实际结果", count)
}
