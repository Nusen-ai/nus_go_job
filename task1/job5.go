package task1

import "fmt"

// 5.给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
// 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。
// 这些数字按从左到右，从最高位到最低位排列。
// 这个大整数不包含任何前导 0。
// 将大整数加 1，并返回结果的数字数组。
func plusOne(digits []int) []int {
	n := len(digits)

	// 从最后一位开始向前遍历
	for i := n - 1; i >= 0; i-- {
		// 当前位加1
		digits[i]++
		// 如果加1后小于10，没有进位，直接返回
		if digits[i] < 10 {
			return digits
		}
		// 如果等于10，产生进位，当前位置0
		digits[i] = 0
	}

	// 如果循环结束后还有进位（即第一位也产生了进位）
	// 需要在数组最前面添加1
	return append([]int{1}, digits...)
}

func Job5Test() {
	fmt.Println(plusOne([]int{1, 2, 3}))    // [1 2 4]
	fmt.Println(plusOne([]int{4, 3, 2, 1})) // [4 3 2 2]
	fmt.Println(plusOne([]int{9}))          // [1 0]
	fmt.Println(plusOne([]int{9, 9}))       // [1 0 0]
}
