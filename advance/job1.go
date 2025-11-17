package advance

import "fmt"

// 1. 定义一个函数，接收整数指针并增加10
func addTen(num *int) {
	*num += 10
}

func Job1TestDemo() {
	value := 5
	fmt.Printf("修改前的值: %d\n", value) // 修改前的值: 5

	addTen(&value)
	fmt.Printf("修改后的值: %d\n", value) // 修改后的值: 15
}
