package task1

import "fmt"

// 3.给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func isValid(s string) bool {
	// 使用栈来解决括号匹配问题
	stack := make([]rune, 0)

	// 定义括号映射关系
	bracketMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	// 遍历字符串中的每个字符
	for _, char := range s {
		// 如果是右括号
		if matchingLeft, isRight := bracketMap[char]; isRight {
			// 检查栈是否为空或者栈顶元素不匹配
			if len(stack) == 0 || stack[len(stack)-1] != matchingLeft {
				return false
			}
			// 匹配成功，弹出栈顶元素
			stack = stack[:len(stack)-1]
		} else {
			// 如果是左括号，压入栈中
			stack = append(stack, char)
		}
	}

	// 最后栈应该为空才表示所有括号都正确匹配
	return len(stack) == 0
}

func Job3Test() {

	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid("{[}"))
	fmt.Println(isValid("{}[]"))
	fmt.Println(isValid("{}}}"))

}
