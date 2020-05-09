package stack

//20. 有效的括号
//https://leetcode-cn.com/problems/valid-parentheses/

//法一：用go自带的链表实现

//法二：map实现
func isValid(s string) bool {
	var stack []byte
	for _, v := range s {
		switch byte(v) {
		case '(', '{', '[':
			stack = append(stack, byte(v))
		case ')', '}', ']':
			if len(stack) > 0 && stack[len(stack)-1] == match(byte(v)) {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func match(s byte) byte {
	switch s {
	case ')':
		s = '('
	case ']':
		s = '['
	case '}':
		s = '{'
	}
	return s
}
