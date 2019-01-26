package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	var ltHead, ltTail *ListNode
	var geHead, geTail *ListNode

	for ; head != nil; head = head.Next {
		if head.Val < x {
			if ltHead == nil {
				ltHead, ltTail = head, head
				continue
			}

			ltTail.Next = head
			ltTail = head
		} else {
			if geHead == nil {
				geHead, geTail = head, head
				continue
			}

			geTail.Next = head
			geTail = head
		}
	}

	if geTail != nil {
		geTail.Next = nil
	}

	if ltTail == nil {
		return geHead
	}

	ltTail.Next = geHead

	return ltHead
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
	nodes := []ListNode{ListNode{1, nil}, ListNode{4, nil}, ListNode{3, nil}, ListNode{2, nil}}
	nodes[0].Next = &nodes[1]
	nodes[1].Next = &nodes[2]
	nodes[2].Next = &nodes[3]

	printList(partition(&nodes[0], 3))
}

func main() {
	test1()
}
