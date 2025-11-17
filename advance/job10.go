package advance

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 10.使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func Job10TestDemo() {
	// 使用int64类型的原子计数器
	var counter int64
	var wg sync.WaitGroup

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// 每个协程对计数器进行1000次递增操作
			for j := 0; j < 1000; j++ {
				// 使用原子操作递增计数器
				atomic.AddInt64(&counter, 1)
			}
			fmt.Printf("协程 %d 完成\n", id)
		}(i)
	}

	// 等待所有协程完成
	wg.Wait()

	// 输出最终结果
	fmt.Printf("最终计数器的值: %d\n", counter)

	// 验证结果是否正确
	expected := int64(10 * 1000)
	fmt.Printf("期望值: %d, 结果正确: %v\n", expected, counter == expected)
}
