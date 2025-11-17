package advance

import (
	"fmt"
	"sync"
)

// 9. 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func Job9TestDemo() {
	var (
		counter int
		mutex   sync.Mutex
		wg      sync.WaitGroup
	)

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 每个协程对计数器进行1000次递增操作
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	// 等待所有协程完成
	wg.Wait()

	// 输出计数器的值
	fmt.Printf("最终计数器的值: %d\n", counter)
}
