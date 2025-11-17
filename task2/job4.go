package task2

import (
	"fmt"
	"sync"
	"time"
)

// 4. 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。

// Task 定义任务类型
type Task func()

// TaskResult 任务执行结果
type TaskResult struct {
	TaskID    int
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
	Error     error
}

// TaskScheduler 任务调度器
type TaskScheduler struct {
	tasks   []Task
	results []TaskResult
	wg      sync.WaitGroup
	mutex   sync.Mutex
}

// NewTaskScheduler 创建新的任务调度器
func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks:   make([]Task, 0),
		results: make([]TaskResult, 0),
	}
}

// AddTask 添加任务
func (ts *TaskScheduler) AddTask(task Task) {
	ts.tasks = append(ts.tasks, task)
}

// executeTask 执行单个任务并记录时间
func (ts *TaskScheduler) executeTask(taskID int, task Task) {
	defer ts.wg.Done()

	startTime := time.Now()

	// 执行任务
	task()

	endTime := time.Now()
	duration := endTime.Sub(startTime)

	// 记录结果
	result := TaskResult{
		TaskID:    taskID,
		StartTime: startTime,
		EndTime:   endTime,
		Duration:  duration,
	}

	ts.mutex.Lock()
	ts.results = append(ts.results, result)
	ts.mutex.Unlock()
}

// Run 并发执行所有任务
func (ts *TaskScheduler) Run() {
	ts.results = make([]TaskResult, 0)

	for i, task := range ts.tasks {
		ts.wg.Add(1)
		go ts.executeTask(i+1, task)
	}

	ts.wg.Wait()
}

// PrintResults 打印执行结果
func (ts *TaskScheduler) PrintResults() {
	fmt.Println("\n=== 任务执行统计 ===")
	fmt.Printf("总任务数: %d\n", len(ts.results))
	fmt.Println("----------------------------")

	var totalDuration time.Duration
	for _, result := range ts.results {
		fmt.Printf("任务 %d:\n", result.TaskID)
		fmt.Printf("  开始时间: %s\n", result.StartTime.Format("15:04:05.000"))
		fmt.Printf("  结束时间: %s\n", result.EndTime.Format("15:04:05.000"))
		fmt.Printf("  执行时间: %v\n", result.Duration)
		fmt.Println("  --------------------")
		totalDuration += result.Duration
	}

	fmt.Printf("总执行时间: %v\n", totalDuration)
}

// 示例任务函数
func sampleTask1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("任务1执行完成")
}

func sampleTask2() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("任务2执行完成")
}

func sampleTask3() {
	time.Sleep(150 * time.Millisecond)
	fmt.Println("任务3执行完成")
}

func sampleTask4() {
	time.Sleep(50 * time.Millisecond)
	fmt.Println("任务4执行完成")
}

func Job4TestDemo() {

	fmt.Println("\n=== 任务调度器 ===")

	// 创建任务调度器
	scheduler := NewTaskScheduler()

	// 添加任务
	scheduler.AddTask(sampleTask1)
	scheduler.AddTask(sampleTask2)
	scheduler.AddTask(sampleTask3)
	scheduler.AddTask(sampleTask4)

	// 添加一些计算密集型任务
	scheduler.AddTask(func() {
		sum := 0
		for i := 0; i < 1000000; i++ {
			sum += i
		}
		fmt.Printf("计算任务完成，结果: %d\n", sum)
	})

	// 执行所有任务
	fmt.Println("开始执行所有任务...")
	start := time.Now()
	scheduler.Run()
	totalTime := time.Since(start)

	// 打印结果
	scheduler.PrintResults()
	fmt.Printf("\n实际总耗时(并发): %v\n", totalTime)
}
