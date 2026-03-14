package main

import (
	"fmt"
	"sync"
	"time"
)

//2 协程

// 定义任务结构体
type Task struct {
	ID       int
	Name     string
	Function func() error
}

// 定义任务结果结构体
type TaskResult struct {
	TaskID    int
	TaskName  string
	Duration  time.Duration
	Err       error
	startTime time.Time
	endTime   time.Time
}

// 定义调度器
type TaskScheduler struct {
	Tasks       []Task
	TaskResults chan TaskResult
	wg          sync.WaitGroup
}

// 创建调度器
func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		Tasks:       make([]Task, 0),
		TaskResults: make(chan TaskResult, 100),
	}
}

// 添加任务
func (ts *TaskScheduler) AddTask(ID int, Name string, Fn func() error) {
	task := &Task{
		ID:       ID,
		Name:     Name,
		Function: Fn,
	}
	ts.Tasks = append(ts.Tasks, *task)
}

// 执行单个任务
func (ts *TaskScheduler) ExecuteTask(task Task) {
	defer ts.wg.Done()
	startTime := time.Now()
	err := task.Function()
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	tr := &TaskResult{
		TaskID:    task.ID,
		TaskName:  task.Name,
		Duration:  duration,
		Err:       err,
		startTime: startTime,
		endTime:   endTime,
	}
	ts.TaskResults <- *tr
}

// 执行所有任务
func (ts *TaskScheduler) ExecuteAll() {
	for _, task := range ts.Tasks {
		ts.wg.Add(1)
		go ts.ExecuteTask(task)
	}
	ts.wg.Wait()
	close(ts.TaskResults)
	fmt.Println("所有任务执行完毕")
}

func task1() error {
	time.Sleep(1 * time.Second)
	fmt.Println("task 1 正在执行")
	return nil
}

func task2() error {
	time.Sleep(2 * time.Second)
	fmt.Println("task 2 正在执行")
	return nil
}

func task3() error {
	time.Sleep(3 * time.Second)
	fmt.Println("task 3 正在执行")
	return fmt.Errorf("task 3 执行失败")
}

func task4() error {
	time.Sleep(2 * time.Second)
	fmt.Println("task 4 正在执行")
	return nil
}

func task5() error {
	time.Sleep(1 * time.Second)
	fmt.Println("task 5 正在执行")
	return nil
}
func main() {

	// 	2.2 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
	// 考察点 ：协程原理、并发任务调度。
	scheduler := NewTaskScheduler()
	scheduler.AddTask(1, "数据处理任务", task1)
	scheduler.AddTask(2, "网络请求任务", task2)
	scheduler.AddTask(3, "文件上传任务", task3)
	scheduler.AddTask(4, "数据库查询任务", task4)
	scheduler.AddTask(5, "邮件发送任务", task5)
	start := time.Now()
	scheduler.ExecuteAll()
	end := time.Now()
	totalTime := end.Sub(start)
	fmt.Println("总共耗时", totalTime)
	fmt.Println()
	fmt.Println()
	fmt.Println("---------------------------打印任务报告-----------------------------")
	for taskResult := range scheduler.TaskResults {
		fmt.Println()
		status := "成功"
		if taskResult.Err != nil {
			status = fmt.Sprintf("失败(%v)", taskResult.Err)
		}
		fmt.Printf("任务ID=%d |名称=%s\n",
			taskResult.TaskID, taskResult.TaskName)
		fmt.Printf("开始时间=%s\n",
			taskResult.startTime.Format("2006-01-02 15:04:05"))
		fmt.Printf("结束时间=%s\n",
			taskResult.endTime.Format("2006-01-02 15:04:05"))
		fmt.Printf("持续时间=%d秒\n",
			int(taskResult.Duration.Seconds()))
		fmt.Printf("状态=%s\n", status)
		fmt.Println()
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	}

}
