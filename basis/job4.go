package basis

import "fmt"

// 4.查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if j >= len(strs[i]) || strs[i][j] != prefix[j] {
				prefix = prefix[:j]
				break
			}
		}
	}
	return prefix
}
func Job4Test() {
	strs1 := []string{"flower", "flow", "flight"}
	fmt.Println(LongestCommonPrefix(strs1)) // fl
	strs2 := []string{"dog", "racecar", "car"}
	fmt.Println(LongestCommonPrefix(strs2)) // ""
	strs3 := []string{"a"}
	fmt.Println(LongestCommonPrefix(strs3)) // a

}
