package task1

import (
	"fmt"
	"sort"
)

// 7.合并重叠的区间
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 按区间的起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}
	current := intervals[0]

	for i := 1; i < len(intervals); i++ {
		// 如果当前区间与下一个区间有重叠
		if current[1] >= intervals[i][0] {
			// 合并区间，取结束位置的最大值
			if intervals[i][1] > current[1] {
				current[1] = intervals[i][1]
			}
		} else {
			// 没有重叠，将当前区间加入结果
			result = append(result, current)
			current = intervals[i]
		}
	}

	// 添加最后一个区间
	result = append(result, current)

	return result
}

func Job7Test() {
	// 测试示例
	intervals1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Printf("输入: %v\n", intervals1)        // 输入: [[1,3],[2,6],[8,10],[15,18]]
	fmt.Printf("输出: %v\n", merge(intervals1)) // 输入: [[1,6],[8,10],[15,18]]

	intervals2 := [][]int{{1, 4}, {4, 5}}
	fmt.Printf("输入: %v\n", intervals2)        // 输入: [[1,4],[4,5]]
	fmt.Printf("输出: %v\n", merge(intervals2)) // 输入: [[1,5]]

	intervals3 := [][]int{{4, 7}, {1, 4}}
	fmt.Printf("输入: %v\n", intervals3)        // 输入: [[4,7],[1,4]]
	fmt.Printf("输出: %v\n", merge(intervals3)) // 输入: [[1,7]]
}
