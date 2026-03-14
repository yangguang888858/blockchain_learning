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

// 4.2 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。

func produce(dataChan chan<- int) {
	defer wg.Done()
	//关闭管道,防止死锁
	defer close(dataChan)
	for i := 1; i <= 100; i++ {
		time.Sleep(100 * time.Millisecond)
		dataChan <- i
		// fmt.Println("写入数据", i)

	}
}

func consume(dataChan <-chan int) {
	defer wg.Done()
	for v := range dataChan {
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("读取到的数据%d\n", v)
	}
}
func main() {

	//4.2 带缓冲的通道
	//4.2.1 若没有消费者读数据,且缓冲区大小<实际数据大小,写会阻塞
	//4.2.2 若没有消费者读数据,且缓冲区大小>=实际数据大小,写不会阻塞
	//4.2.3 若有消费者读数据,且缓冲区大小<实际数据大小,读数据的速度>=写数据的速度,写不会阻塞;读数据的速度<写数据的速度,写会阻塞
	//4.2.4 若有消费者读数据,且缓冲区大小>=实际数据大小,无论读数据的速度如何,写都不会阻塞
	//总之,缓冲区满了写就会阻塞,缓冲区空了读就会阻塞

	var dataChan = make(chan int, 100)
	wg.Add(1)
	go produce(dataChan)
	wg.Add(1)
	go consume(dataChan)
	wg.Wait()
	fmt.Println("主线程退出...")
}
