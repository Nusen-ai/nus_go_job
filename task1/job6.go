package task1

import "fmt"

// 6. 删除有序数组中的重复项
// 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
// 元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
// 考虑 nums 的唯一元素的数量为 k。去重后，返回唯一元素的数量 k。
// nums 的前 k 个元素应包含 排序后 的唯一数字。下标 k - 1 之后的剩余元素可以忽略。
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 使用双指针，i 指向当前不重复元素的位置
	i := 0
	for j := 1; j < len(nums); j++ {
		// 如果发现不同的元素
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	// 返回唯一元素的数量
	return i + 1
}
func Job6Test() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	nums1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums1))

}
