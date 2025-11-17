package advance

import "fmt"

// 2. 接收整数切片指针并将每个元素乘以2
func doubleSlice(slicePtr *[]int) {
	// 通过指针访问切片
	for i := 0; i < len(*slicePtr); i++ {
		(*slicePtr)[i] *= 2
	}
}

func Job2TestDemo() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("修改前的切片: %v\n", numbers)

	doubleSlice(&numbers)
	fmt.Printf("修改后的切片: %v\n", numbers)
}
