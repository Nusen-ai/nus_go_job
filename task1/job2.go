package task1

// 2. 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

import (
	"fmt"
	"strconv"
)

// 方法一：字符串方法
func isPalindromeString(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

// 方法二：数学方法
func isPalindromeMath(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return x == reversed || x == reversed/10
}

func Job2TestPalindromeString() {
	testCases := []int{121, -121, 10, 12321, 0, 1221, 123}

	fmt.Println("测试字符串方法:")
	for _, num := range testCases {
		fmt.Printf("%d: %t\n", num, isPalindromeString(num))
	}

	fmt.Println("\n测试数学方法:")
	for _, num := range testCases {
		fmt.Printf("%d: %t\n", num, isPalindromeMath(num))
	}

}

// 测试字符串方法:
// 121: true
// -121: false
// 10: false
// 12321: true
// 0: true
// 1221: true
// 123: false

// 测试数学方法:
// 121: true
// -121: false
// 10: false
// 12321: true
// 0: true
// 1221: true
// 123: false
