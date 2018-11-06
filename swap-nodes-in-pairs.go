package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	curr := head
	var newHead *ListNode
	var prev *ListNode

	if curr != nil && curr.Next != nil {
		newHead = curr.Next
	} else {
		newHead = head
	}

	for curr != nil && curr.Next != nil {
		first := curr
		second := curr.Next
		nextCur := second.Next

		first.Next = second.Next
		second.Next = first

		if prev != nil {
			prev.Next = second
		}

		curr = nextCur
		prev = first
	}

	return newHead
}

func main() {

}
