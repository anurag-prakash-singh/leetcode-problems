package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseN(head *ListNode, rem int, prev *ListNode, sentinel **ListNode) *ListNode {
	if head == nil || rem < 0 {
		return nil
	}

	if rem == 0 {
		*sentinel = head.Next
		head.Next = prev
		return head
	}

	remCurrHead := head.Next
	head.Next = prev

	return reverseN(remCurrHead, rem-1, head, sentinel)
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	start := head
	var prev *ListNode

	for i := 0; i < m-1; i++ {
		prev = start
		start = start.Next
	}

	var sentinel *ListNode

	mnRevHead := reverseN(start, n-m, prev, &sentinel)

	if prev != nil {
		prev.Next = mnRevHead
	}

	start.Next = sentinel

	if prev != nil {
		return head
	} else {
		return mnRevHead
	}
}

func printList(head *ListNode) {
	fmt.Printf("List: ")

	for ; head != nil; head = head.Next {
		fmt.Printf("%v", head.Val)

		if head.Next != nil {
			fmt.Printf(" -> ")
		}
	}

	fmt.Println()
}

func test1() {
	n5 := &ListNode{5, nil}
	n4 := &ListNode{4, n5}
	n3 := &ListNode{3, n4}
	n2 := &ListNode{2, n3}
	n1 := &ListNode{1, n2}

	mnReversed := reverseBetween(n1, 2, 4)

	printList(mnReversed)
}

func test2() {
	n5 := &ListNode{5, nil}
	n4 := &ListNode{4, n5}
	n3 := &ListNode{3, n4}
	n2 := &ListNode{2, n3}
	n1 := &ListNode{1, n2}

	mnReversed := reverseBetween(n1, 1, 3)

	printList(mnReversed)
}

func test3() {
	n5 := &ListNode{5, nil}
	n4 := &ListNode{4, n5}
	n3 := &ListNode{3, n4}
	n2 := &ListNode{2, n3}
	n1 := &ListNode{1, n2}

	mnReversed := reverseBetween(n1, 1, 5)

	printList(mnReversed)
}

func test4() {
	n5 := &ListNode{5, nil}
	n4 := &ListNode{4, n5}
	n3 := &ListNode{3, n4}
	n2 := &ListNode{2, n3}
	n1 := &ListNode{1, n2}

	mnReversed := reverseBetween(n1, 5, 5)

	printList(mnReversed)
}

func test5() {
	n1 := &ListNode{1, nil}

	mnReversed := reverseBetween(n1, 1, 1)

	printList(mnReversed)
}

func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	test5()
}
