package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func listLen(head *ListNode) int {
	n := 0

	for head != nil {
		head = head.Next
		n++
	}

	return n
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	n := listLen(head)
	var newHead *ListNode
	origHead := head

	if n == 0 {
		return head
	}

	k = k % n

	if k == 0 {
		return head
	}

	var lastNode *ListNode

	for i := 0; i < n; i++ {
		lastNode = head

		if i == n-k-1 {
			newHead = head.Next
			head.Next = nil
			head = newHead
		} else {
			head = head.Next
		}
	}

	lastNode.Next = origHead

	return newHead
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d", head.Val)

		if head.Next != nil {
			fmt.Printf(" ")
		}

		head = head.Next
	}

	fmt.Println()
}

func test1() {
	n3 := &ListNode{3, nil}
	n2 := &ListNode{2, n3}
	n1 := &ListNode{1, n2}

	rotated := rotateRight(n1, 1)
	printList(rotated)
}

func test2() {
	n3 := &ListNode{3, nil}
	n2 := &ListNode{2, n3}
	n1 := &ListNode{1, n2}

	rotated := rotateRight(n1, 2)
	printList(rotated)
}

func test3() {
	n3 := &ListNode{3, nil}
	n2 := &ListNode{2, n3}
	n1 := &ListNode{1, n2}

	rotated := rotateRight(n1, 3)
	printList(rotated)
}

func main() {
	test1()
	test2()
	test3()
}
