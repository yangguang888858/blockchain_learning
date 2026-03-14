package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

//4 管道

// 4.1 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
func generateData(dataChan chan int) {
	defer wg.Done()
	//关闭管道,防止死锁
	defer close(dataChan)
	for i := 1; i <= 10; i++ {
		dataChan <- i
	}
}

func printData(dataChan chan int) {
	defer wg.Done()
	for v := range dataChan {
		fmt.Println(v)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// 4.1 不带缓冲的通道
	// 4.1.1 若没有消费者读数据,写会阻塞
	// 4.1.2 若有消费者读数据,无论读速度如何,写都不会阻塞
	var dataChan = make(chan int)
	wg.Add(1)
	go generateData(dataChan)
	wg.Add(1)
	go printData(dataChan)
	wg.Wait()
	fmt.Println("主线程退出...")
}
