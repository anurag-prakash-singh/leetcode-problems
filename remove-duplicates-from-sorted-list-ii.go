package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	headNext := head.Next
	duplicatesFound := false

	for ; headNext != nil && headNext.Val == head.Val; headNext = headNext.Next {
		duplicatesFound = true
	}

	if duplicatesFound {
		return deleteDuplicates(headNext)
	} else {
		head.Next = deleteDuplicates(headNext)
	}

	return head
}

func printList(head *ListNode) {
	for ; head != nil; head = head.Next {
		fmt.Printf("%d", head.Val)

		if head.Next != nil {
			fmt.Printf(" -> ")
		}
	}

	fmt.Println()
}

func test1() {
	nodes := []ListNode{ListNode{1, nil}, ListNode{2, nil}, ListNode{3, nil}}
	nodes[0].Next = &nodes[1]
	nodes[1].Next = &nodes[2]

	printList(deleteDuplicates(&nodes[0]))
}

func test2() {
	nodes := []ListNode{ListNode{1, nil}, ListNode{2, nil}, ListNode{2, nil}}
	nodes[0].Next = &nodes[1]
	nodes[1].Next = &nodes[2]

	printList(deleteDuplicates(&nodes[0]))
}

func test3() {
	nodes := []ListNode{ListNode{1, nil}, ListNode{2, nil}, ListNode{2, nil}, ListNode{3, nil}}
	nodes[0].Next = &nodes[1]
	nodes[1].Next = &nodes[2]
	nodes[2].Next = &nodes[3]

	printList(deleteDuplicates(&nodes[0]))
}

func test4() {
	nodes := []ListNode{ListNode{1, nil}, ListNode{1, nil}}
	nodes[0].Next = &nodes[1]

	printList(deleteDuplicates(&nodes[0]))
}

func main() {
	// test1()
	// test2()
	// test3()
	test4()
}
