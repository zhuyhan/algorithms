package list

// 2. 两数相加
// https://leetcode-cn.com/problems/add-two-numbers/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//	初始化结果队列
	sumList := new(ListNode)
	node1, node2, node3 := l1, l2, sumList
	plus := 0 //进1
	for node1 != nil || node2 != nil || plus > 0 {
		var num1, num2 int
		if node1 != nil {
			num1 = node1.Val
		}
		if node2 != nil {
			num2 = node2.Val
		}
		sum := num1 + num2 + plus
		//判断是否进一
		plus = sum / 10
		sum = sum % 10

		sumNode := &ListNode{Val: sum}
		node3.Next = sumNode
		if node1 != nil {
			node1 = node1.Next
		}
		if node2 != nil {
			node2 = node2.Next
		}
		node3 = node3.Next
	}
	return sumList.Next
}
