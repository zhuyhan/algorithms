package list

// 141. 环形链表
// https://leetcode-cn.com/problems/linked-list-cycle/
// 快慢自增方法;当快指针=慢指针，说明有环
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for fast != nil || slow != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
