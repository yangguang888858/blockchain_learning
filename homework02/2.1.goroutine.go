package main

import (
	"fmt"
	"sync"
	"time"
)

//2 协程

var (
	wg sync.WaitGroup
)

// 打印奇数
func printJs() {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			time.Sleep(1 * time.Second)
			fmt.Println("打印奇数的协程<<<:", i)
		}
	}
}

// 打印偶数
func printOs() {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("打印偶数的协程>>>:", i)
		}
	}
}
func main() {
	// 2.1 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
	// 考察点 ： go 关键字的使用、协程的并发执行。
	wg.Add(1)
	go printJs()
	wg.Add(1)
	go printOs()
	wg.Wait()
	fmt.Println("主协程退出...")
}
