package middleoftheLinkedList

func MiddleNodeNMethod1(head *ListNode) *ListNode {

	listNode := head
	l := 0
	for listNode != nil {
		l++
		listNode = listNode.Next
	}

	listNode = head
	for i := 0; i < l/2; i++ {
		listNode = listNode.Next
	}

	return listNode
}

func MiddleNodeNMethod2(head *ListNode) *ListNode {

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
