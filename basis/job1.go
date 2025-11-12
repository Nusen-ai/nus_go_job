package basis

import "fmt"

// 只出现一次的数字：
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
// 找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。

// 方案一：异或运算
func singleNumberXOR(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	fmt.Println(result)
	return result
}

// 方案二：map
func singleNumberMap(nums []int) int {
	countMap := make(map[int]int)

	// 第一次遍历：统计每个数字的出现次数
	for _, num := range nums {
		// fmt.Println(num)
		countMap[num]++
		// fmt.Printf("%+v\n", countMap)
	}

	// 第二次遍历：找出出现一次的数字
	for num, count := range countMap {
		if count == 1 {
			fmt.Println(num)
			return num
		}
	}

	return -1 // 理论上不会执行到这里
}
func singleNumberMap1(nums []int) int {
	countMap := make(map[int]int)

	for _, num := range nums {
		if _, exists := countMap[num]; exists {
			delete(countMap, num)
		} else {
			countMap[num] = 1
		}
	}

	// 最后map中剩下的就是只出现一次的元素
	for num := range countMap {
		return num
	}

	return -1
}
func Job1TestDemo() {
	var list []int = []int{2, 2, 1}
	singleNumberXOR(list)
	singleNumberMap(list)
	singleNumberMap1(list)

}
