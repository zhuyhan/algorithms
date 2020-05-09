package list

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head
	for curr != nil {
		if curr.Next != nil && curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}
