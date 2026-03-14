package main

import (
	"fmt"
	"sync"
)

//5 锁机制:实现并发数据安全

var (
	wg sync.WaitGroup
)

// 定义Counter结构体
type Counter struct {
	mu    sync.Mutex
	count int
}

// 定义getCount()方法
func (c *Counter) getCount() int {
	defer c.mu.Unlock()
	c.mu.Lock()
	return c.count
}

// 定义increment()方法
func (c *Counter) increment() {
	defer c.mu.Unlock()
	c.mu.Lock()
	c.count++
}

func main() {
	//5.1 第1种方式:使用sync.mutex有锁方式
	c := &Counter{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 1; i <= 1000; i++ {
				c.increment()
			}
		}()
	}
	wg.Wait()
	fmt.Println("预计结果", 10*1000)
	fmt.Println("实际结果", c.getCount())
}
