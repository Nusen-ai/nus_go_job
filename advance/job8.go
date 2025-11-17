package advance

import (
	"fmt"
	"sync"
)

// 8.实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
func Job8TestDemo() {

	// 创建一个缓冲大小为10的通道
	ch := make(chan int, 10)

	// 使用WaitGroup来等待所有goroutine完成
	var wg sync.WaitGroup

	// 生产者协程 - 发送100个整数
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch) // 发送完成后关闭通道

		for i := 1; i <= 100; i++ {
			fmt.Printf("生产者发送: %d\n", i)
			ch <- i // 发送整数到通道
		}
		fmt.Println("生产者完成")
	}()

	// 消费者协程 - 接收并打印整数
	wg.Add(1)
	go func() {
		defer wg.Done()

		for num := range ch { // 循环接收直到通道关闭
			fmt.Printf("消费者接收: %d\n", num)
		}
		fmt.Println("消费者完成")
	}()

	// 等待所有goroutine完成
	wg.Wait()
	fmt.Println("程序结束")
}
