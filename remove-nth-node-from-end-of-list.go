package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	endPtr := head

	for i := 1; i < n; i++ {
		endPtr = endPtr.Next
	}

	var curr *ListNode
	var prev *ListNode

	for endPtr != nil {
		endPtr = endPtr.Next
		prev = curr

		if curr == nil {
			curr = head
		} else {
			curr = curr.Next
		}
	}

	if curr == head {
		return head.Next
	}

	prev.Next = curr.Next

	return head
}

func main() {

}
