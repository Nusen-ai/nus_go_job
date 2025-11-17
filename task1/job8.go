package task1

import "fmt"

// 8.给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func twoSum(nums []int, target int) []int {
	// 创建一个哈希表来存储数字和对应的索引
	numMap := make(map[int]int)

	for i, num := range nums {
		// 计算需要的补数
		complement := target - num

		// 检查补数是否在哈希表中
		if idx, exists := numMap[complement]; exists {
			return []int{idx, i}
		}

		// 将当前数字和索引存入哈希表
		numMap[num] = i
	}

	// 根据题目假设，总会有一个有效答案，所以这里不会执行到
	return nil
}

func Job8Test() {
	nums := []int{2, 7, 11, 15}
	target := 13
	fmt.Println(twoSum(nums, target)) // [0, 2]
}
