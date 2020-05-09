package list

// 206. 反转链表
// https://leetcode-cn.com/problems/reverse-linked-list/

//方法一：就第逆序
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead *ListNode

	for head.Next != nil {
		tmp := head.Next
		head.Next = newHead
		newHead = head
		head = tmp
	}

	return newHead

}

//方法二：递归
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
