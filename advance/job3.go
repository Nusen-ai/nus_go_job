package advance

import (
	"fmt"
	"sync"
)

func printOddNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("奇数: %d\n", i)
	}
}

func printEvenNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("偶数: %d\n", i)
	}
}

func Job3TestDemo() {
	var wg sync.WaitGroup

	wg.Add(2)

	// 启动奇数打印协程
	go printOddNumbers(&wg)

	// 启动偶数打印协程
	go printEvenNumbers(&wg)

	// 等待所有协程完成
	wg.Wait()

	fmt.Println("所有数字打印完成!")
}
