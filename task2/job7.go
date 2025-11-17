package task2

// 7. 编写一个程序，使用通道实现两个协程之间的通信。
// 一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。

import (
	"fmt"
	"sync"
)

func Job7TestDemo() {
	// 创建一个整型通道，用于协程间通信
	ch := make(chan int)

	// 使用WaitGroup来等待所有协程完成
	var wg sync.WaitGroup
	wg.Add(2) // 等待两个协程

	// 生产者协程：生成1到10的整数并发送到通道
	go func() {
		defer wg.Done() // 协程结束时通知WaitGroup
		defer close(ch) // 关闭通道，通知接收方没有更多数据

		for i := 1; i <= 10; i++ {
			fmt.Printf("发送: %d\n", i)
			ch <- i // 发送整数到通道
		}
		fmt.Println("生产者完成")
	}()

	// 消费者协程：从通道接收整数并打印
	go func() {
		defer wg.Done() // 协程结束时通知WaitGroup

		for num := range ch { // 循环接收直到通道关闭
			fmt.Printf("接收: %d\n", num)
		}
		fmt.Println("消费者完成")
	}()

	// 等待两个协程都完成
	wg.Wait()
	fmt.Println("程序结束")
}
